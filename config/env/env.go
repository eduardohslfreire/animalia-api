package env

import (
	"github.com/eduardohslfreire/animalia-api/pkg/env"
)

var (
	// AppPort application port
	AppPort string

	// HTTPReadTimeout timeout in seconds for api (read)
	HTTPReadTimeout int

	// HTTPRequestTimeout timeout in seconds for api (requests)
	HTTPRequestTimeout int

	// HTTPWriteTimeout timeout in seconds for api (write)
	HTTPWriteTimeout int

	// DbUser database application user
	DbUser string

	// DbPassword database application user password
	DbPassword string

	// DbName database name
	DbName string

	// DbHost database hostname
	DbHost string

	// DbPort database port
	DbPort int

	// DbTimeZone database timezone
	DbTimeZone string

	// DbMaxIdleConns maximum number of idle connections in the database
	DbMaxIdleConns int

	// DbMaxOpenConns maximum number of open connections in the database
	DbMaxOpenConns int

	// DbConnMaxLifeTime maximum time for which an idle connection can be reused in seconds
	DbConnMaxLifeTime int

	// LogLevel determines the log level to be displayed in the stdout: DEBUG, ERROR, FATAL, INFO PANIC ou WARN
	LogLevel string

	// PrometheusPushGateway Prometheus address for sending metrics
	PrometheusPushGateway string

	// RedisExpirationHours data expiration time, in hours, in Redis memory
	RedisExpirationHours int

	// RedisHost cache server hostname
	RedisHost string

	// RedisPassword cache server password
	RedisPassword string
)

func init() {
	env := env.NewEnv()

	AppPort = env.GetString("APP_PORT")
	HTTPReadTimeout = env.GetInt("HTTP_READ_TIMEOUT")
	HTTPRequestTimeout = env.GetInt("HTTP_REQUEST_TIMEOUT")
	HTTPWriteTimeout = env.GetInt("HTTP_WRITE_TIMEOUT")
	DbUser = env.GetString("DB_USER")
	DbPassword = env.GetString("DB_PASSWORD")
	DbName = env.GetString("DB_NAME")
	DbHost = env.GetString("DB_HOST")
	DbPort = env.GetInt("DB_PORT")
	DbTimeZone = env.GetString("DB_TIMEZONE")
	DbMaxIdleConns = env.GetInt("DB_MAX_IDLE_CONNS")
	DbMaxOpenConns = env.GetInt("DB_MAX_OPEN_CONNS")
	DbConnMaxLifeTime = env.GetInt("DB_CONN_MAX_LIFE_TIME")
	LogLevel = env.GetString("LOG_LEVEL")
	PrometheusPushGateway = env.GetString("PROMETHEUS_PUSH_GATEWAY")
	RedisExpirationHours = env.GetInt("REDIS_EXPIRATION_HOURS")
	RedisHost = env.GetString("REDIS_HOST")
	RedisPassword = env.GetString("REDIS_PASSWORD")
}
