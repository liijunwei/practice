package sse

import (
	"bytes"
	"encoding/json"
	"greenlight/internal/assert"
)

type Event struct {
	Type string `json:"type,omitempty"`
	Data []byte `json:"data,omitempty"`
}

func (e *Event) encode(buf *bytes.Buffer) error {
	if e.Type != "" {
		buf.WriteString("event:")
		buf.WriteString(e.Type)
		buf.WriteByte('\n')
	}

	if len(e.Data) > 0 {
		buf.WriteString("data:")
		buf.Write(e.Data)
		buf.WriteByte('\n')
	}

	buf.WriteByte('\n')

	if e.Type == "" && len(e.Data) == 0 {
		return &InvalidEventError{msg: "at least one non zero field required"}
	}

	return nil
}

func (e *Event) String() string {
	data, err := json.Marshal(e)
	assert.NoError(err, "failed to dump SSE event to json") // won't happen in most of the cases

	return string(data)
}
