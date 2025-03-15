package main

import "github.com/0ero-1ne/martha-server/internal/app"

const configPath = "configs/config.ini"

func main() {
	app.Run(configPath)
}
