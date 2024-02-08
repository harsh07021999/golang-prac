package main

import (
	"fmt"
	"log"
	"net/http"
	"urlShrtGo/models"
)

func pingHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "pong")
}

func main() {

	http.HandleFunc("/ping", pingHandler)
	// models
	models.ConnectDB()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
