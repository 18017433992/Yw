//SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import { AggregatorV3Interface } from "@chainlink/contracts/src/v0.8/shared/interfaces/AggregatorV3Interface.sol";


    contract NftAuction is Initializable, UUPSUpgradeable {
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
            address tokenAddress; //支付代币地址 0地址表示ETH，其他地址表示ERC20



        }
receive() external payable {} // 处理纯ETH转账
fallback() external payable {} // 处理未知函数调用

    //状态变量
    // AggregatorV3Interface internal priceETHFeed;
    mapping(address => AggregatorV3Interface) public priceFeeds; //存储不同代币的美元价格，精度为1e8

    mapping(uint256 => Auction) public auctions;
    //下一个拍卖ID
    uint public nextAuctionId;
    //管理员地址
    address public admin;

    function initialize() initializer public{
        admin = msg.sender;
    }
    function setPriceETHFeed(address tokenAddress,address _priceFeed) public {   //设置兑换价格：第一个参数是代币合约地址，                                                                                //第二个参数是喂价地址（类似ETH/USD）
        // require(msg.sender == admin,"Only admin can set priceETHFeed.");
        // priceETHFeed = AggregatorV3Interface(_priceETHFeed);
        priceFeeds[tokenAddress] = AggregatorV3Interface(_priceFeed);//添加或更新喂价地址
    }

       function getChainlinkDataFeedLatestAnswer(address tokenAddress) public view returns (int) {//参数是代币合约地址
        AggregatorV3Interface priceFeed = priceFeeds[tokenAddress];
        // prettier-ignore
        (
            /* uint80 roundId */,
            int256 answer,
            /*uint256 startedAt*/,
            /*uint256 updatedAt*/,
            /*uint80 answeredInRound*/
        ) = priceFeed.latestRoundData();//获取最新兑换后的价格
        return answer;
    }
    //创建拍卖
    function createAuction(uint256 _startPrice,uint256 _duration,address _nftAddress, uint256 _tokenId) public { 
        //只有管理员才可以创建拍卖
        require(msg.sender == admin,"Only admin can create auction.");
        //检查参数
        require(_duration >10,"Duration must be greater than 10S");
        require(_startPrice >0,"Start price must be greater than 0.");
        //转移NFT到合约地址
        IERC721(_nftAddress).approve(address(this), _tokenId);
        //    IERC721(_nftAddress).transferFrom(msg.sender, address(this), _tokenId);
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
            tokenId:_tokenId,
            tokenAddress:address(0) //默认是ETH
});
        nextAuctionId++;

    } 
    //参与拍卖
    function bid(uint256 _auctionId,uint256 amount,address _tokenAddress) public payable{ //参数1：拍卖ID，参数2：出价金额，参数3：代币地址
        //统一的价值尺度，拍卖默认是EHT，则需要把其它代币转化为ETH价值
       //如果是ETH，则直接传递ETH,msg.value是拍卖的ETH数量
       //如果是ERC20，则需要传递代币数量和代币地址
         Auction storage auction = auctions[_auctionId];
        uint payValue;//换算后最终支付的金额
        if(_tokenAddress != address(0)){
             //如果是代币ERC20，则需要转换为想要的价值，如美元
            payValue = amount *uint(getChainlinkDataFeedLatestAnswer(_tokenAddress));//获取ERC20的美元价值      
        }else{
           //如果是ETH，则直接使用msg.value
           amount = msg.value;
           payValue = amount * uint(getChainlinkDataFeedLatestAnswer(address(0)));//获取ETH的美元价值
      } 
      //换算出开始价格和最高价的美元价值
      uint startPrice = auction.startPrice * uint(getChainlinkDataFeedLatestAnswer(auction.tokenAddress));
      uint highestBid = auction.highestBid * uint(getChainlinkDataFeedLatestAnswer(auction.tokenAddress));
        //检查出价是否有效
        require(payValue >= startPrice && payValue > highestBid,"Bid must be higher than the current highest bid.");
        require(!auction.ended,"Auction has ended.");
        require(block.timestamp < auction.startTime + auction.duration,"Auction has ended.");
        //转移ERC20代币到合约
        IERC20(_tokenAddress).transferFrom(msg.sender,address(this),amount);//不需要换算，直接转移代币数量给合约，种类是_tokenAddress的代币
        //如果之前最高价是ETH，则退还ETH
        if(auction.tokenAddress == address(0)){
            payable(auction.highestBidder).transfer(auction.highestBid);
        }else{//如果之前最高价是ERC20代币，则退还ERC20代币
           IERC20(auction.tokenAddress).transfer(auction.highestBidder,auction.highestBid);
        }
        auction.tokenAddress = _tokenAddress;
        auction.highestBid = amount;
        auction.highestBidder = msg.sender;
    }


    //结束拍卖
    function endAuction(uint256 _auctionId) external { 
        Auction storage auction = auctions[_auctionId];
        //获取当前拍卖是否结束
        require(!auction.ended && auction.startTime + auction.duration >= block.timestamp,"Auction has ended.");
        //将NFT转移给最高出价者
        IERC721(auction.nftAddress).safeTransferFrom(admin, auction.highestBidder, auction.tokenId);
        //转移剩余资金转移给卖家
        payable(address(this)).transfer(address(this).balance);
        auction.ended = true;
    }
    function _authorizeUpgrade(address newImplementation) internal override view {
        //只有管理员才可以升级
        require(msg.sender == admin, "Only admin can upgrade.");

    }

    }