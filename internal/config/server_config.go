package config

import "fmt"

type ServerConfig struct {
	Host string `ini:"host"`
	Port string `ini:"port"`
}

func (config ServerConfig) GetAddress() string {
	return fmt.Sprintf("%s:%s", config.Host, config.Port)
}
