package geth_api

import (
	"awesomeProject1/globle"
	"awesomeProject1/models/res"
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"math/big"
	"strconv"
)

func (GethApi) GetBlock(c *gin.Context) {
	param := c.Query("number")
	var blockNumber *big.Int
	if param == "" {
		blockNumber = nil
	} else {
		value, _ := strconv.Atoi(param)
		blockNumber = big.NewInt(int64(value))
	}
	header, _ := ethclient.Dial(globle.Config.GethConfig.GetGatewayAddr())
	//设置高度

	//如果是空的 查询最新一条区块
	//blockNumber = nil
	block, _ := header.BlockByNumber(context.Background(), blockNumber)

	res.OkWithData(map[string]any{
		"区块号":       block.Number(),
		"区块时间戳":     block.Time(),
		"交易的区块hash": block.Hash().Hex(),
		"区块链摘要":     block.Difficulty(),
		"共多少交易":     len(block.Transactions()),
	}, c)
}

func (GethApi) GetTransactions(c *gin.Context) {
	param := c.Query("number")
	var blockNumber *big.Int
	if param == "" {
		blockNumber = nil
	} else {
		value, _ := strconv.Atoi(param)
		blockNumber = big.NewInt(int64(value))
	}

	dial, err := ethclient.Dial(globle.Config.GethConfig.GetGatewayAddr())

	if err != nil {
		logrus.Info(err.Error())
		res.Fail(c)
		return
	}
	number, err := dial.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		logrus.Info(err.Error())
		res.Fail(c)
		return
	}
	mapReturn := make(map[string]any)
	for _, tx := range number.Transactions() {
		mapReturn["交易的hash值"] = tx.Hash().Hex()
		mapReturn["用了多少燃料"] = tx.Gas()
		mapReturn["燃料多少钱"] = tx.GasPrice()
		mapReturn["Nonce"] = tx.Nonce()
		mapReturn["Data"] = tx.Data()
		mapReturn["接收人hash"] = tx.To().Hex()
		chainID, err := dial.ChainID(context.Background())
		if err != nil {
			logrus.Error(err.Error())
		}
		if sender, err := types.Sender(types.NewLondonSigner(chainID), tx); err == nil {
			mapReturn["发送人hash"] = sender.Hex()
		} else {
			logrus.Error(err.Error(), tx.Type())
		}
		receipt, err := dial.TransactionReceipt(context.Background(), tx.Hash())

		mapReturn["收据状态"] = receipt.Status
		mapReturn["收据日志"] = receipt.Logs
	}

	res.OkWithData(mapReturn, c)

}
