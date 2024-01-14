package onelove

import (
	"log"
	"net/http"
	"text/template"
)

// Home handles HTTP requests for the home page and renders the appropriate HTML templates
func Onelove(w http.ResponseWriter, r *http.Request, infos Structure) {
	template, err := template.ParseFiles(
		"./pages/onelove.html",
	)
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, infos)
}

// Home handles HTTP requests for the home page and renders the appropriate HTML templates
func Home(w http.ResponseWriter, r *http.Request, infos Structure) {
	template, err := template.ParseFiles(
		"./pages/index.html",
	)
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, infos)
}

func handler(w http.ResponseWriter, r *http.Request, infos *Structure) {
	truc := r.FormValue("name")
	if truc == "there" {
		http.Redirect(w, r, "/onelove", http.StatusSeeOther)
	} else {
		infos.Page = "MAIN"
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
