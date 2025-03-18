package geth_api

import (
	"awesomeProject1/globle"
	"awesomeProject1/models/res"
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"math/big"
	"strconv"
)

type GethApi struct {
}

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

func (a GethApi) TransactionsByBlock(c *gin.Context) {
	mapReturn := make(map[string]any)
	var test = &struct {
		BlockHash string `json:"blockHash"  binding:"required"`
		TransHash string `json:"transHash"  binding:"required"`
		Index     uint   `json:"index" `
	}{}
	err := c.ShouldBind(test)
	if err != nil {
		errors := res.ValidateErrors(err, test)
		logrus.Error(errors)
		res.FailWithMsg(errors, c)
		return
	}

	client, err := ethclient.Dial(globle.Config.GethConfig.GetGatewayAddr())
	if client == nil {
		logrus.Error(err.Error())
		res.FailWithMsg(err.Error(), c)
		return
	}
	//根据块的hash获取数量
	count, err := client.TransactionCount(context.Background(), common.HexToHash(test.BlockHash))
	if err != nil {
		logrus.Error(err.Error())
		res.FailWithMsg(err.Error(), c)
		return
	}
	//交易的hash
	tx, pending, err := client.TransactionByHash(context.Background(), common.HexToHash(test.TransHash))
	if err != nil {
		logrus.Error(err.Error())
		res.FailWithMsg(err.Error(), c)
		return
	}
	mapReturn["区块交易数量"] = count
	mapReturn["pending"] = pending
	mapReturn["交易的hash值"] = tx.Hash().Hex()
	mapReturn["用了多少燃料"] = tx.Gas()
	mapReturn["燃料多少钱"] = tx.GasPrice()
	mapReturn["Nonce"] = tx.Nonce()
	mapReturn["Data"] = tx.Data()
	mapReturn["接收人hash"] = tx.To().Hex()
	//该区块中的第几条交易 从0开始
	block, err := client.TransactionInBlock(context.Background(), common.HexToHash(test.BlockHash), test.Index)

	if err != nil {
		logrus.Info(err.Error())
		res.Fail(c)
		return
	}
	mapReturn["根据索引获取hash"] = &map[string]any{
		"hash": block.Hash().Hex(),
	}
	res.OkWithData(mapReturn, c)

}
