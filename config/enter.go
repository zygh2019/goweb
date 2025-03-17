package config

type Config struct {
	Mysql      `yaml:"mysql"`
	Logger     `yaml:"logger"`
	System     `yaml:"system"`
	GethConfig `yaml:"geth_gateway" mapstructure:"geth_gateway"`
}

// mysql结构体
