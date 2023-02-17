package simple_prom

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type metricsServer struct {
	Registry *prometheus.Registry
}

var Metrics = metricsServer{Registry: prometheus.NewRegistry()}

func init() {
	Metrics.Registry.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
	)
}

func (metrics metricsServer) NewCounter(opts prometheus.CounterOpts) prometheus.Counter {
	return promauto.With(metrics.Registry).NewCounter(opts)
}

func (metrics metricsServer) NewHistogram(opts prometheus.HistogramOpts) prometheus.Histogram {
	return promauto.With(metrics.Registry).NewHistogram(opts)
}
