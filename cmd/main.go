package main

import (
	"fmt"
	"log"
	"os"

	"github.com/guihbc/rinha-de-backend-2024-q1/internal/app/database"
	"github.com/guihbc/rinha-de-backend-2024-q1/internal/app/http"
	"github.com/valyala/fasthttp"
)

func init() {
	database.Connect()
}

func main() {
	listen := fmt.Sprintf(":%s", os.Getenv("HTTP_PORT"))
	log.Printf("Listening on %s", listen)
	log.Fatal(fasthttp.ListenAndServe(listen, http.GetRouter().Handler))
}
