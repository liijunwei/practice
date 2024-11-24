package sse

import (
	"bytes"
	"encoding/json"
	"greenlight/internal/assert"
	"strconv"
	"time"
)

type Event struct {
	Type    string        `json:"type,omitempty"`
	ID      string        `json:"id,omitempty"`
	Comment string        `json:"comment,omitempty"`
	Data    []byte        `json:"data,omitempty"`
	Retry   time.Duration `json:"retry,omitempty"`
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

	if len(e.ID) > 0 {
		buf.WriteString("id:")
		buf.WriteString(e.ID)
		buf.WriteByte('\n')
	}

	if e.Retry > 0 {
		buf.WriteString("retry:")
		buf.WriteString(strconv.FormatInt(e.Retry.Milliseconds(), 10))
		buf.WriteByte('\n')
	}

	if e.Comment != "" {
		buf.WriteString(":")
		buf.WriteString(e.Comment)
		buf.WriteByte('\n')
	}

	buf.WriteByte('\n')

	if e.Type == "" && len(e.Data) == 0 && e.Comment == "" && e.Retry == 0 {
		return &InvalidEventError{msg: "at least one non zero field required"}
	}

	return nil
}

func (e *Event) String() string {
	data, err := json.Marshal(e)
	assert.NoError(err, "failed to dump SSE event to json") // won't happen in most of the cases

	return string(data)
}
