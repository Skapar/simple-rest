package app

import (
	"github.com/Skapar/simple-rest/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(config *config.Config) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(config.PostgresAddr), &gorm.Config{})
}
