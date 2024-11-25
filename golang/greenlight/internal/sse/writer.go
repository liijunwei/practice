package sse

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// write SSE headers and events
type sseWriter[T any] struct {
	// buffer is reused for encoding events to reduce allocations
	buffer *bytes.Buffer
	respW  http.ResponseWriter
}

func newEventWriter[T any](respW http.ResponseWriter) *sseWriter[T] {
	return &sseWriter[T]{
		buffer: &bytes.Buffer{},
		respW:  respW,
	}
}

// write required SSE headers
func writeHeader(respW http.ResponseWriter) {
	respW.Header().Set("Content-Type", "text/event-stream")
	respW.Header().Set("Cache-Control", "no-store")
	respW.Header().Set("X-Accel-Buffering", "no")
	respW.Header().Set("Connection", "keep-alive")
	respW.WriteHeader(http.StatusOK)
}

func writeMessage(respW http.ResponseWriter, buffer *bytes.Buffer, data any) error {
	dataB, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to encode data: %w", err)
	}

	event := &Event{
		Type: "data",
		Data: dataB,
	}

	return writeEvent(respW, buffer, event)
}

// write an error event to response
func writeError(respW http.ResponseWriter, buffer *bytes.Buffer, err error) error {
	data := []byte(err.Error())

	event := &Event{
		Type: "error",
		Data: data,
	}

	return writeEvent(respW, buffer, event)
}

// encode and write an event to response
func writeEvent(respW http.ResponseWriter, buffer *bytes.Buffer, event *Event) error {
	buffer.Reset()

	if err := event.encode(buffer); err != nil {
		return fmt.Errorf("failed to encode SSE event, event: %s: error: %w", event, err)
	}

	_, err := respW.Write(buffer.Bytes())
	if err != nil {
		return fmt.Errorf("failed to write SSE event to response, event: %s: error: %w: ", event, err)
	}

	if f, ok := respW.(http.Flusher); ok {
		f.Flush()
	}

	return nil
}
