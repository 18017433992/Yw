//SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts/proxy/Clones.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "./EnhancedNftAuction.sol";
import "./PriceOracleManager.sol";

/**
 * @title NftAuctionFactory
 * @dev Factory contract for creating and managing NFT auctions
 * Uses Uniswap V2-like factory pattern with deterministic auction addresses
 */
contract NftAuctionFactory is 
    Initializable, 
    UUPSUpgradeable, 
    OwnableUpgradeable
{
    
    // Auction implementation contract for cloning
    address public auctionImplementation;
    
    // Price oracle manager
    PriceOracleManager public priceOracle;
    
    // Auction tracking
    mapping(bytes32 => address) public getAuction; // Hash -> auction address
    mapping(address => bool) public isAuction; // auction address -> bool
    address[] public allAuctions;
    
    // Auction parameters and statistics
    mapping(address => address[]) public auctionsBySeller;
    mapping(address => address[]) public auctionsByNftContract;
    mapping(address => mapping(uint256 => address)) public auctionByNftToken;
    
    // Fee structure (basis points: 10000 = 100%)
    uint256 public platformFeeRate; // Platform fee (e.g., 250 = 2.5%)
    address public feeRecipient;
    
    // Default auction settings
    uint256 public defaultMinDuration = 1 hours;
    uint256 public defaultMaxDuration = 30 days;
    address[] public defaultAllowedTokens;
    
    // Auction counts and limits
    uint256 public totalAuctions;
    mapping(address => uint256) public userAuctionCount;
    uint256 public maxAuctionsPerUser = 100;
    
    // Events
    event AuctionCreated(
        address indexed seller,
        address indexed nftContract,
        uint256 indexed tokenId,
        address auction,
        uint256 auctionId
    );
    event AuctionImplementationUpdated(address oldImplementation, address newImplementation);
    event PriceOracleUpdated(address oldOracle, address newOracle);
    event PlatformFeeUpdated(uint256 oldFee, uint256 newFee);
    event FeeRecipientUpdated(address oldRecipient, address newRecipient);
    event DefaultTokensUpdated(address[] oldTokens, address[] newTokens);
    event AuctionStarted(address indexed auction, address indexed seller);
    event AuctionEnded(address indexed auction, address indexed winner, uint256 finalPrice);
    
    modifier validNftContract(address nftContract) {
        require(nftContract != address(0), "Invalid NFT contract");
        require(_isContract(nftContract), "Not a contract");
        _;
    }
    
    modifier onlyValidAuction(address auction) {
        require(isAuction[auction], "Not a valid auction");
        _;
    }
    
    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }
    
    function initialize(
        address _auctionImplementation,
        address _priceOracle,
        address _feeRecipient,
        uint256 _platformFeeRate
    ) public initializer {
        __Ownable_init(msg.sender);
        __UUPSUpgradeable_init();
        __ReentrancyGuard_init();
        __Pausable_init();
        
        require(_auctionImplementation != address(0), "Invalid implementation");
        require(_priceOracle != address(0), "Invalid oracle");
        require(_feeRecipient != address(0), "Invalid fee recipient");
        require(_platformFeeRate <= 1000, "Fee too high"); // Max 10%
        
        auctionImplementation = _auctionImplementation;
        priceOracle = PriceOracleManager(_priceOracle);
        feeRecipient = _feeRecipient;
        platformFeeRate = _platformFeeRate;
        
        // Set default allowed tokens (ETH)
        defaultAllowedTokens.push(address(0));
    }
    
    /**
     * @dev Create a new NFT auction
     * @param nftContract The NFT contract address
     * @param tokenId The NFT token ID
     * @param duration Auction duration in seconds
     * @param reservePrice Reserve price in USD (8 decimals)
     * @param allowedTokens Payment tokens allowed (empty array uses defaults)
     * @return auction The created auction contract address
     */
    function createAuction(
        address nftContract,
        uint256 tokenId,
        uint256 duration,
        uint256 reservePrice,
        address[] memory allowedTokens
    ) 
        external 
        nonReentrant 
        whenNotPaused 
        validNftContract(nftContract) 
        returns (address auction) 
    {
        require(duration >= defaultMinDuration && duration <= defaultMaxDuration, "Invalid duration");
        require(reservePrice > 0, "Reserve price must be > 0");
        require(userAuctionCount[msg.sender] < maxAuctionsPerUser, "Too many auctions");
        
        // Verify NFT ownership and approval
        require(IERC721(nftContract).ownerOf(tokenId) == msg.sender, "Not NFT owner");
        require(
            IERC721(nftContract).isApprovedForAll(msg.sender, address(this)) ||
            IERC721(nftContract).getApproved(tokenId) == address(this),
            "Factory not approved"
        );
        
        // Check for existing auction
        require(auctionByNftToken[nftContract][tokenId] == address(0), "Auction already exists");
        
        // Use default tokens if none specified
        address[] memory finalAllowedTokens = allowedTokens.length > 0 ? allowedTokens : defaultAllowedTokens;
        
        // Validate all tokens are supported by oracle
        for (uint256 i = 0; i < finalAllowedTokens.length; i++) {
            require(priceOracle.isSupportedToken(finalAllowedTokens[i]), "Token not supported");
        }
        
        // Create deterministic salt for auction address
        bytes32 salt = _computeSalt(msg.sender, nftContract, tokenId, block.timestamp);
        
        // Check if auction already exists (should not happen with timestamp in salt)
        require(getAuction[salt] == address(0), "Auction already exists");
        
        // Clone auction implementation
        auction = Clones.cloneDeterministic(auctionImplementation, salt);
        
        // Initialize the auction
        EnhancedNftAuction(auction).initialize(
            msg.sender,
            nftContract,
            tokenId,
            duration,
            reservePrice,
            address(priceOracle),
            address(this),
            finalAllowedTokens
        );
        
        // Update mappings
        getAuction[salt] = auction;
        isAuction[auction] = true;
        allAuctions.push(auction);
        auctionsBySeller[msg.sender].push(auction);
        auctionsByNftContract[nftContract].push(auction);
        auctionByNftToken[nftContract][tokenId] = auction;
        
        // Update counters
        totalAuctions++;
        userAuctionCount[msg.sender]++;
        
        emit AuctionCreated(msg.sender, nftContract, tokenId, auction, totalAuctions - 1);
        
        return auction;
    }
    
    /**
     * @dev Batch create multiple auctions
     */
    function batchCreateAuctions(
        address[] memory nftContracts,
        uint256[] memory tokenIds,
        uint256[] memory durations,
        uint256[] memory reservePrices,
        address[][] memory allowedTokensArray
    ) external returns (address[] memory auctions) {
        require(nftContracts.length == tokenIds.length, "Array length mismatch");
        require(nftContracts.length == durations.length, "Array length mismatch");
        require(nftContracts.length == reservePrices.length, "Array length mismatch");
        require(nftContracts.length <= 10, "Too many auctions at once");
        
        auctions = new address[](nftContracts.length);
        
        for (uint256 i = 0; i < nftContracts.length; i++) {
            address[] memory allowedTokens = allowedTokensArray.length > i ? 
                allowedTokensArray[i] : new address[](0);
                
            auctions[i] = createAuction(
                nftContracts[i],
                tokenIds[i],
                durations[i],
                reservePrices[i],
                allowedTokens
            );
        }
        
        return auctions;
    }
    
    /**
     * @dev Start an auction (transfer NFT to auction contract)
     */
    function startAuction(address auction) external onlyValidAuction(auction) {
        EnhancedNftAuction auctionContract = EnhancedNftAuction(auction);
        
        // Get auction info
        (address seller, address nftContract, uint256 tokenId, , , , , , , , , ) = 
            auctionContract.getAuctionInfo();
        
        require(msg.sender == seller, "Only seller can start");
        
        // Start the auction
        auctionContract.startAuction();
        
        emit AuctionStarted(auction, seller);
    }
    
    /**
     * @dev Predict auction address before creation
     */
    function predictAuctionAddress(
        address seller,
        address nftContract,
        uint256 tokenId,
        uint256 timestamp
    ) external view returns (address) {
        bytes32 salt = _computeSalt(seller, nftContract, tokenId, timestamp);
        return Clones.predictDeterministicAddress(auctionImplementation, salt, address(this));
    }
    
    /**
     * @dev Get all auctions
     */
    function getAllAuctions() external view returns (address[] memory) {
        return allAuctions;
    }
    
    /**
     * @dev Get auctions by seller
     */
    function getAuctionsBySeller(address seller) external view returns (address[] memory) {
        return auctionsBySeller[seller];
    }
    
    /**
     * @dev Get auctions by NFT contract
     */
    function getAuctionsByNftContract(address nftContract) external view returns (address[] memory) {
        return auctionsByNftContract[nftContract];
    }
    
    /**
     * @dev Get auction by specific NFT token
     */
    function getAuctionByNftToken(address nftContract, uint256 tokenId) 
        external 
        view 
        returns (address) 
    {
        return auctionByNftToken[nftContract][tokenId];
    }
    
    /**
     * @dev Get paginated auctions
     */
    function getAuctionsPaginated(uint256 offset, uint256 limit) 
        external 
        view 
        returns (address[] memory auctions, uint256 total) 
    {
        total = allAuctions.length;
        
        if (offset >= total) {
            return (new address[](0), total);
        }
        
        uint256 end = offset + limit;
        if (end > total) {
            end = total;
        }
        
        auctions = new address[](end - offset);
        for (uint256 i = offset; i < end; i++) {
            auctions[i - offset] = allAuctions[i];
        }
        
        return (auctions, total);
    }
    
    /**
     * @dev Get active auctions
     */
    function getActiveAuctions() external view returns (address[] memory) {
        address[] memory activeAuctions = new address[](allAuctions.length);
        uint256 count = 0;
        
        for (uint256 i = 0; i < allAuctions.length; i++) {
            if (EnhancedNftAuction(allAuctions[i]).isActive()) {
                activeAuctions[count] = allAuctions[i];
                count++;
            }
        }
        
        // Resize array
        address[] memory result = new address[](count);
        for (uint256 i = 0; i < count; i++) {
            result[i] = activeAuctions[i];
        }
        
        return result;
    }
    
    /**
     * @dev Update auction implementation
     */
    function updateAuctionImplementation(address newImplementation) external onlyOwner {
        require(newImplementation != address(0), "Invalid implementation");
        require(_isContract(newImplementation), "Not a contract");
        
        address oldImplementation = auctionImplementation;
        auctionImplementation = newImplementation;
        
        emit AuctionImplementationUpdated(oldImplementation, newImplementation);
    }
    
    /**
     * @dev Update price oracle
     */
    function updatePriceOracle(address newOracle) external onlyOwner {
        require(newOracle != address(0), "Invalid oracle");
        
        address oldOracle = address(priceOracle);
        priceOracle = PriceOracleManager(newOracle);
        
        emit PriceOracleUpdated(oldOracle, newOracle);
    }
    
    /**
     * @dev Update platform fee
     */
    function updatePlatformFee(uint256 newFeeRate) external onlyOwner {
        require(newFeeRate <= 1000, "Fee too high"); // Max 10%
        
        uint256 oldFee = platformFeeRate;
        platformFeeRate = newFeeRate;
        
        emit PlatformFeeUpdated(oldFee, newFeeRate);
    }
    
    /**
     * @dev Update fee recipient
     */
    function updateFeeRecipient(address newRecipient) external onlyOwner {
        require(newRecipient != address(0), "Invalid recipient");
        
        address oldRecipient = feeRecipient;
        feeRecipient = newRecipient;
        
        emit FeeRecipientUpdated(oldRecipient, newRecipient);
    }
    
    /**
     * @dev Update default allowed tokens
     */
    function updateDefaultAllowedTokens(address[] memory newTokens) external onlyOwner {
        require(newTokens.length > 0, "Must have at least one token");
        
        for (uint256 i = 0; i < newTokens.length; i++) {
            require(priceOracle.isSupportedToken(newTokens[i]), "Token not supported");
        }
        
        address[] memory oldTokens = defaultAllowedTokens;
        defaultAllowedTokens = newTokens;
        
        emit DefaultTokensUpdated(oldTokens, newTokens);
    }
    
    /**
     * @dev Update max auctions per user
     */
    function updateMaxAuctionsPerUser(uint256 newLimit) external onlyOwner {
        require(newLimit > 0, "Limit must be > 0");
        maxAuctionsPerUser = newLimit;
    }
    
    /**
     * @dev Emergency pause
     */
    function pause() external onlyOwner {
        _pause();
    }
    
    function unpause() external onlyOwner {
        _unpause();
    }
    
    /**
     * @dev Compute deterministic salt for auction creation
     */
    function _computeSalt(
        address seller,
        address nftContract,
        uint256 tokenId,
        uint256 timestamp
    ) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(seller, nftContract, tokenId, timestamp));
    }
    
    /**
     * @dev Check if address is a contract
     */
    function _isContract(address addr) internal view returns (bool) {
        uint256 size;
        assembly {
            size := extcodesize(addr)
        }
        return size > 0;
    }
    
    /**
     * @dev Get default allowed tokens
     */
    function getDefaultAllowedTokens() external view returns (address[] memory) {
        return defaultAllowedTokens;
    }
    
    function _authorizeUpgrade(address newImplementation) 
        internal 
        override 
        onlyOwner 
    {}
}