package setting_api

import (
	"awesomeProject1/models/res"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (SettingApi) SettingInfoView(c *gin.Context) {
	//拿到一个参数
	user := c.Query("user")

	//会返回是否空值
	user1, ok := c.GetQuery("user")
	//拿到多个参数
	user3 := c.QueryArray("user")
	//拿到多个参数
	param := c.Param("user_id")
	params := c.Params
	if !ok {
		res.FailWithMsg("错误", c)
		return
	}
	res.OkWithData(map[string]any{
		"user":   user,
		"user2":  user1,
		"user3":  user3,
		"param":  param,
		"params": params,
	}, c)

}

func (SettingApi) SettingInfoView2(c *gin.Context) {
	//没有就空字符串
	user := c.PostForm("user")

	defuser := c.DefaultPostForm("user", "default")
	a, _ := c.MultipartForm()
	res.OkWithData(map[string]any{
		"user":  user,
		"user2": defuser,
		"user3": a,
	}, c)

}

// 原始参数
type Email struct {
	Email string `json:"email" form:"email" binding:"required"`
	User  string `json:"user"`
	Age   int    `json:"age" binding:"required,min=1"`
}

func (SettingApi) SettingInfoView3(c *gin.Context) {

	email := Email{}
	if err := c.ShouldBindJSON(&email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request data",
			"details": err.Error(),
		})
		return
	}
	//必须是form
	if err := c.ShouldBindQuery(&email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request data",
			"details": err.Error(),
		})
		return
	}

	res.OkWithData(email, c)
}
func (SettingApi) SettingInfoView4(c *gin.Context) {

	email := Email{}
	if err := c.ShouldBind(&email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request data",
			"details": err.Error(),
		})
		return
	}
	res.OkWithData(email, c)
}
