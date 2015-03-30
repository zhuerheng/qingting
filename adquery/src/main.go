package main

import (
	"./query"
	"net/http"
)

func main() {
	http.HandleFunc("/add", query.Add)
	http.HandleFunc("/query", query.Query)
	http.ListenAndServe(":11111", nil)
}
