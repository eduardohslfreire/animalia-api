package env

import (
	"github.com/eduardohslfreire/animalia-api/pkg/env"
)

var (
	// AppPort porta da aplicação
	AppPort string

	// HTTPReadTimeout timeout em segundos para a api (read)
	HTTPReadTimeout int

	// HTTPRequestTimeout timeout em segundos para a api (requests)
	HTTPRequestTimeout int

	// HTTPWriteTimeout timeout em segundos para a api (write)
	HTTPWriteTimeout int

	// DbUser usuário de aplicação do banco de dados
	DbUser string

	// DbPassword senha do usuário da aplicação do banco de dados
	DbPassword string

	// DbName nome do banco de dados
	DbName string

	// DbHost hostname do banco de dados
	DbHost string

	// DbPort porta do banco de dados
	DbPort int

	// DbTimeZone timezone do banco de dados
	DbTimeZone string

	// DbMaxIdleConns número máximo de conexões inativas no banco de dados
	DbMaxIdleConns int

	// DbMaxOpenConns número máximo de conexões abertas no banco de dados
	DbMaxOpenConns int

	// DbConnMaxLifeTime tempo máximo para que uma conexão inativas pode ser reutilizada em segundos
	DbConnMaxLifeTime int

	// LogLevel determina o log level a ser exibido no stdout (controle da severidade do log): DEBUG, ERROR, FATAL, INFO PANIC ou WARN
	LogLevel string

	// PrometheusPushGateway endereço do Prometheus para envio das metricas
	PrometheusPushGateway string

	// RedisExpirationHours tempo de expiração dos dados, em horas, na memória do Redis
	RedisExpirationHours int

	// RedisHost nome do host da base de dados Redis
	RedisHost string

	// RedisPassword senha da base de dados Redis
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
