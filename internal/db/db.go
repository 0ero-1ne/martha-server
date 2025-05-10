package db

import (
	"time"

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

	sqlDB, err := connection.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return connection, nil
}
