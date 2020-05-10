package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"net/http/pprof"
	_ "net/http/pprof"
	"strconv"
	"time"
)
var requestCount = prometheus.NewCounter(prometheus.CounterOpts{Name: "app_request_count"})
var httpRequestHistogram = prometheus.NewHistogram(prometheus.HistogramOpts{Buckets:[]float64{0.05, 0.1 , 0.15, 0.2, 0.25,0.3,0.5},Name:"app_response_duration_miliseconds"})

func main() {
	prometheus.MustRegister(requestCount,httpRequestHistogram)
	muxer := mux.NewRouter()
	muxer.PathPrefix("/debug/pprof/").HandlerFunc(pprof.Index)
	muxer.Handle("/metrics",promhttp.Handler())
	muxer.HandleFunc("/raj", func(writer http.ResponseWriter, request *http.Request) {
		tt := time.Now()
		requestCount.Inc()
		writer.WriteHeader(200)
		writer.Write([]byte(getFibonacci(request.URL.Query().Get("fib"))))
		tttt := float64(time.Since(tt).Milliseconds())
		defer httpRequestHistogram.Observe(tttt)
		return

	})
	http.ListenAndServe(":8081",muxer)
}

func getFibonacci(s string )  string {
	fmt.Println("request recieved")
     ii , _ := strconv.Atoi(s)
     return strconv.Itoa(getFib(ii))

}

func getFib(i int) int   {
	if i < 2 {
		return i
	}
	return getFib(i-1)+getFib(i-2 )
}