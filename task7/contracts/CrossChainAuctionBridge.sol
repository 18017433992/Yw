//SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@chainlink/contracts-ccip/src/v0.8/ccip/interfaces/IRouterClient.sol";
import  { Client }  "@chainlink/contracts-ccip/src/v0.8/ccip/libraries/Client.sol";
import "@chainlink/contracts-ccip/src/v0.8/ccip/applications/CCIPReceiver.sol";

/**
 * @title CrossChainAuctionBridge
 * @dev Handles cross-chain auction functionality using Chainlink CCIP
 * Allows users on different chains to participate in auctions
 */
contract CrossChainAuctionBridge is 
    Initializable,
    UUPSUpgradeable,
    OwnableUpgradeable,
    ReentrancyGuardUpgradeable,
    CCIPReceiver 
{
    
    enum MessageType {
        CROSS_CHAIN_BID,
        AUCTION_ENDED,
        NFT_TRANSFER,
        PAYMENT_TRANSFER
    }
    
    struct CrossChainBid {
        address bidder;
        address auctionContract;
        address paymentToken;
        uint256 amount;
        uint256 usdValue;
        uint64 sourceChainSelector;
        bytes32 messageId;
    }
    
    struct CrossChainAuction {
        address auctionContract;
        uint64 homeChainSelector;
        mapping(uint64 => bool) allowedChains;
        uint64[] allowedChainsList;
        bool isActive;
    }
    
    // CCIP Router
    IRouterClient public router;
    
    // Supported chains and their selectors
    mapping(uint64 => bool) public supportedChains;
    mapping(uint64 => address) public bridgeContracts; // Chain selector => bridge contract address
    uint64[] public supportedChainsList;
    
    // Cross-chain auction management
    mapping(address => CrossChainAuction) public crossChainAuctions;
    mapping(bytes32 => CrossChainBid) public crossChainBids;
    mapping(address => mapping(uint64 => uint256)) public chainBidCounts;
    
    // Message tracking
    mapping(bytes32 => bool) public processedMessages;
    
    // Gas limits for different message types
    mapping(MessageType => uint256) public gasLimits;
    
    // Fees and limits
    uint256 public crossChainFeeMultiplier = 150; // 1.5x multiplier for cross-chain fees
    uint256 public maxCrossChainBid = 1000000 * 10**8; // $1M USD limit
    
    // Events
    event CrossChainBidSent(
        bytes32 indexed messageId,
        address indexed bidder,
        address indexed auctionContract,
        uint64 destinationChain,
        uint256 amount
    );
    event CrossChainBidReceived(
        bytes32 indexed messageId,
        address indexed bidder,
        address indexed auctionContract,
        uint64 sourceChain,
        uint256 amount
    );
    event AuctionRegisteredForCrossChain(
        address indexed auctionContract,
        uint64[] allowedChains
    );
    event ChainSupportAdded(uint64 chainSelector, address bridgeContract);
    event ChainSupportRemoved(uint64 chainSelector);
    event CrossChainAuctionEnded(
        address indexed auctionContract,
        address indexed winner,
        uint64 winnerChain
    );
    
    modifier onlySupportedChain(uint64 chainSelector) {
        require(supportedChains[chainSelector], "Chain not supported");
        _;
    }
    
    modifier onlyCrossChainAuction(address auctionContract) {
        require(crossChainAuctions[auctionContract].isActive, "Not cross-chain auction");
        _;
    }
    
    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }
    
    function initialize(address _router) public initializer {
        __Ownable_init(msg.sender);
        __UUPSUpgradeable_init();
        __ReentrancyGuard_init();
        __CCIPReceiver_init(_router);
        
        router = IRouterClient(_router);
        
        // Set default gas limits
        gasLimits[MessageType.CROSS_CHAIN_BID] = 300000;
        gasLimits[MessageType.AUCTION_ENDED] = 200000;
        gasLimits[MessageType.NFT_TRANSFER] = 400000;
        gasLimits[MessageType.PAYMENT_TRANSFER] = 300000;
    }
    
    /**
     * @dev Register an auction for cross-chain participation
     */
    function registerAuctionForCrossChain(
        address auctionContract,
        uint64[] memory allowedChains
    ) external onlyOwner {
        require(auctionContract != address(0), "Invalid auction contract");
        require(allowedChains.length > 0, "Must specify allowed chains");
        
        CrossChainAuction storage auction = crossChainAuctions[auctionContract];
        auction.auctionContract = auctionContract;
        auction.homeChainSelector = _getCurrentChainSelector();
        auction.isActive = true;
        
        // Clear existing allowed chains
        for (uint256 i = 0; i < auction.allowedChainsList.length; i++) {
            auction.allowedChains[auction.allowedChainsList[i]] = false;
        }
        delete auction.allowedChainsList;
        
        // Set new allowed chains
        for (uint256 i = 0; i < allowedChains.length; i++) {
            require(supportedChains[allowedChains[i]], "Chain not supported");
            auction.allowedChains[allowedChains[i]] = true;
            auction.allowedChainsList.push(allowedChains[i]);
        }
        
        emit AuctionRegisteredForCrossChain(auctionContract, allowedChains);
    }
    
    /**
     * @dev Send cross-chain bid
     */
    function sendCrossChainBid(
        uint64 destinationChainSelector,
        address auctionContract,
        address paymentToken,
        uint256 amount,
        uint256 usdValue
    ) external payable nonReentrant onlySupportedChain(destinationChainSelector) {
        require(amount > 0, "Invalid bid amount");
        require(usdValue <= maxCrossChainBid, "Bid exceeds maximum");
        require(bridgeContracts[destinationChainSelector] != address(0), "No bridge on destination");
        
        // Transfer payment token to this contract
        if (paymentToken != address(0)) {
            IERC20(paymentToken).transferFrom(msg.sender, address(this), amount);
        } else {
            require(msg.value >= amount, "Insufficient ETH sent");
        }
        
        // Prepare CCIP message
        bytes memory data = abi.encode(
            MessageType.CROSS_CHAIN_BID,
            msg.sender,
            auctionContract,
            paymentToken,
            amount,
            usdValue,
            _getCurrentChainSelector()
        );
        
        Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
            receiver: abi.encode(bridgeContracts[destinationChainSelector]),
            data: data,
            tokenAmounts: new Client.EVMTokenAmount[](0),
            extraArgs: Client._argsToBytes(
                Client.EVMExtraArgsV1({gasLimit: gasLimits[MessageType.CROSS_CHAIN_BID]})
            ),
            feeToken: address(0) // Pay fees in native token
        });
        
        // Calculate and pay CCIP fees
        uint256 fees = router.getFee(destinationChainSelector, message);
        require(msg.value >= fees + (paymentToken == address(0) ? amount : 0), "Insufficient fee");
        
        bytes32 messageId = router.ccipSend{value: fees}(destinationChainSelector, message);
        
        // Store bid information
        crossChainBids[messageId] = CrossChainBid({
            bidder: msg.sender,
            auctionContract: auctionContract,
            paymentToken: paymentToken,
            amount: amount,
            usdValue: usdValue,
            sourceChainSelector: _getCurrentChainSelector(),
            messageId: messageId
        });
        
        chainBidCounts[auctionContract][destinationChainSelector]++;
        
        emit CrossChainBidSent(messageId, msg.sender, auctionContract, destinationChainSelector, amount);
    }
    
    /**
     * @dev Handle incoming CCIP messages
     */
    function _ccipReceive(Client.Any2EVMMessage memory any2EvmMessage) internal override {
        bytes32 messageId = any2EvmMessage.messageId;
        
        // Prevent duplicate processing
        require(!processedMessages[messageId], "Message already processed");
        processedMessages[messageId] = true;
        
        uint64 sourceChainSelector = any2EvmMessage.sourceChainSelector;
        require(supportedChains[sourceChainSelector], "Source chain not supported");
        
        (
            MessageType messageType,
            address bidder,
            address auctionContract,
            address paymentToken,
            uint256 amount,
            uint256 usdValue,
            uint64 originChain
        ) = abi.decode(any2EvmMessage.data, (MessageType, address, address, address, uint256, uint256, uint64));
        
        if (messageType == MessageType.CROSS_CHAIN_BID) {
            _processCrossChainBid(
                messageId,
                bidder,
                auctionContract,
                paymentToken,
                amount,
                usdValue,
                sourceChainSelector
            );
        } else if (messageType == MessageType.AUCTION_ENDED) {
            _processAuctionEnded(auctionContract, bidder, sourceChainSelector);
        }
        // Add more message type handlers as needed
    }
    
    /**
     * @dev Process incoming cross-chain bid
     */
    function _processCrossChainBid(
        bytes32 messageId,
        address bidder,
        address auctionContract,
        address paymentToken,
        uint256 amount,
        uint256 usdValue,
        uint64 sourceChainSelector
    ) internal onlyCrossChainAuction(auctionContract) {
        CrossChainAuction storage auction = crossChainAuctions[auctionContract];
        require(auction.allowedChains[sourceChainSelector], "Chain not allowed for this auction");
        
        // Store cross-chain bid
        crossChainBids[messageId] = CrossChainBid({
            bidder: bidder,
            auctionContract: auctionContract,
            paymentToken: paymentToken,
            amount: amount,
            usdValue: usdValue,
            sourceChainSelector: sourceChainSelector,
            messageId: messageId
        });
        
        // TODO: Integrate with auction contract to place bid
        // This would require the auction contract to support cross-chain bids
        
        emit CrossChainBidReceived(messageId, bidder, auctionContract, sourceChainSelector, amount);
    }
    
    /**
     * @dev Process auction ended message
     */
    function _processAuctionEnded(
        address auctionContract,
        address winner,
        uint64 winnerChain
    ) internal {
        emit CrossChainAuctionEnded(auctionContract, winner, winnerChain);
        
        // TODO: Handle cross-chain settlement
        // This might involve transferring NFT or payment across chains
    }
    
    /**
     * @dev Add support for a new chain
     */
    function addChainSupport(
        uint64 chainSelector,
        address bridgeContract
    ) external onlyOwner {
        require(!supportedChains[chainSelector], "Chain already supported");
        require(bridgeContract != address(0), "Invalid bridge contract");
        
        supportedChains[chainSelector] = true;
        bridgeContracts[chainSelector] = bridgeContract;
        supportedChainsList.push(chainSelector);
        
        emit ChainSupportAdded(chainSelector, bridgeContract);
    }
    
    /**
     * @dev Remove chain support
     */
    function removeChainSupport(uint64 chainSelector) external onlyOwner {
        require(supportedChains[chainSelector], "Chain not supported");
        
        supportedChains[chainSelector] = false;
        delete bridgeContracts[chainSelector];
        
        // Remove from supported chains list
        for (uint256 i = 0; i < supportedChainsList.length; i++) {
            if (supportedChainsList[i] == chainSelector) {
                supportedChainsList[i] = supportedChainsList[supportedChainsList.length - 1];
                supportedChainsList.pop();
                break;
            }
        }
        
        emit ChainSupportRemoved(chainSelector);
    }
    
    /**
     * @dev Update gas limit for message type
     */
    function updateGasLimit(MessageType messageType, uint256 newLimit) external onlyOwner {
        require(newLimit >= 100000 && newLimit <= 2000000, "Invalid gas limit");
        gasLimits[messageType] = newLimit;
    }
    
    /**
     * @dev Update cross-chain fee multiplier
     */
    function updateCrossChainFeeMultiplier(uint256 newMultiplier) external onlyOwner {
        require(newMultiplier >= 100 && newMultiplier <= 500, "Invalid multiplier"); // 1x to 5x
        crossChainFeeMultiplier = newMultiplier;
    }
    
    /**
     * @dev Update max cross-chain bid
     */
    function updateMaxCrossChainBid(uint256 newMax) external onlyOwner {
        require(newMax > 0, "Max must be > 0");
        maxCrossChainBid = newMax;
    }
    
    /**
     * @dev Get cross-chain fee estimate
     */
    function getCrossChainFee(
        uint64 destinationChainSelector,
        address auctionContract,
        uint256 amount
    ) external view returns (uint256) {
        require(supportedChains[destinationChainSelector], "Chain not supported");
        
        bytes memory data = abi.encode(
            MessageType.CROSS_CHAIN_BID,
            msg.sender,
            auctionContract,
            address(0), // ETH
            amount,
            amount, // Simplified USD value
            _getCurrentChainSelector()
        );
        
        Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
            receiver: abi.encode(bridgeContracts[destinationChainSelector]),
            data: data,
            tokenAmounts: new Client.EVMTokenAmount[](0),
            extraArgs: Client._argsToBytes(
                Client.EVMExtraArgsV1({gasLimit: gasLimits[MessageType.CROSS_CHAIN_BID]})
            ),
            feeToken: address(0)
        });
        
        uint256 baseFee = router.getFee(destinationChainSelector, message);
        return (baseFee * crossChainFeeMultiplier) / 100;
    }
    
    /**
     * @dev Get supported chains
     */
    function getSupportedChains() external view returns (uint64[] memory) {
        return supportedChainsList;
    }
    
    /**
     * @dev Get cross-chain auction info
     */
    function getCrossChainAuctionInfo(address auctionContract) 
        external 
        view 
        returns (
            uint64 homeChainSelector,
            uint64[] memory allowedChains,
            bool isActive
        ) 
    {
        CrossChainAuction storage auction = crossChainAuctions[auctionContract];
        return (
            auction.homeChainSelector,
            auction.allowedChainsList,
            auction.isActive
        );
    }
    
    /**
     * @dev Check if chain is allowed for auction
     */
    function isChainAllowedForAuction(
        address auctionContract,
        uint64 chainSelector
    ) external view returns (bool) {
        return crossChainAuctions[auctionContract].allowedChains[chainSelector];
    }
    
    /**
     * @dev Get current chain selector (implementation depends on network)
     */
    function _getCurrentChainSelector() internal view returns (uint64) {
        // This would return the actual chain selector for the current network
        // For demo purposes, returning a placeholder
        return 1; // Replace with actual implementation
    }
    
    /**
     * @dev Emergency withdrawal function
     */
    function emergencyWithdraw(address token, uint256 amount) external onlyOwner {
        if (token == address(0)) {
            payable(owner()).transfer(amount);
        } else {
            IERC20(token).transfer(owner(), amount);
        }
    }
    
    function _authorizeUpgrade(address newImplementation) 
        internal 
        override 
        onlyOwner 
    {}
    
    // Receive function to handle native token payments
    receive() external payable {}
}