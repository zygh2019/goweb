// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

/**
 * @title 创建合约必须写版本号 并且 需要用contract 关键字表示构建了一个合约
 *
 * @author
 * @notice
 */
contract Hello {
    //stack 1024个slot
    //每个sot 是每个slot 字节

    //memory evm 自身的内存
    //storage 链上永久存储
    //
    //memory evm 自身的内存
    //状态变量相当于常量
    /**
     * @notice 与java 类似 不过public 在变量后面
     */
    string public hello = "";
    /**
     * @notice 8位
     */
    int8 public i = -1;
    /**
     * @notice 只能是正数
     */
    uint public a2 = 1 * 2 ** 256 - 1;
    bool public b = true;
    address public addr = 0xD918f784bf508B20e3c13be414483cA5E305F589;
    /**
     * @dev 字节
     */
    bytes2 public b2;
    enum staus {
        Actity,
        Stop
    }

    int[] public arry;
    string[] public arr2;
    bool[] public bools;
    /**
     * 结构体
     */
    struct Name {
        string name;
    }
    Name name = Name("aa");

    Name name1 = Name({name: "aaa"});

    /**
     * @dev 以太坊转账的地址 必须显式转换为 address payable
     */
    address payable payableAddress = payable(addr); // address 转 address payable
    /**
     * @dev 获取余额
     */
    uint256 balance = addr.balance;

    /**
     * 返回当前合约地址
     */
    function getAddress() public view returns (address) {
        return address(this); // 返回当前合约的地址
    }

    function destroyContract(address payable recipient) public {
        selfdestruct(recipient); // 销毁合约并发送以太币
    }
}

/**
 * @title type(Hello).name: 获取合约的名字。
type(Hello).creationCode: 获取创建合约的字节码。
type(Hello).runtimeCode: 获取合约运行时的字节码。
 * @author 
 * @notice 
 */
//获取合约的类型
contract HelloType {
    function getContractInfo()
        public
        pure
        returns (string memory, bytes memory, bytes memory)
    {
        return (
            type(Hello).name,
            type(Hello).creationCode,
            type(Hello).runtimeCode
        );
    }
}

/**
 * @title 因为合约账户是有有合约大小的 所有需要用这个判断是否为合约账户
 * @author
 * @notice
 */
contract AddressChecker {
    function isContract(address addr) internal view returns (bool) {
        uint256 size;

        assembly {
            size := extcodesize(addr)
        } // 获取地址的代码大小
        return size > 0; // 大于 0 说明是合约地址
    }

    function getSelect() public pure returns (bytes4) {
        return this.AddressChecker.getSelect;
    }
}

contract Limit {
    modifier greaterThanZero(uint256 _value, uint256 _value1) {
        require(_value > 0, "Value must be greater than zero");
        _;
    }

    function setValue(uint256 _num) public greaterThanZero(_num) {
        // 执行函数逻辑
    }

    function setaa() public returns (uint a, uint b, uint c) {
        a = 1;
        b = 2;
        return (a, b, 2);
    }

    function aaa() public {
        (uint a, , ) = setaa();
    }
}
