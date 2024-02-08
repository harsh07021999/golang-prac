package controllers

import (
	"fmt"
	"net/http"
)

func PingHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "pong")
}
