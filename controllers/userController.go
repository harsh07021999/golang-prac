package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"urlShrtGo/dtos"
	"urlShrtGo/filter"
	"urlShrtGo/models"
)

func GetEmail(res http.ResponseWriter, req *http.Request) {

	if req.Method != "GET" {
		http.Error(res, "Method not supported", http.StatusBadRequest)
		return
	}

	tokenString := req.Header.Get("token")

	if err := filter.VerifyToken(tokenString); err != nil {
		log.Fatal("Token verification filed", err)
		http.Error(res, "Token verification filed", http.StatusUnauthorized)
	}

	// var ureq dtos.UserRequest
	var name = req.URL.Query().Get("name")
	// decoder := json.NewDecoder(req.Body)
	// if err := decoder.Decode(&ureq); err != nil {
	// 	log.Fatal("Json decode failed", err)
	// 	http.Error(res, "Invalid JSON format", http.StatusBadRequest)
	// }
	// defer req.Body.Close()

	var email []string
	err := models.DB.Select(&email, "SELECT email FROM users WHERE name = $1", name)
	if err != nil {
		log.Fatal("DB query failed", err)
		http.Error(res, "INternal Server Error", http.StatusInternalServerError)
	}
	var response dtos.UserResponse
	response.Email = email[0]
	json.NewEncoder(res).Encode(response)
}
