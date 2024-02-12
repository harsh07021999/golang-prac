package main

import (
	"log"
	"net/http"
	"urlShrtGo/controllers"
	"urlShrtGo/models"
)

func main() {

	http.HandleFunc("/ping", controllers.PingHandler)
	http.HandleFunc("/login", controllers.LoginHandler)
	http.HandleFunc("/signup", controllers.SignupHandler)
	http.HandleFunc("/user", controllers.GetEmail)
	// models
	models.ConnectDB()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
