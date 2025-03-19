### get started

simple shorturl app with sqlite3+sqlc

features:
1. create short url
2. redirect to original url
3. analyse short url access
4. test single sqlite3 db max qps(HOW?)

go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
guide: https://docs.sqlc.dev/en/latest/tutorials/getting-started-sqlite.html
metrics: https://prometheus.io/docs/tutorials/instrumenting_http_server_in_go

```sh
rm /tmp/shorturl-app.db; sqlc generate && go run .
```

### shorturl stat - create

- golang + sqlite3(1 connection)
- localhost(toxiproxy applies 40(-20,+20)ms latency)
    - jmeter run 100 threads for 20s
    - throughput: 2230 request/s (create new shorturl)
    - p95 latency: 59 ms

- localhost
    - jmeter run 100 threads for 20s
    - throughput: 4462 request/s (create new shorturl)
    - p95 latency: 65 ms
    - p50 latency: 13 ms

### shorturl stat - visit & redirect

- golang + sqlite3(1 connection)
- localhost(toxiproxy applies 40(-20,+20)ms latency)
    - jmeter run 100 threads for 20s
    - throughput: 258 request/s (visit shorturl)
    - p95 latency: 569 ms
    - p50 latency: 280 ms

- localhost
    - jmeter run 100 threads for 20s
    - throughput: 296 request/s (visit shorturl)
    - p95 latency: 520 ms
    - p50 latency: 235 ms

too slow, possible reasons:
- visiting example.com too slow
    - redirect to `http://localhost:28080/metrics` with 40(-20,+20)ms latency
    -
    - throughput: 448 request/s (visit shorturl)
    - p95 latency: 41 ms
    - p50 latency: 59 ms
    -
    - latency improvement is obvious, but not the throughput

- db connection pool too small
    - pool: 1->100
    -
    - throughput: 332 request/s (visit shorturl)
    - p95 latency: 306 ms
    - p50 latency: 265 ms
    -
    - improvement is quite limited

- separate read and write looks ok, yet still limited

### TODOs

- TODO try replace check existense with redis bloom filter to see whether it improves create performance
- TODO think about how to scale the app with go + sqlite
- TODO try replace sqlite with postgres
- TODO practice more about prometheus+grafana and report custom metrics
