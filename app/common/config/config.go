// Package config parses environment variable into usable structs
package config

import (
	"fmt"

	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
	pkgerrs "github.com/pkg/errors"
)

// Config defines configuration to create a client.
type Config struct {
	Env    string `env:"ENV"`
	Server Server
	JWT    JWT

	// Dependencies section
	PostgreSQL PostgreSQL
	RabbitMQ   RabbitMQ
}

// Server holds configuration for server
type Server struct {
	BootPath string `env:"SERVER_BOOT_PATH"`
}

// JWT struct is config to JWT
type JWT struct {
	Secret string `env:"JWT_SECRET"`
}

// PostgreSQL struct is config to PostgreSQL client
type PostgreSQL struct {
	Host            string `env:"POSTGRESQL_HOST,default=localhost"`
	Port            string `env:"POSTGRESQL_PORT,default=5432"`
	User            string `env:"POSTGRESQL_USER"`
	Password        string `env:"POSTGRESQL_PASSWORD"`
	DBName          string `env:"POSTGRESQL_DBNAME"`
	MaxOpenConns    string `env:"POSTGRESQL_MAX_OPEN_CONNS,default=5"`
	MaxConnLifetime string `env:"POSTGRESQL_MAX_CONN_LIFETIME,default=10m"`
	MaxIdleLifetime string `env:"POSTGRESQL_MAX_IDLE_LIFETIME,default=5m"`
}

// RabbitMQ struct is config to RabbitMQ client
type RabbitMQ struct {
	Host     string `env:"RABBITMQ_HOST,default=localhost"`
	Port     string `env:"RABBITMQ_PORT,default=5672"`
	User     string `env:"RABBITMQ_USER"`
	Password string `env:"RABBITMQ_PASSWORD"`
}

// Load loads and creates an instance of Config from given env file path or existing env variables.
func Load(envpath string) (Config, error) {
	// just skip loading env files if it is not exists, env files only used in local dev
	_ = godotenv.Load(envpath)
	return decodeConfig()
}

// Overload loads and creates an instance of Config from given env file path or existing env variables.
// This WILL OVERRIDE an env variable that already exists - consider the .env file to forcefully set all vars.
func Overload(envpath string) (Config, error) {
	// just skip loading env files if it is not exists, env files only used in local dev
	_ = godotenv.Overload(envpath)
	return decodeConfig()
}

func decodeConfig() (Config, error) {
	var config Config
	if err := envdecode.Decode(&config); err != nil {
		return Config{}, pkgerrs.WithStack(err)
	}
	return config, nil
}

// IsProduction returns true if the environment is production
func (cfg Config) IsProduction() bool {
	return cfg.Env == "production"
}

// ConnectionString return database connection string
func (pg PostgreSQL) ConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable pool_max_conns=%s pool_max_conn_lifetime=%s pool_max_conn_idle_time=%s",
		pg.Host,
		pg.Port,
		pg.User,
		pg.Password,
		pg.DBName,
		pg.MaxOpenConns,
		pg.MaxConnLifetime,
		pg.MaxIdleLifetime,
	)
}

// ConnectionString return rabbitmq connection string
func (rb RabbitMQ) ConnectionString() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s",
		rb.User,
		rb.Password,
		rb.Host,
		rb.Port,
	)
}
