package main

import (
	"net/http"
	"html/template"
	"github.com/configcat-labs/ab-testing-in-go-sample/ampli"
)

var homePageTemplate = template.Must(template.ParseFiles("index.html"))
var storePageTemplate = template.Must(template.ParseFiles("store.html"))

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	homePageTemplate.Execute(w, nil)
}

func storePageHandler(w http.ResponseWriter, r *http.Request) {
	ampli.Instance.PageView("user@example.com")
	storePageTemplate.Execute(w, nil)
}

func main() {
	ampli.Instance.Load(ampli.LoadOptions{
    Environment: ampli.EnvironmentDevelopment,
		Client: ampli.LoadClientOptions{
			APIKey: "f0f6506c41bf8cc3a341cd4febf17293",
		},
})

	port := "3000"

	mux := http.NewServeMux()

	mux.HandleFunc("/", homePageHandler)
	mux.HandleFunc("/store", storePageHandler)
	http.ListenAndServe(":"+port, mux)
}