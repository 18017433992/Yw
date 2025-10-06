const { ethers,deployments} = require("hardhat")
const { expect } = require("chai");
describe("Test upgrade", async function () { 
    it("Should deploy the contract", async function () { 
        //1.部署业务合约
        await deployments.fixture(["deployNftAuction"]);
        const nftauctionProxy = await deployments.get("NftAuctionProxy");

        //2.调用createAuction创建拍卖
        //int256 _startPrice,uint256 _duration,address _nftAddress, uint256 _tokenId
         const nftAuction = await ethers.getContractAt("NftAuction", nftauctionProxy.address);
         await nftAuction.createAuction(
            ethers.parseEther("0.01"),
             100*1000,
             ethers.ZeroAddress,
             1
            );
        const  auction1 =   await nftAuction.auctions(0)
    
        //console.log("读取合约的auction1的startTime:",await auction1.startTime);

         console.log("读取合约的auction1:",auction1);
        //3.升级合约 
        await deployments.fixture(["upgradeNftAuction"]);
        //4.读取合约的auction[0]
        const auction2 =await nftAuction.auctions(0)  

        console.log("读取合约的auction2:",auction2);
        //console.log("读取合约的auction1的startTime:",await auction1.startTime);
        //console.log("读取合约的auction2的startTime:",await auction2.startTime);
      //检查开始时间是否一致
        const nftAuctionV2 = await ethers.getContractAt("NftAuctionV2", nftauctionProxy.address);//升级后合约地址不变
        expect( auction2.startTime).to.equal(auction1.startTime);
         const hello = await nftAuctionV2.testHello();
        console.log("hello:",hello);
        
})
})
