package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	labels = []string{"route", "method"}

	httpRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_request_count",
			Help: "Total HTTP Requests",
		},
		labels,
	)
	httpDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration",
			Help:    "HTTP Request Duration Distribution",
			Buckets: prometheus.ExponentialBuckets(1, 10, 4),
		},
		labels,
	)
	httpSummary = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "http_request_summary",
			Help:       "HTTP Request Summary",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.95: 0.005, 0.99: 0.001},
		},
		labels,
	)
)

func main() {

	rand.Seed(time.Now().UnixNano())

	prometheus.MustRegister(httpRequests)
	prometheus.MustRegister(httpDuration)
	prometheus.MustRegister(httpSummary)

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/index", IndexHandler)
	r.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(":8080", r))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	n := rand.Intn(1000) // n will be between 0 and 10
	time.Sleep(time.Duration(n) * time.Millisecond)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome")

	end := time.Since(start).Milliseconds()

	httpDuration.With(prometheus.Labels{"route": r.RequestURI, "method": r.Method}).Observe(float64(end))
	httpSummary.With(prometheus.Labels{"route": r.RequestURI, "method": r.Method}).Observe(float64(end))
	httpRequests.With(prometheus.Labels{"route": r.RequestURI, "method": r.Method}).Inc()
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	n := rand.Intn(1000)
	time.Sleep(time.Duration(n) * time.Millisecond)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Index")

	end := time.Since(start).Milliseconds()

	httpDuration.With(prometheus.Labels{"route": r.RequestURI, "method": r.Method}).Observe(float64(end))
	httpSummary.With(prometheus.Labels{"route": r.RequestURI, "method": r.Method}).Observe(float64(end))
	httpRequests.With(prometheus.Labels{"route": r.RequestURI, "method": r.Method}).Inc()
}
