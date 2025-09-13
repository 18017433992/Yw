// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

    //初始化代币名称、标识
    //分配初始化代币余额给合约创建者  
   //定义存储地址和余额的mapping 地址1  金额
   //定义存储授权的最大转账金额  嵌套mapping 地址1 地址2  金额
   //定义获取账户余额 BalaceOf
   //获取精度   getBigDeciaml
   //定义获取授权的转账剩余额度
    //定义转账函数   transfer

    //定义授权函数   approve
    //定义授权转账函数  transferfrom ->approve ->transfer
    //定义事件
    //触发事件

//测试场景
// 1：查询初始化合约所属账户余额
// 2、转账 
// 3、授权协议转账金额 
// 4、协议转账 
// 5.查看剩余协议转账余额 
// 6.协议转账超过协议金额

//0x18E5BC63c979B3A96a85EDA34fCAE563A2bd7455,0xB08e68a277A5e042f1EC33954cB619B6E9D99711,10000
//0x18E5BC63c979B3A96a85EDA34fCAE563A2bd7455,0x1b25ad2DD14a7D492d9edBBC579dde3b2eA90c8d,10000


contract cat { 

   mapping (address => uint256) public   _balances; //定义存储地址和余额的mapping 地址1  金额
   mapping (address => mapping (address => uint256)) public  _allowances; //定义存储授权的最大转账金额  嵌套mapping 地址1 地址2  金额
    event TransferEvent  (address indexed from, address indexed to, uint256 value);
    event ApproveEvent  (address indexed owner, address indexed spender, uint256 value);

    uint256 public  _totalSupply = 100*10**18; //代币数量
    string private _name;
    string private _symbol;
    constructor (string memory name,string memory symbol) {  //代币名称 代币标识
         _name = name;
         _symbol = symbol;
        _mint(msg.sender, _totalSupply);//将代币初始化给合约所属者
    }


//初始化代币
    function _mint(address account, uint256 value) internal   {
        if (account == address(0)) {
            revert("address can not be zero");
        }
        _balances[account] = value;
    }

  //获取账户余额
  function balanceOf(address account) external  view   returns(uint256) {
        return  _balances[account];
  }
//定义对外转账函数
function transfer(address to, uint256 value) external     returns (bool success){
     _transfer(msg.sender, to, value);
     return true ;
      
}
//定义对外授权函数
function approve(address spender, uint256  value)public    returns (bool success)  {
     _approve(msg.sender,spender, value);
    

     return  true;
}
//定义对外授权转账函数
function transferFrom(address from, address to, uint256 value)public   returns (bool success) {
           if (_allowances[from][to] < value){
                revert("_allowances must be then value");
           }
            _transfer(from, to, value);
             if(_allowances[from][to] != 0){
        _allowances[from][to] = _allowances[from][to] - value;
       }
            return  true;

}
//定义获取剩余转账余额
function allowance(address owner,address spender)public view  returns (uint256 remaining) {
     return  _allowances[owner][spender];
}

//对内实现转账函数
function _transfer(address from ,address to , uint256 value) private   {
               if(from == address(0)){
        revert("owner.address can not be zero");
      }
      if(to == address(0)){
        revert("to.address can not be zero");
      }
      if(value == 0){
        revert("value can not be zero");
      }
       _balances[from] = _balances[from] - value;
      _balances[to] = _balances[to] + value;
     
      emit TransferEvent(from,to,value);
}

//对内实现转账授权金额
function _approve(address owner, address spender, uint256 value) private      {
    require(owner != address(0), "ERC20: approve from the zero address");
    require(spender != address(0), "ERC20: approve to the zero address");
    _allowances[owner][spender] = value;
     emit   ApproveEvent(owner,spender,value);
 }



}