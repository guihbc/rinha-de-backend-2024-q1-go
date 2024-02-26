package database

import (
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func config() *pgxpool.Config {
	databaseConfig, err := pgxpool.ParseConfig(os.Getenv("DB_URL"))
	if err != nil {
		log.Fatalln("Failed to parse pool config")
	}

	databaseConfig.MaxConns = int32(256)
	databaseConfig.MinConns = int32(10)
	databaseConfig.MaxConnLifetime = time.Hour
	databaseConfig.MaxConnIdleTime = time.Minute
	databaseConfig.HealthCheckPeriod = time.Second * 30
	databaseConfig.ConnConfig.ConnectTimeout = time.Second * 5

	return databaseConfig
}
