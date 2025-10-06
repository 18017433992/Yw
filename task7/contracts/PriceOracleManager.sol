//SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@chainlink/contracts/src/v0.8/shared/interfaces/AggregatorV3Interface.sol";

/**
 * @title PriceOracleManager
 * @dev Manages Chainlink price feeds for multiple tokens with USD conversion
 * Supports ETH and ERC20 tokens with proper decimal handling
 */
contract PriceOracleManager is Initializable, UUPSUpgradeable, OwnableUpgradeable {
    
    struct PriceFeedInfo {
        AggregatorV3Interface priceFeed;
        uint8 decimals;
        bool isActive;
        string description;
    }
    
    // Token address => Price feed info
    mapping(address => PriceFeedInfo) public priceFeeds;
    
    // ETH address (0x0) for ETH price feed
    address public constant ETH_ADDRESS = address(0);
    
    // Supported tokens list
    address[] public supportedTokens;
    mapping(address => bool) public isSupportedToken;
    
    // Price staleness threshold (24 hours)
    uint256 public constant PRICE_STALENESS_THRESHOLD = 24 * 60 * 60;
    
    // Events
    event PriceFeedAdded(address indexed token, address indexed priceFeed, string description);
    event PriceFeedUpdated(address indexed token, address indexed priceFeed, string description);
    event PriceFeedRemoved(address indexed token);
    event PriceFeedStatusChanged(address indexed token, bool isActive);
    
    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }
    
    function initialize() public initializer {
        __Ownable_init(msg.sender);
        __UUPSUpgradeable_init();
    }
    
    /**
     * @dev Add or update a price feed for a token
     * @param token Token address (use address(0) for ETH)
     * @param priceFeed Chainlink price feed address
     * @param description Human readable description
     */
    function setPriceFeed(
        address token,
        address priceFeed,
        string memory description
    ) external onlyOwner {
        require(priceFeed != address(0), "Invalid price feed address");
        
        AggregatorV3Interface feed = AggregatorV3Interface(priceFeed);
        
        // Validate the price feed by calling it
        try feed.latestRoundData() returns (
            uint80,
            int256 price,
            uint256,
            uint256 updatedAt,
            uint80
        ) {
            require(price > 0, "Invalid price from feed");
            require(updatedAt > 0, "Price feed not updated");
        } catch {
            revert("Price feed validation failed");
        }
        
        uint8 decimals = feed.decimals();
        bool isNewToken = !isSupportedToken[token];
        
        priceFeeds[token] = PriceFeedInfo({
            priceFeed: feed,
            decimals: decimals,
            isActive: true,
            description: description
        });
        
        if (isNewToken) {
            supportedTokens.push(token);
            isSupportedToken[token] = true;
            emit PriceFeedAdded(token, priceFeed, description);
        } else {
            emit PriceFeedUpdated(token, priceFeed, description);
        }
    }
    
    /**
     * @dev Remove a price feed
     */
    function removePriceFeed(address token) external onlyOwner {
        require(isSupportedToken[token], "Token not supported");
        
        delete priceFeeds[token];
        isSupportedToken[token] = false;
        
        // Remove from supportedTokens array
        for (uint256 i = 0; i < supportedTokens.length; i++) {
            if (supportedTokens[i] == token) {
                supportedTokens[i] = supportedTokens[supportedTokens.length - 1];
                supportedTokens.pop();
                break;
            }
        }
        
        emit PriceFeedRemoved(token);
    }
    
    /**
     * @dev Toggle price feed active status
     */
    function setPriceFeedStatus(address token, bool isActive) external onlyOwner {
        require(isSupportedToken[token], "Token not supported");
        priceFeeds[token].isActive = isActive;
        emit PriceFeedStatusChanged(token, isActive);
    }
    
    /**
     * @dev Get latest price for a token in USD (8 decimals)
     * @param token Token address (use address(0) for ETH)
     * @return price Latest price in USD with 8 decimals
     */
    function getLatestPrice(address token) external view returns (int256 price) {
        PriceFeedInfo memory feedInfo = priceFeeds[token];
        require(feedInfo.isActive, "Price feed not active");
        
        (
            uint80 roundId,
            int256 rawPrice,
            uint256 startedAt,
            uint256 updatedAt,
            uint80 answeredInRound
        ) = feedInfo.priceFeed.latestRoundData();
        
        require(rawPrice > 0, "Invalid price");
        require(updatedAt > 0, "Price not updated");
        require(block.timestamp - updatedAt <= PRICE_STALENESS_THRESHOLD, "Price too stale");
        require(answeredInRound >= roundId, "Stale price round");
        
        // Normalize to 8 decimals
        if (feedInfo.decimals != 8) {
            if (feedInfo.decimals > 8) {
                price = rawPrice / int256(10 ** (feedInfo.decimals - 8));
            } else {
                price = rawPrice * int256(10 ** (8 - feedInfo.decimals));
            }
        } else {
            price = rawPrice;
        }
    }
    
    /**
     * @dev Convert token amount to USD value
     * @param token Token address
     * @param amount Token amount (in token's native decimals)
     * @param tokenDecimals Token's decimal places
     * @return usdValue USD value with 8 decimals
     */
    function convertToUSD(
        address token,
        uint256 amount,
        uint8 tokenDecimals
    ) external view returns (uint256 usdValue) {
        int256 price = this.getLatestPrice(token);
        require(price > 0, "Invalid price");
        
        // Calculate USD value: (amount * price) / (10^tokenDecimals)
        // Result has 8 decimals (price decimals)
        usdValue = (amount * uint256(price)) / (10 ** tokenDecimals);
    }
    
    /**
     * @dev Convert USD value to token amount
     * @param token Token address
     * @param usdValue USD value with 8 decimals
     * @param tokenDecimals Token's decimal places
     * @return amount Token amount
     */
    function convertFromUSD(
        address token,
        uint256 usdValue,
        uint8 tokenDecimals
    ) external view returns (uint256 amount) {
        int256 price = this.getLatestPrice(token);
        require(price > 0, "Invalid price");
        
        // Calculate token amount: (usdValue * 10^tokenDecimals) / price
        amount = (usdValue * (10 ** tokenDecimals)) / uint256(price);
    }
    
    /**
     * @dev Compare two bids in different tokens by converting to USD
     * @param token1 First token address
     * @param amount1 First token amount
     * @param decimals1 First token decimals
     * @param token2 Second token address
     * @param amount2 Second token amount
     * @param decimals2 Second token decimals
     * @return result 1 if first bid is higher, -1 if second is higher, 0 if equal
     */
    function compareBids(
        address token1,
        uint256 amount1,
        uint8 decimals1,
        address token2,
        uint256 amount2,
        uint8 decimals2
    ) external view returns (int8 result) {
        uint256 usdValue1 = this.convertToUSD(token1, amount1, decimals1);
        uint256 usdValue2 = this.convertToUSD(token2, amount2, decimals2);
        
        if (usdValue1 > usdValue2) {
            result = 1;
        } else if (usdValue1 < usdValue2) {
            result = -1;
        } else {
            result = 0;
        }
    }
    
    /**
     * @dev Get price feed information
     */
    function getPriceFeedInfo(address token) 
        external 
        view 
        returns (
            address priceFeedAddress,
            uint8 decimals,
            bool isActive,
            string memory description
        ) 
    {
        PriceFeedInfo memory feedInfo = priceFeeds[token];
        return (
            address(feedInfo.priceFeed),
            feedInfo.decimals,
            feedInfo.isActive,
            feedInfo.description
        );
    }
    
    /**
     * @dev Get all supported tokens
     */
    function getSupportedTokens() external view returns (address[] memory) {
        return supportedTokens;
    }
    
    /**
     * @dev Get number of supported tokens
     */
    function getSupportedTokensCount() external view returns (uint256) {
        return supportedTokens.length;
    }
    
    /**
     * @dev Check if price data is fresh
     */
    function isPriceFresh(address token) external view returns (bool) {
        if (!isSupportedToken[token] || !priceFeeds[token].isActive) {
            return false;
        }
        
        (, , , uint256 updatedAt, ) = priceFeeds[token].priceFeed.latestRoundData();
        return (block.timestamp - updatedAt) <= PRICE_STALENESS_THRESHOLD;
    }
    
    /**
     * @dev Batch get prices for multiple tokens
     */
    function batchGetPrices(address[] calldata tokens) 
        external 
        view 
        returns (int256[] memory prices) 
    {
        prices = new int256[](tokens.length);
        for (uint256 i = 0; i < tokens.length; i++) {
            prices[i] = this.getLatestPrice(tokens[i]);
        }
    }
    
    function _authorizeUpgrade(address newImplementation) 
        internal 
        override 
        onlyOwner 
    {}
}