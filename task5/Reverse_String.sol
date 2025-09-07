// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract reverse {
    //反转字符串
function reverString(string memory str)public pure   returns(string memory)  {
        uint length = bytes(str).length;
        bytes memory strBytes = new bytes(length);
        for (uint i = 0; i < length; i++) {
            strBytes[i] = bytes(str)[length - i - 1];
        }
        return string(strBytes);
}
//将整数转化为罗马数字
function intToRoman(uint256 num) public pure returns (string memory) {
        require(num > 0 && num < 4000, "Number must be between 1 and 3999");
        
        // 定义罗马数字符号和对应的值
        string[13] memory romanSymbols = ["M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"];
        uint256[13] memory romanValues = [uint256(1000), 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1];  
        bytes memory result = new bytes(0);
        
        // 逐步构建罗马数字
        for (uint256 i = 0; i < romanValues.length; i++) {
            while (num >= romanValues[i]) {
                num -= romanValues[i];
                result = abi.encodePacked(result, romanSymbols[i]);
            }
        }
        
        return string(result);
    }





}

