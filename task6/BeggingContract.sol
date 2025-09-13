// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

// 合约应包含以下功能：
// 一个 mapping 来记录每个捐赠者的捐赠金额。地址、金额
// 一个 donate 函数，允许用户向合约发送以太币，并记录捐赠信息。
// 一个 withdraw 函数，允许合约所有者提取所有资金。 校验是否合约所有者
// 一个 getDonation 函数，允许查询某个地址的捐赠金额。
// 使用 payable 修饰符和 address.transfer 实现支付和提款。
// 捐赠事件：添加 Donation 事件，记录每次捐赠的地址和金额。
// 捐赠排行榜：实现一个功能，显示捐赠金额最多的前 3 个地址。
// 时间限制：添加一个时间限制，只有在特定时间段内才能捐赠。

contract BeggingContract {
  address public owner;  //定义合约所有者
   uint256 public startTime;
    uint256 public endTime;
    
  event Donation(address indexed donator, uint256 amount);

  
    constructor() {
        owner =  msg.sender ;//调用者为合约所有者
        startTime = block.timestamp;
        endTime = startTime + 1 days;
    }
   receive() external payable {   //收款逻辑 
   } 
  fallback() external payable {     // fallback 收款逻辑 

   }

    mapping (address=>uint256) public  _donates;  //定义存储捐赠者地址和捐赠金额
       address[] public _arrdonates;
           modifier onlyActivePeriod() {
        require(block.timestamp >= startTime && block.timestamp <= endTime, "Donation period has ended");
        _;
    }
    modifier onlyOwner() {  //判断调用者是否为合约所有者
        require(msg.sender == owner,"Only owner can call this function");
        _;
    }
    function donate() public  onlyActivePeriod payable  returns (bool success){  //捐赠函数
            require(msg.sender.balance >= 0,"Donation amount must be greater than 0");
            _donates[msg.sender] +=  msg.value; //记录捐赠者地址和捐赠金额
            _arrdonates.push(msg.sender);
            emit   Donation(msg.sender,msg.value);//触发监听事件

            return true;
    }
  function getDonation(address account) public  view returns (uint256 amount){
        return  _donates[account] ;//查询捐赠者捐赠的金额
  }

  function withdraw(uint amount) public onlyOwner returns (bool success) {  //合约所有者提取金额
             require(address(this).balance >= amount,"balance not enough");
             payable(owner).transfer(amount); //合约所有者提取金额,不对捐赠者的金额做扣减
             return true;
  }
//获取捐赠最大的三个地址(冒泡排序)
 function getOneSecondDonate()public returns(address add_donate1,address add_donate2,address add_donate3) {
             
             for(uint256 i =0;i<_arrdonates.length-1;i++){
               for(uint256 j = 0;j<_arrdonates.length-1-i;j++){
                 if(_donates[_arrdonates[j]]<_donates[_arrdonates[j+1]]){
                    address temp =  _arrdonates[j];
                      _arrdonates[j]=   _arrdonates[j+1]; 
                      _arrdonates[j+1] = temp ;
                }
               }
               
         
               } add_donate1 =  _arrdonates[0];
                 add_donate2 =  _arrdonates[1];
                 add_donate3 =  _arrdonates[2];
               return (add_donate1,add_donate2,add_donate3);

 }
}

// 0x5B38Da6a701c568545dCfcB03FcB875f56beddC4 1
// 0xAb8483F64d9C6d1EcF9b849Ae677dD3315835cb2  2
// 0x4B20993Bc481177ec7E8f571ceCaE8A9e22C02db   3
// 0x78731D3Ca6b7E34aC0F824c42a7cC18A495cabaB  4