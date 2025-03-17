package api

import (
	"awesomeProject1/api/geth_api"
	"awesomeProject1/api/setting_api"
)

type ApiGroup struct {
	setting_api.SettingApi
	geth_api.GethApi
}

var ApiGroupApp = new(ApiGroup)
