// SPDX-License-Identifier: MIT
pragma solidity ^0.8;
contract Revert_init {
 // 使用常量数组存储罗马数字符号和对应的值    

        // 将整数转换为罗马数字
    function intToRoman(uint256 num) public pure returns (string memory) {

 string[13] memory ROMAN_SYMBOLS = ["M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"];
 uint256[13] memory ROMAN_VALUES = [uint256(1000), 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1];  


    require(num > 0 && num < 4000, "Number must be between 1 and 3999");
        
        bytes memory result = new bytes(0);
        
        // 使用贪心算法构建罗马数字
        for (uint256 i = 0; i < ROMAN_VALUES.length; i++) {
            while (num >= ROMAN_VALUES[i]) {
                num -= ROMAN_VALUES[i];
                result = abi.encodePacked(result, ROMAN_SYMBOLS[i]);
            }
        }
        
        return string(result);
    }
}


