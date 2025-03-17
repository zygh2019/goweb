package routers

import (
	"awesomeProject1/globle"
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(globle.Config.System.Env)
	router := gin.Default()
	//顶层
	group := router.Group("api")
	routerGroup := RouterGroup{group}

	routerGroup.SettingRouters()
	routerGroup.EthClientRouters()
	return router
}
