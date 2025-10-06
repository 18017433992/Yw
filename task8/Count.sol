// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract Count {
  event CountNumber(uint256 id);
  uint256  id;
  string public version;


  constructor(string memory _version) {
    version = _version;
  }

  function CountNmb() public  {
     id++;
     emit CountNumber(id);
  }
}