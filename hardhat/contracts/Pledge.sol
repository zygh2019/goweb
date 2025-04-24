// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.10;

import "./Ipledge.sol";

contract Pledge is Ipledge {
    //合约部署者
    address private owner;
    //提供者
    address private finalProposer;
    //委员会
    address[] private finalcommitees;
    //质押池
    mapping(address => uint) private stakingPool;

    address[] private stakingPoolAddresss;

    mapping(address => bool) private whileList;

    /**
     * _表是函数执行位置
     */
    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can call this");
        _;
    }
    modifier onlyProposer() {
        require(msg.sender == owner, "Only owner can call this");
        _;
    }
    modifier onlyWhile() {
        require(whileList[msg.sender], "onlyWhile can call this");
        _;
    }

    //只有白名单才可以质押 其他人不可以
    receive() external payable onlyWhile {
        //记录该发送人的质押池
        stakingPool[msg.sender] += msg.value;
        stakingPoolAddresss.push(msg.sender);
    }

    function unstaking(uint amount) external override onlyWhile {
        //解质押
        stakingPool[msg.sender] -= amount;
    }

    function selectProposer() external override onlyOwner {
        require(stakingPoolAddresss.length > 1, "stakingPool is min");
        //随机一个索引
        uint randomIndex = uint(
            keccak256(
                abi.encodePacked(block.timestamp, block.prevrandao, msg.sender)
            )
        ) % stakingPoolAddresss.length;
        //选出验证者
        finalProposer = stakingPoolAddresss[randomIndex];

        for (uint i = 0; i < stakingPoolAddresss.length; i++) {
            if (finalProposer != stakingPoolAddresss[i]) {
                finalcommitees.push(stakingPoolAddresss[i]);
            }
        }
    }

    function proposeBlock() external override onlyProposer onlyWhile {}

    function rewardLiu() external override onlyProposer onlyWhile {}

    function getContractAddress() external override onlyWhile {}

    function getBlockHight() external override onlyWhile returns (uint) {}

    function getBlockTransfore(
        uint hight
    ) external override onlyWhile returns (transforeList[] memory) {}

    function addWhile(address apovAddress) external override {}
}
