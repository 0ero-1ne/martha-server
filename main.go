package main

import (
	"server/config"
	"server/dbconfig"
)

func main() {
	cfg := config.Init()
	dbconfig.InitDatabase(cfg)
}
