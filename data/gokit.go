// Basic Go Kit service setup
package main

import (
	"context"
	"net/http"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

type StringService struct{}

func (StringService) Uppercase(_ context.Context, s string) (string, error) {
	return strings.ToUpper(s), nil
}

func makeUppercaseEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(string)
		v, err := svc.Uppercase(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func main() {
	logger := log.NewLogfmtLogger(log.StdlibWriter{})
	svc := StringService{}
	uppercaseHandler := httptransport.NewServer(
		makeUppercaseEndpoint(svc),
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var s string
			if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
				return nil, err
			}
			return s, nil
		},
		func(_ context.Context, w http.ResponseWriter, response interface{}) error {
			return json.NewEncoder(w).Encode(response)
		},
	)
	http.Handle("/uppercase", uppercaseHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	http.ListenAndServe(":8080", nil)
}







// Logging middleware for Go Kit service
package main

import (
	"context"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func loggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			start := time.Now()
			response, err := next(ctx, request)
			logger.Log("request", request, "response", response, "took", time.Since(start))
			return response, err
		}
	}
}

func main() {
	logger := log.NewLogfmtLogger(log.StdlibWriter{})
	svc := StringService{}
	uppercaseEndpoint := loggingMiddleware(logger)(makeUppercaseEndpoint(svc))

	uppercaseHandler := httptransport.NewServer(
		uppercaseEndpoint,
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var s string
			if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
				return nil, err
			}
			return s, nil
		},
		func(_ context.Context, w http.ResponseWriter, response interface{}) error {
			return json.NewEncoder(w).Encode(response)
		},
	)

	http.Handle("/uppercase", uppercaseHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	http.ListenAndServe(":8080", nil)
}







// Basic Go Kit client example
package main

import (
	"context"
	"fmt"
	"net/url"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := log.NewLogfmtLogger(log.StdlibWriter{})
	u, _ := url.Parse("http://localhost:8080/uppercase")
	client := httptransport.NewClient(
		"POST",
		u,
		func(_ context.Context, r *http.Request, request interface{}) error {
			s := request.(string)
			r.Body = ioutil.NopCloser(strings.NewReader(s))
			return nil
		},
		func(_ context.Context, r *http.Response) (interface{}, error) {
			var response string
			if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
				return nil, err
			}
			return response, nil
		},
	).Endpoint()

	response, err := client(context.Background(), "hello")
	if err != nil {
		logger.Log("err", err)
		return
	}
	fmt.Println(response)
}






// Instrumentation middleware for Go Kit service
package main

import (
	"context"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func instrumentationMiddleware(requestCount metrics.Counter, duration metrics.Histogram) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			defer func(begin time.Time) {
				requestCount.Add(1)
				duration.Observe(time.Since(begin).Seconds())
			}(time.Now())
			return next(ctx, request)
		}
	}
}

func main() {
	logger := log.NewLogfmtLogger(log.StdlibWriter{})
	requestCount := prometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "example",
		Subsystem: "string_service",
		Name:      "request_count",
	}, []string{})
	duration := prometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "example",
		Subsystem: "string_service",
		Name:      "duration_seconds",
	}, []string{})

	svc := StringService{}
	uppercaseEndpoint := instrumentationMiddleware(requestCount, duration)(makeUppercaseEndpoint(svc))

	uppercaseHandler := httptransport.NewServer(
		uppercaseEndpoint,
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var s string
			if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
				return nil, err
			}
			return s, nil
		},
		func(_ context.Context, w http.ResponseWriter, response interface{}) error {
			return json.NewEncoder(w).Encode(response)
		},
	)

	http.Handle("/uppercase", uppercaseHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	http.ListenAndServe(":8080", nil)
}






// Circuit breaker middleware for Go Kit service
package main

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd/circuitbreaker"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/sony/gobreaker"
)

func main() {
	logger := log.NewLogfmtLogger(log.StdlibWriter{})
	svc := StringService{}
	breaker := circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))

	uppercaseEndpoint := breaker(makeUppercaseEndpoint(svc))

	uppercaseHandler := httptransport.NewServer(
		uppercaseEndpoint,
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var s string
			if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
				return nil, err
			}
			return s, nil
		},
		func(_ context.Context, w http.ResponseWriter, response interface{}) error {
			return json.NewEncoder(w).Encode(response)
		},
	)

	http.Handle("/uppercase", uppercaseHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	http.ListenAndServe(":8080", nil)
}






// Using retry middleware with Go Kit service
package main

import (
	"context"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd/retry"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := log.NewLogfmtLogger(log.StdlibWriter{})
	svc := StringService{}
	retryEndpoint := retry.NewEndpoint(
		makeUppercaseEndpoint(svc),
		retry.RetryOptions{Max: 3, Interval: 5 * time.Second},
	)

	uppercaseHandler := httptransport.NewServer(
		retryEndpoint,
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var s string
			if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
				return nil, err
			}
			return s, nil
		},
		func(_ context.Context, w http.ResponseWriter, response interface{}) error {
			return json.NewEncoder(w).Encode(response)
		},
	)

	http.Handle("/uppercase", uppercaseHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	http.ListenAndServe(":8080", nil)
}






// Adding custom request context in Go Kit service
package main

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := log.NewLogfmtLogger(log.StdlibWriter{})
	svc := StringService{}
	uppercaseEndpoint := makeUppercaseEndpoint(svc)

	uppercaseHandler := httptransport.NewServer(
		uppercaseEndpoint,
		func(_ context.Context, r *http.Request) (interface{}, error) {
			ctx := context.WithValue(r.Context(), "requestID", "123")
			r = r.WithContext(ctx)
			var s string
			if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
				return nil, err
			}
			return s, nil
		},
		func(_ context.Context, w http.ResponseWriter, response interface{}) error {
			return json.NewEncoder(w).Encode(response)
		},
	)

	http.Handle("/uppercase", uppercaseHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	http.ListenAndServe(":8080", nil)
}






// Adding error handling middleware to Go Kit service
package main

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func errorMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			response, err := next(ctx, request)
			if err != nil {
				logger.Log("error", err)
			}
			return response, err
		}
	}
}

func main() {
	logger := log.NewLogfmtLogger(log.StdlibWriter{})
	svc := StringService{}
	uppercaseEndpoint := errorMiddleware(logger)(makeUppercaseEndpoint(svc))

	uppercaseHandler := httptransport.NewServer(
		uppercaseEndpoint,
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var s string
			if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
				return nil, err
			}
			return s, nil
		},
		func(_ context.Context, w http.ResponseWriter, response interface{}) error {
			return json.NewEncoder(w).Encode(response)
		},
	)

	http.Handle("/uppercase", uppercaseHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	http.ListenAndServe(":8080", nil)
}







// Adding authentication middleware to Go Kit service
package main

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simple token-based auth
		token := r.Header.Get("Authorization")
		if token != "valid-token" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	logger := log.NewLogfmtLogger(log.StdlibWriter{})
	svc := StringService{}
	uppercaseEndpoint := makeUppercaseEndpoint(svc)

	uppercaseHandler := httptransport.NewServer(
		uppercaseEndpoint,
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var s string
			if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
				return nil, err
			}
			return s, nil
		},
		func(_ context.Context, w http.ResponseWriter, response interface{}) error {
			return json.NewEncoder(w).Encode(response)
		},
	)

	http.Handle("/uppercase", authMiddleware(uppercaseHandler))
	logger.Log("msg", "HTTP", "addr", ":8080")
	http.ListenAndServe(":8080", nil)
}






// Adding timeout middleware to Go Kit service
package main

import (
	"context"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func timeoutMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
		defer cancel()
		return next(ctx, request)
	}
}

func main() {
	logger := log.NewLogfmtLogger(log.StdlibWriter{})
	svc := StringService{}
	uppercaseEndpoint := timeoutMiddleware(makeUppercaseEndpoint(svc))

	uppercaseHandler := httptransport.NewServer(
		uppercaseEndpoint,
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var s string
			if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
				return nil, err
			}
			return s, nil
		},
		func(_ context.Context, w http.ResponseWriter, response interface{}) error {
			return json.NewEncoder(w).Encode(response)
		},
	)

	http.Handle("/uppercase", uppercaseHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	http.ListenAndServe(":8080", nil)
}







// Using multiple endpoints in Go Kit service
package main

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := log.NewLogfmtLogger(log.StdlibWriter{})
	svc := StringService{}
	uppercaseEndpoint := makeUppercaseEndpoint(svc)
	countEndpoint := makeCountEndpoint(svc)

	uppercaseHandler := httptransport.NewServer(
		uppercaseEndpoint,
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var s string
			if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
				return nil, err
			}
			return s, nil
		},
		func(_ context.Context, w http.ResponseWriter, response interface{}) error {
			return json.NewEncoder(w).Encode(response)
		},
	)

	countHandler := httptransport.NewServer(
		countEndpoint,
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var s string
			if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
				return nil, err
			}
			return s, nil
		},
		func(_ context.Context, w http.ResponseWriter, response interface{}) error {
			return json.NewEncoder(w).Encode(response)
		},
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	http.ListenAndServe(":8080", nil)
}





// Adding tracing middleware to Go Kit service
package main

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	zipkin "github.com/openzipkin/zipkin-go"
	zipkinHTTP "github.com/openzipkin/zipkin-go/middleware/http"
	zipkinReporter "github.com/openzipkin/zipkin-go/reporter/http"
)

func main() {
	logger := log.NewLogfmtLogger(log.StdlibWriter{})
	svc := StringService{}
	tracer, _ := zipkin.NewTracer(zipkinReporter.NewReporter("http://localhost:9411/api/v2/spans"))
	uppercaseEndpoint := zipkinHTTP.NewServerMiddleware(tracer)(makeUppercaseEndpoint(svc))

	uppercaseHandler := httptransport.NewServer(
		uppercaseEndpoint,
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var s string
			if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
				return nil, err
			}
			return s, nil
		},
		func(_ context.Context, w http.ResponseWriter, response interface{}) error {
			return json.NewEncoder(w).Encode(response)
		},
	)

	http.Handle("/uppercase", uppercaseHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	http.ListenAndServe(":8080", nil)
}






// Custom request/response encoder and decoder in Go Kit service
package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	httptransport "github.com.go-kit/kit/transport/http"
)

type request struct {
	S string `json:"s"`
}

type response struct {
	V string `json:"v"`
}

func main() {
	logger := log.NewLogfmtLogger(log.StdlibWriter{})
	svc := StringService{}
	uppercaseEndpoint := makeUppercaseEndpoint(svc)

	uppercaseHandler := httptransport.NewServer(
		uppercaseEndpoint,
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var req request
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				return nil, err
			}
			return req, nil
		},
		func(_ context.Context, w http.ResponseWriter, response interface{}) error {
			return json.NewEncoder(w).Encode(response)
		},
	)

	http.Handle("/uppercase", uppercaseHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	http.ListenAndServe(":8080", nil)
}





// Service discovery with Consul in Go Kit service
package main

import (
	"context"
	"net"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/consul/consulsd"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/hashicorp/consul/api"
)

func main() {
	logger := log.NewLogfmtLogger(log.StdlibWriter{})
	consulClient, _ := api.NewClient(api.DefaultConfig())
	client := consulsd.NewClient(consulClient)
	registrar := consulsd.NewRegistrar(client, &consul.Registration{
		ID:      "uppercase",
		Name:    "uppercase",
		Address: "localhost",
		Port:    8080,
	}, logger)

	registrar.Register()
	defer registrar.Deregister()

	svc := StringService{}
	uppercaseEndpoint := makeUppercaseEndpoint(svc)

	uppercaseHandler := httptransport.NewServer(
		uppercaseEndpoint,
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var s string
			if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
				return nil, err
			}
			return s, nil
		},
		func(_ context.Context, w http.ResponseWriter, response interface{}) error {
			return json.NewEncoder(w).Encode(response)
		},
	)

	http.Handle("/uppercase", uppercaseHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	http.ListenAndServe(":8080", nil)
}




// Service discovery with etcd in Go Kit service
package main

import (
	"context"
	"net/http"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd/etcdv3"
	httptransport "github.com/go-kit/kit/transport/http"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	logger := log.NewLogfmtLogger(log.StdlibWriter{})
	client, _ := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})

	registrar := etcdv3.NewRegistrar(client, etcdv3.Service{
		Key:   "/services/uppercase",
		Value: "http://localhost:8080",
	}, logger)

	registrar.Register()
	defer registrar.Deregister()

	svc := StringService{}
	uppercaseEndpoint := makeUppercaseEndpoint(svc)

	uppercaseHandler := httptransport.NewServer(
		uppercaseEndpoint,
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var s string
			if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
				return nil, err
			}
			return s, nil
		},
		func(_ context.Context, w http.ResponseWriter, response interface{}) error {
			return json.NewEncoder(w).Encode(response)
		},
	)

	http.Handle("/uppercase", uppercaseHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	http.ListenAndServe(":8080", nil)
}






// Adding rate limiting middleware to Go Kit service
package main

import (
	"context"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/ratelimit"
	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/time/rate"
)

func main() {
	logger := log.NewLogfmtLogger(log.StdlibWriter{})
	svc := StringService{}
	limiter := rate.NewLimiter(rate.Every(1*time.Second), 1)
	uppercaseEndpoint := ratelimit.NewErroringLimiter(limiter)(makeUppercaseEndpoint(svc))

	uppercaseHandler := httptransport.NewServer(
		uppercaseEndpoint,
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var s string
			if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
				return nil, err
			}
			return s, nil
		},
		func(_ context.Context, w http.ResponseWriter, response interface{}) error {
			return json.NewEncoder(w).Encode(response)
		},
	)

	http.Handle("/uppercase", uppercaseHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	http.ListenAndServe(":8080", nil)
}







// Adding Prometheus metrics to Go Kit service
package main

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func main() {
	logger := log.NewLogfmtLogger(log.StdlibWriter{})
	requestCount := prometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "example",
		Subsystem: "string_service",
		Name:      "request_count",
	}, []string{})
	svc := StringService{}
	uppercaseEndpoint := prometheus.NewCounterMiddleware(requestCount)(makeUppercaseEndpoint(svc))

	uppercaseHandler := httptransport.NewServer(
		uppercaseEndpoint,
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var s string
			if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
				return nil, err
			}
			return s, nil
		},
		func(_ context.Context, w http.ResponseWriter, response interface{}) error {
			return json.NewEncoder(w).Encode(response)
		},
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/metrics", stdprometheus.Handler())
	logger.Log("msg", "HTTP", "addr", ":8080")
	http.ListenAndServe(":8080", nil)
}






// Adding throttling to Go Kit service
package main

import (
	"context"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/throttle"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := log.NewLogfmtLogger(log.StdlibWriter{})
	svc := StringService{}
	upperEndpoint := throttle.Throttle(1, 1*time.Second)(makeUppercaseEndpoint(svc))

	uppercaseHandler := httptransport.NewServer(
		upperEndpoint,
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var s string
			if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
				return nil, err
			}
			return s, nil
		},
		func(_ context.Context, w http.ResponseWriter, response interface{}) error {
			return json.NewEncoder(w).Encode(response)
		},
	)

	http.Handle("/uppercase", uppercaseHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	http.ListenAndServe(":8080", nil)
}





// Implementing graceful shutdown in Go Kit service
package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := log.NewLogfmtLogger(log.StdlibWriter{})
	svc := StringService{}
	uppercaseEndpoint := makeUppercaseEndpoint(svc)

	uppercaseHandler := httptransport.NewServer(
		uppercaseEndpoint,
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var s string
			if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
				return nil, err
			}
			return s, nil
		},
		func(_ context.Context, w http.ResponseWriter, response interface{}) error {
			return json.NewEncoder(w).Encode(response)
		},
	)

	http.Handle("/uppercase", uppercaseHandler)
	server := &http.Server{Addr: ":8080"}

	go func() {
		logger.Log("msg", "HTTP", "addr", ":8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log("error", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	server.Shutdown(ctx)
	logger.Log("msg", "Server gracefully stopped")
}




// Dynamic client-side load balancing in Go Kit service
package main

import (
	"context"
	"net/http"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd/lb"
	"github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/consul/consulsd"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/hashicorp/consul/api"
)

func main() {
	logger := log.NewLogfmtLogger(log.StdlibWriter{})
	consulClient, _ := api.NewClient(api.DefaultConfig())
	client := consulsd.NewClient(consulClient)

	instancer := consulsd.NewInstancer(client, logger, "uppercase", []string{}, true)
	endpointer := consulsd.NewEndpointer(instancer, makeUppercaseFactory(), logger)
	balancer := lb.NewRoundRobin(endpointer)

	svc := StringService{}
	uppercaseEndpoint := makeUppercaseEndpoint(svc)

	uppercaseHandler := httptransport.NewServer(
		uppercaseEndpoint,
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var s string
			if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
				return nil, err
			}
			return s, nil
		},
		func(_ context.Context, w http.ResponseWriter, response interface{}) error {
			return json.NewEncoder(w).Encode(response)
		},
	)

	http.Handle("/uppercase", httptransport.NewServer(balancer.Endpoint(), nil, nil))
	logger.Log("msg", "HTTP", "addr", ":8080")
	http.ListenAndServe(":8080", nil)
}






// Adding Zipkin tracing to Go Kit service
package main

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	zipkin "github.com/openzipkin/zipkin-go"
	zipkinHTTP "github.com/openzipkin/zipkin-go/middleware/http"
	zipkinReporter "github.com/openzipkin/zipkin-go/reporter/http"
)

func main() {
	logger := log.NewLogfmtLogger(log.StdlibWriter{})
	tracer, _ := zipkin.NewTracer(zipkinReporter.NewReporter("http://localhost:9411/api/v2/spans"))
	svc := StringService{}
	uppercaseEndpoint := zipkinHTTP.NewServerMiddleware(tracer)(makeUppercaseEndpoint(svc))

	uppercaseHandler := httptransport.NewServer(
		uppercaseEndpoint,
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var s string
			if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
				return nil, err
			}
			return s, nil
		},
		func(_ context.Context, w http.ResponseWriter, response interface{}) error {
			return json.NewEncoder(w).Encode(response)
		},
	)

	http.Handle("/uppercase", uppercaseHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	http.ListenAndServe(":8080", nil)
}






// Adding a custom HTTP handler in Go Kit service
package main

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := log.NewLogfmtLogger(log.StdlibWriter{})
	svc := StringService{}
	uppercaseEndpoint := makeUppercaseEndpoint(svc)

	uppercaseHandler := httptransport.NewServer(
		uppercaseEndpoint,
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var s string
			if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
				return nil, err
			}
			return s, nil
		},
		func(_ context.Context, w http.ResponseWriter, response interface{}) error {
			return json.NewEncoder(w).Encode(response)
		},
		httptransport.ServerBefore(func(ctx context.Context, r *http.Request) context.Context {
			// Add custom logic before handling request
			return ctx
		}),
		httptransport.ServerAfter(func(ctx context.Context, w http.ResponseWriter) context.Context {
			// Add custom logic after handling request
			return ctx
		}),
	)

	http.Handle("/uppercase", uppercaseHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	http.ListenAndServe(":8080", nil)
}






// Adding circuit breaker to Go Kit service
package main

import (
	"context"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd/lb"
	"github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/consul/consulsd"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/sony/gobreaker"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/hashicorp/consul/api"
)

func main() {
	logger := log.NewLogfmtLogger(log.StdlibWriter{})
	consulClient, _ := api.NewClient(api.DefaultConfig())
	client := consulsd.NewClient(consulClient)

	instancer := consulsd.NewInstancer(client, logger, "uppercase", []string{}, true)
	endpointer := consulsd.NewEndpointer(instancer, makeUppercaseFactory(), logger)
	breaker := circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))
	balancer := lb.NewRoundRobin(endpointer)

	svc := StringService{}
	uppercaseEndpoint := breaker(makeUppercaseEndpoint(svc))

	uppercaseHandler := httptransport.NewServer(
		uppercaseEndpoint,
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var s string
			if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
				return nil, err
			}
			return s, nil
		},
		func(_ context.Context, w http.ResponseWriter, response interface{}) error {
			return json.NewEncoder(w).Encode(response)
		},
	)

	http.Handle("/uppercase", httptransport.NewServer(balancer.Endpoint(), nil, nil))
	logger.Log("msg", "HTTP", "addr", ":8080")
	http.ListenAndServe(":8080", nil)
}





// Adding retry with circuit breaker to Go Kit service
package main

import (
	"context"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd/lb"
	"github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/consul/consulsd"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/retry"
	"github.com/sony/gobreaker"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/hashicorp/consul/api"
)

func main() {
	logger := log.NewLogfmtLogger(log.StdlibWriter{})
	consulClient, _ := api.NewClient(api.DefaultConfig())
	client := consulsd.NewClient(consulClient)

	instancer := consulsd.NewInstancer(client, logger, "uppercase", []string{}, true)
	endpointer := consulsd.NewEndpointer(instancer, makeUppercaseFactory(), logger)
	breaker := circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))
	retryEndpoint := retry.NewEndpoint(breaker(makeUppercaseEndpoint(svc)), retry.RetryOptions{Max: 3, Interval: 5 * time.Second})
	balancer := lb.NewRoundRobin(endpointer)

	svc := StringService{}
	uppercaseEndpoint := retryEndpoint

	uppercaseHandler := httptransport.NewServer(
		uppercaseEndpoint,
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var s string
			if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
				return nil, err
			}
			return s, nil
		},
		func(_ context.Context, w http.ResponseWriter, response interface{}) error {
			return json.NewEncoder(w).Encode(response)
		},
	)

	http.Handle("/uppercase", httptransport.NewServer(balancer.Endpoint(), nil, nil))
	logger.Log("msg", "HTTP", "addr", ":8080")
	http.ListenAndServe(":8080", nil)
}





// Adding retry with circuit breaker to Go Kit service
package main

import (
	"context"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd/lb"
	"github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/consul/consulsd"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/retry"
	"github.com/sony/gobreaker"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/hashicorp/consul/api"
)

func main() {
	logger := log.NewLogfmtLogger(log.StdlibWriter{})
	consulClient, _ := api.NewClient(api.DefaultConfig())
	client := consulsd.NewClient(consulClient)

	instancer := consulsd.NewInstancer(client, logger, "uppercase", []string{}, true)
	endpointer := consulsd.NewEndpointer(instancer, makeUppercaseFactory(), logger)
	breaker := circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))
	retryEndpoint := retry.NewEndpoint(breaker(makeUppercaseEndpoint(svc)), retry.RetryOptions{Max: 3, Interval: 5 * time.Second})
	balancer := lb.NewRoundRobin(endpointer)

	svc := StringService{}
	uppercaseEndpoint := retryEndpoint

	uppercaseHandler := httptransport.NewServer(
		uppercaseEndpoint,
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var s string
			if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
				return nil, err
			}
			return s, nil
		},
		func(_ context.Context, w http.ResponseWriter, response interface{}) error {
			return json.NewEncoder(w).Encode(response)
		},
	)

	http.Handle("/uppercase", httptransport.NewServer(balancer.Endpoint(), nil, nil))
	logger.Log("msg", "HTTP", "addr", ":8080")
	http.ListenAndServe(":8080", nil)
}






// Adding custom middleware to Go Kit service
package main

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := log.NewLogfmtLogger(log.StdlibWriter{})
	svc := StringService{}
	uppercaseEndpoint := makeUppercaseEndpoint(svc)

	uppercaseHandler := httptransport.NewServer(
		uppercaseEndpoint,
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var s string
			if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
				return nil, err
			}
			return s, nil
		},
		func(_ context.Context, w http.ResponseWriter, response interface{}) error {
			return json.NewEncoder(w).Encode(response)
		},
		httptransport.ServerBefore(func(ctx context.Context, r *http.Request) context.Context {
			// Custom middleware logic
			return ctx
		}),
	)

	http.Handle("/uppercase", uppercaseHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	http.ListenAndServe(":8080", nil)
}





// Setting up a JSON RPC server with Go Kit
package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport/http/jsonrpc"
)

func main() {
	logger := log.NewLogfmtLogger(log.StdlibWriter{})
	svc := StringService{}
	uppercaseEndpoint := makeUppercaseEndpoint(svc)

	uppercaseHandler := jsonrpc.NewServer(
		uppercaseEndpoint,
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var s string
			if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
				return nil, err
			}
			return s, nil
		},
		func(_ context.Context, w http.ResponseWriter, response interface{}) error {
			return json.NewEncoder(w).Encode(response)
		},
	)

	http.Handle("/rpc/uppercase", uppercaseHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	http.ListenAndServe(":8080", nil)
}




// Adding timeout middleware to Go Kit service
package main

import (
	"context"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := log.NewLogfmtLogger(log.StdlibWriter{})
	svc := StringService{}
	uppercaseEndpoint := makeUppercaseEndpoint(svc)
	uppercaseEndpoint = tracing.ServerMiddleware(log.NewNopLogger())(uppercaseEndpoint)

	uppercaseHandler := httptransport.NewServer(
		uppercaseEndpoint,
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var s string
			if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
				return nil, err
			}
			return s, nil
		},
		func(_ context.Context, w http.ResponseWriter, response interface{}) error {
			return json.NewEncoder(w).Encode(response)
		},
		httptransport.ServerBefore(func(ctx context.Context, r *http.Request) context.Context {
			ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
			go func() {
				<-ctx.Done()
				if ctx.Err() == context.DeadlineExceeded {
					cancel()
				}
			}()
			return ctx
		}),
	)

	http.Handle("/uppercase", uppercaseHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	http.ListenAndServe(":8080", nil)
}







// Adding caching middleware to Go Kit service
package main

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

type cache struct {
	mu    sync.Mutex
	store map[string]string
}

func (c *cache) Get(key string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.store[key]
	return val, ok
}

func (c *cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[key] = value
}

func main() {
	logger := log.NewLogfmtLogger(log.StdlibWriter{})
	svc := StringService{}
	c := &cache{store: make(map[string]string)}

	uppercaseEndpoint := makeUppercaseEndpoint(svc)
	uppercaseEndpoint = func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			req := request.(string)
			if val, ok := c.Get(req); ok {
				return val, nil
			}
			resp, err := next(ctx, req)
			if err == nil {
				c.Set(req, resp.(string))
			}
			return resp, err
		}
	}(uppercaseEndpoint)

	uppercaseHandler := httptransport.NewServer(
		uppercaseEndpoint,
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var s string
			if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
				return nil, err
			}
			return s, nil
		},
		func(_ context.Context, w http.ResponseWriter, response interface{}) error {
			return json.NewEncoder(w).Encode(response)
		},
	)

	http.Handle("/uppercase", uppercaseHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	http.ListenAndServe(":8080", nil)
}







// Adding rate limiting middleware to Go Kit service
package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/ratelimit"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := log.NewLogfmtLogger(log.StdlibWriter{})
	svc := StringService{}
	uppercaseEndpoint := makeUppercaseEndpoint(svc)
	limiter := ratelimit.NewTokenBucketLimiter(ratelimit.NewBucketWithRate(1, 1))

	uppercaseEndpoint = limiter(uppercaseEndpoint)

	uppercaseHandler := httptransport.NewServer(
		uppercaseEndpoint,
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var s string
			if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
				return nil, err
			}
			return s, nil
		},
		func(_ context.Context, w http.ResponseWriter, response interface{}) error {
			return json.NewEncoder(w).Encode(response)
		},
	)

	http.Handle("/uppercase", uppercaseHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	http.ListenAndServe(":8080", nil)
}






// Adding request logging middleware to Go Kit service
package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := log.NewLogfmtLogger(log.StdlibWriter{})
	svc := StringService{}
	uppercaseEndpoint := makeUppercaseEndpoint(svc)

	uppercaseEndpoint = func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			logger.Log("msg", "request received")
			defer logger.Log("msg", "request completed")
			return next(ctx, request)
		}
	}(uppercaseEndpoint)

	uppercaseHandler := httptransport.NewServer(
		uppercaseEndpoint,
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var s string
			if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
				return nil, err
			}
			return s, nil
		},
		func(_ context.Context, w http.ResponseWriter, response interface{}) error {
			return json.NewEncoder(w).Encode(response)
		},
	)

	http.Handle("/uppercase", uppercaseHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	http.ListenAndServe(":8080", nil)
}





// Adding metrics middleware to Go Kit service
package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := log.NewLogfmtLogger(log.StdlibWriter{})
	svc := StringService{}
	uppercaseEndpoint := makeUppercaseEndpoint(svc)
	counter := prometheus.NewCounterFrom(prometheus.CounterOpts{
		Namespace: "example",
		Name:      "uppercase_request_count",
		Help:      "Number of requests received for uppercase endpoint",
	}, []string{})

	uppercaseEndpoint = func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			counter.Add(1)
			return next(ctx, request)
		}
	}(uppercaseEndpoint)

	uppercaseHandler := httptransport.NewServer(
		uppercaseEndpoint,
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var s string
			if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
				return nil, err
			}
			return s, nil
		},
		func(_ context.Context, w http.ResponseWriter, response interface{}) error {
			return json.NewEncoder(w).Encode(response)
		},
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/metrics", promhttp.Handler())
	logger.Log("msg", "HTTP", "addr", ":8080")
	http.ListenAndServe(":8080", nil)
}






// Adding timeout and retry middleware to Go Kit service
package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/retry"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := log.NewLogfmtLogger(log.StdlibWriter{})
	svc := StringService{}
	uppercaseEndpoint := makeUppercaseEndpoint(svc)

	uppercaseEndpoint = func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			return retry.Retry(3, time.Second, 2, next)(ctx, request)
		}
	}(uppercaseEndpoint)

	uppercaseHandler := httptransport.NewServer(
		uppercaseEndpoint,
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var s string
			if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
				return nil, err
			}
			if s == "retry" {
				return nil, errors.New("temporary error, please retry")
			}
			return s, nil
		},
		func(_ context.Context, w http.ResponseWriter, response interface{}) error {
			return json.NewEncoder(w).Encode(response)
		},
		httptransport.ServerErrorEncoder(func(ctx context.Context, err error, w http.ResponseWriter) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}),
	)

	http.Handle("/uppercase", uppercaseHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	http.ListenAndServe(":8080", nil)
}





// Adding custom response headers in Go Kit service
package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := log.NewLogfmtLogger(log.StdlibWriter{})
	svc := StringService{}
	uppercaseEndpoint := makeUppercaseEndpoint(svc)

	uppercaseHandler := httptransport.NewServer(
		uppercaseEndpoint,
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var s string
			if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
				return nil, err
			}
			return s, nil
		},
		func(_ context.Context, w http.ResponseWriter, response interface{}) error {
			w.Header().Set("X-Custom-Header", "Value")
			return json.NewEncoder(w).Encode(response)
		},
	)

	http.Handle("/uppercase", uppercaseHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	http.ListenAndServe(":8080", nil)
}




// Adding request ID middleware to Go Kit service
package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

type contextKey string

const requestIDKey = contextKey("requestID")

func main() {
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	svc := StringService{}
	uppercaseEndpoint := makeUppercaseEndpoint(svc)

	uppercaseEndpoint = func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			reqID := ctx.Value(requestIDKey).(string)
			logger.Log("requestID", reqID, "msg", "request received")
			defer logger.Log("requestID", reqID, "msg", "request completed")
			return next(ctx, request)
		}
	}(uppercaseEndpoint)

	uppercaseHandler := httptransport.NewServer(
		uppercaseEndpoint,
		func(ctx context.Context, r *http.Request) (interface{}, error) {
			reqID := r.Header.Get("X-Request-ID")
			ctx = context.WithValue(ctx, requestIDKey, reqID)
			var s string
			if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
				return nil, err
			}
			return s, nil
		},
		func(_ context.Context, w http.ResponseWriter, response interface{}) error {
			return json.NewEncoder(w).Encode(response)
		},
	)

	http.Handle("/uppercase", uppercaseHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	http.ListenAndServe(":8080", nil)
}





// Adding circuit breaker middleware to Go Kit service
package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/sony/gobreaker"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := log.NewLogfmtLogger(log.StdlibWriter{})
	svc := StringService{}
	uppercaseEndpoint := makeUppercaseEndpoint(svc)
	cb := circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "uppercase",
		Timeout: 5 * time.Second,
	}))

	uppercaseEndpoint = cb(uppercaseEndpoint)

	uppercaseHandler := httptransport.NewServer(
		uppercaseEndpoint,
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var s string
			if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
				return nil, err
			}
			if s == "break" {
				return nil, errors.New("service unavailable")
			}
			return s, nil
		},
		func(_ context.Context, w http.ResponseWriter, response interface{}) error {
			return json.NewEncoder(w).Encode(response)
		},
		httptransport.ServerErrorEncoder(func(ctx context.Context, err error, w http.ResponseWriter) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}),
	)

	http.Handle("/uppercase", uppercaseHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	http.ListenAndServe(":8080", nil)
}





// Implementing service discovery in Go Kit service
package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

type Discovery struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func main() {
	logger := log.NewLogfmtLogger(log.StdlibWriter{})
	svc := StringService{}
	uppercaseEndpoint := makeUppercaseEndpoint(svc)
	discovery := Discovery{Host: "localhost", Port: "8080"}

	uppercaseHandler := httptransport.NewServer(
		uppercaseEndpoint,
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var s string
			if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
				return nil, err
			}
			return s, nil
		},
		func(_ context.Context, w http.ResponseWriter, response interface{}) error {
			w.Header().Set("X-Service-Host", discovery.Host)
			w.Header().Set("X-Service-Port", discovery.Port)
			return json.NewEncoder(w).Encode(response)
		},
	)

	http.Handle("/uppercase", uppercaseHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	http.ListenAndServe(":8080", nil)
}







// Implementing event-driven communication in Go Kit service
package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

type Event struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func main() {
	logger := log.NewLogfmtLogger(log.StdlibWriter{})
	svc := StringService{}
	uppercaseEndpoint := makeUppercaseEndpoint(svc)
	eventBus := make(chan Event)

	uppercaseHandler := httptransport.NewServer(
		uppercaseEndpoint,
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var s string
			if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
				return nil, err
			}
			eventBus <- Event{Type: "uppercase", Message: s}
			return s, nil
		},
		func(_ context.Context, w http.ResponseWriter, response interface{}) error {
			return json.NewEncoder(w).Encode(response)
		},
	)

	go func() {
		for {
			select {
			case event := <-eventBus:
				logger.Log("event", event.Type, "message", event.Message)
			}
		}
	}()

	http.Handle("/uppercase", uppercaseHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	http.ListenAndServe(":8080", nil)
}





// Logging with context in Go Kit service
package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := log.NewLogfmtLogger(log.StdlibWriter{})
	svc := StringService{}
	uppercaseEndpoint := makeUppercaseEndpoint(svc)

	uppercaseHandler := httptransport.NewServer(
		uppercaseEndpoint,
		func(ctx context.Context, r *http.Request) (interface{}, error) {
			logger := log.With(logger, "method", "uppercase")
			var s string
			if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
				logger.Log("error", err)
				return nil, err
			}
			logger.Log("msg", "received", "string", s)
			return s, nil
		},
		func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
			logger := log.With(logger, "method", "uppercase")
			logger.Log("msg", "sending response")
			return json.NewEncoder(w).Encode(response)
		},
	)

	http.Handle("/uppercase", uppercaseHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	http.ListenAndServe(":8080", nil)
}







// Implementing graceful shutdown with context in Go Kit service
package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := log.NewLogfmtLogger(log.StdlibWriter{})
	svc := StringService{}
	uppercaseEndpoint := makeUppercaseEndpoint(svc)

	uppercaseHandler := httptransport.NewServer(
		uppercaseEndpoint,
		func(ctx context.Context, r *http.Request) (interface{}, error) {
			var s string
			if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
				return nil, err
			}
			return s, nil
		},
		func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
			return json.NewEncoder(w).Encode(response)
		},
	)

	http.Handle("/uppercase", uppercaseHandler)

	go func() {
		logger.Log("msg", "HTTP", "addr", ":8080")
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			logger.Log("msg", "HTTP server stopped")
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	logger.Log("msg", "shutting down HTTP server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.Log("error", err)
	}
}






// Adding timeout middleware to Go Kit service
package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := log.NewLogfmtLogger(log.StdlibWriter{})
	svc := StringService{}
	uppercaseEndpoint := makeUppercaseEndpoint(svc)
	timeout := 3 * time.Second

	uppercaseEndpoint = func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			ctx, cancel := context.WithTimeout(ctx, timeout)
			defer cancel()
			return next(ctx, request)
		}
	}(uppercaseEndpoint)

	uppercaseHandler := httptransport.NewServer(
		uppercaseEndpoint,
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var s string
			if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
				return nil, err
			}
			return s, nil
		},
		func(_ context.Context, w http.ResponseWriter, response interface{}) error {
			return json.NewEncoder(w).Encode(response)
		},
		httptransport.ServerErrorEncoder(func(ctx context.Context, err error, w http.ResponseWriter) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}),
	)

	http.Handle("/uppercase", uppercaseHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	http.ListenAndServe(":8080", nil)
}







// Adding timeout middleware with default value to Go Kit service
package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

const defaultTimeout = 3 * time.Second

func main() {
	logger := log.NewLogfmtLogger(log.StdlibWriter{})
	svc := StringService{}
	uppercaseEndpoint := makeUppercaseEndpoint(svc)

	uppercaseEndpoint = func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			timeout, ok := ctx.Deadline()
			if !ok {
				ctx, _ = context.WithTimeout(ctx, defaultTimeout)
			} else {
				logger.Log("msg", "using request deadline as timeout", "deadline", timeout)
			}
			return next(ctx, request)
		}
	}(uppercaseEndpoint)

	uppercaseHandler := httptransport.NewServer(
		uppercaseEndpoint,
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var s string
			if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
				return nil, err
			}
			return s, nil
		},
		func(_ context.Context, w http.ResponseWriter, response interface{}) error {
			return json.NewEncoder(w).Encode(response)
		},
		httptransport.ServerErrorEncoder(func(ctx context.Context, err error, w http.ResponseWriter) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}),
	)

	http.Handle("/uppercase", uppercaseHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	http.ListenAndServe(":8080", nil)
}







// Adding context value middleware to Go Kit service
package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

type contextKey string

const userIDKey = contextKey("userID")

func main() {
	logger := log.NewLogfmtLogger(log.StdlibWriter{})
	svc := StringService{}
	uppercaseEndpoint := makeUppercaseEndpoint(svc)

	uppercaseEndpoint = func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			userID := ctx.Value(userIDKey).(string)
			logger.Log("userID", userID, "msg", "request received")
			defer logger.Log("userID", userID, "msg", "request completed")
			return next(ctx, request)
		}
	}(uppercaseEndpoint)

	uppercaseHandler := httptransport.NewServer(
		uppercaseEndpoint,
		func(ctx context.Context, r *http.Request) (interface{}, error) {
			userID := r.Header.Get("X-User-ID")
			ctx = context.WithValue(ctx, userIDKey, userID)
			var s string
			if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
				return nil, err
			}
			return s, nil
		},
		func(_ context.Context, w http.ResponseWriter, response interface{}) error {
			return json.NewEncoder(w).Encode(response)
		},
	)

	http.Handle("/uppercase", uppercaseHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	http.ListenAndServe(":8080", nil)
}




// Adding request ID middleware with UUID generation to Go Kit service
package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	uuid "github.com/satori/go.uuid"
)

type contextKey string

const requestIDKey = contextKey("requestID")

func main() {
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	svc := StringService{}
	uppercaseEndpoint := makeUppercaseEndpoint(svc)

	uppercaseEndpoint = func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			reqID := ctx.Value(requestIDKey).(string)
			logger.Log("requestID", reqID, "msg", "request received")
			defer logger.Log("requestID", reqID, "msg", "request completed")
			return next(ctx, request)
		}
	}(uppercaseEndpoint)

	uppercaseHandler := httptransport.NewServer(
		uppercaseEndpoint,
		func(ctx context.Context, r *http.Request) (interface{}, error) {
			reqID := uuid.NewV4().String()
			ctx = context.WithValue(ctx, requestIDKey, reqID)
			var s string
			if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
				return nil, err
			}
			return s, nil
		},
		func(_ context.Context, w http.ResponseWriter, response interface{}) error {
			return json.NewEncoder(w).Encode(response)
		},
	)

	http.Handle("/uppercase", uppercaseHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	http.ListenAndServe(":8080", nil)
}





// Adding structured logging middleware to Go Kit service
package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := log.NewLogfmtLogger(log.StdlibWriter{})
	svc := StringService{}
	uppercaseEndpoint := makeUppercaseEndpoint(svc)

	uppercaseEndpoint = func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			logger.Log("msg", "request received", "request", request)
			defer logger.Log("msg", "request completed")
			return next(ctx, request)
		}
	}(uppercaseEndpoint)

	uppercaseHandler := httptransport.NewServer(
		uppercaseEndpoint,
		func(ctx context.Context, r *http.Request) (interface{}, error) {
			var s string
			if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
				return nil, err
			}
			return s, nil
		},
		func(_ context.Context, w http.ResponseWriter, response interface{}) error {
			logger.Log("msg", "sending response", "response", response)
			return json.NewEncoder(w).Encode(response)
		},
	)

	http.Handle("/uppercase", uppercaseHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	http.ListenAndServe(":8080", nil)
}




