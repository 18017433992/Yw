const { ethers } = require("hardhat");
const fs = require("fs");
const path = require("path");

/**
 * Demo script for NFT Auction Marketplace
 * This script demonstrates the complete workflow of creating and participating in NFT auctions
 */
async function main() {
  console.log("🎭 NFT Auction Marketplace Demo");
  console.log("================================\n");
  
  // Get signers
  const [deployer, seller, bidder1, bidder2] = await ethers.getSigners();
  
  console.log("👥 Participants:");
  console.log(`├── Deployer: ${deployer.address}`);
  console.log(`├── Seller: ${seller.address}`);
  console.log(`├── Bidder 1: ${bidder1.address}`);
  console.log(`└── Bidder 2: ${bidder2.address}\n`);
  
  try {
    // Load deployed contracts (from deployment summary)
    const deploymentPath = path.resolve(__dirname, "../deployments", hre.network.name, "deployment-summary.json");
    
    if (!fs.existsSync(deploymentPath)) {
      console.log("❌ Deployment summary not found. Please deploy contracts first.");
      console.log("   Run: npx hardhat deploy --tags AuctionMarketplace");
      return;
    }
    
    const deploymentData = JSON.parse(fs.readFileSync(deploymentPath, 'utf8'));
    const contracts = deploymentData.contracts;
    
    // Get contract instances
    const nftContract = await ethers.getContractAt(
      "EnhancedERC721", 
      contracts.EnhancedERC721.address
    );
    
    const auctionFactory = await ethers.getContractAt(
      "NftAuctionFactory", 
      contracts.NftAuctionFactory.proxy
    );
    
    const priceOracle = await ethers.getContractAt(
      "PriceOracleManager", 
      contracts.PriceOracleManager.proxy
    );
    
    const testToken = await ethers.getContractAt(
      "TestERC20", 
      contracts.TestERC20.address
    );
    
    console.log("📋 Loaded Contracts:");
    console.log(`├── NFT Contract: ${await nftContract.getAddress()}`);
    console.log(`├── Auction Factory: ${await auctionFactory.getAddress()}`);
    console.log(`├── Price Oracle: ${await priceOracle.getAddress()}`);
    console.log(`└── Test Token: ${await testToken.getAddress()}\n`);
    
    // Demo Step 1: Check current NFT ownership
    console.log("🎨 Step 1: Check NFT Ownership");
    console.log("─────────────────────────────");
    
    try {
      const totalSupply = await nftContract.totalSupply();
      console.log(`Total NFTs minted: ${totalSupply}`);
      
      for (let i = 1; i <= Number(totalSupply); i++) {
        const owner = await nftContract.ownerOf(i);
        const tokenURI = await nftContract.tokenURI(i);
        console.log(`├── NFT #${i}: Owner ${owner} (${tokenURI})`);
      }
    } catch (error) {
      console.log("⚠️  No NFTs found or error reading NFTs");
    }
    
    // Demo Step 2: Transfer NFT to seller if needed
    console.log("\n📤 Step 2: Prepare NFT for Auction");
    console.log("─────────────────────────────────");
    
    const nftId = 1;
    const currentOwner = await nftContract.ownerOf(nftId);
    
    if (currentOwner !== seller.address) {
      console.log(`Transferring NFT #${nftId} from ${currentOwner} to seller...`);
      await nftContract.connect(deployer).transferFrom(currentOwner, seller.address, nftId);
      console.log("✅ NFT transferred to seller");
    } else {
      console.log("✅ NFT already owned by seller");
    }
    
    // Approve factory to manage seller's NFTs
    const isApprovedForAll = await nftContract.isApprovedForAll(seller.address, await auctionFactory.getAddress());
    if (!isApprovedForAll) {
      console.log("Approving factory for NFT management...");
      await nftContract.connect(seller).setApprovalForAll(await auctionFactory.getAddress(), true);
      console.log("✅ Factory approved for NFT management");
    }
    
    // Demo Step 3: Distribute test tokens to bidders
    console.log("\n💰 Step 3: Distribute Test Tokens");
    console.log("─────────────────────────────────");
    
    const testTokenAmount = ethers.parseUnits("5000", 6); // 5000 TUSDC
    
    await testToken.connect(deployer).transfer(bidder1.address, testTokenAmount);
    await testToken.connect(deployer).transfer(bidder2.address, testTokenAmount);
    
    const bidder1Balance = await testToken.balanceOf(bidder1.address);
    const bidder2Balance = await testToken.balanceOf(bidder2.address);
    
    console.log(`├── Bidder 1 TUSDC balance: ${ethers.formatUnits(bidder1Balance, 6)}`);
    console.log(`└── Bidder 2 TUSDC balance: ${ethers.formatUnits(bidder2Balance, 6)}`);
    
    // Demo Step 4: Check price feeds
    console.log("\n📊 Step 4: Check Price Feeds");
    console.log("───────────────────────────");
    
    try {
      const ethPrice = await priceOracle.getLatestPrice(ethers.ZeroAddress);
      console.log(`├── ETH/USD Price: $${ethers.formatUnits(ethPrice, 8)}`);
      
      const tokenPrice = await priceOracle.getLatestPrice(await testToken.getAddress());
      console.log(`└── TUSDC/USD Price: $${ethers.formatUnits(tokenPrice, 8)}`);
    } catch (error) {
      console.log("⚠️  Price feeds not configured for this network");
    }
    
    // Demo Step 5: Create an auction
    console.log("\n🏛️  Step 5: Create Auction");
    console.log("─────────────────────────");
    
    const auctionDuration = 3600; // 1 hour
    const reservePrice = ethers.parseUnits("100", 8); // $100 USD
    const allowedTokens = [ethers.ZeroAddress, await testToken.getAddress()]; // ETH and TUSDC
    
    console.log("Creating auction with parameters:");
    console.log(`├── NFT: #${nftId}`);
    console.log(`├── Duration: ${auctionDuration} seconds (1 hour)`);
    console.log(`├── Reserve Price: $${ethers.formatUnits(reservePrice, 8)}`);
    console.log(`└── Allowed Tokens: ETH, TUSDC`);
    
    const createTx = await auctionFactory.connect(seller).createAuction(
      await nftContract.getAddress(),
      nftId,
      auctionDuration,
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
    console.log(`✅ Auction created at: ${auctionAddress}`);
    
    // Get auction contract instance
    const auction = await ethers.getContractAt("EnhancedNftAuction", auctionAddress);
    
    // Demo Step 6: Start the auction
    console.log("\n🚀 Step 6: Start Auction");
    console.log("───────────────────────");
    
    await auctionFactory.connect(seller).startAuction(auctionAddress);
    console.log("✅ Auction started - NFT transferred to auction contract");
    
    // Verify NFT transfer
    const newNftOwner = await nftContract.ownerOf(nftId);
    console.log(`├── NFT #${nftId} now owned by: ${newNftOwner}`);
    console.log(`└── Auction address: ${auctionAddress}`);
    
    // Demo Step 7: Place bids
    console.log("\n💸 Step 7: Place Bids");
    console.log("────────────────────");
    
    // Bidder 1 bids with ETH
    const ethBidAmount = ethers.parseEther("0.06"); // Should be ~$120 at $2000/ETH
    console.log(`Bidder 1 placing ETH bid: ${ethers.formatEther(ethBidAmount)} ETH`);
    
    await auction.connect(bidder1).placeBid(
      ethers.ZeroAddress, 
      ethBidAmount, 
      { value: ethBidAmount }
    );
    console.log("✅ ETH bid placed by Bidder 1");
    
    // Bidder 2 bids with TUSDC
    const tokenBidAmount = ethers.parseUnits("150", 6); // $150 TUSDC
    console.log(`Bidder 2 placing TUSDC bid: ${ethers.formatUnits(tokenBidAmount, 6)} TUSDC`);
    
    await testToken.connect(bidder2).approve(auctionAddress, tokenBidAmount);
    await auction.connect(bidder2).placeBid(
      await testToken.getAddress(),
      tokenBidAmount
    );
    console.log("✅ TUSDC bid placed by Bidder 2");
    
    // Demo Step 8: Check auction status
    console.log("\n📈 Step 8: Auction Status");
    console.log("────────────────────────");
    
    const auctionInfo = await auction.getAuctionInfo();
    const bidHistory = await auction.getBidHistory();
    
    console.log("Current auction state:");
    console.log(`├── Highest Bidder: ${auctionInfo.highestBidder}`);
    console.log(`├── Highest Bid (USD): $${ethers.formatUnits(auctionInfo.highestBidUSD, 8)}`);
    console.log(`├── Total Bids: ${auctionInfo.bidCount}`);
    console.log(`├── Started: ${auctionInfo.started}`);
    console.log(`├── Ended: ${auctionInfo.ended}`);
    console.log(`└── Active: ${await auction.isActive()}`);
    
    console.log("\nBid History:");
    for (let i = 0; i < bidHistory.length; i++) {
      const bid = bidHistory[i];
      const tokenSymbol = bid.paymentToken === ethers.ZeroAddress ? "ETH" : "TUSDC";
      const decimals = bid.paymentToken === ethers.ZeroAddress ? 18 : 6;
      console.log(`├── Bid ${i + 1}: ${ethers.formatUnits(bid.amount, decimals)} ${tokenSymbol} (~$${ethers.formatUnits(bid.usdValue, 8)}) by ${bid.bidder}`);
    }
    
    // Demo Step 9: Show factory statistics
    console.log("\n📊 Step 9: Factory Statistics");
    console.log("────────────────────────────");
    
    const totalAuctions = await auctionFactory.totalAuctions();
    const activeAuctions = await auctionFactory.getActiveAuctions();
    const sellerAuctions = await auctionFactory.getAuctionsBySeller(seller.address);
    
    console.log(`├── Total Auctions Created: ${totalAuctions}`);
    console.log(`├── Active Auctions: ${activeAuctions.length}`);
    console.log(`└── Seller's Auctions: ${sellerAuctions.length}`);
    
    // Demo Step 10: Instructions for ending auction
    console.log("\n⏰ Step 10: Next Steps");
    console.log("────────────────────");
    console.log("To complete the auction demo:");
    console.log("1. Wait for auction duration to expire (1 hour)");
    console.log("2. Call auction.endAuction() to finalize");
    console.log("3. The highest bidder will receive the NFT");
    console.log("4. The seller will receive the winning bid amount");
    
    console.log("\n🎉 Demo completed successfully!");
    console.log("The auction is now live and accepting bids.");
    
    // Save demo results
    const demoResults = {
      timestamp: new Date().toISOString(),
      network: hre.network.name,
      participants: {
        deployer: deployer.address,
        seller: seller.address,
        bidder1: bidder1.address,
        bidder2: bidder2.address
      },
      auction: {
        address: auctionAddress,
        nftId: nftId,
        reservePrice: ethers.formatUnits(reservePrice, 8),
        duration: auctionDuration,
        allowedTokens: allowedTokens
      },
      bids: bidHistory.map(bid => ({
        bidder: bid.bidder,
        paymentToken: bid.paymentToken,
        amount: bid.amount.toString(),
        usdValue: ethers.formatUnits(bid.usdValue, 8),
        timestamp: bid.timestamp.toString()
      }))
    };
    
    const demoResultsPath = path.resolve(__dirname, "../demo-results.json");
    fs.writeFileSync(demoResultsPath, JSON.stringify(demoResults, null, 2));
    console.log(`\n📁 Demo results saved to: ${demoResultsPath}`);
    
  } catch (error) {
    console.error("\n❌ Demo failed:", error);
    console.error("\nTroubleshooting:");
    console.error("1. Ensure contracts are deployed: npx hardhat deploy");
    console.error("2. Check network configuration");
    console.error("3. Verify account balances");
    console.error("4. Check price feed configuration");
  }
}

// Handle script execution
if (require.main === module) {
  main()
    .then(() => process.exit(0))
    .catch((error) => {
      console.error(error);
      process.exit(1);
    });
}

module.exports = main;