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


# ./stress.sh
# ./stress.sh 1 1
# ./stress.sh 1 1 30

# for sqlite, 1 connection is enough
wrk -t ${1:-1} -c ${2:-1} -d ${3:-10s} -s ./stress.lua --latency http://localhost:8080/shorturl

sqlite3 /tmp/shorturl-app.db "PRAGMA wal_checkpoint(FULL);"
sqlite3 /tmp/shorturl-app.db "PRAGMA integrity_check;"

# prepare datasource
sqlite3 /tmp/shorturl-app-backup.db <<EOF
.mode csv
.output /tmp/shorturls.csv
SELECT * FROM shorturls;
.output stdout
EOF
