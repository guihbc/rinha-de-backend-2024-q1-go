package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func config() *pgxpool.Config {
	databaseConfig, err := pgxpool.ParseConfig(os.Getenv("DB_URL"))
	if err != nil {
		log.Fatalln("Failed to parse pool config")
	}

	databaseConfig.MaxConns = int32(156)
	databaseConfig.MinConns = int32(10)
	databaseConfig.MaxConnLifetime = time.Hour
	databaseConfig.MaxConnIdleTime = time.Minute
	databaseConfig.HealthCheckPeriod = time.Second * 30
	databaseConfig.ConnConfig.ConnectTimeout = time.Second * 5

	databaseConfig.BeforeAcquire = func(ctx context.Context, c *pgx.Conn) bool {
		log.Println("Before Acquiring the connection pool")
		return true
	}

	databaseConfig.AfterRelease = func(c *pgx.Conn) bool {
		log.Println("After releasing the connection pool")
		return true
	}

	databaseConfig.BeforeClose = func(c *pgx.Conn) {
		log.Println("Closed the connection")
	}

	return databaseConfig
}
