package sse

import (
	"expvar"
)

// global variable
var sseMetrics metrics = newMetrics()

type metrics struct {
	totalConnections      *expvar.Int
	namedTotalConnections *expvar.Map
}

func newMetrics() metrics {
	metrics := metrics{
		totalConnections:      expvar.NewInt("total_active_sse_connections"),
		namedTotalConnections: expvar.NewMap("named_active_sse_connections"),
	}

	return metrics
}

func (m metrics) record(sseName string) func() {
	m.totalConnections.Add(1)
	m.namedTotalConnections.Add(sseName, 1)

	return func() {
		m.totalConnections.Add(-1)
		m.namedTotalConnections.Add(sseName, -1)
	}
}
