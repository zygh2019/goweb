package routers

import (
	"awesomeProject1/api"
)

func (r RouterGroup) SettingRouters() {
	settingApi := api.ApiGroupApp.SettingApi
	//下面一层
	group := r.Group("/setting")
	group.GET("get/:user_id/:a", settingApi.SettingInfoView)
	group.GET("get2", settingApi.SettingInfoView2)
	group.POST("get3", settingApi.SettingInfoView3)
	group.POST("get4", settingApi.SettingInfoView4)
}
