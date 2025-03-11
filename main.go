package main

import (
	"awesomeProject1/core"
	"awesomeProject1/globle"
	"log"
)

func main() {
	core.InitConf()
	core.InitGorm()
	log.Println(globle.DB)
}
