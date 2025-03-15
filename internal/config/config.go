package config

import (
	"gopkg.in/ini.v1"
)

type DatabaseConfig interface {
	GetDSN() string
}

type Config struct {
	PostgresConfig PostgresConfig `ini:"database.postgres"`
	ServerConfig   ServerConfig   `ini:"server"`
}

func Init(configPath string) Config {
	iniData, err := ini.Load(configPath)

	if err != nil {
		panic("Can not load config.ini file: " + err.Error())
	}

	var config Config
	err = iniData.MapTo(&config)

	if err != nil {
		panic("Can not parse ini data: " + err.Error())
	}

	return config
}
