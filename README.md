# prometheus-statuspage-pusher

A fork of [beekpr/prometheus-statuspage-pusher](https://github.com/beekpr/prometheus-statuspage-pusher) with minor enhancements and focus on components instead of metrics.

The difference compared to other forks is the focus on updating components rather than system metrics.

## Usage

```
Usage of ./prometheus-statuspage-pusher:
  -prom string
    	URL of Prometheus server (default: "http://localhost:9090")
  -apikey string
    	Statuspage API key
  -pageid string
    	Statuspage page ID
  -config string
    	Query config file (default: "queries.yaml")
  -interval string
    	Metric push interval (default: 300s / 5m)
  -loglevel string
    	Log level accepted by Logrus, for example, "error", "warn", "info", "debug", ... (default: "info")

Alternatively, you can use environment variables instead, e.g. PROM, APIKEY, PAGEID, ...
```

## Config:

Syntax:

```
componentID: prometheus-expression
```

The prometheus-expression needs to return a single element vector, like:
Where 1 = "operationa", else "major outage"

```
abcdef: avg(up{job="web"})
```
