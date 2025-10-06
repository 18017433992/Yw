const { ethers, upgrades } = require("hardhat");
const fs = require("fs");
const path = require("path");

module.exports = async ({ getNamedAccounts, deployments }) => {
  const { save } = deployments;
  const { deployer } = await getNamedAccounts();
  
  console.log("🚀 Starting NFT Auction Marketplace deployment...");
  console.log("Deployer address:", deployer);
  
  // Deployment configuration
  const config = {
    platformFeeRate: 250, // 2.5%
    maxAuctionsPerUser: 100,
    // Mock Chainlink price feeds (replace with real addresses in production)
    ethUsdFeed: "0x694AA1769357215DE4FAC081bf1f309aDC325306", // ETH/USD on Sepolia
    // For local testing, we'll deploy mock price feeds
  };
  
  let deploymentData = {
    deployer: deployer,
    timestamp: new Date().toISOString(),
    network: hre.network.name,
    contracts: {}
  };
  
  try {
    // Step 1: Deploy Price Oracle Manager
    console.log("\n📊 Deploying Price Oracle Manager...");
    const PriceOracleManager = await ethers.getContractFactory("PriceOracleManager");
    const priceOracleProxy = await upgrades.deployProxy(
      PriceOracleManager, 
      [], 
      { initializer: "initialize" }
    );
    await priceOracleProxy.waitForDeployment();
    
    const priceOracleAddress = await priceOracleProxy.getAddress();
    const priceOracleImplAddress = await upgrades.erc1967.getImplementationAddress(priceOracleAddress);
    
    deploymentData.contracts.PriceOracleManager = {
      proxy: priceOracleAddress,
      implementation: priceOracleImplAddress
    };
    
    console.log("✅ Price Oracle Manager deployed:");
    console.log("   Proxy:", priceOracleAddress);
    console.log("   Implementation:", priceOracleImplAddress);
    
    await save("PriceOracleManager", {
      abi: PriceOracleManager.interface.format("json"),
      address: priceOracleAddress,
    });
    
    // Step 2: Deploy Enhanced NFT Auction Implementation
    console.log("\n🎯 Deploying Enhanced NFT Auction Implementation...");
    const EnhancedNftAuction = await ethers.getContractFactory("EnhancedNftAuction");
    const auctionImplementation = await EnhancedNftAuction.deploy();
    await auctionImplementation.waitForDeployment();
    
    const auctionImplAddress = await auctionImplementation.getAddress();
    deploymentData.contracts.EnhancedNftAuction = {
      implementation: auctionImplAddress
    };
    
    console.log("✅ Enhanced NFT Auction Implementation deployed at:", auctionImplAddress);
    
    // Step 3: Deploy NFT Auction Factory
    console.log("\n🏭 Deploying NFT Auction Factory...");
    const NftAuctionFactory = await ethers.getContractFactory("NftAuctionFactory");
    const factoryProxy = await upgrades.deployProxy(
      NftAuctionFactory,
      [
        auctionImplAddress,
        priceOracleAddress,
        deployer, // fee recipient
        config.platformFeeRate
      ],
      { initializer: "initialize" }
    );
    await factoryProxy.waitForDeployment();
    
    const factoryAddress = await factoryProxy.getAddress();
    const factoryImplAddress = await upgrades.erc1967.getImplementationAddress(factoryAddress);
    
    deploymentData.contracts.NftAuctionFactory = {
      proxy: factoryAddress,
      implementation: factoryImplAddress
    };
    
    console.log("✅ NFT Auction Factory deployed:");
    console.log("   Proxy:", factoryAddress);
    console.log("   Implementation:", factoryImplAddress);
    
    await save("NftAuctionFactory", {
      abi: NftAuctionFactory.interface.format("json"),
      address: factoryAddress,
    });
    
    // Step 4: Deploy Enhanced ERC721 NFT
    console.log("\n🖼️  Deploying Enhanced ERC721 NFT...");
    const EnhancedERC721 = await ethers.getContractFactory("EnhancedERC721");
    const nftContract = await EnhancedERC721.deploy(
      "Auction NFT Collection",
      "ANFT",
      deployer
    );
    await nftContract.waitForDeployment();
    
    const nftAddress = await nftContract.getAddress();
    deploymentData.contracts.EnhancedERC721 = {
      address: nftAddress
    };
    
    console.log("✅ Enhanced ERC721 NFT deployed at:", nftAddress);
    
    await save("EnhancedERC721", {
      abi: EnhancedERC721.interface.format("json"),
      address: nftAddress,
    });
    
    // Step 5: Deploy Test ERC20 Token
    console.log("\n💰 Deploying Test ERC20 Token...");
    const TestERC20 = await ethers.getContractFactory("TestERC20");
    const testToken = await TestERC20.deploy(
      "Test USD Coin",
      "TUSDC",
      6, // 6 decimals like USDC
      ethers.parseUnits("1000000", 6), // 1M tokens
      deployer
    );
    await testToken.waitForDeployment();
    
    const testTokenAddress = await testToken.getAddress();
    deploymentData.contracts.TestERC20 = {
      address: testTokenAddress
    };
    
    console.log("✅ Test ERC20 Token deployed at:", testTokenAddress);
    
    await save("TestERC20", {
      abi: TestERC20.interface.format("json"),
      address: testTokenAddress,
    });
    
    // Step 6: Setup Price Feeds (for local testing, we'll skip real Chainlink feeds)
    if (hre.network.name === "localhost" || hre.network.name === "hardhat") {
      console.log("\n⚠️  Local network detected - skipping Chainlink price feed setup");
      console.log("   In production, configure real Chainlink price feeds");
    } else {
      console.log("\n🔗 Setting up Chainlink Price Feeds...");
      // Set ETH/USD price feed
      if (config.ethUsdFeed) {
        const tx = await priceOracleProxy.setPriceFeed(
          ethers.ZeroAddress, // ETH
          config.ethUsdFeed,
          "ETH/USD"
        );
        await tx.wait();
        console.log("✅ ETH/USD price feed configured");
      }
    }
    
    // Step 7: Configure NFT Contract
    console.log("\n⚙️  Configuring NFT Contract...");
    const setFactoryTx = await nftContract.setAuctionFactory(factoryAddress);
    await setFactoryTx.wait();
    console.log("✅ Auction factory set on NFT contract");
    
    // Step 8: Mint some test NFTs
    console.log("\n🎨 Minting test NFTs...");
    const testTokenURIs = [
      "https://example.com/nft/1.json",
      "https://example.com/nft/2.json",
      "https://example.com/nft/3.json"
    ];
    
    const batchMintTx = await nftContract.batchMint(deployer, testTokenURIs);
    await batchMintTx.wait();
    console.log(`✅ Minted ${testTokenURIs.length} test NFTs`);
    
    // Step 9: Save deployment data
    const deploymentFilePath = path.resolve(__dirname, "../deployments", hre.network.name, "deployment-summary.json");
    
    // Ensure directory exists
    const deploymentDir = path.dirname(deploymentFilePath);
    if (!fs.existsSync(deploymentDir)) {
      fs.mkdirSync(deploymentDir, { recursive: true });
    }
    
    // Add contract ABIs to deployment data
    deploymentData.contracts.PriceOracleManager.abi = PriceOracleManager.interface.format("json");
    deploymentData.contracts.NftAuctionFactory.abi = NftAuctionFactory.interface.format("json");
    deploymentData.contracts.EnhancedERC721.abi = EnhancedERC721.interface.format("json");
    deploymentData.contracts.TestERC20.abi = TestERC20.interface.format("json");
    
    fs.writeFileSync(deploymentFilePath, JSON.stringify(deploymentData, null, 2));
    
    console.log("\n🎉 Deployment completed successfully!");
    console.log("📁 Deployment data saved to:", deploymentFilePath);
    
    console.log("\n📋 Contract Summary:");
    console.log("├── Price Oracle Manager:", priceOracleAddress);
    console.log("├── Auction Factory:", factoryAddress);
    console.log("├── NFT Contract:", nftAddress);
    console.log("└── Test Token:", testTokenAddress);
    
    console.log("\n🚀 Next steps:");
    console.log("1. Configure price feeds for production use");
    console.log("2. Set up cross-chain bridge if needed");
    console.log("3. Test auction creation and bidding");
    console.log("4. Configure platform fees and settings");
    
  } catch (error) {
    console.error("\n❌ Deployment failed:", error);
    throw error;
  }
};

module.exports.tags = ["AuctionMarketplace", "Full"];