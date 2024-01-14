package onelove

import (
	"net/http"
)

type Structure struct {
	Page string
}

// Root initializes the game structure and sets up the handlers
func Root() {
	var structu Structure
	structu.Page = "FIRST"

	// Define the root handler
	http.HandleFunc("/onelove", func(w http.ResponseWriter, r *http.Request) {
		Onelove(w, r, structu)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Home(w, r, structu)
	})
	http.HandleFunc("/a", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, &structu)
	})
}
