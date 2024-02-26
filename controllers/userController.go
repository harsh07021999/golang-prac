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

func GetEmail(res http.ResponseWriter, req *http.Request) {

	if req.Method != "GET" {
		http.Error(res, "Method not supported", http.StatusBadRequest)
		return
	}

	tokenString := req.Header.Get("token")

	if err := filter.VerifyToken(tokenString); err != nil {
		util.Error.Println(res, "Token verification filed as :- ", http.StatusUnauthorized)
		res.Write([]byte(err.Error()))
		return
	}
	var name = req.URL.Query().Get("name")

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
