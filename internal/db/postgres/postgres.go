package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/0ero-1ne/martha-server/internal/config"
)

type PostgresDatabase struct {
	config config.PostgresConfig
}

func NewPostgresDatabase(config config.PostgresConfig) PostgresDatabase {
	return PostgresDatabase{
		config: config,
	}
}

func (database PostgresDatabase) Connect() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(database.config.GetDSN()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, err
}
