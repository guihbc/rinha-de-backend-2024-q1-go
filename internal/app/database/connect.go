package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

var conn *pgx.Conn

func Connect() {
	connection, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	log.Println("Database connected!")
	conn = connection
}

func GetConn() *pgx.Conn {
	return conn
}
