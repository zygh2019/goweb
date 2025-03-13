package routers

import (
	"awesomeProject1/api"
	"github.com/gin-gonic/gin"
)

type SettingRouter struct {
}

// 可以传参数的中间件
func groupMiddle(text string) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func (r RouterGroup) SettingRouters() {
	settingApi := api.ApiGroupApp.SettingApi
	//下面一层
	group := r.Group("/setting")
	group.Use(groupMiddle("test"))
	group.GET("get/:user_id/:a", settingApi.SettingInfoView)
	group.GET("get2", settingApi.SettingInfoView2)
	group.POST("get3", settingApi.SettingInfoView3)
	group.POST("get4", settingApi.SettingInfoView3)
	group.POST("insert", settingApi.SettingInfoView5)
	group.POST("insert1", settingApi.SettingInfoView6)
	group.POST("insert2", settingApi.SettingInfoView7)
	group.POST("insert8", settingApi.SettingInfoView8)
}
