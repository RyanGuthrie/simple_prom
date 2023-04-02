# simple_prom
Simple wrapper around the GoLang client for Prometheus

## Usage
go.mod

    require github.com/RyanGuthrie/simple_prom

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

## Releasing a new version

Run the `tag.sh` script

    $ ./tag v1.0.9

That'll ensure everything has been committed, and tag the current revision