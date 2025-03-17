package main

import (
	"awesomeProject1/core"
	"awesomeProject1/globle"
	"awesomeProject1/routers"
)

func main() {

	core.InitConf()
	core.InitLogger()
	core.InitGorm()
	server := routers.InitRouter()
	globle.Log.Info("启动完成", globle.Config.System.GetAddr())
	server.Run(globle.Config.System.GetAddr())

}
