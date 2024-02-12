package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"urlShrtGo/dtos"
	"urlShrtGo/filter"
	"urlShrtGo/models"
)

func LoginHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

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

	tokenString, err := filter.CreateToken(login.Name)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}
	var loginToken dtos.LoginResponse
	loginToken.Token = tokenString
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(loginToken)
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

	fmt.Println(signUp.Name)
	fmt.Println(signUp.Email)
	fmt.Println(signUp.Password)

	if _, err := models.DB.NamedExec("INSERT INTO users (name, email, password) VALUES(:name, :email, :password)", signUp); err != nil {
		log.Fatal("Database query error", err)
	}
}
