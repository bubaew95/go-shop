package conf

import (
	"flag"
	"os"
	"strconv"
)

type ServerConfig struct {
	Port     uint
	Database *DatabaseConfig
	Debug    bool
}

type DatabaseConfig struct {
	Driver                  string
	Dsn                     string
	ConnMaxLifetimeInMinute int
	MaxOpenConns            int
	MaxIdleConns            int
}

func NewServerConfig() *ServerConfig {
	port := flag.Uint("port", 8080, "http server port")
	flag.Parse()

	debug, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		debug = false
	}

	maxLifeInMinute, err := strconv.Atoi(os.Getenv("DATABASE_MAX_LIFE_IN_MINUTE"))
	if err != nil {
		maxLifeInMinute = 3
	}

	maxOpenConns, err := strconv.Atoi(os.Getenv("DATABASE_MAX_OPEN_CONNS"))
	if err != nil {
		maxOpenConns = 10
	}

	maxIdleConns, err := strconv.Atoi(os.Getenv("DATABASE_MAX_IDLE_CONNS"))
	if err != nil {
		maxIdleConns = 1
	}

	return &ServerConfig{
		Port: *port,
		Database: &DatabaseConfig{
			Driver:                  os.Getenv("DATABASE_DRIVER"),
			Dsn:                     os.Getenv("DATABASE_DSN"),
			ConnMaxLifetimeInMinute: maxLifeInMinute,
			MaxOpenConns:            maxOpenConns,
			MaxIdleConns:            maxIdleConns,
		},
		Debug: debug,
	}
}
