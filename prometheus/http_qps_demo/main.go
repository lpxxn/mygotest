package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func init() {
	prometheus.MustRegister(httpRequestCount)
	prometheus.MustRegister(httpRequestDuration)
}

func main() {
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().UnixNano() / int64(time.Millisecond))
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/test", handler)
	if err := http.ListenAndServe(":5112", &myService{}); err != nil {
		panic(err)
	}
}

var httpRequestCount = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_request_count",
		Help: "http request count",
	},
	[]string{"endpoint", "contentLength"},
)

var httpRequestDuration = prometheus.NewSummaryVec(
	prometheus.SummaryOpts{
		Name: "http_request_duration",
		Help: "http request duration",
	},
	[]string{"endpoint"},
)

type myService struct {
}

func (*myService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	path := r.URL.Path
	httpRequestCount.WithLabelValues(path, strconv.FormatInt(r.ContentLength, 10)).Inc()
	httpRequestCount.WithLabelValues("total", "0").Inc()
	http.DefaultServeMux.ServeHTTP(w, r)

	elapsed := (float64)(time.Since(start) / time.Millisecond)
	httpRequestDuration.WithLabelValues(path).Observe(elapsed)
}

func handler(w http.ResponseWriter, r *http.Request) {
	n := rand.Intn(100)
	if n >= 95 {
		time.Sleep(100 * time.Millisecond)
	} else {
		time.Sleep(50 * time.Millisecond)
	}

	fmt.Fprint(w, "hello world")
}
