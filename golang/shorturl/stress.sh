#!/bin/bash
# hey -n 300 -c 100 -m POST -d '{"original": "https://example.com/$(gdate +%s%N)"}' -H "Content-Type: application/json" http://localhost:8080/shorturl
# hey -n 300 -c 100 -m POST -d "{\"original\": \"https://example.com/$(gdate +%s%N)-${RANDOM}\"}" -H "Content-Type: application/json" http://localhost:8080/shorturl

# for i in $(seq 1 3000); do
#     payload=$(printf '{"original": "https://example.com/%s-%s"}' "$(gdate +%s%N)" "$RANDOM")
#     # 这里每个 hey 实例只发送 1 个请求
#     hey -n 1 -c 1 -m POST -d "${payload}" -H "Content-Type: application/json" http://localhost:8080/shorturl &
# done
# wait

# TODO try better approach with hey
# or use wrk+lua


wrk -t5 -c10 -d10s -s ./stress.lua --latency http://localhost:8080/shorturl
