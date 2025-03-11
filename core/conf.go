package core

import (
	"awesomeProject1/config"
	"awesomeProject1/globle"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

/*
*
读取yaml文件的配置
*/
func InitConf() {
	const ConfigFile = "setting.yaml"
	//读取配置yaml文件 并且赋值
	c := &config.Config{}
	file, err := ioutil.ReadFile(ConfigFile)

	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(file, c)
	if err != nil {
		log.Fatalln(err.Error())
	}
	//给全局使用的
	globle.Config = c

}
