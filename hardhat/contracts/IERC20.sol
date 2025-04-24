// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.10;


interface IERC20 {
    function totalSupply() external view returns (uint256); // 返回总供应量
    function balanceOf(address account) external view returns (uint256); // 查询指定账户余额
    function transfer(address recipient, uint256 amount) external returns (bool); // 代币转账
    function allowance(address owner, address spender) external view returns (uint256); // 查询授权额度
    function approve(address spender, uint256 amount) external returns (bool); // 授权指定账户的额度
    function transferFrom(address sender, address recipient, uint256 amount) external returns (bool); // 从授权账户转账
    // 转账和授权事件
    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(address indexed owner, address indexed spender, uint256 value) ;

}
