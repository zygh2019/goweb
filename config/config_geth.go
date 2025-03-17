package config

import "fmt"

type GethConfig struct {
	Url string `yaml:"url"`
	Key string `yaml:"key"`
}

func (s *GethConfig) GetGatewayAddr() string {
	return fmt.Sprintf("%s/%s", s.Url, s.Key)
}
