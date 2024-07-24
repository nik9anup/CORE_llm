// This code starts an HTTP server on port 8080 with three endpoints:
// - "/secured/handle" redirects using http.RedirectHandler to "/login".
// - "/secured/hadlefunc" redirects using http.Redirect to "/login".
// - "/login" responds with a message "Welcome user! Please login!".
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Server is starting...")

	http.Handle("/secured/handle", http.RedirectHandler("/login", http.StatusTemporaryRedirect))
	http.HandleFunc("/secured/hadlefunc", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
	})
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome user! Please login!\n")
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}