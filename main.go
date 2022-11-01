package main

import (
	"net/http"
	"html/template"
)

var homePageTemplate = template.Must(template.ParseFiles("index.html"))
var storePageTemplate = template.Must(template.ParseFiles("store.html"))

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	homePageTemplate.Execute(w, nil)
}

func storePageHandler(w http.ResponseWriter, r *http.Request) {
	storePageTemplate.Execute(w, nil)
}

func main() {
	port := "3000"

	mux := http.NewServeMux()

	mux.HandleFunc("/", homePageHandler)
	mux.HandleFunc("/store", storePageHandler)
	http.ListenAndServe(":"+port, mux)
}