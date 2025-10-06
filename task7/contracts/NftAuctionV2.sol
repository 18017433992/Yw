//SPDX-License-Identifier: MIT
pragma solidity ^0.8;
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

contract NftAuctionV2 is Initializable {
    //结构体
    struct Auction{
        address seller; //拍卖者
        uint256 startPrice;//开始价格
        uint256 duration;//持续时间
        bool ended; //是否结束
        address highestBidder;//最高出价者
        uint256 highestBid;//最高出价
        uint256 startTime;//开始时间
        address nftAddress;//拍卖NFT地址
        uint256 tokenId;//NFT ID


    }
   //状态变量
   mapping(uint256 => Auction) public auctions;
   //下一个拍卖ID
   uint public nextAuctionId;
   //管理员地址
   address public admin;

   function initialize() initializer public{
       admin = msg.sender;
   }
   //创建拍卖
   function createAuction(uint256 _startPrice,uint256 _duration,address _nftAddress, uint256 _tokenId) public { 
     //只有管理员才可以创建拍卖
      require(msg.sender == admin,"Only admin can create auction.");
      //检查参数
       require(_duration >1000*60,"Duration must be greater than 0.");
       require(_startPrice >0,"Start price must be greater than 0.");
       //创建拍卖
        auctions[nextAuctionId] = Auction({
           seller:msg.sender,
           startPrice:_startPrice,
           duration:_duration,
           ended:false,
           highestBidder:address(0),
           highestBid:0,
           startTime:block.timestamp,
           nftAddress:_nftAddress,
           tokenId:_tokenId
       });
       nextAuctionId++;

   } 
//参与拍卖
function bid(uint256 _auctionId) public payable{ 
     Auction storage auction = auctions[_auctionId];
     require(!auction.ended,"Auction has ended.");
     require(block.timestamp < auction.startTime + auction.duration,"Auction has ended.");
     require(msg.value > auction.highestBid && msg.value >auction.startPrice,"Bid must be higher than the current highest bid.");
       if(auction.highestBidder != address(0)){
           payable(auction.highestBidder).transfer(auction.highestBid);
       }
       auction.highestBidder = msg.sender;
       auction.highestBid = msg.value;


}


 function testHello() public pure returns(string memory){ 
     return "hello";
 }

}