package config

type Config struct {
	Mysql  `yaml:"mysql"`
	Logger `yaml:"logger"`
	System `yaml:"system"`
}

// mysql结构体
