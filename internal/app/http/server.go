package http

import (
	"log"
	"os"

	"github.com/valyala/fasthttp"
)

func RunServer() {
	listen := os.Getenv("LISTEN_HTTP_PORT")
	log.Printf("Listening on %s", listen)
	log.Fatal(fasthttp.ListenAndServe(listen, getRouter().Handler))
}
