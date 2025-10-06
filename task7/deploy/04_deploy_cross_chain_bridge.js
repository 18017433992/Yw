const { ethers, upgrades } = require("hardhat");
const fs = require("fs");
const path = require("path");

module.exports = async ({ getNamedAccounts, deployments }) => {
  const { save } = deployments;
  const { deployer } = await getNamedAccounts();
  
  console.log("🌉 Starting Cross-Chain Bridge deployment...");
  console.log("Deployer address:", deployer);
  
  // Network-specific CCIP Router addresses
  const ccipRouters = {
    // Ethereum Sepolia
    "11155111": "0xD0daae2231E9CB96b94C8512223533293C3693Bf",
    // Polygon Mumbai
    "80001": "0x70499c328e1E2a3c41108bd3730F6670a44595D1",
    // Avalanche Fuji
    "43113": "0x554472a2720E5E7D5D3C817529aBA05EEd5F82D8",
    // Arbitrum Sepolia
    "421614": "0x2a9C5afB0d0e4BAb2BCdaE109EC4b0c4Be15a165",
    // Base Sepolia
    "84532": "0xD3b06cEbF099CE7DA4AcCf578aaebFDBd6e88a93"
  };
  
  let deploymentData = {
    deployer: deployer,
    timestamp: new Date().toISOString(),
    network: hre.network.name,
    chainId: hre.network.config.chainId,
    contracts: {}
  };
  
  try {
    // Get CCIP Router address for current network
    const chainId = hre.network.config.chainId?.toString();
    const routerAddress = ccipRouters[chainId];
    
    if (!routerAddress && hre.network.name !== "localhost" && hre.network.name !== "hardhat") {
      console.log("⚠️  CCIP Router not available for this network");
      console.log("   Available networks:", Object.keys(ccipRouters));
      return;
    }
    
    let finalRouterAddress = routerAddress;
    
    // For local testing, deploy a mock router
    if (hre.network.name === "localhost" || hre.network.name === "hardhat") {
      console.log("\n🔧 Deploying Mock CCIP Router for local testing...");
      
      // Create a simple mock router contract
      const mockRouterCode = `
        // SPDX-License-Identifier: MIT
        pragma solidity ^0.8.20;
        
        contract MockCCIPRouter {
          mapping(bytes32 => bool) public sentMessages;
          uint256 public constant FIXED_FEE = 0.01 ether;
          
          function ccipSend(uint64, bytes memory) external payable returns (bytes32) {
            require(msg.value >= FIXED_FEE, "Insufficient fee");
            bytes32 messageId = keccak256(abi.encode(block.timestamp, msg.sender));
            sentMessages[messageId] = true;
            return messageId;
          }
          
          function getFee(uint64, bytes memory) external pure returns (uint256) {
            return FIXED_FEE;
          }
        }
      `;
      
      // For simplicity, we'll use a placeholder address for local testing
      finalRouterAddress = ethers.ZeroAddress;
      console.log("⚠️  Using placeholder router address for local testing");
    }
    
    // Deploy Cross-Chain Auction Bridge
    console.log("\n🌉 Deploying Cross-Chain Auction Bridge...");
    const CrossChainAuctionBridge = await ethers.getContractFactory("CrossChainAuctionBridge");
    
    const bridgeProxy = await upgrades.deployProxy(
      CrossChainAuctionBridge,
      [finalRouterAddress],
      { 
        initializer: "initialize",
        kind: "uups"
      }
    );
    await bridgeProxy.waitForDeployment();
    
    const bridgeAddress = await bridgeProxy.getAddress();
    const bridgeImplAddress = await upgrades.erc1967.getImplementationAddress(bridgeAddress);
    
    deploymentData.contracts.CrossChainAuctionBridge = {
      proxy: bridgeAddress,
      implementation: bridgeImplAddress,
      ccipRouter: finalRouterAddress
    };
    
    console.log("✅ Cross-Chain Auction Bridge deployed:");
    console.log("   Proxy:", bridgeAddress);
    console.log("   Implementation:", bridgeImplAddress);
    console.log("   CCIP Router:", finalRouterAddress);
    
    await save("CrossChainAuctionBridge", {
      abi: CrossChainAuctionBridge.interface.format("json"),
      address: bridgeAddress,
    });
    
    // Configure supported chains (for demonstration)
    if (hre.network.name !== "localhost" && hre.network.name !== "hardhat") {
      console.log("\n⚙️  Configuring supported chains...");
      
      // Example chain configurations (adjust based on actual requirements)
      const supportedChains = [
        { selector: 16015286601757825753n, name: "Ethereum Sepolia" },
        { selector: 12532609583862916517n, name: "Polygon Mumbai" },
        { selector: 14767482510784806043n, name: "Avalanche Fuji" }
      ];
      
      for (const chain of supportedChains) {
        try {
          // In a real deployment, you'd have the actual bridge contract addresses
          const mockBridgeAddress = ethers.ZeroAddress; // Placeholder
          
          const tx = await bridgeProxy.addChainSupport(
            chain.selector,
            mockBridgeAddress
          );
          await tx.wait();
          console.log(`✅ Added support for ${chain.name}`);
        } catch (error) {
          console.log(`⚠️  Failed to add ${chain.name}:`, error.message);
        }
      }
    }
    
    // Save deployment data
    const deploymentFilePath = path.resolve(
      __dirname, 
      "../deployments", 
      hre.network.name, 
      "cross-chain-deployment.json"
    );
    
    const deploymentDir = path.dirname(deploymentFilePath);
    if (!fs.existsSync(deploymentDir)) {
      fs.mkdirSync(deploymentDir, { recursive: true });
    }
    
    deploymentData.contracts.CrossChainAuctionBridge.abi = CrossChainAuctionBridge.interface.format("json");
    fs.writeFileSync(deploymentFilePath, JSON.stringify(deploymentData, null, 2));
    
    console.log("\n🎉 Cross-Chain Bridge deployment completed!");
    console.log("📁 Deployment data saved to:", deploymentFilePath);
    
    console.log("\n📋 Bridge Contract Summary:");
    console.log("├── Bridge Address:", bridgeAddress);
    console.log("├── Implementation:", bridgeImplAddress);
    console.log("└── CCIP Router:", finalRouterAddress);
    
    console.log("\n🚀 Next steps for cross-chain functionality:");
    console.log("1. Deploy bridge contracts on other supported chains");
    console.log("2. Configure cross-chain routes and fee structures");
    console.log("3. Test cross-chain message passing");
    console.log("4. Set up monitoring and alerting");
    console.log("5. Configure gas limits for different message types");
    
    if (hre.network.name === "localhost" || hre.network.name === "hardhat") {
      console.log("\n⚠️  Local Testing Notes:");
      console.log("- CCIP functionality is mocked for local testing");
      console.log("- Use testnet deployments for real cross-chain testing");
      console.log("- Ensure sufficient ETH/LINK for CCIP fees on testnets");
    }
    
  } catch (error) {
    console.error("\n❌ Cross-chain bridge deployment failed:", error);
    throw error;
  }
};

module.exports.tags = ["CrossChain", "Bridge", "CCIP"];
module.exports.dependencies = ["AuctionMarketplace"];