package db

import (
	"time"

	"github.com/0ero-1ne/martha-server/internal/models"
	"gorm.io/gorm"
)

type Database interface {
	Connect() (*gorm.DB, error)
}

func InitDatabase(db Database) (*gorm.DB, error) {
	connection, err := db.Connect()

	if err != nil {
		return nil, err
	}

	err = migrate(connection)

	if err != nil {
		return nil, err
	}

	sqlDB, err := connection.DB()

	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return connection, nil
}

func migrate(connection *gorm.DB) error {
	err := connection.AutoMigrate(
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
		return err
	}

	return nil
}
