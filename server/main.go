package main

import (
	"log"
	"net/http"

	"predict-rgr/backend/api"
)

func main() {
	router := api.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}