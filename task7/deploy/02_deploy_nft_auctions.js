const {ethers,upgrades} = require("hardhat");
const fs = require("fs");
const path = require("path");
module.exports = async ({getNamedAccounts, deployments}) =>
{ 
    const {save} = deployments;
    const {deployer} = await getNamedAccounts();
    console.log("部署用户地址:", deployer);


    //读取proxyNftAuction.json
    const storePath = path.resolve(__dirname,"./proxyNftauction.json");
    const storeData = fs.readFileSync(storePath, "utf-8");
    const{proxyAdress,implAdress,abi} = JSON.parse(storeData);
    //升级版业务合约
    const NftAuctionV2 = await ethers.getContractFactory("NftAuctionV2");
    //升级代理合约
    const NftAuctionProxyV2 = await upgrades.upgradeProxy(proxyAdress, NftAuctionV2);
    await NftAuctionProxyV2.waitForDeployment();
    const proxyAdressV2 = await NftAuctionProxyV2.getAddress();


    await save("NftAuctionProxyV2",{
        address:proxyAdressV2,
        abi:abi
    })
};
module.exports.tags = ["upgradeNftAuction"];