// Import the Prometheus package
package main

import (
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)



// Create a Counter metric
package main

import (
    "github.com/prometheus/client_golang/prometheus"
)

var requestsTotal = prometheus.NewCounter(prometheus.CounterOpts{
    Name: "requests_total",
    Help: "Total number of requests.",
})




// Register a Counter metric
package main

import (
    "github.com/prometheus/client_golang/prometheus"
)

func init() {
    prometheus.MustRegister(requestsTotal)
}





// Increment a Counter metric
package main

import (
    "github.com/prometheus/client_golang/prometheus"
)

func main() {
    requestsTotal.Inc()
}





// Create a Gauge metric
package main

import (
    "github.com/prometheus/client_golang/prometheus"
)

var temperatureGauge = prometheus.NewGauge(prometheus.GaugeOpts{
    Name: "temperature",
    Help: "Current temperature.",
})





// Set a Gauge metric
package main

import (
    "github.com/prometheus/client_golang/prometheus"
)

func main() {
    temperatureGauge.Set(23.5)
}





// Increment a Gauge metric
package main

import (
    "github.com/prometheus/client_golang/prometheus"
)

func main() {
    temperatureGauge.Inc()
}





// Decrement a Gauge metric
package main

import (
    "github.com/prometheus/client_golang/prometheus"
)

func main() {
    temperatureGauge.Dec()
}





// Create a Histogram metric
package main

import (
    "github.com/prometheus/client_golang/prometheus"
)

var requestDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
    Name:    "request_duration_seconds",
    Help:    "Histogram of request duration.",
    Buckets: prometheus.DefBuckets,
})





// Observe a value in Histogram metric
package main

import (
    "github.com/prometheus/client_golang/prometheus"
)

func main() {
    requestDuration.Observe(0.23)
}





// Create a Summary metric
package main

import (
    "github.com/prometheus/client_golang/prometheus"
)

var responseSize = prometheus.NewSummary(prometheus.SummaryOpts{
    Name: "response_size_bytes",
    Help: "Summary of response sizes.",
})





// Observe a value in Summary metric
package main

import (
    "github.com/prometheus/client_golang/prometheus"
)

func main() {
    responseSize.Observe(512)
}





// Create a CounterVec metric
package main

import (
    "github.com/prometheus/client_golang/prometheus"
)

var requestsTotalVec = prometheus.NewCounterVec(
    prometheus.CounterOpts{
        Name: "requests_total",
        Help: "Total number of requests.",
    },
    []string{"method", "endpoint"},
)





// Register a CounterVec metric
package main

import (
    "github.com/prometheus/client_golang/prometheus"
)

func init() {
    prometheus.MustRegister(requestsTotalVec)
}






// Increment a CounterVec metric
package main

import (
    "github.com/prometheus/client_golang/prometheus"
)

func main() {
    requestsTotalVec.WithLabelValues("GET", "/home").Inc()
}





// Create a GaugeVec metric
package main

import (
    "github.com/prometheus/client_golang/prometheus"
)

var temperatureGaugeVec = prometheus.NewGaugeVec(
    prometheus.GaugeOpts{
        Name: "temperature",
        Help: "Current temperature.",
    },
    []string{"location"},
)





// Register a GaugeVec metric
package main

import (
    "github.com/prometheus/client_golang/prometheus"
)

func init() {
    prometheus.MustRegister(temperatureGaugeVec)
}





// Set a value in GaugeVec metric
package main

import (
    "github.com/prometheus/client_golang/prometheus"
)

func main() {
    temperatureGaugeVec.WithLabelValues("server_room").Set(22.3)
}





// Create a HistogramVec metric
package main

import (
    "github.com/prometheus/client_golang/prometheus"
)

var requestDurationVec = prometheus.NewHistogramVec(
    prometheus.HistogramOpts{
        Name:    "request_duration_seconds",
        Help:    "Histogram of request duration.",
        Buckets: prometheus.DefBuckets,
    },
    []string{"method", "endpoint"},
)





// Register a HistogramVec metric
package main

import (
    "github.com/prometheus/client_golang/prometheus"
)

func init() {
    prometheus.MustRegister(requestDurationVec)
}





// Observe a value in HistogramVec metric
package main

import (
    "github.com/prometheus/client_golang/prometheus"
)

func main() {
    requestDurationVec.WithLabelValues("GET", "/home").Observe(0.34)
}





// Create a SummaryVec metric
package main

import (
    "github.com/prometheus/client_golang/prometheus"
)

var responseSizeVec = prometheus.NewSummaryVec(
    prometheus.SummaryOpts{
        Name: "response_size_bytes",
        Help: "Summary of response sizes.",
    },
    []string{"endpoint"},
)





// Register a SummaryVec metric
package main

import (
    "github.com/prometheus/client_golang/prometheus"
)

func init() {
    prometheus.MustRegister(responseSizeVec)
}





// Observe a value in SummaryVec metric
package main

import (
    "github.com/prometheus/client_golang/prometheus"
)

func main() {
    responseSizeVec.WithLabelValues("/home").Observe(420)
}





// Create a custom Collector
package main

import (
    "github.com/prometheus/client_golang/prometheus"
)

type MyCollector struct {
    fooMetric *prometheus.Desc
}

func NewMyCollector() *MyCollector {
    return &MyCollector{
        fooMetric: prometheus.NewDesc("foo_metric", "Description of foo metric", nil, nil),
    }
}

func (collector *MyCollector) Describe(ch chan<- *prometheus.Desc) {
    ch <- collector.fooMetric
}

func (collector *MyCollector) Collect(ch chan<- prometheus.Metric) {
    ch <- prometheus.MustNewConstMetric(collector.fooMetric, prometheus.GaugeValue, 1.0)
}





// Register a custom Collector
package main

import (
    "github.com/prometheus/client_golang/prometheus"
)

func init() {
    prometheus.MustRegister(NewMyCollector())
}





// Create an HTTP handler for metrics
package main

import (
    "net/http"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
    http.Handle("/metrics", promhttp.Handler())
    http.ListenAndServe(":8080", nil)
}






// Start an HTTP server to expose metrics
package main

import (
    "log"
    "net/http"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
    http.Handle("/metrics", promhttp.Handler())
    log.Fatal(http.ListenAndServe(":8080", nil))
}






// Create a constant Counter metric
package main

import (
    "github.com/prometheus/client_golang/prometheus"
)

var constCounter = prometheus.NewCounter(prometheus.CounterOpts{
    Name: "const_counter",
    Help: "A counter that is set to a constant value.",
})





// Set a constant Counter metric
package main

import (
    "github.com/prometheus/client_golang/prometheus"
)

func main() {
    constCounter.Add(100)
}





// Use PushGateway to push metrics
package main

import (
    "log"
    "github.com/prometheus/client_golang/prometheus/push"
    "github.com/prometheus/client_golang/prometheus"
)

var requestsTotal = prometheus.NewCounter(prometheus.CounterOpts{
    Name: "requests_total",
    Help: "Total number of requests.",
})

func main() {
    prometheus.MustRegister(requestsTotal)
    pusher := push.New("http://localhost:9091", "my_job").Collector(requestsTotal)
    if err := pusher.Push(); err != nil {
        log.Fatal("Could not push to PushGateway:", err)
    }
}





// Create a CounterFunc metric
package main

import (
    "runtime"
    "github.com/prometheus/client_golang/prometheus"
)

var goRoutineCounter = prometheus.NewCounterFunc(prometheus.CounterOpts{
    Name: "go_goroutines",
    Help: "Number of goroutines.",
}, func() float64 {
    return float64(runtime.NumGoroutine())
})





// Register a CounterFunc metric
package main

import (
    "github.com/prometheus/client_golang/prometheus"
)

func init() {
    prometheus.MustRegister(goRoutineCounter)
}





// Create a GaugeFunc metric
package main

import (
    "runtime"
    "github.com/prometheus/client_golang/prometheus"
)

var memoryUsageGauge = prometheus.NewGaugeFunc(prometheus.GaugeOpts{
    Name: "memory_usage_bytes",
    Help: "Memory usage in bytes.",
}, func() float64 {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    return float64(m.Alloc)
})





// Register a GaugeFunc metric
package main

import (
    "github.com/prometheus/client_golang/prometheus"
)

func init() {
    prometheus.MustRegister(memoryUsageGauge)
}





// Create a Summary with Objectives
package main

import (
    "github.com/prometheus/client_golang/prometheus"
)

var latencySummary = prometheus.NewSummary(prometheus.SummaryOpts{
    Name:       "latency_seconds",
    Help:       "Summary of latencies.",
    Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
})





// Observe a value in Summary with Objectives
package main

import (
    "github.com/prometheus/client_golang/prometheus"
)

func main() {
    latencySummary.Observe(1.2)
}





// Create a CounterVec with dynamic labels
package main

import (
    "github.com/prometheus/client_golang/prometheus"
)

var errorsCounterVec = prometheus.NewCounterVec(
    prometheus.CounterOpts{
        Name: "errors_total",
        Help: "Total number of errors.",
    },
    []string{"type"},
)





// Register a CounterVec with dynamic labels
package main

import (
    "github.com/prometheus/client_golang/prometheus"
)

func init() {
    prometheus.MustRegister(errorsCounterVec)
}





// Increment a CounterVec with dynamic labels
package main

import (
    "github.com/prometheus/client_golang/prometheus"
)

func main() {
    errorsCounterVec.WithLabelValues("network").Inc()
}


