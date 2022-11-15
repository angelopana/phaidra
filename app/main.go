package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/angelopana/phaidra/client"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	_ "github.com/prometheus/client_golang/prometheus/promhttp"
)

type ServiceRequests struct {
	Url string `json:"url,omitempty"`
}

type http_get struct {
	Url  string `json:"url,omitempty"`
	Code int    `json:"code,omitempty"`
}

type Managers struct {
	Metrics *client.Metrics
}

// metricsHandler
func (m *Managers) metricsHandler(w http.ResponseWriter, r *http.Request) {

	promhttp.Handler()
	fmt.Fprintf(w, "HEEE")
}

// requestHandler middleware
func (m *Managers) requestHandler(w http.ResponseWriter, r *http.Request) {

	var req ServiceRequests

	// Get method type
	methodType := r.Method

	// Decode body of user request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Fprintf(w, "Error decoding request: %v", err)
		return
	}
	status, err := http.NewRequest(methodType, req.Url, nil)
	if err != nil {
		fmt.Fprintf(w, "Error in %v request to %v", r.Method, req.Url)
		return
	}

	client := &http.Client{}
	// send request
	resp, _ := client.Do(status)
	// get status code
	code := strconv.Itoa(resp.StatusCode)

	// increment each user requests and set values
	m.Metrics.Counter.With(prometheus.Labels{
		"url":  req.Url,
		"code": code,
	}).Inc()

	fmt.Fprintf(w, "%v", resp.StatusCode)

}

func main() {

	pom := client.InitPrometheusMetrics()
	prometheus.MustRegister(pom.Counter)

	managers := &Managers{
		Metrics: pom,
	}

	// Listen and monitor for user requests
	server := http.NewServeMux()
	server.HandleFunc("/", managers.requestHandler)

	// Listen and response for user metrics requests
	// metrics := http.NewServeMux()
	// metrics.HandleFunc("/metrics", managers.metricsHandler)

	// golang light wieght thread for user requests
	go func() {
		http.ListenAndServe("localhost:8080", server)

	}()

	http.ListenAndServe(":9095", managers.metricsHandler)

}
