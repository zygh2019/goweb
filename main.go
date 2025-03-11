package main

import (
	"awesomeProject1/core"
	"awesomeProject1/globle"
	"github.com/sirupsen/logrus"
)

func main() {
	core.InitConf()
	core.InitGorm()
	globle.Log = core.InitLogger()

	logrus.Infof("init default logger")
}
