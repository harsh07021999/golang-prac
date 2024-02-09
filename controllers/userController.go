package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"urlShrtGo/dtos"
	"urlShrtGo/models"
)

func LoginHandler(res http.ResponseWriter, req *http.Request) {

	if req.Method != "POST" {
		http.Error(res, "Method not supported", http.StatusBadRequest)
		return
	}
	var login dtos.LoginRequest
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&login); err != nil {
		log.Fatal("Json decode error", err)
	}
	defer req.Body.Close()

	fmt.Println(login.Name, login.Password)
}

func SignupHandler(res http.ResponseWriter, req *http.Request) {

	if req.Method != "POST" {
		http.Error(res, "Method not supported", http.StatusBadRequest)
		return
	}
	var signUp dtos.SignUpRequest
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&signUp); err != nil {
		log.Fatal("Json decode error", err)
	}
	defer req.Body.Close()

	fmt.Println(signUp.Name, signUp.Email, signUp.Password)

	if _, err := models.DB.Query("INSERT INTO users (name, email, password) VALUES(:name, :email, :password)"); err != nil {
		log.Fatal("Database query error", err)
	}
}
