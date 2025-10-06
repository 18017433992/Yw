//SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";
import "./PriceOracleManager.sol";

/**
 * @title EnhancedNftAuction
 * @dev Individual auction contract with enhanced Chainlink oracle integration
 * Supports multiple payment tokens with real-time USD conversion
 */
contract EnhancedNftAuction is 
    Initializable, 
    UUPSUpgradeable, 
    OwnableUpgradeable
{
    
    struct Bid {
        address bidder;
        address paymentToken;
        uint256 amount;
        uint256 usdValue; // USD value at time of bid (8 decimals)
        uint256 timestamp;
    }
    
    struct AuctionInfo {
        // Basic auction info
        address seller;
        address nftContract;
        uint256 tokenId;
        uint256 startTime;
        uint256 endTime;
        uint256 reservePrice; // In USD (8 decimals)
        
        // Status
        bool started;
        bool ended;
        bool cancelled;
        
        // Bidding info
        Bid highestBid;
        uint256 bidCount;
        
        // Settings
        uint256 minBidIncrement; // In USD (8 decimals)
        uint256 bidExtensionTime; // Time to extend if bid placed near end
        address[] allowedTokens; // Allowed payment tokens
    }
    
    // Auction information
    AuctionInfo public auction;
    
    // Oracle manager for price feeds
    PriceOracleManager public priceOracle;
    
    // Factory contract that created this auction
    address public factory;
    
    // Bid history
    Bid[] public bidHistory;
    mapping(address => uint256) public bidderRefunds;
    
    // Constants
    uint256 public constant MINIMUM_AUCTION_DURATION = 1 hours;
    uint256 public constant MAXIMUM_AUCTION_DURATION = 30 days;
    uint256 public constant DEFAULT_BID_EXTENSION = 10 minutes;
    uint256 public constant MINIMUM_BID_INCREMENT = 50 * 10**6; // $0.50 in USD (8 decimals)
    
    // Events
    event AuctionCreated(
        address indexed seller,
        address indexed nftContract,
        uint256 indexed tokenId,
        uint256 startTime,
        uint256 endTime,
        uint256 reservePrice
    );
    event BidPlaced(
        address indexed bidder,
        address indexed paymentToken,
        uint256 amount,
        uint256 usdValue,
        uint256 bidIndex
    );
    event AuctionEnded(
        address indexed winner,
        address indexed paymentToken,
        uint256 amount,
        uint256 usdValue
    );
    event AuctionCancelled();
    event AuctionExtended(uint256 newEndTime);
    event RefundClaimed(address indexed bidder, address token, uint256 amount);
    
    modifier onlyFactory() {
        require(msg.sender == factory, "Only factory can call");
        _;
    }
    
    modifier auctionActive() {
        require(auction.started && !auction.ended && !auction.cancelled, "Auction not active");
        require(block.timestamp >= auction.startTime, "Auction not started");
        require(block.timestamp <= auction.endTime, "Auction ended");
        _;
    }
    
    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }
    
    function initialize(
        address _seller,
        address _nftContract,
        uint256 _tokenId,
        uint256 _duration,
        uint256 _reservePrice,
        address _priceOracle,
        address _factory,
        address[] memory _allowedTokens
    ) public initializer {
        __Ownable_init(_seller);
        __UUPSUpgradeable_init();
        
        require(_seller != address(0), "Invalid seller");
        require(_nftContract != address(0), "Invalid NFT contract");
        require(_duration >= MINIMUM_AUCTION_DURATION && _duration <= MAXIMUM_AUCTION_DURATION, "Invalid duration");
        require(_reservePrice > 0, "Reserve price must be > 0");
        require(_priceOracle != address(0), "Invalid oracle");
        require(_factory != address(0), "Invalid factory");
        require(_allowedTokens.length > 0, "No payment tokens allowed");
        
        // Verify NFT ownership
        require(IERC721(_nftContract).ownerOf(_tokenId) == _seller, "Seller not owner");
        
        priceOracle = PriceOracleManager(_priceOracle);
        factory = _factory;
        
        // Validate allowed tokens have price feeds
        for (uint256 i = 0; i < _allowedTokens.length; i++) {
            require(priceOracle.isSupportedToken(_allowedTokens[i]), "Token not supported by oracle");
        }
        
        auction = AuctionInfo({
            seller: _seller,
            nftContract: _nftContract,
            tokenId: _tokenId,
            startTime: block.timestamp,
            endTime: block.timestamp + _duration,
            reservePrice: _reservePrice,
            started: false,
            ended: false,
            cancelled: false,
            highestBid: Bid(address(0), address(0), 0, 0, 0),
            bidCount: 0,
            minBidIncrement: MINIMUM_BID_INCREMENT,
            bidExtensionTime: DEFAULT_BID_EXTENSION,
            allowedTokens: _allowedTokens
        });
        
        emit AuctionCreated(_seller, _nftContract, _tokenId, auction.startTime, auction.endTime, _reservePrice);
    }
    
    /**
     * @dev Start the auction (transfer NFT to contract)
     */
    function startAuction() external {
        require(msg.sender == auction.seller || msg.sender == factory, "Only seller or factory");
        require(!auction.started, "Already started");
        require(!auction.cancelled, "Auction cancelled");
        require(block.timestamp <= auction.endTime, "Auction expired");
        
        // Transfer NFT to this contract
        IERC721(auction.nftContract).transferFrom(auction.seller, address(this), auction.tokenId);
        auction.started = true;
    }
    
    /**
     * @dev Place a bid with specified payment token
     */
    function placeBid(address paymentToken, uint256 amount) 
        external 
        payable   
        auctionActive 
    {
        require(amount > 0, "Bid amount must be > 0");
        require(_isTokenAllowed(paymentToken), "Payment token not allowed");
        
        uint8 tokenDecimals;
        uint256 actualAmount;
        
        if (paymentToken == address(0)) {
            // ETH payment
            require(msg.value > 0, "Must send ETH");
            actualAmount = msg.value;
            tokenDecimals = 18;
        } else {
            // ERC20 payment
            require(msg.value == 0, "Don't send ETH for token payment");
            actualAmount = amount;
            tokenDecimals = IERC20Metadata(paymentToken).decimals();
            
            // Transfer tokens to this contract
            IERC20(paymentToken).transferFrom(msg.sender, address(this), actualAmount);
        }
        
        // Convert bid to USD value
        uint256 usdValue = priceOracle.convertToUSD(paymentToken, actualAmount, tokenDecimals);
        require(usdValue >= auction.reservePrice, "Bid below reserve price");
        
        // Check minimum bid increment
        if (auction.bidCount > 0) {
            require(
                usdValue >= auction.highestBid.usdValue + auction.minBidIncrement,
                "Bid increment too small"
            );
        }
        
        // Refund previous highest bidder
        if (auction.highestBid.bidder != address(0)) {
            _addRefund(auction.highestBid.bidder, auction.highestBid.paymentToken, auction.highestBid.amount);
        }
        
        // Create new bid
        Bid memory newBid = Bid({
            bidder: msg.sender,
            paymentToken: paymentToken,
            amount: actualAmount,
            usdValue: usdValue,
            timestamp: block.timestamp
        });
        
        auction.highestBid = newBid;
        auction.bidCount++;
        bidHistory.push(newBid);
        
        // Extend auction if bid placed near end
        if (block.timestamp + auction.bidExtensionTime > auction.endTime) {
            auction.endTime = block.timestamp + auction.bidExtensionTime;
            emit AuctionExtended(auction.endTime);
        }
        
        emit BidPlaced(msg.sender, paymentToken, actualAmount, usdValue, bidHistory.length - 1);
    }
    
    /**
     * @dev End the auction and transfer NFT to winner
     */
    function endAuction() external {
        require(auction.started, "Auction not started");
        require(!auction.ended && !auction.cancelled, "Auction already ended");
        require(block.timestamp > auction.endTime, "Auction still active");
        
        auction.ended = true;
        
        if (auction.highestBid.bidder != address(0)) {
            // Transfer NFT to winner
            IERC721(auction.nftContract).safeTransferFrom(
                address(this), 
                auction.highestBid.bidder, 
                auction.tokenId
            );
            
            // Transfer payment to seller (minus any fees)
            _transferPaymentToSeller(
                auction.highestBid.paymentToken,
                auction.highestBid.amount
            );
            
            emit AuctionEnded(
                auction.highestBid.bidder,
                auction.highestBid.paymentToken,
                auction.highestBid.amount,
                auction.highestBid.usdValue
            );
        } else {
            // No bids - return NFT to seller
            IERC721(auction.nftContract).safeTransferFrom(
                address(this), 
                auction.seller, 
                auction.tokenId
            );
            
            emit AuctionEnded(address(0), address(0), 0, 0);
        }
    }
    
    /**
     * @dev Cancel auction (only before first bid)
     */
    function cancelAuction() external {
        require(msg.sender == auction.seller || msg.sender == owner(), "Only seller or owner");
        require(auction.started, "Auction not started");
        require(!auction.ended, "Auction already ended");
        require(auction.bidCount == 0, "Cannot cancel with bids");
        
        auction.cancelled = true;
        
        // Return NFT to seller
        IERC721(auction.nftContract).safeTransferFrom(
            address(this), 
            auction.seller, 
            auction.tokenId
        );
        
        emit AuctionCancelled();
    }
    
    /**
     * @dev Claim refund for losing bids
     */
    function claimRefund() external {
        uint256 ethRefund = bidderRefunds[msg.sender];
        require(ethRefund > 0, "No refund available");
        
        bidderRefunds[msg.sender] = 0;
        
        // For now, only handle ETH refunds. ERC20 refunds would need separate tracking
        if (ethRefund > 0) {
            payable(msg.sender).transfer(ethRefund);
            emit RefundClaimed(msg.sender, address(0), ethRefund);
        }
    }
    
    /**
     * @dev Get auction status
     */
    function getAuctionInfo() external view returns (
        address seller,
        address nftContract,
        uint256 tokenId,
        uint256 startTime,
        uint256 endTime,
        uint256 reservePrice,
        bool started,
        bool ended,
        bool cancelled,
        uint256 bidCount,
        address highestBidder,
        uint256 highestBidUSD
    ) {
        return (
            auction.seller,
            auction.nftContract,
            auction.tokenId,
            auction.startTime,
            auction.endTime,
            auction.reservePrice,
            auction.started,
            auction.ended,
            auction.cancelled,
            auction.bidCount,
            auction.highestBid.bidder,
            auction.highestBid.usdValue
        );
    }
    
    /**
     * @dev Get bid history
     */
    function getBidHistory() external view returns (Bid[] memory) {
        return bidHistory;
    }
    
    /**
     * @dev Get allowed payment tokens
     */
    function getAllowedTokens() external view returns (address[] memory) {
        return auction.allowedTokens;
    }
    
    /**
     * @dev Internal function to check if token is allowed
     */
    function _isTokenAllowed(address token) internal view returns (bool) {
        for (uint256 i = 0; i < auction.allowedTokens.length; i++) {
            if (auction.allowedTokens[i] == token) {
                return true;
            }
        }
        return false;
    }
    
    /**
     * @dev Internal function to add refund
     */
    function _addRefund(address bidder, address token, uint256 amount) internal {
        if (token == address(0)) {
            // ETH refund
            bidderRefunds[bidder] += amount;
        } else {
            // ERC20 refund - transfer immediately
            IERC20(token).transfer(bidder, amount);
        }
    }
    
    /**
     * @dev Internal function to transfer payment to seller
     */
    function _transferPaymentToSeller(address token, uint256 amount) internal {
        // In a production system, you might deduct marketplace fees here
        if (token == address(0)) {
            payable(auction.seller).transfer(amount);
        } else {
            IERC20(token).transfer(auction.seller, amount);
        }
    }
    

    
    /**
     * @dev Update minimum bid increment
     */
    function setMinBidIncrement(uint256 newIncrement) external onlyOwner {
        require(newIncrement >= MINIMUM_BID_INCREMENT, "Increment too small");
        auction.minBidIncrement = newIncrement;
    }
    
    /**
     * @dev Check if auction is active
     */
    function isActive() external view returns (bool) {
        return auction.started && 
               !auction.ended && 
               !auction.cancelled && 
               block.timestamp >= auction.startTime && 
               block.timestamp <= auction.endTime;
    }
    
    function _authorizeUpgrade(address newImplementation) 
        internal 
        override 
        onlyOwner 
    {}
    
    // Receive function to handle direct ETH transfers
    receive() external payable {
        // Allow direct ETH transfers for simplicity
    }
}