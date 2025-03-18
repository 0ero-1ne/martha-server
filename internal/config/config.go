package config

import (
	"fmt"

	"gopkg.in/ini.v1"
)

type Config struct {
	PostgresConfig PostgresConfig `ini:"database.postgres"`
	ServerConfig   ServerConfig   `ini:"server"`
	JWTConfig      JWTConfig      `ini:"jwt"`
}

func Init(configPath string) Config {
	iniData, err := ini.Load(configPath)

	if err != nil {
		panic(fmt.Sprintf("Can not load %s config file: %s", configPath, err.Error()))
	}

	var config Config
	err = iniData.MapTo(&config)

	if err != nil {
		panic("Can not parse ini data: " + err.Error())
	}

	return config
}
