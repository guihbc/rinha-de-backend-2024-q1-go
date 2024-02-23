package main

import (
	"github.com/guihbc/rinha-de-backend-2024-q1/internal/app/database"
	"github.com/guihbc/rinha-de-backend-2024-q1/internal/app/http"
)

func main() {
	database.Connect()
	http.RunServer()
}
