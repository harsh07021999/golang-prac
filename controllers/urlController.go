package controllers

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"urlShrtGo/dtos"
	"urlShrtGo/filter"
	"urlShrtGo/models"
	"urlShrtGo/util"
)

func GetRealUrl(res http.ResponseWriter, req *http.Request) {

	tokenString := req.Header.Get("token")

	if err := filter.VerifyToken(tokenString); err != nil {
		util.Error.Println(res, "Token verification filed as :- ", http.StatusUnauthorized)
		res.Write([]byte(err.Error()))
		return
	}

	var urlHash = req.URL.Path

	var url []string

	err := models.DB.Select(&url, "SELECT url from urlmap where hash = #1", urlHash)
	if err != nil {
		util.Error.Println("Error fetching url ", err)
		var errString = "Error fetching url " + err.Error()
		json.NewEncoder(res).Encode(errString)
		return
	}
	var responseURL = url[0]
	json.NewEncoder(res).Encode(responseURL)

}

func PostRealUrl(res http.ResponseWriter, req *http.Request) {

	tokenString := req.Header.Get("token")

	if err := filter.VerifyToken(tokenString); err != nil {
		util.Error.Println(res, "Token verification filed as :- ", http.StatusUnauthorized)
		res.Write([]byte(err.Error()))
		return
	}
	var shortReq dtos.ShortUrlRequest
	if err := json.NewDecoder(req.Body).Decode(&shortReq); err != nil {
		util.Error.Println("Json decode failed", http.StatusBadRequest)
		http.Error(res, "Json decode failed "+err.Error(), http.StatusBadRequest)
	}
	defer req.Body.Close()

	var domain = "http://localhost/"

	if !util.CheckEmptyOrBlankString(shortReq.CustomHash) {
		var oldhash []string = nil
		err := models.DB.Select(&oldhash, "SELECT hash FROM urlhash where hash = $1", shortReq.CustomHash)
		if err != nil {
			util.Error.Fatal("DB query failed", err)
			http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		}
		if oldhash == nil {
			if _, err := models.DB.NamedExec("INSERT INTO urlhash (hash, url, uid) VALUES(:name, :email, :password)", shortReq.CustomHash); err != nil {
				util.Error.Println("Database query error", err)
			}
		}
	} else {
		constHashSize := 8
		salt := make([]byte, 16)
		rand.Read(salt)
		data := append([]byte(shortReq.OriginalUrl), salt...)
		hash := sha256.New()
		hash.Write(data)
		hashedbytes := hash.Sum(nil)
		constLenHash := hex.EncodeToString(hashedbytes[:constHashSize/2])
		util.Debug.Println(constLenHash)
		domain = domain + constLenHash
	}
	json.NewEncoder(res).Encode(domain)

}
