const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules");

module.exports = buildModule("LiuTokenModul", (m) => {
    const name = "LiuToken";       // 代币名称
    const symbol = "LIU";          // 代币符号
    const decimals = 18;           // 小数位数
    const initialSupply = 1000000; // 初始供应量 (1百万)

    const contract = m.contract("LiuToken", [name, symbol, decimals, initialSupply]);

    return { contract };
});
