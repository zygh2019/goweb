package geth_api

import (
	"awesomeProject1/globle"
	"awesomeProject1/models/res"
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/sha3"
	"log"
	"math"
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
		"区块号":         block.Number(),
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

type Purse struct {
	Id               uint   `json:"id" gorm:"primary_key"`
	PrivateKey       string `json:"privateKey16" gorm:"size:255"`
	PublicKey        string `json:"publicKey16" gorm:"size:255"`
	PublicKeyAddress string `json:"publicKey16Address" gorm:"size:255"`
}

func (a GethApi) CreatePurse(c *gin.Context) {
	//生成一个随机的私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	//私钥转换成字节
	privateKeyBytes := crypto.FromECDSA(privateKey)
	//转换为16进制
	privateKeyx16 := hexutil.Encode(privateKeyBytes)[2:]
	//生成公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	//去除不需要的自负
	publicKey16 := hexutil.Encode(publicKeyBytes)[4:]

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	//创建表
	globle.DB.AutoMigrate(&Purse{})
	//创建
	purse := &Purse{PrivateKey: privateKeyx16, PublicKey: publicKey16, PublicKeyAddress: address}
	globle.DB.Create(purse)
	res.OkWithData(purse, c)

}

func (a GethApi) GetBalance(c *gin.Context) {
	value := c.Query("address")
	client, err := ethclient.Dial(globle.Config.GethConfig.GetGatewayAddr())
	account := common.HexToAddress(value)
	//nil为最新区块的金额
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	/**
	以太坊中的数字是使用尽可能小的单位来处理的，因为它们是定点精度，在ETH中它是wei。要读取ETH值，您必须做计算wei/10^18。因为我们正在处理大数，我们得导入原生的Gomath和math/big包。这是您做的转换。
	*/
	res.OkWithData(BigInt2StringBalance(balance), c)
}

// 进行转换
func BigInt2StringBalance(balance *big.Int) *big.Float {
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	return ethValue
}

type TransferLog struct {
	Id          uint   `json:"id" gorm:"primary_key"`
	FromAddress string `json:"fromAddress" gorm:"size:100"`
	ToAddress   string `json:"toAddress" gorm:"size:100"`
	GasLimit    uint64 `json:"gasLimit" gorm:"size:100"`
	GasPrice    string `json:"gasPrice" gorm:"size:100"`
	WeiValue    string `json:"weiValue" gorm:"size:100"`
	EthAmount   string `json:"ethAmount" gorm:"size:100"`
}

func (a GethApi) Transfer(c *gin.Context) {
	param := &struct {
		Sender   string `json:"sender" binding:"required"`
		Receiver string `json:"receiver" binding:"required"`
		Amount   string `json:"amount" binding:"required"`
	}{}
	err := c.ShouldBind(param)
	if err != nil {
		errors := res.ValidateErrors(err, param)
		logrus.Error(errors)
		res.FailWithMsg(errors, c)
		return
	}
	client, err := ethclient.Dial(globle.Config.GethConfig.GetGatewayAddr())

	//3.计算publickey
	privateKey, err := crypto.HexToECDSA(param.Sender)
	if err != nil {
		logrus.Error(err)
		res.FailWithMsg(err.Error(), c)
		return
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	//公要地址
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	//校验金额
	//1，校验金额
	balance, err := client.BalanceAt(context.Background(), fromAddress, nil)
	if err != nil {
		logrus.Info(err.Error())
		res.FailWithMsg(err.Error(), c)
		return
	}
	stringBalance := BigInt2StringBalance(balance)
	amount := new(big.Float)
	_, success := amount.SetString(param.Amount)
	if !success {
		logrus.Error(success)
		res.Fail(c)
		return
	}
	// 2.比较 amount 和 stringBalance
	if stringBalance.Cmp(amount) < 0 {
		res.FailWithMsg("钱不充足", c)
		return
	}

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	gasLimit := uint64(21000) // in units
	// 定义 1 ETH 对应的 wei 值
	oneETHInWei := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	//转换成float类型
	flowtypeWEi := new(big.Float).SetInt(oneETHInWei)
	//将param转换成float
	paramAmount := new(big.Float)
	paramAmount.SetString(param.Amount)
	//计算最终的wei值
	weiAmount := new(big.Int)
	flowtypeWEi.Mul(flowtypeWEi, paramAmount).Int(weiAmount)
	//要发送的人的地址
	toAddress := common.HexToAddress(param.Receiver)
	tx := types.NewTransaction(nonce, toAddress, weiAmount, gasLimit, gasPrice, nil)

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		logrus.Error(err)
		res.FailWithMsg(err.Error(), c)
		return
	}
	//交易进行签名
	signedTx, err := types.SignTx(tx, types.NewLondonSigner(chainID), privateKey)
	if err != nil {
		logrus.Error(err)
		res.FailWithMsg(err.Error(), c)
		return
	}
	//进行发送
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		logrus.Error(err)
		res.FailWithMsg(err.Error(), c)
		return
	}
	globle.DB.AutoMigrate(&TransferLog{})

	transferLog := &TransferLog{
		FromAddress: fromAddress.String(),
		ToAddress:   toAddress.String(),
		GasLimit:    gasLimit,
		GasPrice:    gasPrice.String(),
		WeiValue:    weiAmount.String(),
		EthAmount:   paramAmount.String(),
	}
	globle.DB.Create(transferLog)
	res.OkWithData(transferLog, c)

}

func (a GethApi) TokenTransfer(c *gin.Context) {

	param := &struct {
		Sender   string `json:"sender" binding:"required"`
		Receiver string `json:"receiver" binding:"required"`
		Amount   string `json:"amount" binding:"required"`
	}{}
	err := c.ShouldBind(param)
	if err != nil {
		errors := res.ValidateErrors(err, param)
		logrus.Error(errors)
		res.FailWithMsg(errors, c)
		return

	}
	client, err := ethclient.Dial(globle.Config.GethConfig.GetGatewayAddr())

	privateKey, err := crypto.HexToECDSA(param.Sender)
	if err != nil {
		logrus.Error(err)
		res.FailWithMsg(err.Error(), c)
		return
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	value := big.NewInt(0) // in wei (0 eth)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(gasPrice) // 23256
	toAddress := common.HexToAddress(param.Receiver)

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress)) // 0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d
	amount := new(big.Int)
	amount.SetString("100000000000000000000", 10) // 1000 tokens
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount)) // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	})
	if err != nil {
		logrus.Error(err)
		res.FailWithMsg(err.Error(), c)
		return
	}
	fmt.Println(gasLimit) // 23256
	tokenAddress := common.HexToAddress(globle.Config.GethConfig.Contracts)
	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		logrus.Error(err)
		res.FailWithMsg(err.Error(), c)
		return
	}

	signedTx, err := types.SignTx(tx, types.NewLondonSigner(chainID), privateKey)
	if err != nil {
		logrus.Error(err)
		res.FailWithMsg(err.Error(), c)
		return
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		logrus.Error(err)
		res.FailWithMsg(err.Error(), c)
		return
	}

	res.OkWithData(map[string]any{
		"value": value.String(),
	}, c)

}
