package db

import (
	"fmt"
	"time"

	"github.com/eduardohslfreire/animalia-api/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Postgresql ...
type Postgresql struct {
	Database
	Logger *logger.GenericLogger
}

// Open represent a factory of database.
// Will initialize db session.
func (p *Postgresql) Open() (*gorm.DB, error) {
	var err error
	var db *gorm.DB

	p.Logger = logger.NewGenericLogger()

	connectionString := p.GetDNS()

	db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tab_", // table name prefix, table for `User` would be `tab_users`
			SingularTable: true,   // use singular table name, table for `User` would be `user` with this option enabled
		},
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	})

	p.DB, err = db.DB()
	if err != nil {
		p.Logger.LogIt("ERROR", fmt.Sprintf("[DATABASE-SESSION-ERROR] - Failed to get session with database. %s", err.Error()), nil)
		return nil, err
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	p.DB.SetMaxIdleConns(p.MaxIdleConns)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	p.DB.SetMaxOpenConns(p.MaxOpenConns)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	p.DB.SetConnMaxLifetime(time.Second * time.Duration(p.ConnMaxLifetime))

	p.Logger.LogIt("INFO", "[DATABASE-INIT] - Connection to database started", nil)

	return db, err
}

// GetDNS ...
func (p *Postgresql) GetDNS() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=%s", p.Host, p.Port, p.User, p.Password, p.Name, p.TimeZone)
}
