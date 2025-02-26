

### shorturl stat

- golang + sqlite3(1 connection)
- localhost(toxiproxy applies 40(-20,+20)ms latency)
    - jmeter run 100 threads for 20s
    - throughput: 2230 request/s (create new shorturl)
    - p95 latency: 59ms

- localhost
    - jmeter run 100 threads for 20s
    - throughput: 4462 request/s (create new shorturl)
    - p95 latency: 65ms
    - p50 latency: 13ms
