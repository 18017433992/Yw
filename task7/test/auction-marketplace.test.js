const { expect } = require("chai");
const { ethers, upgrades } = require("hardhat");
const { time } = require("@nomicfoundation/hardhat-network-helpers");

describe("NFT Auction Marketplace", function () {
  let priceOracle;
  let auctionFactory;
  let nftContract;
  let testToken;
  let auctionImplementation;
  
  let owner;
  let seller;
  let bidder1;
  let bidder2;
  let feeRecipient;
  
  const PLATFORM_FEE_RATE = 250; // 2.5%
  const ETH_USD_PRICE = 2000 * 10**8; // $2000 USD (8 decimals)
  const TOKEN_USD_PRICE = 1 * 10**8; // $1 USD (8 decimals)
  
  beforeEach(async function () {
    [owner, seller, bidder1, bidder2, feeRecipient] = await ethers.getSigners();
    
    // Deploy Mock Price Feed for testing
    const MockPriceFeed = await ethers.getContractFactory("MockV3Aggregator");
    
    // Deploy Price Oracle Manager
    const PriceOracleManager = await ethers.getContractFactory("PriceOracleManager");
    priceOracle = await upgrades.deployProxy(PriceOracleManager, [], {
      initializer: "initialize"
    });
    await priceOracle.waitForDeployment();
    
    // Deploy mock price feeds
    const ethPriceFeed = await MockPriceFeed.deploy(8, ETH_USD_PRICE);
    const tokenPriceFeed = await MockPriceFeed.deploy(8, TOKEN_USD_PRICE);
    
    // Set up price feeds
    await priceOracle.setPriceFeed(ethers.ZeroAddress, await ethPriceFeed.getAddress(), "ETH/USD");
    
    // Deploy Enhanced NFT Auction Implementation
    const EnhancedNftAuction = await ethers.getContractFactory("EnhancedNftAuction");
    auctionImplementation = await EnhancedNftAuction.deploy();
    await auctionImplementation.waitForDeployment();
    
    // Deploy NFT Auction Factory
    const NftAuctionFactory = await ethers.getContractFactory("NftAuctionFactory");
    auctionFactory = await upgrades.deployProxy(
      NftAuctionFactory,
      [
        await auctionImplementation.getAddress(),
        await priceOracle.getAddress(),
        feeRecipient.address,
        PLATFORM_FEE_RATE
      ],
      { initializer: "initialize" }
    );
    await auctionFactory.waitForDeployment();
    
    // Deploy Enhanced ERC721 NFT
    const EnhancedERC721 = await ethers.getContractFactory("EnhancedERC721");
    nftContract = await EnhancedERC721.deploy(
      "Test NFT",
      "TNFT",
      owner.address
    );
    await nftContract.waitForDeployment();
    
    // Deploy Test ERC20 Token
    const TestERC20 = await ethers.getContractFactory("TestERC20");
    testToken = await TestERC20.deploy(
      "Test USDC",
      "TUSDC",
      6,
      ethers.parseUnits("1000000", 6),
      owner.address
    );
    await testToken.waitForDeployment();
    
    // Set up price feed for test token
    await priceOracle.setPriceFeed(await testToken.getAddress(), await tokenPriceFeed.getAddress(), "TUSDC/USD");
    
    // Configure NFT contract
    await nftContract.setAuctionFactory(await auctionFactory.getAddress());
    
    // Mint test NFTs
    await nftContract.mint(seller.address, "https://example.com/nft/1.json");
    await nftContract.mint(seller.address, "https://example.com/nft/2.json");
    
    // Approve factory for NFT transfers
    await nftContract.connect(seller).setApprovalForAll(await auctionFactory.getAddress(), true);
    
    // Distribute test tokens
    await testToken.transfer(bidder1.address, ethers.parseUnits("10000", 6));
    await testToken.transfer(bidder2.address, ethers.parseUnits("10000", 6));
  });
  
  describe("Price Oracle Manager", function () {
    it("Should set and get price feeds correctly", async function () {
      const latestPrice = await priceOracle.getLatestPrice(ethers.ZeroAddress);
      expect(latestPrice).to.equal(ETH_USD_PRICE);
    });
    
    it("Should convert token amounts to USD correctly", async function () {
      const ethAmount = ethers.parseEther("1"); // 1 ETH
      const usdValue = await priceOracle.convertToUSD(ethers.ZeroAddress, ethAmount, 18);
      expect(usdValue).to.equal(ETH_USD_PRICE); // $2000
    });
    
    it("Should compare bids in different tokens", async function () {
      const ethAmount = ethers.parseEther("0.5"); // 0.5 ETH = $1000
      const tokenAmount = ethers.parseUnits("1000", 6); // 1000 TUSDC = $1000
      
      const result = await priceOracle.compareBids(
        ethers.ZeroAddress, ethAmount, 18,
        await testToken.getAddress(), tokenAmount, 6
      );
      
      expect(result).to.equal(0); // Equal values
    });
  });
  
  describe("NFT Auction Factory", function () {
    it("Should create auction successfully", async function () {
      const duration = 3600; // 1 hour
      const reservePrice = 100 * 10**8; // $100 USD
      const allowedTokens = [ethers.ZeroAddress]; // ETH only
      
      const tx = await auctionFactory.connect(seller).createAuction(
        await nftContract.getAddress(),
        1,
        duration,
        reservePrice,
        allowedTokens
      );
      
      const receipt = await tx.wait();
      const event = receipt.logs.find(log => {
        try {
          const parsed = auctionFactory.interface.parseLog(log);
          return parsed.name === "AuctionCreated";
        } catch {
          return false;
        }
      });
      
      expect(event).to.not.be.undefined;
      
      const totalAuctions = await auctionFactory.totalAuctions();
      expect(totalAuctions).to.equal(1);
    });
    
    it("Should fail to create auction without NFT ownership", async function () {
      const duration = 3600;
      const reservePrice = 100 * 10**8;
      const allowedTokens = [ethers.ZeroAddress];
      
      await expect(
        auctionFactory.connect(bidder1).createAuction(
          await nftContract.getAddress(),
          1,
          duration,
          reservePrice,
          allowedTokens
        )
      ).to.be.revertedWith("Not NFT owner");
    });
    
    it("Should get auctions by seller", async function () {
      const duration = 3600;
      const reservePrice = 100 * 10**8;
      const allowedTokens = [ethers.ZeroAddress];
      
      await auctionFactory.connect(seller).createAuction(
        await nftContract.getAddress(),
        1,
        duration,
        reservePrice,
        allowedTokens
      );
      
      const sellerAuctions = await auctionFactory.getAuctionsBySeller(seller.address);
      expect(sellerAuctions.length).to.equal(1);
    });
  });
  
  describe("NFT Auction", function () {
    let auctionAddress;
    let auction;
    
    beforeEach(async function () {
      const duration = 3600; // 1 hour
      const reservePrice = 100 * 10**8; // $100 USD
      const allowedTokens = [ethers.ZeroAddress, await testToken.getAddress()];
      
      const tx = await auctionFactory.connect(seller).createAuction(
        await nftContract.getAddress(),
        1,
        duration,
        reservePrice,
        allowedTokens
      );
      
      const receipt = await tx.wait();
      const event = receipt.logs.find(log => {
        try {
          const parsed = auctionFactory.interface.parseLog(log);
          return parsed.name === "AuctionCreated";
        } catch {
          return false;
        }
      });
      
      auctionAddress = event.args.auction;
      auction = await ethers.getContractAt("EnhancedNftAuction", auctionAddress);
      
      // Start the auction
      await auctionFactory.connect(seller).startAuction(auctionAddress);
    });
    
    it("Should start auction and transfer NFT", async function () {
      const nftOwner = await nftContract.ownerOf(1);
      expect(nftOwner).to.equal(auctionAddress);
      
      const isActive = await auction.isActive();
      expect(isActive).to.be.true;
    });
    
    it("Should accept ETH bids", async function () {
      const bidAmount = ethers.parseEther("0.1"); // 0.1 ETH = $200
      
      await auction.connect(bidder1).placeBid(
        ethers.ZeroAddress,
        bidAmount,
        { value: bidAmount }
      );
      
      const auctionInfo = await auction.getAuctionInfo();
      expect(auctionInfo.highestBidder).to.equal(bidder1.address);
    });
    
    it("Should accept ERC20 token bids", async function () {
      const bidAmount = ethers.parseUnits("200", 6); // 200 TUSDC = $200
      
      await testToken.connect(bidder1).approve(auctionAddress, bidAmount);
      await auction.connect(bidder1).placeBid(
        await testToken.getAddress(),
        bidAmount
      );
      
      const auctionInfo = await auction.getAuctionInfo();
      expect(auctionInfo.highestBidder).to.equal(bidder1.address);
    });
    
    it("Should handle bid competition correctly", async function () {
      // First bid: 0.1 ETH = $200
      const bid1Amount = ethers.parseEther("0.1");
      await auction.connect(bidder1).placeBid(
        ethers.ZeroAddress,
        bid1Amount,
        { value: bid1Amount }
      );
      
      // Second bid: 300 TUSDC = $300 (higher)
      const bid2Amount = ethers.parseUnits("300", 6);
      await testToken.connect(bidder2).approve(auctionAddress, bid2Amount);
      await auction.connect(bidder2).placeBid(
        await testToken.getAddress(),
        bid2Amount
      );
      
      const auctionInfo = await auction.getAuctionInfo();
      expect(auctionInfo.highestBidder).to.equal(bidder2.address);
    });
    
    it("Should reject bids below reserve price", async function () {
      const lowBidAmount = ethers.parseEther("0.01"); // 0.01 ETH = $20 (below $100 reserve)
      
      await expect(
        auction.connect(bidder1).placeBid(
          ethers.ZeroAddress,
          lowBidAmount,
          { value: lowBidAmount }
        )
      ).to.be.revertedWith("Bid below reserve price");
    });
    
    it("Should extend auction on late bids", async function () {
      const initialEndTime = (await auction.getAuctionInfo()).endTime;
      
      // Fast forward to near end
      await time.increaseTo(Number(initialEndTime) - 300); // 5 minutes before end
      
      const bidAmount = ethers.parseEther("0.1");
      await auction.connect(bidder1).placeBid(
        ethers.ZeroAddress,
        bidAmount,
        { value: bidAmount }
      );
      
      const newEndTime = (await auction.getAuctionInfo()).endTime;
      expect(newEndTime).to.be.gt(initialEndTime);
    });
    
    it("Should end auction and transfer NFT to winner", async function () {
      // Place winning bid
      const bidAmount = ethers.parseEther("0.1");
      await auction.connect(bidder1).placeBid(
        ethers.ZeroAddress,
        bidAmount,
        { value: bidAmount }
      );
      
      // Fast forward past auction end
      const auctionInfo = await auction.getAuctionInfo();
      await time.increaseTo(Number(auctionInfo.endTime) + 1);
      
      // End auction
      await auction.endAuction();
      
      const nftOwner = await nftContract.ownerOf(1);
      expect(nftOwner).to.equal(bidder1.address);
    });
  });
  
  describe("Enhanced ERC721", function () {
    it("Should mint NFTs with correct token URIs", async function () {
      const tokenURI = "https://example.com/nft/test.json";
      await nftContract.mint(owner.address, tokenURI);
      
      const totalMinted = await nftContract.totalMinted();
      const retrievedURI = await nftContract.tokenURI(totalMinted);
      
      expect(retrievedURI).to.equal(tokenURI);
    });
    
    it("Should batch mint multiple NFTs", async function () {
      const tokenURIs = [
        "https://example.com/nft/batch1.json",
        "https://example.com/nft/batch2.json",
        "https://example.com/nft/batch3.json"
      ];
      
      const tx = await nftContract.batchMint(owner.address, tokenURIs);
      const receipt = await tx.wait();
      
      const event = receipt.logs.find(log => {
        try {
          const parsed = nftContract.interface.parseLog(log);
          return parsed.name === "BatchMinted";
        } catch {
          return false;
        }
      });
      
      expect(event).to.not.be.undefined;
      expect(event.args.tokenIds.length).to.equal(3);
    });
    
    it("Should set and get royalty information", async function () {
      const tokenId = 1;
      const royaltyRecipient = owner.address;
      const royaltyFraction = 500; // 5%
      
      await nftContract.setTokenRoyalty(tokenId, royaltyRecipient, royaltyFraction);
      
      const salePrice = ethers.parseEther("1");
      const [recipient, royaltyAmount] = await nftContract.royaltyInfo(tokenId, salePrice);
      
      expect(recipient).to.equal(royaltyRecipient);
      expect(royaltyAmount).to.equal(salePrice * BigInt(royaltyFraction) / 10000n);
    });
  });
  
  describe("Integration Tests", function () {
    it("Should handle complete auction lifecycle", async function () {
      // Create auction
      const duration = 3600;
      const reservePrice = 100 * 10**8; // $100
      const allowedTokens = [ethers.ZeroAddress];
      
      const createTx = await auctionFactory.connect(seller).createAuction(
        await nftContract.getAddress(),
        1,
        duration,
        reservePrice,
        allowedTokens
      );
      
      const createReceipt = await createTx.wait();
      const createEvent = createReceipt.logs.find(log => {
        try {
          const parsed = auctionFactory.interface.parseLog(log);
          return parsed.name === "AuctionCreated";
        } catch {
          return false;
        }
      });
      
      const auctionAddress = createEvent.args.auction;
      const auction = await ethers.getContractAt("EnhancedNftAuction", auctionAddress);
      
      // Start auction
      await auctionFactory.connect(seller).startAuction(auctionAddress);
      
      // Place bids
      const bid1 = ethers.parseEther("0.1"); // $200
      await auction.connect(bidder1).placeBid(ethers.ZeroAddress, bid1, { value: bid1 });
      
      const bid2 = ethers.parseEther("0.15"); // $300
      await auction.connect(bidder2).placeBid(ethers.ZeroAddress, bid2, { value: bid2 });
      
      // Fast forward and end auction
      const auctionInfo = await auction.getAuctionInfo();
      await time.increaseTo(Number(auctionInfo.endTime) + 1);
      await auction.endAuction();
      
      // Verify final state
      const nftOwner = await nftContract.ownerOf(1);
      expect(nftOwner).to.equal(bidder2.address);
      
      const finalAuctionInfo = await auction.getAuctionInfo();
      expect(finalAuctionInfo.ended).to.be.true;
    });
  });
});