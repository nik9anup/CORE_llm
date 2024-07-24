// This code starts an HTTP server on port 8080 that renders a template file "template.tpl"
// and serves it when accessing the root ("/") endpoint.
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	fmt.Println("Server is starting...")
	tpl, err := template.ParseFiles("template.tpl")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tpl.Execute(w, "John Doe")
		if err != nil {
			panic(err)
		}
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}