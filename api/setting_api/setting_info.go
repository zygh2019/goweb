package setting_api

import (
	"awesomeProject1/globle"
	"awesomeProject1/models/res"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
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

type UserDetail struct {
	ID uint64 `json:"id" form:"id" gorm:"primary_key"`
	//不然就是空字符串
	Name   string `json:"name" form:"name"  gorm:"size:10" binding:"required"`
	Remark string `json:"remark" form:"remark"  gorm:"type:varchar(10)" binding:"required"`
	//可以存null
	EmailStr *string `json:"emailStr" form:"emailStr"  gorm:"type:varchar(10)"`
	//不能为空默认值是999
	Password string `json:"password" form:"password" gorm:"type:varchar(10);default:999;comment:密码"  binding:"required"`
}
type UserDetailVO struct {
	ID uint64 `json:"id" form:"id" gorm:"primary_key"`
	//不然就是空字符串
	Name   string `json:"name" form:"name"  gorm:"size:10" binding:"required"`
	Remark string `json:"remark" form:"remark"  gorm:"type:varchar(10)" binding:"required"`
	//可以存null
	EmailStr *string `json:"emailStr" form:"emailStr"  gorm:"type:varchar(10)"`
	//不能为空默认值是999
	Password string `json:"password" form:"password" gorm:"type:varchar(10);default:999;comment:密码"  binding:"required"`
}

// 钩子
func (u UserDetail) BeforeCreate(scope *gorm.DB) error {

	return nil
}
func (SettingApi) SettingInfoView5(c *gin.Context) {
	user := []UserDetail{}

	err := c.ShouldBind(&user)
	if err != nil {
		logrus.Error(err)
		res.FailWithMsg(err.Error(), c)
		return
	}
	globle.DB.AutoMigrate(&UserDetail{})
	result := globle.DB.Debug().Create(&user)
	userDetailOne := UserDetail{ID: 2}
	globle.DB.Debug().Take(&userDetailOne, "id = ?", 2)
	globle.DB.Debug().Take(&userDetailOne)
	//查询多条记录
	globle.DB.Debug().Find(&userDetailOne)
	//globle.DB.Debug().Take(&userDetailOne)
	//globle.DB.Debug().First(&userDetailOne)
	//globle.DB.Debug().Last(&userDetailOne)

	user1 := []UserDetail{}
	globle.DB.Debug().Find(&user1, []int{2}).Update("name", "我是name2")
	globle.DB.Debug().Find(&user1, []int{2}).Updates(UserDetail{Name: "100", Password: "2033"})
	use3 := []UserDetail{}
	use4 := []UserDetailVO{}

	globle.DB.Debug().Order("id desc").Where("id > ? and name = ?", 1, "admin2").Find(&use3).Scan(&use4)

	//globle.DB.Debug().Where("id > ? and name = ?", 1, "我是name2").Delete(&UserDetail{})
	res.OkWithData(map[string]any{
		"ID":           user,
		"RowsAffected": result.RowsAffected,
		"Error":        result.Error,
		"data":         userDetailOne,
		"wheredat":     use3,
		"use4":         use4,
	}, c)
}
