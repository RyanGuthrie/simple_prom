package simple_prom

import (
	"log"
	"net/http"
	"time"

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

func (metrics metricsServer) NewGauge(opts prometheus.GaugeOpts) prometheus.Gauge {
	return promauto.With(metrics.registry).NewGauge(opts)
}

// NewGaugeFunc creates a gauge and every 60 seconds, queries the `f` func for the current value to publish.  This is good for applying long-lived
// gauges that can be updated via an isolated function.  The underlying Go Routine can be stopped by closing the `chan` returned
func (metrics metricsServer) NewGaugeFunc(opts prometheus.GaugeOpts, updateFrequency time.Duration, f func() float64) chan bool {
	gauge := promauto.With(metrics.registry).NewGauge(opts)

	closeChan := make(chan bool)

	go func() {
		for {
			select {
			case <-closeChan:
				log.Printf("Shutting down gauge for %v\n", gauge)
				break
			default:
				gauge.Set(f())
				time.Sleep(updateFrequency)
			}
		}
	}()

	return closeChan
}
