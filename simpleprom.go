package simple_prom

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type metricsServer struct {
	// HTTP Handler for exposing Prometheus Metrics (recommended to bind to the /metrics endpoint)
	Handler http.Handler

	registry *prometheus.Registry
}

// Metrics - container used for creating counters, histograms, gauges, etc...
var Metrics = metricsServer{registry: prometheus.NewRegistry()}

func init() {
	Metrics.registry.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
	)

	Metrics.Handler = promhttp.HandlerFor(Metrics.registry, promhttp.HandlerOpts{Registry: Metrics.registry})
}

func (metrics metricsServer) NewCounter(opts prometheus.CounterOpts) prometheus.Counter {
	return promauto.With(metrics.registry).NewCounter(opts)
}

func (metrics metricsServer) NewHistogram(opts prometheus.HistogramOpts) prometheus.Histogram {
	return promauto.With(metrics.registry).NewHistogram(opts)
}

func (metrics metricsServer) NewSummary(opts prometheus.SummaryOpts) prometheus.Summary {
	return promauto.With(metrics.registry).NewSummary(opts)
}
