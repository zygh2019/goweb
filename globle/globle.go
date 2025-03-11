package globle

import (
	"awesomeProject1/config"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	DB     *gorm.DB
)
