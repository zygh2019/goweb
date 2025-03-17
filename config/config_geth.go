package config

import "strconv"

type Mysql struct {
	Host            string `yaml:"host"`
	Config          string `yaml:"config"`
	Port            int    `yaml:"port"`
	Db              string `yaml:"db"`
	Username        string `yaml:"username"`
	Password        string `yaml:"password"` //日志登记
	MaxIdleConns    int    `yaml:"max_idle_conns"`
	MaxOpenConns    int    `yaml:"max_open_conns"`
	ConnMaxLifetime int    `yaml:"conn_max_lifetime"`
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.Db + "?"
}
