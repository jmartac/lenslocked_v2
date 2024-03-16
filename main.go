package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact page</h1>To get in touch, please send an email to <a href=\"mailto:my@email.com\">my email</a>")
}

type Router struct{}

func (Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		homeHandler(w, req)
	case "/contact":
		contactHandler(w, req)
	default:
		http.Error(w, "Not f Found ¯\\_(ツ)_/¯", http.StatusNotFound)
	}
}

func main() {
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", &Router{})
}
