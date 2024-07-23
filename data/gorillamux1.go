// Basic Routing with Gorilla Mux

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new router instance
	r := mux.NewRouter()

	// Define a handler function
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Gorilla Mux!")
	}

	// Register the handler function for the "/" route
	r.HandleFunc("/", handler)

	// Start the HTTP server
	http.ListenAndServe(":8080", r)
}





// Handling Route Parameters with Gorilla Mux

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new router instance
	r := mux.NewRouter()

	// Define a handler function for "/hello/{name}"
	handler := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		fmt.Fprintf(w, "Hello, %s!", name)
	}

	// Register the handler function for the "/hello/{name}" route
	r.HandleFunc("/hello/{name}", handler)

	// Start the HTTP server
	http.ListenAndServe(":8080", r)
}





// Example 1: Serving Static Files with Gorilla Mux

package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Serve static files from the "static" directory
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":8080", r)
}





// Example 2: Using Middleware with Gorilla Mux

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing middleware...")
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()

	// Apply middleware to all routes
	r.Use(middleware)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Gorilla Mux with Middleware!")
	})

	http.ListenAndServe(":8080", r)
}





// Example 3: Using Subrouters with Gorilla Mux

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Subrouter for API routes
	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "API Endpoint: List of Users")
	})

	api.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "API Endpoint: List of Products")
	})

	http.ListenAndServe(":8080", r)
}





// Example 4: Handling Query Parameters with Gorilla Mux

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		queryParams := r.URL.Query()
		searchTerm := queryParams.Get("q")
		fmt.Fprintf(w, "Search Query: %s", searchTerm)
	})

	http.ListenAndServe(":8080", r)
}





// Example 5: Using Route Constraints with Gorilla Mux

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/user/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		fmt.Fprintf(w, "User ID: %s", id)
	})

	http.ListenAndServe(":8080", r)
}





// Example 6: Custom 404 Handler with Gorilla Mux

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Custom 404 handler
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Custom 404 Not Found")
	})

	http.ListenAndServe(":8080", r)
}





// Example 7: Named Routes with Gorilla Mux

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/articles/{category}/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		category := vars["category"]
		id := vars["id"]
		fmt.Fprintf(w, "Category: %s, ID: %s", category, id)
	}).Methods("GET").Name("articleRoute")

	// Generating a URL using the named route
	url, err := r.Get("articleRoute").URL("category", "technology", "id", "42")
	if err == nil {
		fmt.Println("Generated URL:", url.String())
	}

	http.ListenAndServe(":8080", r)
}





// Example 8: Route Prefixes with Gorilla Mux

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Subrouter for all routes prefixed with "/admin"
	admin := r.PathPrefix("/admin").Subrouter()

	admin.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Admin Dashboard")
	})

	admin.HandleFunc("/settings", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Admin Settings")
	})

	http.ListenAndServe(":8080", r)
}





// Example 9: Custom Middleware for Logging with Gorilla Mux

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		fmt.Printf("[%s] %s %s\n", r.Method, r.RequestURI, time.Since(start))
	})
}

func main() {
	r := mux.NewRouter()

	// Apply logging middleware
	r.Use(loggingMiddleware)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Gorilla Mux with Logging Middleware!")
	})

	http.ListenAndServe(":8080", r)
}





// Example 10: Route Middleware for Authentication with Gorilla Mux

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate authentication check (e.g., JWT verification)
		token := r.Header.Get("Authorization")
		if token != "valid_token" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()

	// Apply authentication middleware to specific route
	r.HandleFunc("/protected", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Authenticated Route")
	}).Methods("GET").Name("protectedRoute").Middleware(authMiddleware)

	http.ListenAndServe(":8080", r)
}





// Example 11: Custom Error Handling with Gorilla Mux

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func errorHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Custom Error Message", http.StatusInternalServerError)
}

func main() {
	r := mux.NewRouter()

	// Register a handler for an error endpoint
	r.HandleFunc("/error", errorHandler)

	http.ListenAndServe(":8080", r)
}





// Example 12: Route Groups with Middleware using Gorilla Mux

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Logging Middleware: ", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()

	// Group of routes with shared middleware
	api := r.PathPrefix("/api").Subrouter()
	api.Use(loggingMiddleware)

	api.HandleFunc("/endpoint1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "API Endpoint 1")
	})

	api.HandleFunc("/endpoint2", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "API Endpoint 2")
	})

	http.ListenAndServe(":8080", r)
}





// Example 13: Handling HTTP Methods with Gorilla Mux

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Route with specific HTTP method
	r.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			fmt.Fprintf(w, "Form Submitted via POST")
		} else {
			fmt.Fprintf(w, "Method not allowed")
		}
	}).Methods(http.MethodGet, http.MethodPost)

	http.ListenAndServe(":8080", r)
}





// Example 14: Serving Single Page Applications (SPA) with Gorilla Mux

package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Serve SPA from the "build" directory
	spa := http.FileServer(http.Dir("build"))
	r.PathPrefix("/").Handler(spa)

	// Handle 404 errors by serving the SPA
	r.NotFoundHandler = spa

	http.ListenAndServe(":8080", r)
}





// Example 15: Custom Middleware for Structured Logging with Gorilla Mux

package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Pass request to the next handler
		next.ServeHTTP(w, r)

		// Log request details
		log.Printf("[%s] %s %s %v", r.Method, r.RequestURI, r.RemoteAddr, time.Since(start))
	})
}

func main() {
	r := mux.NewRouter()

	// Apply logging middleware to all routes
	r.Use(loggingMiddleware)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Gorilla Mux with Structured Logging!"))
	})

	http.ListenAndServe(":8080", r)
}





// Example 16: URL Path Prefix and Stripping with Gorilla Mux

package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Define a route with a path prefix
	r.HandleFunc("/files/{filepath:.+}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		filepath := vars["filepath"]
		fmt.Fprintf(w, "File Path: %s", filepath)
	}).Methods("GET")

	// Serve static files from the "files" directory
	r.PathPrefix("/files/").Handler(http.StripPrefix("/files/", http.FileServer(http.Dir("files"))))

	http.ListenAndServe(":8080", r)
}





// Example 17: Middleware Chain with Gorilla Mux

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func middleware1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing middleware 1...")
		next.ServeHTTP(w, r)
	})
}

func middleware2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing middleware 2...")
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()

	// Chain multiple middlewares
	r.Use(middleware1)
	r.Use(middleware2)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Gorilla Mux with Middleware Chain!")
	})

	http.ListenAndServe(":8080", r)
}





// Example 18: Subrouter with Middleware Chain using Gorilla Mux

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Logging Middleware: ", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate authentication check
		token := r.Header.Get("Authorization")
		if token != "valid_token" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()

	// Subrouter with middleware chain
	api := r.PathPrefix("/api").Subrouter()
	api.Use(loggingMiddleware)
	api.Use(authMiddleware)

	api.HandleFunc("/resource", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Accessed Protected Resource")
	})

	http.ListenAndServe(":8080", r)
}





// Example 19: Custom Middleware for CORS with Gorilla Mux

package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow CORS from any origin with certain headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()

	// Apply CORS middleware
	r.Use(corsMiddleware)

	r.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Data endpoint with CORS enabled"))
	}).Methods("GET")

	http.ListenAndServe(":8080", r)
}





// Example 20: Middleware to Set Response Headers with Gorilla Mux

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func setHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set custom response headers
		w.Header().Set("X-Server-Version", "1.0")
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()

	// Apply headers middleware
	r.Use(setHeadersMiddleware)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Gorilla Mux with Custom Headers!")
	})

	http.ListenAndServe(":8080", r)
}





// Example 21: Custom Error Handling with Logging using Gorilla Mux

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func errorHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Custom Error Message", http.StatusInternalServerError)
	log.Printf("Error: %s %s", r.Method, r.URL.Path)
}

func main() {
	r := mux.NewRouter()

	// Register a handler for an error endpoint
	r.HandleFunc("/error", errorHandler)

	http.ListenAndServe(":8080", r)
}





// Example 22: Route Prefix with Authentication using Gorilla Mux

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate authentication check
		token := r.Header.Get("Authorization")
		if token != "valid_token" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()

	// Protected routes under "/secure" path
	secure := r.PathPrefix("/secure").Subrouter()
	secure.Use(authMiddleware)

	secure.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Accessed Protected Profile")
	})

	secure.HandleFunc("/settings", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Accessed Protected Settings")
	})

	http.ListenAndServe(":8080", r)
}





// Example 23: URL Path with Regular Expressions using Gorilla Mux

package main

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Route with regular expression path constraint
	r.HandleFunc("/product/{category:[a-z]+}/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		category := vars["category"]
		id := vars["id"]
		fmt.Fprintf(w, "Category: %s, ID: %s", category, id)
	}).Methods("GET")

	http.ListenAndServe(":8080", r)
}





// Example 24: Request Logging with Context using Gorilla Mux

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)

		// Pass request to the next handler
		next.ServeHTTP(w, r)

		log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
	})
}

func main() {
	r := mux.NewRouter()

	// Apply logging middleware to all routes
	r.Use(loggingMiddleware)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Gorilla Mux with Request Logging!")
	})

	http.ListenAndServe(":8080", r)
}





// Example 25: Handling JSON Requests and Responses with Gorilla Mux

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func main() {
	r := mux.NewRouter()

	// JSON response example
	r.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userID := vars["id"]

		// Simulate fetching user data from a database
		user := User{
			ID:       1,
			Username: "john_doe",
			Email:    "john.doe@example.com",
		}

		// Convert user struct to JSON
		jsonData, err := json.Marshal(user)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Set content type header and write JSON response
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}).Methods("GET")

	http.ListenAndServe(":8080", r)
}





// Example 26: Route Variables and Path Encoding with Gorilla Mux

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Route with URL variables and path encoding
	r.HandleFunc("/article/{title}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		fmt.Fprintf(w, "Article Title: %s", title)
	}).Methods("GET")

	http.ListenAndServe(":8080", r)
}





// Example 27: Request Logging with Response Time using Gorilla Mux

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)

		// Pass request to the next handler
		next.ServeHTTP(w, r)

		log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
	})
}

func main() {
	r := mux.NewRouter()

	// Apply logging middleware to all routes
	r.Use(loggingMiddleware)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Gorilla Mux with Request Logging and Response Time!")
	})

	http.ListenAndServe(":8080", r)
}





// Example 28: Handling Form Data with Gorilla Mux

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Handle form submission
	r.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		username := r.Form.Get("username")
		password := r.Form.Get("password")

		fmt.Fprintf(w, "Submitted Form Data: Username - %s, Password - %s", username, password)
	}).Methods("POST")

	http.ListenAndServe(":8080", r)
}






// Example 29: IP Filtering Middleware with Gorilla Mux

package main

import (
	"fmt"
	"net"
	"net/http"

	"github.com/gorilla/mux"
)

func ipFilterMiddleware(next http.Handler) http.Handler {
	allowedIPs := []string{"127.0.0.1", "::1"} // Allowed IPs

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Check if IP is in the allowed list
		allowed := false
		for _, allowedIP := range allowedIPs {
			if ip == allowedIP {
				allowed = true
				break
			}
		}

		if !allowed {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()

	// Apply IP filtering middleware to all routes
	r.Use(ipFilterMiddleware)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Gorilla Mux with IP Filtering Middleware!")
	})

	http.ListenAndServe(":8080", r)
}





// Example 30: Handling File Uploads with Gorilla Mux

package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	// Parse multipart form data
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Get file from form data
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Create a new file on the server
	f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer f.Close()

	// Copy file content from uploaded file to the new file
	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File uploaded successfully: %s", handler.Filename)
}

func main() {
	r := mux.NewRouter()

	// Handle file upload
	r.HandleFunc("/upload", uploadFileHandler).Methods("POST")

	http.ListenAndServe(":8080", r)
}





// Example 31: URL Path with Query Parameters using Gorilla Mux

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Route with URL path and query parameters
	r.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		queryParams := r.URL.Query()
		searchTerm := queryParams.Get("q")
		fmt.Fprintf(w, "Search Query: %s", searchTerm)
	}).Methods("GET")

	http.ListenAndServe(":8080", r)
}





// Example 32: Subrouter with Custom Middleware using Gorilla Mux

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Logging Middleware: ", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()

	// Subrouter with custom middleware
	api := r.PathPrefix("/api").Subrouter()
	api.Use(loggingMiddleware)

	api.HandleFunc("/endpoint1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "API Endpoint 1")
	})

	api.HandleFunc("/endpoint2", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "API Endpoint 2")
	})

	http.ListenAndServe(":8080", r)
}





// Example 33: Route with Custom Handler Function using Gorilla Mux

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func customHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Custom Handler Function")
}

func main() {
	r := mux.NewRouter()

	// Route with a custom handler function
	r.HandleFunc("/custom", customHandler)

	http.ListenAndServe(":8080", r)
}





// Example 34: Rate Limiting Middleware with Gorilla Mux

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/time/rate"
)

func rateLimitMiddleware(next http.Handler) http.Handler {
	limiter := rate.NewLimiter(rate.Every(time.Second), 10) // 10 requests per second

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, "Rate Limit Exceeded", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()

	// Apply rate limiting middleware to all routes
	r.Use(rateLimitMiddleware)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Gorilla Mux with Rate Limiting Middleware!")
	})

	http.ListenAndServe(":8080", r)
}





// Example 35: CSRF Protection Middleware with Gorilla Mux

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// CSRF protection middleware
	csrfMiddleware := csrf.Protect([]byte("32-byte-long-auth-key"))

	// Apply CSRF middleware to all routes
	r.Use(csrfMiddleware)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Generate CSRF token
		token := csrf.Token(r)
		fmt.Fprintf(w, "CSRF Token: %s", token)
	})

	http.ListenAndServe(":8080", r)
}





// Example 36: Route with Custom Response Headers using Gorilla Mux

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func customHeadersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Custom-Header", "Value")
	fmt.Fprintf(w, "Response with Custom Headers")
}

func main() {
	r := mux.NewRouter()

	// Route with custom response headers
	r.HandleFunc("/custom-headers", customHeadersHandler)

	http.ListenAndServe(":8080", r)
}





// Example 37: Authentication Middleware with JWT using Gorilla Mux

package main

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var jwtKey = []byte("secret_key")

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate JWT token validation
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Check token signing method etc.
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()

	// Apply authentication middleware to specific route
	r.HandleFunc("/protected", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Authenticated Route")
	}).Methods("GET").Name("protectedRoute").Middleware(authMiddleware)

	http.ListenAndServe(":8080", r)
}





// Example 38: Handling Multiple Parameters in URL Path using Gorilla Mux

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Route with multiple parameters in URL path
	r.HandleFunc("/product/{category}/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		category := vars["category"]
		id := vars["id"]
		fmt.Fprintf(w, "Category: %s, ID: %s", category, id)
	}).Methods("GET")

	http.ListenAndServe(":8080", r)
}





// Example 39: Serving Gzip Compressed Content with Gorilla Mux

package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Serve gzip compressed content
	r.HandleFunc("/compressed", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Encoding", "gzip")
		http.ServeFile(w, r, "example.txt.gz")
	})

	http.ListenAndServe(":8080", r)
}





// Example 40: Metrics and Monitoring Middleware with Gorilla Mux

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Histogram of request duration in seconds.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)
)

func init() {
	prometheus.MustRegister(requestDuration)
}

func metricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Pass request to the next handler
		next.ServeHTTP(w, r)

		duration := time.Since(start).Seconds()
		requestDuration.WithLabelValues(r.Method, r.URL.Path).Observe(duration)
	})
}

func main() {
	r := mux.NewRouter()

	// Metrics endpoint for Prometheus
	r.Handle("/metrics", promhttp.Handler())

	// Apply metrics middleware to all routes
	r.Use(metricsMiddleware)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Gorilla Mux with Metrics Middleware!")
	})

	http.ListenAndServe(":8080", r)
}





// Example 41: Serving Static Files with Gorilla Mux

package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Serve static files from the "static" directory
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":8080", r)
}





// Example 42: CORS Middleware with Gorilla Mux

package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()

	// CORS middleware
	corsHandler := cors.Default().Handler

	// Apply CORS middleware to all routes
	r.Use(corsHandler)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Gorilla Mux with CORS Middleware!"))
	})

	http.ListenAndServe(":8080", r)
}





// Example 43: Custom Error Handling Middleware with Gorilla Mux

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func errorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()

	// Apply custom error handling middleware to all routes
	r.Use(errorMiddleware)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Simulate a panic for demonstration
		panic("something went wrong")
	})

	http.ListenAndServe(":8080", r)
}





// Example 44: Request Timeout Handling Middleware with Gorilla Mux

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func timeoutMiddleware(timeout time.Duration) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.TimeoutHandler(next, timeout, "Request Timeout")
	}
}

func main() {
	r := mux.NewRouter()

	// Apply timeout middleware to all routes with a 5-second timeout
	r.Use(timeoutMiddleware(5 * time.Second))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(3 * time.Second) // Simulate some work
		fmt.Fprintf(w, "Request completed within timeout")
	})

	http.ListenAndServe(":8080", r)
}





// Example 45: Handling WebSockets with Gorilla Mux

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	for {
		// Echo incoming WebSocket messages
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			return
		}
		if err := conn.WriteMessage(messageType, p); err != nil {
			return
		}
	}
}

func main() {
	r := mux.NewRouter()

	// WebSocket endpoint
	r.HandleFunc("/ws", websocketHandler)

	http.ListenAndServe(":8080", r)
}





// Example 46: Custom 404 Not Found Page with Gorilla Mux

package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Custom 404 Page Not Found"))
}

func main() {
	r := mux.NewRouter()

	// Handle 404 Not Found with custom handler
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Gorilla Mux with Custom 404 Page Handling!"))
	})

	http.ListenAndServe(":8080", r)
}





// Example 47: Secure Headers Middleware with Gorilla Mux

package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/secure"
)

func secureHeadersMiddleware(next http.Handler) http.Handler {
	secureMiddleware := secure.New(secure.Options{
		FrameDeny:             true,
		ContentTypeNosniff:    true,
		BrowserXssFilter:      true,
		ContentSecurityPolicy: "default-src 'self'",
	})

	return secureMiddleware.Handler(next)
}

func main() {
	r := mux.NewRouter()

	// Apply secure headers middleware to all routes
	r.Use(secureHeadersMiddleware)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Gorilla Mux with Secure Headers Middleware!"))
	})

	http.ListenAndServe(":8080", r)
}





// Example 48: Session Management Middleware with Gorilla Mux

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("secret"))

func sessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session-name")

		// Set session values
		if session.Values["authenticated"] == nil {
			session.Values["authenticated"] = false
		}

		// Save the session
		session.Save(r, w)

		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()

	// Apply session management middleware to all routes
	r.Use(sessionMiddleware)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session-name")
		if session.Values["authenticated"].(bool) {
			fmt.Fprintf(w, "Authenticated User")
		} else {
			fmt.Fprintf(w, "Unauthenticated User")
		}
	})

	http.ListenAndServe(":8080", r)
}





// Example 49: Handling Redirects with Gorilla Mux

package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Redirect from "/old" to "/new"
	r.HandleFunc("/old", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/new", http.StatusMovedPermanently)
	})

	r.HandleFunc("/new", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Redirected from /old"))
	})

	http.ListenAndServe(":8080", r)
}





// Example 50: Integration with Negroni Middleware using Gorilla Mux

package main

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Example Negroni middleware
	n := negroni.Classic()

	// Use Negroni middleware with Gorilla Mux router
	n.UseHandler(r)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Gorilla Mux with Negroni Middleware!")
	})

	http.ListenAndServe(":8080", n)
}






//Example 51: Basic Routing with Gorilla Mux

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new router instance
	r := mux.NewRouter()

	// Define a handler function
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Gorilla Mux!")
	}

	// Register the handler function for the "/" route
	r.HandleFunc("/", handler)

	// Start the HTTP server
	http.ListenAndServe(":8080", r)
}





//Example 52: Handling Route Parameters with Gorilla Mux

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new router instance
	r := mux.NewRouter()

	// Define a handler function for "/hello/{name}"
	handler := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		fmt.Fprintf(w, "Hello, %s!", name)
	}

	// Register the handler function for the "/hello/{name}" route
	r.HandleFunc("/hello/{name}", handler)

	// Start the HTTP server
	http.ListenAndServe(":8080", r)
}
