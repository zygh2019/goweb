// SPDX-License-Identifier: Unlicense
// 定义版本
pragma solidity ^0.8.10;

interface Ipledge {
    event stakingEvent(address pledgor, uint amount);
    event unStakingEvent(address pledgor, uint amount);
    struct transforeList {
        address from;
        address to;
        uint amount;
    }

    /**
     *
     * @param apovAddress  增加白名单
     */
    function addWhile(address apovAddress) external;

    /**
     * 解质押方法
     */
    function unstaking(uint amount) external;

    /**
     * 选举提供者
     */
    function selectProposer() external;

    /**
     * 提供者打包区块
     */
    function proposeBlock() external;

    /**
     * 奖励代币
     */
    function rewardLiu() external;

    /**
     * 获取合约地址
     */
    function getContractAddress() external;

    /**
     * 获取当前块的高度
     */
    function getBlockHight() external returns (uint);

    /**
     * 获取当前块的交易
     */
    function getBlockTransfore(
        uint hight
    ) external returns (transforeList[] memory);
}
