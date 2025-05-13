package main

import (
	"common"
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
)

var (
	httpAddr = common.EnvString("HTTP_ADDR", ":3000")
)

func main() {
	mux := http.NewServeMux()
	handler := NewHandler()
	handler.registerRoutes(mux)

	log.Print("Starting http server at %s", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start http server")
	}

}
