package setting_api

import "github.com/gin-gonic/gin"

func (SettingApi) SettingInfoView(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "SettingInfoView"})
}
