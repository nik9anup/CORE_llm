// Basic routing with Chi.
package main

import (
    "net/http"

    "github.com/go-chi/chi"
)

func main() {
    r := chi.NewRouter()
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, Chi!"))
    })
    http.ListenAndServe(":8080", r)
}



// Handling route parameters with Chi.
package main

import (
    "net/http"

    "github.com/go-chi/chi"
)

func main() {
    r := chi.NewRouter()
    r.Get("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
        id := chi.URLParam(r, "id")
        w.Write([]byte("User ID: " + id))
    })
    http.ListenAndServe(":8080", r)
}




// Serving static files with Chi.
package main

import (
    "net/http"

    "github.com/go-chi/chi"
)

func main() {
    r := chi.NewRouter()
    r.Get("/*", http.FileServer(http.Dir("./static")).ServeHTTP)
    http.ListenAndServe(":8080", r)
}





// Applying middleware with Chi.
package main

import (
    "net/http"

    "github.com/go-chi/chi"
    "github.com/go-chi/chi/middleware"
)

func main() {
    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, Chi!"))
    })
    http.ListenAndServe(":8080", r)
}




// Using sub-routers with Chi.
package main

import (
    "net/http"

    "github.com/go-chi/chi"
)

func main() {
    r := chi.NewRouter()
    r.Route("/admin", func(r chi.Router) {
        r.Get("/", func(w http.ResponseWriter, r *http.Request) {
            w.Write([]byte("Admin Dashboard"))
        })
        r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
            w.Write([]byte("List of Users"))
        })
    })
    http.ListenAndServe(":8080", r)
}





// Handling different HTTP methods with Chi.
package main

import (
    "net/http"

    "github.com/go-chi/chi"
)

func main() {
    r := chi.NewRouter()
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("GET Request"))
    })
    r.Post("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("POST Request"))
    })
    r.Put("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("PUT Request"))
    })
    r.Delete("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("DELETE Request"))
    })
    http.ListenAndServe(":8080", r)
}





// Grouping routes with Chi.
package main

import (
    "net/http"

    "github.com/go-chi/chi"
)

func main() {
    r := chi.NewRouter()
    r.Route("/users", func(r chi.Router) {
        r.Get("/", func(w http.ResponseWriter, r *http.Request) {
            w.Write([]byte("List of Users"))
        })
        r.Post("/", func(w http.ResponseWriter, r *http.Request) {
            w.Write([]byte("Create User"))
        })
    })
    http.ListenAndServe(":8080", r)
}





// Validating URL parameters with Chi.
package main

import (
    "net/http"

    "github.com/go-chi/chi"
    "github.com/go-chi/chi/middleware"
    "gopkg.in/go-playground/validator.v9"
)

type User struct {
    ID   string `validate:"uuid"`
    Name string `validate:"required"`
}

func main() {
    r := chi.NewRouter()
    r.Use(middleware.ValidateURLParams(
        validator.New(),
    ))

    r.Get("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
        id := chi.URLParam(r, "id")
        w.Write([]byte("User ID: " + id))
    })

    r.Post("/user", func(w http.ResponseWriter, r *http.Request) {
        var user User
        if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte("Bad request"))
            return
        }
        validate := validator.New()
        if err := validate.Struct(user); err != nil {
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte("Validation error"))
            return
        }
        w.Write([]byte("User created"))
    })

    http.ListenAndServe(":8080", r)
}





// Implementing custom middleware with Chi.
package main

import (
    "net/http"

    "github.com/go-chi/chi"
    "github.com/go-chi/chi/middleware"
)

func MyCustomMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Custom middleware logic here
        next.ServeHTTP(w, r)
    })
}

func main() {
    r := chi.NewRouter()
    r.Use(MyCustomMiddleware)
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, Chi!"))
    })
    http.ListenAndServe(":8080", r)
}





// Implementing graceful shutdown with Chi.
package main

import (
    "context"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/go-chi/chi"
)

func main() {
    r := chi.NewRouter()
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, Chi!"))
    })

    srv := &http.Server{
        Addr:    ":8080",
        Handler: r,
    }

    go func() {
        if err := srv.ListenAndServe(); err != nil {
            if err == http.ErrServerClosed {
                return
            }
            panic(err)
        }
    }()

    // Graceful shutdown
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := srv.Shutdown(ctx); err != nil {
        panic(err)
    }
}




// Implementing HTTP/2 server push with Chi.
package main

import (
    "net/http"

    "github.com/go-chi/chi"
)

func main() {
    r := chi.NewRouter()
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        pusher, ok := w.(http.Pusher)
        if ok {
            if err := pusher.Push("/styles.css", nil); err != nil {
                http.Error(w, "Failed to push: "+err.Error(), http.StatusInternalServerError)
                return
            }
        }
        w.Write([]byte("Hello, Chi!"))
    })

    http.ListenAndServe(":8080", r)
}




// Implementing rate limiting with Chi.
package main

import (
    "net/http"

    "github.com/go-chi/chi"
    "github.com/go-chi/chi/middleware"
    "golang.org/x/time/rate"
)

func main() {
    r := chi.NewRouter()
    r.Use(middleware.NewRateLimit(
        rate.NewLimiter(rate.Limit(5), 10),
    ))

    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, Chi!"))
    })

    http.ListenAndServe(":8080", r)
}




// Implementing basic authentication with Chi.
package main

import (
    "net/http"

    "github.com/go-chi/chi"
    "github.com/go-chi/chi/middleware"
)

func main() {
    r := chi.NewRouter()
    r.Use(middleware.BasicAuth("user", "password"))

    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Authenticated"))
    })

    http.ListenAndServe(":8080", r)
}




// Implementing JWT authentication with Chi.
package main

import (
    "net/http"

    "github.com/go-chi/chi"
    "github.com/go-chi/jwtauth"
)

func main() {
    r := chi.NewRouter()

    jwt := jwtauth.New("HS256", []byte("secret"), nil)

    r.Group(func(r chi.Router) {
        r.Use(jwtauth.Verifier(jwt))
        r.Use(jwtauth.Authenticator)

        r.Get("/protected", func(w http.ResponseWriter, r *http.Request) {
            w.Write([]byte("Authenticated"))
        })
    })

    http.ListenAndServe(":8080", r)
}




// Handling WebSocket connections with Chi.
package main

import (
    "net/http"

    "github.com/go-chi/chi"
    "github.com/go-chi/chi/middleware"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func main() {
    r := chi.NewRouter()
    r.Use(middleware.Logger)

    r.Get("/ws", func(w http.ResponseWriter, r *http.Request) {
        conn, err := upgrader.Upgrade(w, r, nil)
        if err != nil {
            http.Error(w, "Could not open WebSocket connection", http.StatusBadRequest)
            return
        }
        defer conn.Close()

        for {
            _, _, err := conn.ReadMessage()
            if err != nil {
                return
            }
        }
    })

    http.ListenAndServe(":8080", r)
}





// Handling WebSocket connections with Chi.
package main

import (
    "net/http"

    "github.com/go-chi/chi"
    "github.com/go-chi/chi/middleware"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func main() {
    r := chi.NewRouter()
    r.Use(middleware.Logger)

    r.Get("/ws", func(w http.ResponseWriter, r *http.Request) {
        conn, err := upgrader.Upgrade(w, r, nil)
        if err != nil {
            http.Error(w, "Could not open WebSocket connection", http.StatusBadRequest)
            return
        }
        defer conn.Close()

        for {
            _, _, err := conn.ReadMessage()
            if err != nil {
                return
            }
        }
    })

    http.ListenAndServe(":8080", r)
}




// Handling request context with Chi.
package main

import (
    "context"
    "net/http"

    "github.com/go-chi/chi"
)

func main() {
    r := chi.NewRouter()

    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        ctx := r.Context()
        ctx = context.WithValue(ctx, "key", "value")
        // Use ctx in subsequent handlers
        w.Write([]byte("Context handled"))
    })

    http.ListenAndServe(":8080", r)
}




// Implementing a custom 404 handler with Chi.
package main

import (
    "net/http"

    "github.com/go-chi/chi"
)

func main() {
    r := chi.NewRouter()

    r.NotFound(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte("Custom 404 Not Found"))
    })

    http.ListenAndServe(":8080", r)
}




// Implementing context cancellation with Chi.
package main

import (
    "context"
    "net/http"
    "time"

    "github.com/go-chi/chi"
)

func main() {
    r := chi.NewRouter()

    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
        defer cancel()

        // Perform operations with ctx
        w.Write([]byte("Context cancelled after 5 seconds"))
    })

    http.ListenAndServe(":8080", r)
}




// Using multiple middlewares with Chi.
package main

import (
    "net/http"

    "github.com/go-chi/chi"
    "github.com/go-chi/chi/middleware"
)

func main() {
    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)

    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, Chi!"))
    })

    http.ListenAndServe(":8080", r)
}




// Handling request headers with Chi.
package main

import (
    "net/http"

    "github.com/go-chi/chi"
)

func main() {
    r := chi.NewRouter()

    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        headerValue := r.Header.Get("Header-Name")
        w.Write([]byte("Header Value: " + headerValue))
    })

    http.ListenAndServe(":8080", r)
}




// Setting response headers with Chi.
package main

import (
    "net/http"

    "github.com/go-chi/chi"
)

func main() {
    r := chi.NewRouter()

    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Custom-Header", "value")
        w.Write([]byte("Response with custom header"))
    })

    http.ListenAndServe(":8080", r)
}




// Handling errors with Chi.
package main

import (
    "net/http"

    "github.com/go-chi/chi"
)

func main() {
    r := chi.NewRouter()

    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        // Simulate an error
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    })

    http.ListenAndServe(":8080", r)
}




// Serving a Single Page Application (SPA) with Chi.
package main

import (
    "net/http"

    "github.com/go-chi/chi"
    "github.com/go-chi/chi/middleware"
)

func main() {
    r := chi.NewRouter()

    // Serve static files
    fs := http.FileServer(http.Dir("./static"))
    r.Handle("/*", http.StripPrefix("/", fs))

    // SPA routing
    r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "./static/index.html")
    })

    http.ListenAndServe(":8080", r)
}




// Handling cookies with Chi.
package main

import (
    "net/http"

    "github.com/go-chi/chi"
)

func main() {
    r := chi.NewRouter()

    r.Get("/set-cookie", func(w http.ResponseWriter, r *http.Request) {
        cookie := http.Cookie{
            Name:  "username",
            Value: "john_doe",
        }
        http.SetCookie(w, &cookie)
        w.Write([]byte("Cookie set"))
    })

    r.Get("/get-cookie", func(w http.ResponseWriter, r *http.Request) {
        cookie, err := r.Cookie("username")
        if err != nil {
            w.Write([]byte("No cookie found"))
            return
        }
        w.Write([]byte("Cookie value: " + cookie.Value))
    })

    http.ListenAndServe(":8080", r)
}




// Handling request timeout with Chi.
package main

import (
    "context"
    "net/http"
    "time"

    "github.com/go-chi/chi"
)

func main() {
    r := chi.NewRouter()

    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
        defer cancel()

        select {
        case <-time.After(3 * time.Second):
            w.Write([]byte("Operation completed"))
        case <-ctx.Done():
            w.Write([]byte("Request timed out"))
        }
    })

    http.ListenAndServe(":8080", r)
}





// Handling redirects with Chi.
package main

import (
    "net/http"

    "github.com/go-chi/chi"
)

func main() {
    r := chi.NewRouter()

    r.Get("/redirect", func(w http.ResponseWriter, r *http.Request) {
        http.Redirect(w, r, "/destination", http.StatusTemporaryRedirect)
    })

    r.Get("/destination", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Redirected to destination"))
    })

    http.ListenAndServe(":8080", r)
}





// Handling form data with Chi.
package main

import (
    "net/http"

    "github.com/go-chi/chi"
)

func main() {
    r := chi.NewRouter()

    r.Post("/form", func(w http.ResponseWriter, r *http.Request) {
        err := r.ParseForm()
        if err != nil {
            http.Error(w, "Failed to parse form", http.StatusBadRequest)
            return
        }
        username := r.Form.Get("username")
        password := r.Form.Get("password")
        // Handle form data
        w.Write([]byte("Received form data"))
    })

    http.ListenAndServe(":8080", r)
}





// Implementing request logging with Chi.
package main

import (
    "net/http"
    "os"

    "github.com/go-chi/chi"
    "github.com/go-chi/chi/middleware"
)

func main() {
    r := chi.NewRouter()
    r.Use(middleware.Logger)

    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, Chi!"))
    })

    http.ListenAndServe(":8080", r)
}




// Implementing gzip compression with Chi.
package main

import (
    "net/http"

    "github.com/go-chi/chi"
    "github.com/go-chi/chi/middleware"
)

func main() {
    r := chi.NewRouter()
    r.Use(middleware.Compress(5))

    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, Chi!"))
    })

    http.ListenAndServe(":8080", r)
}




// Chaining custom middleware with Chi.
package main

import (
    "net/http"

    "github.com/go-chi/chi"
)

func FirstMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // First middleware logic
        next.ServeHTTP(w, r)
    })
}

func SecondMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Second middleware logic
        next.ServeHTTP(w, r)
    })
}

func main() {
    r := chi.NewRouter()
    r.Use(FirstMiddleware)
    r.Use(SecondMiddleware)

    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, Chi!"))
    })

    http.ListenAndServe(":8080", r)
}





// Implementing custom error handling middleware with Chi.
package main

import (
    "net/http"

    "github.com/go-chi/chi"
    "github.com/go-chi/chi/middleware"
)

func CustomErrorHandler(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if r := recover(); r != nil {
                http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            }
        }()
        next.ServeHTTP(w, r)
    })
}

func main() {
    r := chi.NewRouter()
    r.Use(CustomErrorHandler)

    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        // Simulate an error
        panic("Something went wrong")
    })

    http.ListenAndServe(":8080", r)
}





// Using context-based middleware with Chi.
package main

import (
    "net/http"
    "time"

    "github.com/go-chi/chi"
)

func TimeoutMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
        defer cancel()

        r = r.WithContext(ctx)
        next.ServeHTTP(w, r)
    })
}

func main() {
    r := chi.NewRouter()
    r.Use(TimeoutMiddleware)

    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        time.Sleep(3 * time.Second) // Simulate a long operation
        w.Write([]byte("Operation completed"))
    })

    http.ListenAndServe(":8080", r)
}





// Streaming responses with Chi.
package main

import (
    "net/http"
    "strconv"
    "time"

    "github.com/go-chi/chi"
)

func main() {
    r := chi.NewRouter()

    r.Get("/stream", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/event-stream")
        w.Header().Set("Cache-Control", "no-cache")
        w.Header().Set("Connection", "keep-alive")

        ticker := time.NewTicker(1 * time.Second)
        defer ticker.Stop()

        for {
            select {
            case <-ticker.C:
                w.Write([]byte("data: " + strconv.FormatInt(time.Now().Unix(), 10) + "\n\n"))
                w.(http.Flusher).Flush()
            case <-r.Context().Done():
                return
            }
        }
    })

    http.ListenAndServe(":8080", r)
}





// Handling file uploads with Chi.
package main

import (
    "fmt"
    "net/http"
    "os"
    "path/filepath"

    "github.com/go-chi/chi"
)

func main() {
    r := chi.NewRouter()

    r.Post("/upload", func(w http.ResponseWriter, r *http.Request) {
        file, handler, err := r.FormFile("file")
        if err != nil {
            http.Error(w, "Failed to get file", http.StatusBadRequest)
            return
        }
        defer file.Close()

        // Save uploaded file
        filePath := filepath.Join("./uploads", handler.Filename)
        f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
        if err != nil {
            http.Error(w, "Failed to save file", http.StatusInternalServerError)
            return
        }
        defer f.Close()
        _, err = io.Copy(f, file)
        if err != nil {
            http.Error(w, "Failed to save file", http.StatusInternalServerError)
            return
        }

        w.Write([]byte(fmt.Sprintf("File %s uploaded successfully", handler.Filename)))
    })

    http.ListenAndServe(":8080", r)
}




// Parsing and validating requests with Chi.
package main

import (
    "encoding/json"
    "net/http"

    "github.com/go-chi/chi"
    "github.com/go-chi/chi/middleware"
    "gopkg.in/go-playground/validator.v9"
)

type User struct {
    ID   string `json:"id" validate:"required"`
    Name string `json:"name" validate:"required"`
}

func main() {
    r := chi.NewRouter()
    r.Use(middleware.Logger)

    r.Post("/user", func(w http.ResponseWriter, r *http.Request) {
        var user User
        err := json.NewDecoder(r.Body).Decode(&user)
        if err != nil {
            http.Error(w, "Invalid request body", http.StatusBadRequest)
            return
        }

        validate := validator.New()
        if err := validate.Struct(user); err != nil {
            http.Error(w, "Validation error", http.StatusBadRequest)
            return
        }

        // Process valid user
        w.Write([]byte("User created"))
    })

    http.ListenAndServe(":8080", r)
}





// Implementing a reverse proxy with Chi.
package main

import (
    "net/http"
    "net/http/httputil"
    "net/url"

    "github.com/go-chi/chi"
)

func main() {
    r := chi.NewRouter()

    r.Get("/proxy", func(w http.ResponseWriter, r *http.Request) {
        targetURL := "http://example.com"
        url, err := url.Parse(targetURL)
        if err != nil {
            http.Error(w, "Failed to parse target URL", http.StatusInternalServerError)
            return
        }

        proxy := httputil.NewSingleHostReverseProxy(url)
        proxy.ServeHTTP(w, r)
    })

    http.ListenAndServe(":8080", r)
}





// Implementing a custom response writer with Chi.
package main

import (
    "net/http"
    "bytes"
    "log"

    "github.com/go-chi/chi"
)

type customResponseWriter struct {
    http.ResponseWriter
    statusCode int
    buffer     *bytes.Buffer
}

func (w *customResponseWriter) WriteHeader(statusCode int) {
    w.statusCode = statusCode
    w.ResponseWriter.WriteHeader(statusCode)
}

func (w *customResponseWriter) Write(data []byte) (int, error) {
    if w.statusCode == 0 {
        w.statusCode = http.StatusOK
    }
    return w.buffer.Write(data)
}

func (w *customResponseWriter) Flush() {
    _, err := w.ResponseWriter.Write(w.buffer.Bytes())
    if err != nil {
        log.Printf("Failed to flush response: %v", err)
    }
}

func main() {
    r := chi.NewRouter()

    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        customWriter := &customResponseWriter{
            ResponseWriter: w,
            buffer:         &bytes.Buffer{},
        }

        customWriter.WriteHeader(http.StatusCreated)
        customWriter.Write([]byte("Custom response"))
        customWriter.Flush()
    })

    http.ListenAndServe(":8080", r)
}





// Implementing OAuth authentication with Chi.
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "net/http"

    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"

    "github.com/go-chi/chi"
)

var googleOAuthConfig = oauth2.Config{
    ClientID:     "CLIENT_ID",
    ClientSecret: "CLIENT_SECRET",
    Endpoint:     google.Endpoint,
    RedirectURL:  "http://localhost:8080/callback",
    Scopes:       []string{"profile", "email"},
}

func main() {
    r := chi.NewRouter()

    r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
        url := googleOAuthConfig.AuthCodeURL("state")
        http.Redirect(w, r, url, http.StatusTemporaryRedirect)
    })

    r.Get("/callback", func(w http.ResponseWriter, r *http.Request) {
        code := r.URL.Query().Get("code")
        token, err := googleOAuthConfig.Exchange(context.Background(), code)
        if err != nil {
            http.Error(w, fmt.Sprintf("Failed to exchange token: %v", err), http.StatusBadRequest)
            return
        }

        // Use token to get user information
        client := googleOAuthConfig.Client(context.Background(), token)
        profileInfo, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
        if err != nil {
            http.Error(w, fmt.Sprintf("Failed to get user profile: %v", err), http.StatusBadRequest)
            return
        }

        defer profileInfo.Body.Close()

        var profile map[string]interface{}
        if err := json.NewDecoder(profileInfo.Body).Decode(&profile); err != nil {
            http.Error(w, fmt.Sprintf("Failed to decode profile: %v", err), http.StatusInternalServerError)
            return
        }

        fmt.Fprintf(w, "Profile: %+v", profile)
    })

    http.ListenAndServe(":8080", r)
}










