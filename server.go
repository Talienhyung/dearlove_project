package main

import (
	"net/http"
	onelove "onelove/function"
)

func main() {
	onelove.Root()
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static", fs))
	http.ListenAndServe(":8080", nil)
}
