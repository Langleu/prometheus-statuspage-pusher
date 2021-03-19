# prometheus-statuspage-pusher

# Usage

```
Usage of ./prometheus-statuspage-pusher:
  -c string
    	Query config file (default "queries.yaml")
  -i duration
    	Metric push interval (default 30s)
  -pu string
    	URL of Prometheus API (default "http://localhost:9091/prometheus")
  -si string
    	Statuspage page ID
  -st string
    	Statuspage Oauth token (defaults to value of STATUSPAGE_TOKEN environment variable)
  -su string
    	URL of Statuspage API (default "https://api.statuspage.io")
  -prometheusPort int
      Port to serve Prometheus metrics from (default 9095)
```

## Config:

Syntax:

```
metricID: prometheus-expression
```

The prometheus-expression needs to return a single element vector, like:

```
abcdef: avg(up{job="web"})
```
