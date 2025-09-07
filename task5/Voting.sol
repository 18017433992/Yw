// SPDX-License-Identifier: MIT
pragma solidity ^0.8;
//创建一个Voting 合约
contract Voting {
     //候选人列表   
     string []  public candidates =  ["Alice","Bob"];

    //添加候选人 
function setcandidates(string memory candidate) public {
        candidates.push(candidate);
}
     //mappint存储候选人票数
     mapping(string => uint256) public votecount;
     //投票vote
    function vote (string memory candidate) public {
           votecount[candidate] +=1;
    }
    //获取投票getVotes
    function getVotes(string memory name)public view returns(uint256){
            return votecount[name];
    }
    //重置所有投票结果
    function resetVotes() public {
        for (uint256 i = 0; i < candidates.length; i++) {
            votecount[candidates[i]] = 0;
        }
    }
    }


