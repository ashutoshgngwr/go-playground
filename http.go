package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
)

var name string

func requestHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintln(w, "Your name is", name)
	case http.MethodPost:
		content, _ := ioutil.ReadAll(r.Body)
		name = string(content)
		fmt.Fprintln(w, "OK")
	}
}

func main() {
	http.HandleFunc("/name", requestHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
