# simple_prom
Simple wrapper around the GoLang client for Prometheus

## Usage
go.mod

    require github.com/RyanGuthrie/simple_prom latest

### Exposing `/metrics` endpoint

Add the `simple_prom.Metrics.Handler` handler to the HTTP server, eg:

    mux.Handle("/metrics", simple_prom.Metrics.Handler)

### Creating new metrics

Creating new metrics, for example a counter:

	requestCounter := simple_prom.Metrics.NewCounter(prometheus.CounterOpts{
		Namespace: "http",
		Subsystem: "overall",
		Name:      "requests",
		Help:      "Number of HTTP requests made overall across all endpoints",
	})
