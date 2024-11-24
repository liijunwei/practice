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
func (w *sseWriter[T]) writeHeader() {
	w.respW.Header().Set("Content-Type", "text/event-stream")
	w.respW.Header().Set("Cache-Control", "no-store")
	w.respW.Header().Set("X-Accel-Buffering", "no")
	w.respW.Header().Set("Connection", "keep-alive")
	w.respW.WriteHeader(http.StatusOK)
}

func (w *sseWriter[T]) writeMessage(data T) error {
	dataB, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to encode data: %w", err)
	}

	event := &Event{
		Type: "data",
		Data: dataB,
	}

	return w.writeEvent(event)
}

// write an error event to response
func (w *sseWriter[T]) writeError(err error) error {
	dataB, encodeErr := []byte(err.Error()), error(nil)

	if encodeErr != nil {
		return fmt.Errorf("error encode data:%w", err)
	}

	event := &Event{
		Type: "error", // TODO check further for SSE spec
		Data: dataB,
	}

	return w.writeEvent(event)
}

// encode and write an event to response
func (w *sseWriter[T]) writeEvent(event *Event) error {
	w.buffer.Reset()

	if err := event.encode(w.buffer); err != nil {
		return fmt.Errorf("failed to encode SSE event, event: %s: error: %w", event, err)
	}

	_, err := w.respW.Write(w.buffer.Bytes())
	if err != nil {
		return fmt.Errorf("failed to write SSE event to response, event: %s: error: %w: %w", event, err)
	}

	if f, ok := w.respW.(http.Flusher); ok {
		f.Flush()
	}

	return nil
}
