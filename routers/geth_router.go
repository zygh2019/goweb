package routers

import (
	"awesomeProject1/api"
)

type EthClientRouter struct {
}


func (r RouterGroup) EthClientRouters() {
	ethClientApi := api.ApiGroupApp.EthClientApi
	//下面一层
	group := r.Group("/eth_client")
	group.GET("getBlockInfo", ethClientApi.)

}
