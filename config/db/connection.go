package db

import (
	"github.com/eduardohslfreire/animalia-api/config/env"
	"github.com/eduardohslfreire/animalia-api/pkg/db"
	"gorm.io/gorm"
)

// InitDatabase ...
func InitDatabase() (*gorm.DB, error) {
	database := new(db.Postgresql)
	database.Host = env.DbHost
	database.Port = env.DbPort
	database.User = env.DbUser
	database.Password = env.DbPassword
	database.Name = env.DbName
	database.TimeZone = env.DbTimeZone
	database.MaxIdleConns = env.DbMaxIdleConns
	database.ConnMaxLifetime = env.DbConnMaxLifeTime
	database.MaxOpenConns = env.DbMaxOpenConns

	return database.Open()
}
