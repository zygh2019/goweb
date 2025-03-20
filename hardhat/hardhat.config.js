require("@nomicfoundation/hardhat-toolbox");


module.exports = {
  solidity: "0.8.28",
  networks: {
    ganache: {
      url: 'http://127.0.0.1:7546', // Ganache 的 RPC 地址
      chainId: 1337, // Ganache 默认的 chainId 是 1337
      accounts: [
        // 你可以在这里添加 Ganache 中的账户私钥
        '0x5991cd76f703ee21d0cc940986a6513f349774e401a2ff2fe62f0ebe39f2c66c',
        // 更多账户...
      ],
    },
  },
};


