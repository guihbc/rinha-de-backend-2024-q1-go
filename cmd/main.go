package main

import (
	"log"
	"os"

	"github.com/guihbc/rinha-de-backend-2024-q1/internal/app/database"
	"github.com/guihbc/rinha-de-backend-2024-q1/internal/app/http"
	"github.com/valyala/fasthttp"
)

func main() {
	database.Connect()

	listen := os.Getenv("LISTEN_HTTP_PORT")
	log.Printf("Listening on %s", listen)
	log.Fatal(fasthttp.ListenAndServe(listen, http.GetRouter().Handler))
}
