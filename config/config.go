package config

import "gopkg.in/ini.v1"

type Config struct {
	Database struct {
		DSN string `ini:"host"`
	} `ini:"database"`
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
