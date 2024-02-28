package controllers

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strings"
	"urlShrtGo/dtos"
	"urlShrtGo/filter"
	"urlShrtGo/models"
	"urlShrtGo/util"
)

type HashReq struct {
	Hash string `db:"hash"`
	URL  string `db:"url"`
	UID  int64  `db:"uid"`
}

func GetRealUrl(res http.ResponseWriter, req *http.Request) {

	tokenString := req.Header.Get("token")

	if err := filter.VerifyToken(tokenString); err != nil {
		util.Error.Println(res, "Token verification filed as :- ", http.StatusUnauthorized)
		res.Write([]byte(err.Error()))
		return
	}

	var urlHash = strings.Split(req.URL.Path, "/")[1]

	var url []string

	err := models.DB.Select(&url, "SELECT url from urlhash where hash = $1", urlHash)
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

	id, err := filter.GetTokenId(tokenString)
	if err != nil {
		util.Error.Println(err.Error())
	}

	if !util.CheckEmptyOrBlankString(shortReq.CustomHash) {
		var oldhash []string
		err := models.DB.Select(&oldhash, "SELECT hash FROM urlhash where hash = $1", shortReq.CustomHash)
		if err != nil {
			util.Error.Fatal("DB query failed", err)
			http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		}
		if oldhash == nil {
			var datas HashReq
			datas.UID = int64(id)
			datas.URL = shortReq.OriginalUrl
			datas.Hash = shortReq.CustomHash

			if _, err := models.DB.NamedExec("INSERT INTO urlhash (hash, url, uid) VALUES(:hash, :url, :uid)", &datas); err != nil {
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

		var datas HashReq
		datas.UID = int64(id)
		datas.URL = shortReq.OriginalUrl
		datas.Hash = constLenHash

		if _, err := models.DB.NamedExec("INSERT INTO urlhash (hash, url, uid) VALUES(:hash, :url, :uid)", &datas); err != nil {
			util.Error.Println("Database query error", err)
		}
		json.NewEncoder(res).Encode(datas.Hash)
	}

}
