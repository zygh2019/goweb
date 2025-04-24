// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.10;

import {IERC20} from "./IERC20.sol";


contract LiuToken is IERC20 {
    string public name;  // 代币名称
    string public symbol;    // 代币符号
    uint8 public decimals;       // 小数位数，标准为18位
    uint256 private _totalSupply;     // 总供应量
    
    // 账户余额映射
    mapping(address => uint256) private _balances;
    // 授权映射：存储授权的额度
    mapping(address => mapping(address => uint256)) private _allowances;

    address public owner; // 合约拥有者

    // 修饰器：仅合约拥有者可调用的函数
    modifier onlyOwner() {
        require(msg.sender == owner, "Not the contract owner");
        _;
    }

    // 构造函数，初始化代币的名称、符号、小数位数和初始供应量
    constructor(string memory _name, string memory _symbol, uint8 _decimals, uint256 initialSupply) {
        owner = msg.sender;
        name = _name;
        symbol = _symbol;
        decimals = _decimals;
        _mint(owner, initialSupply * 10 ** _decimals); // 铸造代币并分配给合约拥有者
    }

    // 返回总供应量
    function totalSupply() external view override returns (uint256) {
        return _totalSupply;
    }

    // 查询账户余额
    function balanceOf(address account) external view override returns (uint256) {
        return _balances[account];
    }

    // 代币转账
    function transfer(address recipient, uint256 amount) external override returns (bool) {
        require(_balances[msg.sender] >= amount, "Insufficient balance"); // 确保余额充足
        _balances[msg.sender] -= amount;
        _balances[recipient] += amount;
        emit Transfer(msg.sender, recipient, amount); // 触发转账事件
        return true;
    }

    // 查询授权额度
    function allowance(address owner, address spender) external view override returns (uint256) {
        return _allowances[owner][spender];
    }

    // 授权指定账户的额度
    function approve(address spender, uint256 amount) external override returns (bool) {
        _allowances[msg.sender][spender] = amount;
        emit Approval(msg.sender, spender, amount); // 触发授权事件
        return true;
    }

    // 从授权账户转账
    function transferFrom(address sender, address recipient, uint256 amount) external override returns (bool) {
        require(_balances[sender] >= amount, "Insufficient balance");
        require(_allowances[sender][msg.sender] >= amount, "Allowance exceeded");
        _balances[sender] -= amount;
        _balances[recipient] += amount;
        _allowances[sender][msg.sender] -= amount;
        emit Transfer(sender, recipient, amount); // 触发转账事件
        return true;
    }

    // 铸造新代币，仅限合约拥有者
    function mint(address to, uint256 amount) public onlyOwner {
        _mint(to, amount);
    }

    // 销毁代币
    function burn(uint256 amount) public {
        require(_balances[msg.sender] >= amount, "Insufficient balance");
        _balances[msg.sender] -= amount;
        _totalSupply -= amount;
        emit Transfer(msg.sender, address(0), amount); // 触发销毁事件
    }

    // 内部函数：铸造新代币
    function _mint(address to, uint256 amount) internal {
        _totalSupply += amount;
        _balances[to] += amount;
        emit Transfer(address(0), to, amount); // 触发铸造事件
    }


}
