package dbconfig

import (
	"server/config"
	"server/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func InitDatabase(config config.Config) {
	if database != nil {
		return
	}

	connect(config)
	migrate()
}

func connect(config config.Config) {
	var err error
	database, err = gorm.Open(postgres.Open(config.PostgresConfig.GetDSN()), &gorm.Config{})

	if err != nil {
		panic("Can not connect to database: " + err.Error())
	}
}

func migrate() {
	err := database.AutoMigrate(&models.Author{}, &models.Book{}, &models.Tag{})

	if err != nil {
		panic("Can not migrate models: " + err.Error())
	}
}

func GetDB() *gorm.DB {
	return database
}
