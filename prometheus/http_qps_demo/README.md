
```
# The job name is added as a label `job=<job_name>` to any timeseries scraped from this config.
  - job_name: 'go-api-metrics-test'
    metrics_path: "/metrics"

    scrape_interval: 5s

    static_configs:
      - targets: ['http://127.0.0.1:5112']
copy

```


```
http_request_count{contentLength="0",endpoint="total",instance="192.168.1.27:5112",job="go-api-metrics-test"}
```


```
http_request_duration{endpoint="/test",instance="192.168.1.27:5112",job="go-api-metrics-test",quantile="0.99"}
```

