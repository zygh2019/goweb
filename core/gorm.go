package core

import (
	"awesomeProject1/globle"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func InitGorm() *gorm.DB {
	//判断配置是否存在
	if "" == globle.Config.Mysql.Host {
		globle.Log.Error("globle.Config.Mysql.Host is null")
		return nil
	}
	//链接地址
	dsn := globle.Config.Mysql.Dsn()
	var mysqlLogger logger.Interface
	//判断环境配置日志级别
	if globle.Config.System.Env == "dev" {
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		mysqlLogger = logger.Default.LogMode(logger.Error)
	}
	//开启gorm 并设置自定义日志配置
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	//如果开启失败 报错 并返回
	if err != nil {
		globle.Log.Error("数据库连接失败 ", err)
		return nil
	}
	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(globle.Config.Mysql.MaxIdleConns)
	sqlDb.SetMaxOpenConns(globle.Config.Mysql.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Hour * time.Duration(globle.Config.Mysql.ConnMaxLifetime))
	globle.DB = db
	globle.Log.Info("数据库和gorm初始化完成")
	return db
}
