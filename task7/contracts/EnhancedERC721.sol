//SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/Counters.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";

/**
 * @title EnhancedERC721
 * @dev Enhanced ERC721 contract optimized for auction marketplace
 * Features: batch minting, royalty support, auction integration
 */
contract EnhancedERC721 is ERC721Enumerable, ERC721URIStorage, Ownable, ReentrancyGuard {
    using Counters for Counters.Counter;
    
    Counters.Counter private _tokenIds;
    
    // Royalty information
    struct RoyaltyInfo {
        address recipient;
        uint96 royaltyFraction; // Basis points (10000 = 100%)
    }
    
    mapping(uint256 => RoyaltyInfo) private _tokenRoyalties;
    RoyaltyInfo private _defaultRoyalty;
    
    // Auction marketplace contract address
    address public auctionFactory;
    
    // Events
    event TokenMinted(address indexed to, uint256 indexed tokenId, string tokenURI);
    event BatchMinted(address indexed to, uint256[] tokenIds, string[] tokenURIs);
    event RoyaltySet(uint256 indexed tokenId, address recipient, uint96 royaltyFraction);
    event DefaultRoyaltySet(address recipient, uint96 royaltyFraction);
    event AuctionFactorySet(address indexed newFactory);
    
    constructor(
        string memory name,
        string memory symbol,
        address _owner
    ) ERC721(name, symbol) Ownable(_owner) {
        // Set default royalty to 2.5%
        _setDefaultRoyalty(_owner, 250);
    }
    
    /**
     * @dev Mint a single NFT
     */
    function mint(
        address to,
        string memory tokenURI
    ) public onlyOwner returns (uint256) {
        _tokenIds.increment();
        uint256 tokenId = _tokenIds.current();
        
        _mint(to, tokenId);
        _setTokenURI(tokenId, tokenURI);
        
        emit TokenMinted(to, tokenId, tokenURI);
        return tokenId;
    }
    
    /**
     * @dev Batch mint NFTs for efficiency
     */
    function batchMint(
        address to,
        string[] memory tokenURIs
    ) public onlyOwner returns (uint256[] memory) {
        uint256 length = tokenURIs.length;
        require(length > 0 && length <= 100, "Invalid batch size");
        
        uint256[] memory tokenIds = new uint256[](length);
        
        for (uint256 i = 0; i < length; i++) {
            _tokenIds.increment();
            uint256 tokenId = _tokenIds.current();
            
            _mint(to, tokenId);
            _setTokenURI(tokenId, tokenURIs[i]);
            tokenIds[i] = tokenId;
        }
        
        emit BatchMinted(to, tokenIds, tokenURIs);
        return tokenIds;
    }
    
    /**
     * @dev Set auction factory contract
     */
    function setAuctionFactory(address _auctionFactory) external onlyOwner {
        require(_auctionFactory != address(0), "Invalid factory address");
        auctionFactory = _auctionFactory;
        emit AuctionFactorySet(_auctionFactory);
    }
    
    /**
     * @dev Approve auction factory to manage all tokens (for convenience)
     */
    function approveFactory() external {
        require(auctionFactory != address(0), "Factory not set");
        setApprovalForAll(auctionFactory, true);
    }
    
    /**
     * @dev Set royalty for specific token
     */
    function setTokenRoyalty(
        uint256 tokenId,
        address recipient,
        uint96 royaltyFraction
    ) external onlyOwner {
        require(_exists(tokenId), "Token does not exist");
        require(royaltyFraction <= 1000, "Royalty too high"); // Max 10%
        
        _tokenRoyalties[tokenId] = RoyaltyInfo(recipient, royaltyFraction);
        emit RoyaltySet(tokenId, recipient, royaltyFraction);
    }
    
    /**
     * @dev Set default royalty for all tokens
     */
    function setDefaultRoyalty(address recipient, uint96 royaltyFraction) external onlyOwner {
        require(royaltyFraction <= 1000, "Royalty too high"); // Max 10%
        _setDefaultRoyalty(recipient, royaltyFraction);
    }
    
    /**
     * @dev Internal function to set default royalty
     */
    function _setDefaultRoyalty(address recipient, uint96 royaltyFraction) internal {
        _defaultRoyalty = RoyaltyInfo(recipient, royaltyFraction);
        emit DefaultRoyaltySet(recipient, royaltyFraction);
    }
    
    /**
     * @dev Get royalty information for a token
     */
    function royaltyInfo(uint256 tokenId, uint256 salePrice) 
        external 
        view 
        returns (address, uint256) 
    {
        RoyaltyInfo memory royalty = _tokenRoyalties[tokenId];
        
        if (royalty.recipient == address(0)) {
            royalty = _defaultRoyalty;
        }
        
        uint256 royaltyAmount = (salePrice * royalty.royaltyFraction) / 10000;
        return (royalty.recipient, royaltyAmount);
    }
    
    /**
     * @dev Get total number of minted tokens
     */
    function totalMinted() external view returns (uint256) {
        return _tokenIds.current();
    }
    
    /**
     * @dev Get tokens owned by address
     */
    function tokensOfOwner(address owner) external view returns (uint256[] memory) {
        uint256 tokenCount = balanceOf(owner);
        uint256[] memory tokenIds = new uint256[](tokenCount);
        
        for (uint256 i = 0; i < tokenCount; i++) {
            tokenIds[i] = tokenOfOwnerByIndex(owner, i);
        }
        
        return tokenIds;
    }
    
    // Required overrides
    function _beforeTokenTransfer(
        address from,
        address to,
        uint256 tokenId,
        uint256 batchSize
    ) internal override(ERC721, ERC721Enumerable) {
        super._beforeTokenTransfer(from, to, tokenId, batchSize);
    }
    
    function _burn(uint256 tokenId) internal override(ERC721, ERC721URIStorage) {
        super._burn(tokenId);
        delete _tokenRoyalties[tokenId];
    }
    
    function tokenURI(uint256 tokenId)
        public
        view
        override(ERC721, ERC721URIStorage)
        returns (string memory)
    {
        return super.tokenURI(tokenId);
    }
    
    function supportsInterface(bytes4 interfaceId)
        public
        view
        override(ERC721, ERC721Enumerable, ERC721URIStorage)
        returns (bool)
    {
        return super.supportsInterface(interfaceId);
    }
}