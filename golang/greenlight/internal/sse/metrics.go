package sse

import (
	"expvar"
)

type metrics struct {
	totalConnections *expvar.Int
}

func newMetrics() metrics {
	metrics := metrics{
		totalConnections: expvar.NewInt("total_active_sse_connections"),
	}

	return metrics
}

func (m metrics) record() func() {
	m.totalConnections.Add(1)

	return func() {
		m.totalConnections.Add(-1)
	}
}
