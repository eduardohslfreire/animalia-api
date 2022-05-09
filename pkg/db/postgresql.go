package db

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Postgresql ...
type Postgresql struct {
	Database
}

// Open represent a factory of database.
// Will initialize db session.
func (p *Postgresql) Open() (*gorm.DB, error) {
	var err error
	var db *gorm.DB

	connectionString := p.GetDNS()

	db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tab_", // table name prefix, table for `User` would be `tab_users`
			SingularTable: true,   // use singular table name, table for `User` would be `user` with this option enabled
		},
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
		/*
			uncomment to debug

			Logger: logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags),
				logger.Config{
					LogLevel: logger.Info, // Log level Info will output everything
				},
			),*/
	})

	p.DB, err = db.DB()
	if err != nil {
		return nil, err
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	p.DB.SetMaxIdleConns(p.MaxIdleConns)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	p.DB.SetMaxOpenConns(p.MaxOpenConns)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	p.DB.SetConnMaxLifetime(time.Second * time.Duration(p.ConnMaxLifetime))

	return db, err
}

// GetDNS representa a recuperação do acesso ao banco
func (p *Postgresql) GetDNS() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=%s", p.Host, p.Port, p.User, p.Password, p.Name, p.TimeZone)
}
