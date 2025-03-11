package routers

import (
	"awesomeProject1/api"
)

func (r RouterGroup) SettingRouters() {
	settingApi := api.ApiGroupApp.SettingApi
	//下面一层
	group := r.Group("/setting")
	group.GET("get", settingApi.SettingInfoView)
}
