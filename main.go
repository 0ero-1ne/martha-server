package main

import (
	"server/config"
	"server/db"
)

func main() {
	cfg := config.Init()
	db.InitDatabase(cfg)
}
