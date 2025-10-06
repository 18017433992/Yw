    const{ethers,deployments} = require("hardhat");
    const { expect } = require("chai");
    const { parseEther } = require("ethers");
    describe("Test NftAuction", async function () { 
        it("Should be ok", async function () { 
            await main();
        })
    });

    async function main() { 
    await  deployments.fixture(['deployNftAuction']);//部署业务合约
    const nftAuctionProxy = await deployments.get("NftAuctionProxy");//获取代理合约
    const[signer,buyer] = await ethers.getSigners();//获取主钱包地址
    
    //1.部署ERC721合约
    const TestERC721 = await ethers.getContractFactory("TestERC721");
    const testERC721 = await TestERC721.deploy();
    await testERC721.waitForDeployment();
    const testERC721Address = await testERC721.getAddress();
    console.log("TestERC721:", testERC721Address);
    //1.2.铸造10个NFT
    for(let i=0;i<10;i++){
        await testERC721.mint(signer.address,i+1);
    };
    console.log("NFT已铸造完毕");
 
    //2.调用createAuction方法创建拍卖
    const nftAuction = await ethers.getContractAt("NftAuction",nftAuctionProxy.address);//拿到合约实例
       //给代理合约授权
    await testERC721.connect(signer).setApprovalForAll(nftAuctionProxy.address,true);
    const tokenId = 1;
    await nftAuction.createAuction(
        ethers.parseEther("0.01"),
        15,
        testERC721Address, //NFT合约地址
        tokenId
    );
    console.log("拍卖已创建");
    console.log("拍卖前的归属：",signer.address)

    //3.出价
    //await testERC721.connect(buyer).approve(nftAuctionProxy.address,tokenId);//批准合约操作NFT
    await nftAuction.connect(buyer).bid(0,{value:parseEther("0.02")});
        console.log("正在出价...");

    //4.等待10秒，结束拍卖
        //await new Promise(resolve => setTimeout(resolve, 14*1000));
        
        await nftAuction.connect(signer).endAuction(0);

        console.log("正在等待拍卖结束...");
    //5.验证结果
    const auctionsResult = await nftAuction.auctions(0);   
    console.log("拍卖结果:",auctionsResult);
    expect(auctionsResult.highestBid).to.equal(parseEther("0.02"));
    expect(auctionsResult.highestBidder).to.equal(buyer.address);
    //6.查看NFT归属
    const owner = await testERC721.ownerOf(tokenId);
    console.log("NFT归属:",owner);
    expect(owner).to.equal(buyer.address);

    console.log("singer:",signer.address);
    console.log("buyer:",buyer.address);
    
    }