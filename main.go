package main

import (
	"log"
	"net/http"
	"urlShrtGo/controllers"
	"urlShrtGo/models"
)

func main() {

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controllers.GetRealUrl(w, r)
		case http.MethodPost:
			controllers.PostRealUrl(w, r)
		default:
			http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		}
	}))
	mux.Handle("/ping", http.HandlerFunc(controllers.PingHandler))
	mux.Handle("/login", http.HandlerFunc(controllers.LoginHandler))
	mux.Handle("/signup", http.HandlerFunc(controllers.SignupHandler))
	mux.Handle("/user", http.HandlerFunc(controllers.GetEmail))

	// http.HandleFunc("/ping", controllers.PingHandler)
	// http.HandleFunc("/login", controllers.LoginHandler)
	// http.HandleFunc("/signup", controllers.SignupHandler)
	// http.HandleFunc("/user", controllers.GetEmail)

	// models
	models.ConnectDB()

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
