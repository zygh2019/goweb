package api

import "awesomeProject1/api/setting_api"

type ApiGroup struct {
	setting_api.SettingApi
}

var ApiGroupApp = new(ApiGroup)
