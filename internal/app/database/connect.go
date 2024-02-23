package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var conn *pgxpool.Pool

func Connect() {
	maxWaitTime := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), maxWaitTime)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			log.Fatalf("It was not possible to connect to the database within the time limit %s", maxWaitTime)
		default:
			connectionPool, err := pgxpool.NewWithConfig(context.Background(), config())
			if err == nil {
				log.Println("Database config parsed with success!")

				err = connectionPool.Ping(context.Background())
				if err == nil {
					log.Println("Database connected!")
					conn = connectionPool
					return
				}
			}

			fmt.Println("Error connecting to the database", err)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func GetConn() *pgxpool.Pool {
	return conn
}
