package db

import (
	"database/sql"

	"gorm.io/gorm"
)

//IDatabase is an interface of db
type IDatabase interface {
	Open() (*gorm.DB, error)
	GetDNS() string
}

// Database ...
type Database struct {
	DB              *sql.DB
	User            string
	Password        string
	Name            string
	Host            string
	TimeZone        string
	Port            int
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime int
}
