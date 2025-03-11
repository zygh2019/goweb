package config

type config struct {
	Mysql  `yaml:"mysql"`
	Logger `yaml:"logger"`
	System `yaml:"system"`
}

// mysql结构体
