package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/angelopana/phaidra/client"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type ServiceRequests struct {
	Url string `json:"url,omitempty"`
}

type Managers struct {
	Metrics *client.Metrics
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

	//Create a new request from given user body data
	status, err := http.NewRequest(methodType, req.Url, nil)
	if err != nil {
		fmt.Fprintf(w, "Error in %v request to %v", r.Method, req.Url)
		return
	}

	//Client request manager
	client := &http.Client{}
	// send request
	resp, _ := client.Do(status)
	code := strconv.Itoa(resp.StatusCode) //get returned status code

	// increment each user requests and set values
	m.Metrics.Counter.With(prometheus.Labels{
		"url":  req.Url,
		"code": code,
	}).Inc()

}

func main() {

	customeMetric := client.InitPrometheusMetrics()

	prometheusRegistry := prometheus.NewRegistry()
	prometheusRegistry.Register(customeMetric.Counter)

	prometheus.MustRegister()
	managers := &Managers{
		Metrics: customeMetric,
	}

	// Listen and monitor for user requests
	server := http.NewServeMux()
	server.HandleFunc("/", managers.requestHandler)

	// golang light wieght thread for user requests
	go func() {
		http.ListenAndServe("0.0.0.0:8080", server)

	}()

	http.ListenAndServe("0.0.0.0:9095", promhttp.HandlerFor(prometheusRegistry, promhttp.HandlerOpts{Registry: prometheusRegistry}))

	fmt.Println("SERVICE DEPLOYED...")

}
