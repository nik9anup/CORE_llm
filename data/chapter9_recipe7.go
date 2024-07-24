// This code sets up an HTTP server on port 8080:
// - "/welcome" serves the content of "welcome.txt" file.
// - "/html/" serves static files from the "html" directory, stripping "/html" prefix.
package main

import (
	"net/http"
)

func main() {

	fileSrv := http.FileServer(http.Dir("html"))
	fileSrv = http.StripPrefix("/html", fileSrv)

	http.HandleFunc("/welcome", serveWelcome)
	http.Handle("/html/", fileSrv)
	http.ListenAndServe(":8080", nil)
}

func serveWelcome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "welcome.txt")
}