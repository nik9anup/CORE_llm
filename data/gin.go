// Initialize a Gin router and start the server on port 8080.
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.Run() // listen and serve on 0.0.0.0:8080
}



// Define a simple GET route that responds with a JSON message.
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    r.Run() // listen and serve on 0.0.0.0:8080
}



// Define a simple POST route that responds with a success message.
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.POST("/submit", func(c *gin.Context) {
        c.String(200, "Posted successfully")
    })
    r.Run() // listen and serve on 0.0.0.0:8080
}



// Extract a path parameter and respond with a greeting message.
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.GET("/user/:name", func(c *gin.Context) {
        name := c.Param("name")
        c.String(200, "Hello %s", name)
    })
    r.Run() // listen and serve on 0.0.0.0:8080
}



// Extract query parameters and respond with a personalized greeting.
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.GET("/welcome", func(c *gin.Context) {
        firstname := c.DefaultQuery("firstname", "Guest")
        lastname := c.Query("lastname")
        c.String(200, "Hello %s %s", firstname, lastname)
    })
    r.Run() // listen and serve on 0.0.0.0:8080
}



// Group routes under a common path prefix.
package main

import "github.com/gin-gonic/gin"

func loginEndpoint(c *gin.Context) {
    c.String(200, "Login")
}

func submitEndpoint(c *gin.Context) {
    c.String(200, "Submit")
}

func main() {
    r := gin.Default()
    api := r.Group("/api")
    {
        api.GET("/login", loginEndpoint)
        api.GET("/submit", submitEndpoint)
    }
    r.Run() // listen and serve on 0.0.0.0:8080
}



// Respond with a JSON object.
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.GET("/json", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "hey", "status": 200})
    })
    r.Run() // listen and serve on 0.0.0.0:8080
}



// Respond with an XML object.
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.GET("/xml", func(c *gin.Context) {
        c.XML(200, gin.H{"message": "hey", "status": 200})
    })
    r.Run() // listen and serve on 0.0.0.0:8080
}



// Render an HTML template.
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.LoadHTMLGlob("templates/*")
    r.GET("/index", func(c *gin.Context) {
        c.HTML(200, "index.tmpl", gin.H{
            "title": "Main website",
        })
    })
    r.Run() // listen and serve on 0.0.0.0:8080
}



// Serve static files from a directory.
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.Static("/assets", "./assets")
    r.Run() // listen and serve on 0.0.0.0:8080
}



// Use default logger and recovery middleware.
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.Use(gin.Logger())
    r.Use(gin.Recovery())
    r.Run() // listen and serve on 0.0.0.0:8080
}



// Define and use a custom middleware.
package main

import "github.com/gin-gonic/gin"

func AuthRequired() gin.HandlerFunc {
    return func(c *gin.Context) {
        // logic here
        c.Next()
    }
}

func main() {
    r := gin.Default()
    r.Use(AuthRequired())
    r.Run() // listen and serve on 0.0.0.0:8080
}



// Handle a form POST request.
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.POST("/form_post", func(c *gin.Context) {
        message := c.PostForm("message")
        nick := c.DefaultPostForm("nick", "anonymous")
        c.JSON(200, gin.H{"status": "posted", "message": message, "nick": nick})
    })
    r.Run() // listen and serve on 0.0.0.0:8080
}



// Handle a file upload request.
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.POST("/upload", func(c *gin.Context) {
        file, _ := c.FormFile("file")
        c.SaveUploadedFile(file, "/tmp/"+file.Filename)
        c.String(200, "File uploaded successfully")
    })
    r.Run() // listen and serve on 0.0.0.0:8080
}



// Redirect to an external URL.
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.GET("/redirect", func(c *gin.Context) {
        c.Redirect(301, "http://www.google.com/")
    })
    r.Run() // listen and serve on 0.0.0.0:8080
}



// Define a custom handler for unknown routes.
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.NoRoute(func(c *gin.Context) {
        c.JSON(404, gin.H{"message": "Not Found"})
    })
    r.Run() // listen and serve on 0.0.0.0:8080
}



// Recover from panics to avoid crashing the server.
package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
)

func main() {
    defer func() {
        if err := recover(); err != nil {
            fmt.Println("Recovered from panic:", err)
        }
    }()

    r := gin.Default()
    r.GET("/panic", func(c *gin.Context) {
        panic("An unexpected error occurred!")
    })
    r.Run() // listen and serve on 0.0.0.0:8080
}



// Bind JSON payload to a struct.
package main

import (
    "github.com/gin-gonic/gin"
)

type Login struct {
    User     string `json:"user" binding:"required"`
    Password string `json:"password" binding:"required"`
}

func main() {
    r := gin.Default()
    r.POST("/login", func(c *gin.Context) {
        var json Login
        if err := c.ShouldBindJSON(&json); err != nil {
            c.JSON(400, gin.H{"error": err.Error()})
            return
        }
        c.JSON(200, gin.H{"status": "you are logged in"})
    })
    r.Run // listen and serve on 0.0.0.0:8080
}



// Bind form data to a struct.
package main

import (
    "github.com/gin-gonic/gin"
)

type Login struct {
    User     string `form:"user" binding:"required"`
    Password string `form:"password" binding:"required"`
}

func main() {
    r := gin.Default()
    r.POST("/login", func(c *gin.Context) {
        var form Login
        if err := c.ShouldBind(&form); err != nil {
            c.JSON(400, gin.H{"error": err.Error()})
            return
        }
        c.JSON(200, gin.H{"status": "you are logged in"})
    })
    r.Run() // listen and serve on 0.0.0.0:8080
}



// Bind URI parameters to a struct.
package main

import (
    "github.com/gin-gonic/gin"
)

type Person struct {
    Name string `uri:"name" binding:"required"`
    ID   string `uri:"id" binding:"required,uuid"`
}

func main() {
    r := gin.Default()
    r.GET("/person/:name/:id", func(c *gin.Context) {
        var person Person
        if err := c.ShouldBindUri(&person); err != nil {
            c.JSON(400, gin.H{"error": err.Error()})
            return
        }
        c.JSON(200, gin.H{"name": person.Name, "uuid": person.ID})
    })
    r.Run() // listen and serve on 0.0.0.0:8080
}



// Define a custom validation function.
package main

import (
    "gopkg.in/go-playground/validator.v8"
    "github.com/gin-gonic/gin"
)

type Booking struct {
    CheckIn  string `form:"check_in" binding:"required,bookabledate"`
    CheckOut string `form:"check_out" binding:"required,gtfield=CheckIn"`
}

func bookableDate(
    v *validator.Validate, topStruct interface{}, currentStructField interface{},
    field interface{}, param string) bool {
    // Custom validation logic here.
    return true
}

func main() {
    r := gin.Default()
    if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
        v.RegisterValidation("bookabledate", bookableDate)
    }

    r.POST("/book", func(c *gin.Context) {
        var b Booking
        if err := c.ShouldBind(&b); err != nil {
            c.JSON(400, gin.H{"error": err.Error()})
            return
        }
        c.JSON(200, gin.H{"message": "Booking dates are valid!"})
    })

    r.Run() // listen and serve on 0.0.0.0:8080
}




// Set and retrieve a context key.
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.Use(func(c *gin.Context) {
        c.Set("example", "12345")
        c.Next()
    })

    r.GET("/context", func(c *gin.Context) {
        example := c.MustGet("example").(string)
        c.JSON(200, gin.H{"example": example})
    })

    r.Run() // listen and serve on 0.0.0.0:8080
}



// Use a context copy for asynchronous operations.
package main

import (
    "log"
    "time"

    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.GET("/long_async", func(c *gin.Context) {
        cCp := c.Copy()
        go func() {
            time.Sleep(5 * time.Second)
            log.Println("Done! in path " + cCp.Request.URL.Path)
        }()
        c.String(200, "working")
    })

    r.Run() // listen and serve on 0.0.0.0:8080
}



// Serve a file for download.
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.GET("/download", func(c *gin.Context) {
        c.File("/tmp/file.txt")
    })
    r.Run() // listen and serve on 0.0.0.0:8080
}




// Respond with an HTML string.
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.GET("/html", func(c *gin.Context) {
        c.Data(200, "text/html; charset=utf-8", []byte("<html><body>Hello, World!</body></html>"))
    })
    r.Run() // listen and serve on 0.0.0.0:8080
}




// Use a custom render function.
package main

import (
    "github.com/gin-gonic/gin"
)

type JSONP struct {
    Callback string
    Data     interface{}
}

func (r JSONP) Render(w http.ResponseWriter) error {
    if callback := r.Callback; callback != "" {
        w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
        _, err := w.Write([]byte(callback + "("))
        if err != nil {
            return err
        }
        json.NewEncoder(w).Encode(r.Data)
        _, err = w.Write([]byte(");"))
        return err
    }
    return nil
}

func (r JSONP) WriteContentType(w http.ResponseWriter) {
    w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
}

func main() {
    r := gin.Default()
    r.GET("/JSONP", func(c *gin.Context) {
        c.Render(200, JSONP{Callback: "callback", Data: gin.H{"foo": "bar"}})
    })
    r.Run() // listen and serve on 0.0.0.0:8080
}




// Log incoming requests using the default writer.
package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.Use(gin.LoggerWithWriter(gin.DefaultWriter))
    r.Run() // listen and serve on 0.0.0.0:8080
}




// Use secure middleware to enforce SSL.
package main

import (
    "github.com/gin-gonic/gin"
    "github.com/unrolled/secure"
)

func main() {
    r := gin.Default()
    r.Use(func() gin.HandlerFunc {
        return secure.New(secure.Config{
            SSLRedirect: true,
        }).Handler
    }())
    r.Run() // listen and serve on 0.0.0.0:8080
}




// Apply rate limiting to requests.
package main

import (
    "github.com/didip/tollbooth"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    limiter := tollbooth.NewLimiter(1, nil)

    r.Use(func(c *gin.Context) {
        httpError := tollbooth.LimitByRequest(limiter, c.Writer, c.Request)
        if httpError != nil {
            c.String(httpError.StatusCode, httpError.Message)
            c.Abort()
            return
        }
        c.Next()
    })

    r.Run() // listen and serve on 0.0.0.0:8080
}




// Use basic authentication for routes.
package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.Use(gin.BasicAuth(gin.Accounts{
        "foo": "bar",
    }))
    r.Run() // listen and serve on 0.0.0.0:8080
}




// Set custom template functions.
package main

import (
    "github.com/gin-gonic/gin"
    "html/template"
    "time"
)

func main() {
    r := gin.Default()
    r.SetFuncMap(template.FuncMap{
        "formatAsDate": func(t time.Time) string {
            return t.Format("2006-01-02")
        },
    })
    r.LoadHTMLGlob("templates/*")
    r.GET("/date", func(c *gin.Context) {
        c.HTML(200, "date.tmpl", gin.H{
            "now": time.Now(),
        })
    })
    r.Run() // listen and serve on 0.0.0.0:8080
}




// Handle file upload with progress information.
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    r := gin.Default()
    r.POST("/upload", func(c *gin.Context) {
        file, err := c.FormFile("file")
        if err != nil {
            c.String(http.StatusBadRequest, "Bad request")
            return
        }

        c.String(http.StatusOK, "Uploaded successfully %s", file.Filename)
    })
    r.Run() // listen and serve on 0.0.0.0:8080
}




// Set custom headers in the response.
package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.GET("/headers", func(c *gin.Context) {
        c.Header("Content-Type", "application/json")
        c.Header("X-Custom-Header", "value")
        c.JSON(200, gin.H{"message": "Headers set"})
    })
    r.Run() // listen and serve on 0.0.0.0:8080
}




// Apply GZIP middleware to compress responses.
package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/gzip"
)

func main() {
    r := gin.Default()
    r.Use(gzip.Gzip(gzip.DefaultCompression))
    r.GET("/gzip", func(c *gin.Context) {
        c.String(200, "This is a compressed response")
    })
    r.Run() // listen and serve on 0.0.0.0:8080
}




// Add request ID middleware to track requests.
package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/requestid"
)

func main() {
    r := gin.Default()
    r.Use(requestid.New())
    r.GET("/request_id", func(c *gin.Context) {
        id := requestid.Get(c)
        c.JSON(200, gin.H{"request_id": id})
    })
    r.Run() // listen and serve on 0.0.0.0:8080
}




// Implement rate limiting using Gin middleware.
package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/limiter"
    "github.com/ulule/limiter/v3"
    "github.com/ulule/limiter/v3/drivers/store/memory"
)

func main() {
    r := gin.Default()
    rate := limiter.Rate{
        Period: 1 * time.Second,
        Limit:  1,
    }
    store := memory.NewStore()
    instance := limiter.New(store, rate)
    r.Use(limiter.NewMiddleware(instance))
    r.GET("/limited", func(c *gin.Context) {
        c.String(200, "This is a rate-limited route")
    })
    r.Run() // listen and serve on 0.0.0.0:8080
}




// Use localization middleware for multi-language support.
package main

import (
    "github.com/gin-gonic/gin"
    "github.com/nicksnyder/go-i18n/v2/i18n"
    "golang.org/x/text/language"
    "github.com/gin-contrib/multitemplate"
)

func main() {
    r := gin.Default()
    bundle := i18n.NewBundle(language.English)
    bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
    localizer := i18n.NewLocalizer(bundle, "en")

    r.Use(func(c *gin.Context) {
        c.Set("localizer", localizer)
        c.Next()
    })

    r.GET("/hello", func(c *gin.Context) {
        localizer := c.MustGet("localizer").(*i18n.Localizer)
        message := localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Hello"})
        c.String(200, message)
    })

    r.Run() // listen and serve on 0.0.0.0:8080
}




// Version routes using URL prefixes.
package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    v1 := r.Group("/v1")
    {
        v1.GET("/hello", func(c *gin.Context) {
            c.String(200, "Hello from v1")
        })
    }
    v2 := r.Group("/v2")
    {
        v2.GET("/hello", func(c *gin.Context) {
            c.String(200, "Hello from v2")
        })
    }
    r.Run() // listen and serve on 0.0.0.0:8080
}




// Serve custom error pages for different status codes.
package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.NoRoute(func(c *gin.Context) {
        c.HTML(404, "404.html", nil)
    })
    r.LoadHTMLGlob("templates/*")
    r.Run() // listen and serve on 0.0.0.0:8080
}




// Stream a response to the client.
package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.GET("/stream", func(c *gin.Context) {
        c.Stream(func(w io.Writer) bool {
            for i := 0; i < 10; i++ {
                fmt.Fprintf(w, "data %d\n", i)
                time.Sleep(1 * time.Second)
            }
            return false
        })
    })
    r.Run() // listen and serve on 0.0.0.0:8080
}




// Use a custom log formatter for the logger.
package main

import (
    "github.com/gin-gonic/gin"
    "time"
)

func main() {
    r := gin.Default()
    r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
        return fmt.Sprintf("[%s] %s %s %d %s\n",
            param.TimeStamp.Format(time.RFC822),
            param.ClientIP,
            param.Method,
            param.StatusCode,
            param.Latency,
        )
    }))
    r.GET("/ping", func(c *gin.Context) {
        c.String(200, "pong")
    })
    r.Run() // listen and serve on 0.0.0.0:8080
}




// Validate route parameters using middleware.
package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/gin/binding"
    "gopkg.in/go-playground/validator.v8"
)

func main() {
    r := gin.Default()
    if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
        v.RegisterValidation("custom", func(
            v *validator.Validate, topStruct interface{}, field interface{},
            param string) bool {
            return field.(string) == "custom"
        })
    }

    r.GET("/validate/:param", func(c *gin.Context) {
        param := c.Param("param")
        if param != "custom" {
            c.JSON(400, gin.H{"error": "invalid parameter"})
            return
        }
        c.JSON(200, gin.H{"message": "valid parameter"})
    })

    r.Run() // listen and serve on 0.0.0.0:8080
}




// Implement IP whitelisting for access control.
package main

import (
    "github.com/gin-gonic/gin"
    "net"
)

func main() {
    r := gin.Default()
    r.Use(func(c *gin.Context) {
        whitelist := []string{"127.0.0.1", "192.168.1.1"}
        clientIP := net.ParseIP(c.ClientIP())
        for _, ip := range whitelist {
            if clientIP.Equal(net.ParseIP(ip)) {
                c.Next()
                return
            }
        }
        c.JSON(403, gin.H{"error": "IP not whitelisted"})
        c.Abort()
    })
    r.GET("/protected", func(c *gin.Context) {
        c.String(200, "You have access")
    })
    r.Run() // listen and serve on 0.0.0.0:8080
}




// Apply request throttling using a simple counter.
package main

import (
    "github.com/gin-gonic/gin"
    "time"
)

func main() {
    r := gin.Default()
    var counter int
    r.Use(func(c *gin.Context) {
        counter++
        if counter > 100 {
            c.JSON(429, gin.H{"error": "too many requests"})
            c.Abort()
            return
        }
        time.AfterFunc(1*time.Minute, func() {
            counter--
        })
        c.Next()
    })
    r.GET("/throttled", func(c *gin.Context) {
        c.String(200, "This is a throttled route")
    })
    r.Run() // listen and serve on 0.0.0.0:8080
}




// Implement Server-Sent Events (SSE) with Gin.
package main

import (
    "github.com/gin-gonic/gin"
    "time"
)

func main() {
    r := gin.Default()
    r.GET("/sse", func(c *gin.Context) {
        c.Stream(func(w io.Writer) bool {
            for i := 0; i < 10; i++ {
                fmt.Fprintf(w, "data: %d\n\n", i)
                time.Sleep(1 * time.Second)
            }
            return false
        })
    })
    r.Run() // listen and serve on 0.0.0.0:8080
}




// Apply CORS middleware to allow cross-origin requests.
package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

func main() {
    r := gin.Default()
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://example.com"},
        AllowMethods:     []string{"GET", "POST"},
        AllowHeaders:     []string{"Origin"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))
    r.GET("/cors", func(c *gin.Context) {
        c.String(200, "CORS enabled")
    })
    r.Run() // listen and serve on 0.0.0.0:8080
}




// Serve static files from a directory.
package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.Static("/static", "./static")
    r.Run() // listen and serve on 0.0.0.0:8080
}




// Manage user sessions with Gin.
package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/sessions"
    "github.com/gin-contrib/sessions/cookie"
)

func main() {
    r := gin.Default()
    store := cookie.NewStore([]byte("secret"))
    r.Use(sessions.Sessions("mysession", store))

    r.GET("/session", func(c *gin.Context) {
        session := sessions.Default(c)
        session.Set("foo", "bar")
        session.Save()
        c.JSON(200, gin.H{"message": "session saved"})
    })

    r.Run() // listen and serve on 0.0.0.0:8080
}




// Add a health check endpoint to monitor service status.
package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.GET("/health", func(c *gin.Context) {
        c.String(200, "OK")
    })
    r.Run() // listen and serve on 0.0.0.0:8080
}




// Implement simple cache middleware.
package main

import (
    "github.com/gin-gonic/gin"
    "github.com/patrickmn/go-cache"
    "time"
)

func main() {
    r := gin.Default()
    c := cache.New(5*time.Minute, 10*time.Minute)
    r.Use(func(c *gin.Context) {
        if data, found := c.Get(c.Request.RequestURI); found {
            c.JSON(200, data)
            c.Abort()
            return
        }
        c.Next()
        c.Set(c.Request.RequestURI, c.Writer.Body)
    })

    r.GET("/cached", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "This is a cached response"})
    })

    r.Run() // listen and serve on 0.0.0.0:8080
}






















