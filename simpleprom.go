package simple_prom

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

type metricsServer struct {
	// HTTP Handler for exposing Prometheus Metrics (recommended to bind to the /metrics endpoint)
	Handler *http.Handler

	registry *prometheus.Registry
	factory  *promauto.Factory
}

// Metrics - container used for creating counters, histograms, gauges, etc...
var Metrics = metricsServer{registry: prometheus.NewRegistry()}

func init() {
	Metrics.registry.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
	)

	factory := promauto.With(Metrics.registry)
	Metrics.factory = &factory

	metricsHandler := promhttp.HandlerFor(Metrics.registry, promhttp.HandlerOpts{Registry: Metrics.registry})
	Metrics.Handler = &metricsHandler
}

func (metrics metricsServer) NewCounter(opts prometheus.CounterOpts) prometheus.Counter {
	return metrics.factory.NewCounter(opts)
}

func (metrics metricsServer) NewHistogram(opts prometheus.HistogramOpts) prometheus.Histogram {
	return metrics.factory.NewHistogram(opts)
}
