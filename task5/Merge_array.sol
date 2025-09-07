// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract mergearray {
//将两个数组有序合并
     function mergerlist(uint[] memory arr1,uint[] memory arr2) public pure  returns(uint[] memory) {  
           uint[] memory result = new uint[](arr1.length+arr2.length);  
           uint i=0 ;
           uint j=0 ;
           uint k=0 ;  
           while(i<arr1.length && j<arr2.length) {  
               if(arr1[i]<=arr2[j]) {  
                   result[k]=arr1[i];  
                   i++;  
               }else{  
                   result[k]=arr2[j];  
                   j++;  
}
               k++;  
           }  
           while(i<arr1.length) {
               result[k]=arr1[i];
               i++;
               k++;
           }
           while(j<arr2.length) {
               result[k]=arr2[j];
               j++;
               k++;
           }
         return  result;
}
//折半查找
function halfIndex(uint256 [] memory arr,uint256  target) public pure returns(int256) {
        uint left = 0;
        uint right = arr.length - 1;

        // 进行二分查找
        while (left <= right) {
            uint mid = left + (right - left) / 2;

            // 检查中间元素
            if (arr[mid] == target) {
                // 返回目标值的索引
                return int(mid);
            }

            // 如果目标值小于中间元素，缩小查找范围到左半部分
            if (arr[mid] > target) {
                right = mid - 1;
            }
                // 如果目标值大于中间元素，缩小查找范围到右半部分
            else {
                left = mid + 1;
            }
        }

        // 如果没有找到目标值，返回 -1
        return -1;
    }


    
}






