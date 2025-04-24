package routers

import (
	"awesomeProject1/globle"

	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	*gin.RouterGroup
}

//	func errorMiddleware() gin.HandlerFunc {
//		return func(c *gin.Context) {
//			switch c.ContentType() {
//			case "application/json":
//				var anyy any
//				// 请求前逻辑
//				err := c.ShouldBind(&anyy)
//				if err != nil {
//					errinfo := ValidateErrors(err, anyy)
//					if errinfo != "" {
//						res.FailWithMsg(errinfo, c)
//						c.Abort()
//						return
//					}
//				}
//			}
//			// 继续处理请求
//			c.Next()
//		}
//	}
func InitRouter() *gin.Engine {
	gin.SetMode(globle.Config.System.Env)
	router := gin.Default()
	//顶层
	group := router.Group("api")
	//group.Use(errorMiddleware())
	routerGroup := RouterGroup{group}
	
	routerGroup.SettingRouters()
	routerGroup.EthClientRouters()
	return router
}

