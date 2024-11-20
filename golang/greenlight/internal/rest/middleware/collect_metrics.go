package middleware

import (
	"expvar"
	"net/http"
	"strconv"

	"github.com/felixge/httpsnoop"
)

func CollectMetrics(next http.Handler) http.Handler {
	// concurrent safe
	totalRequestsReceived := expvar.NewInt("total_requestes_received")
	totalResponseSent := expvar.NewInt("total_response_sent")
	totalProcessingTimeInMilliseconds := expvar.NewInt("total_processing_time_ms")
	totalResponseSentByStatus := expvar.NewMap("total_response_sent_by_status")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		totalRequestsReceived.Add(1)

		metrics := httpsnoop.CaptureMetrics(next, w, r)

		totalResponseSent.Add(1)
		totalProcessingTimeInMilliseconds.Add(metrics.Duration.Milliseconds())
		totalResponseSentByStatus.Add(strconv.Itoa(metrics.Code), 1)
	})
}
