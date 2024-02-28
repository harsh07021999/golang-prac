package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"urlShrtGo/dtos"
	"urlShrtGo/filter"
	"urlShrtGo/models"
	"urlShrtGo/util"
)

type LoginBO struct {
	password string
	id       int64
}

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

	var stored LoginBO

	stmt, err := models.DB.Prepare("SELECT password, id FROM users WHERE name = $1 ")
	if err != nil {
		log.Fatal("DB query failed", err)
		http.Error(res, "INternal Server Error", http.StatusInternalServerError)
	}
	defer stmt.Close()

	err = stmt.QueryRow(login.Name).Scan(&stored.password, &stored.id)
	if err != nil {
		if err == sql.ErrNoRows {
			util.Error.Println("Error :- ", err)
			http.Error(res, "Invalid username or password", http.StatusUnauthorized)
			return
		} else {
			util.Error.Println("Error :- ", err)
			http.Error(res, "Database error", http.StatusInternalServerError)
			return
		}
	}

	if stored.password != login.Password {
		http.Error(res, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	tokenString, err := filter.CreateToken(stored.id)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		util.Error.Println(err)
	}

	var loginToken dtos.LoginResponse
	loginToken.Token = tokenString
	util.Debug.Println(filter.GetTokenId(tokenString))
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		util.Error.Println(err)
	}
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
