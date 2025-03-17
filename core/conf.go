package core

import (
	"awesomeProject1/config"
	"awesomeProject1/globle"
	"github.com/spf13/viper"
)

/*
*
读取yaml文件的配置
*/
func InitConf() {
	viper.SetConfigName("setting")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	error := viper.ReadInConfig()
	if error != nil {
		panic(error)
	}
	c := config.Config{}
	error = viper.Unmarshal(&c)
	if error != nil {

		panic(error)
	}
	//读取配置yaml文件 并且赋值

	//给全局使用的
	globle.Config = &c

}
