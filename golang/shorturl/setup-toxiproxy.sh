#!/bin/bash

# toxiproxy-server # in another shell

function setup() {
  toxiproxy-cli create -l localhost:28080 -u localhost:8080 http_proxy
  toxiproxy-cli toxic add -t latency -a latency=40 -a jitter=20 http_proxy
  toxiproxy-cli list
  echo "change jmeter to use localhost:28080"
}

function cleanup() {
  toxiproxy-cli delete http_proxy
  toxiproxy-cli list
}

if [[ "" = "$1" ]]; then
  setup
else
  cleanup
fi
