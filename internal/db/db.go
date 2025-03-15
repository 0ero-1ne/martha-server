package db

import (
	"github.com/0ero-1ne/martha/internal/config"
	"github.com/0ero-1ne/martha/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func InitDatabase(config config.DatabaseConfig) {
	if database != nil {
		return
	}

	connect(config)
	migrate()
}

func connect(config config.DatabaseConfig) {
	var err error
	database, err = gorm.Open(postgres.Open(config.GetDSN()), &gorm.Config{})

	if err != nil {
		panic("Can not connect to database: " + err.Error())
	}
}

func migrate() {
	err := database.AutoMigrate(
		&models.Author{},
		&models.Book{},
		&models.Tag{},
		&models.User{},
		&models.Chapter{},
		&models.Comment{},
		&models.CommentRate{},
		&models.BookRate{},
	)

	if err != nil {
		panic("Can not migrate models: " + err.Error())
	}
}

func GetDB() *gorm.DB {
	return database
}
