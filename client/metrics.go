package client

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	Counter *prometheus.CounterVec
}

func InitPrometheusMetrics() *Metrics {
	var m Metrics
	counter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_get",
			Help: "Number of requests submitted by user",
		}, []string{"url", "code"},
	)

	// create metric counter
	m.Counter = counter
	// prometheus.MustRegister(m.Counter)
	return &m
}
