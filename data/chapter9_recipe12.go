// This code starts an HTTP server on port 8080 that responds with "Hello world!" after a delay.
// It handles graceful shutdown using OS signals (SIGINT) to stop the server and waits up to 30 seconds for connections to close.
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(10 * time.Second)
		fmt.Fprintln(w, "Hello world!")
	})

	srv := &http.Server{Addr: ":8080", Handler: mux}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("Server error: %s\n", err)
		}
	}()

	log.Println("Server listening on : " + srv.Addr)

	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)

	<-stopChan // wait for SIGINT
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(
		context.Background(),
		30*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("Server gracefully stopped")
}