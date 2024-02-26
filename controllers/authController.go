package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"urlShrtGo/dtos"
	"urlShrtGo/filter"
	"urlShrtGo/models"
	"urlShrtGo/util"
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
		util.Error.Println("Json decode error", err)
	}
	defer req.Body.Close()

	tokenString, err := filter.CreateToken(login.Name)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		util.Error.Println(err)
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

	util.Debug.Println(signUp.Name)
	util.Debug.Println(signUp.Email)
	util.Debug.Println(signUp.Password)

	if _, err := models.DB.NamedExec("INSERT INTO users (name, email, password) VALUES(:name, :email, :password)", signUp); err != nil {
		log.Fatal("Database query error", err)
	}
}
