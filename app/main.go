package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ServiceRequests struct {
	Url string `json:"url,omitempty"`
}

type MetricsResponse struct {
	Http_get       http_get `json:"http_get,omitempty"`
	RequestCounter int      `json:"request_counter,omitempty"`
}

type http_get struct {
	Url  string `json:"url,omitempty"`
	Code int    `json:"code,omitempty"`
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {}

// requestHandler middleware
func requestHandler(w http.ResponseWriter, r *http.Request) {

	var req ServiceRequests

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Fprintf(w, "Error decoding request: %v", err)
		return
	}

	http.NewRequest("POST", req.Url, nil)

}

func main() {

	// Listen and monitor for user requests
	server := http.NewServeMux()
	server.HandleFunc("/", requestHandler)

	// Listen and response for user metrics requests
	metrics := http.NewServeMux()
	metrics.HandleFunc("/metrics", metricsHandler)

	// golang light wieght thread for user requests
	go func() {
		http.ListenAndServe("localhost:8080", server)
	}()

	http.ListenAndServe("localhost:9095", metrics)
}
