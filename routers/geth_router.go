package routers

import (
	"awesomeProject1/api"
)

func (r RouterGroup) EthClientRouters() {
	gethApi := api.ApiGroupApp.GethApi
	//下面一层
	group := r.Group("/eth_client")

	group.GET("/getBlockInfo", gethApi.GetBlock)
	group.GET("/transactions", gethApi.GetTransactions)
	group.POST("/transactionsByBlock", gethApi.TransactionsByBlock)
	/**
	创建钱包
	*/
	group.GET("/createPurse", gethApi.CreatePurse)
	/**
	查询yu饿
	*/
	group.GET("/getBalance", gethApi.GetBalance)
}
