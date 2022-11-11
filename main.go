package main

import (
	"html/template"
	"net/http"

	"github.com/configcat-labs/ab-testing-in-go-sample/ampli"
	"github.com/configcat/go-sdk/v7"
)

var client = configcat.NewClient("ScDaCD8ETUuG7wYo3BdP2A/5s96HBVckk-RzI-iVf-zRA")

var user = &configcat.UserData{Email: "email1@example.com"}

var homePageTemplate = template.Must(template.ParseFiles("index.html"))
var newHomePageTemplate = template.Must(template.ParseFiles("newIndex.html"))
var storePageTemplate = template.Must(template.ParseFiles("store.html"))


func homePageHandler(w http.ResponseWriter, r *http.Request) {
	isNewHomePageEnabled := client.GetBoolValue("newhomepage", false, user)

	if isNewHomePageEnabled {

		newHomePageTemplate.Execute(w, nil)
		} else {
		// Show the old home page
		homePageTemplate.Execute(w, nil)
	}

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

	fs := http.FileServer(http.Dir("assets"))

	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", homePageHandler)
	mux.HandleFunc("/store", storePageHandler)
	http.ListenAndServe(":"+port, mux)
}