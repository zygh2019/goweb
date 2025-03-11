package setting_api

import (
	"awesomeProject1/models/res"
	"github.com/gin-gonic/gin"
)

func (SettingApi) SettingInfoView(c *gin.Context) {
	res.OkWithData(map[string]any{
		"id": 123333,
	}, c)

}
