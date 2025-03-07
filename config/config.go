package config

import (
	"gopkg.in/ini.v1"
)

type Config struct {
	PostgresConfig PostgresConfig `ini:"database.postgres"`
}

func Init() Config {
	iniData, err := ini.Load("config.ini")

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
