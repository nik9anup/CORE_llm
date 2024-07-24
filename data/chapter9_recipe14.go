// This code defines an HTTP server that responds with "Hello world" and demonstrates handling form data:
// - It logs the request form data before and after calling req.ParseForm().
// - It prints the value of "param1" from the parsed form.
// - It serves "Hello world" as the response to any incoming request.
package main

import (
	"fmt"
	"net/http"
)

type StringServer string

func (s StringServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf("Prior ParseForm: %v\n", req.Form)
	req.ParseForm()
	fmt.Printf("Post ParseForm: %v\n", req.Form)
	fmt.Println("Param1 is : " + req.Form.Get("param1"))
	fmt.Printf("PostForm : %v\n", req.PostForm)
	rw.Write([]byte(string(s)))
}

func createServer(addr string) http.Server {
	return http.Server{
		Addr:    addr,
		Handler: StringServer("Hello world"),
	}
}

func main() {
	s := createServer(":8080")
	fmt.Println("Server is starting...")
	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}