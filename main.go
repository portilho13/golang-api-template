package main

import (
	"log"
	"net/http"

	"github.com/portilho13/golang-api-template/api"
)

const IP string = "127.0.0.1:1234"


func main() {
	mux := api.InitializeRoutes();

	if err := http.ListenAndServe(IP, mux); err != nil {
		log.Fatal(err)
	}
}