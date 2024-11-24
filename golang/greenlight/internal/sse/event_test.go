package sse

import (
	"bytes"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestEvent(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		event      Event
		expectData string
		expectErr  any
	}{
		{
			name: "happy path: type only",
			event: Event{
				Type: "testevent",
			},
			expectData: "event:testevent\n\n",
			expectErr:  nil,
		},
		{
			name: "happy path: data only",
			event: Event{
				Data: []byte("testdata"),
			},
			expectData: "data:testdata\n\n",
			expectErr:  nil,
		},
		{
			name: "happy path: comment only",
			event: Event{
				Comment: "comment",
			},
			expectData: ":comment\n\n",
			expectErr:  nil,
		},
		{
			name:       "happy path: only retry",
			event:      Event{Retry: time.Second},
			expectData: "retry:1000\n\n",
			expectErr:  nil,
		},
		{
			name: "happy path: type and data",
			event: Event{
				Type: "testevent",
				Data: []byte("testdata"),
			},
			expectData: "event:testevent\ndata:testdata\n\n",
			expectErr:  nil,
		},
		{
			name: "happy path: all field",
			event: Event{
				Type:    "testevent",
				ID:      "1",
				Comment: "comment",
				Data:    []byte("testdata"),
				Retry:   time.Millisecond * 20,
			},
			expectData: "event:testevent\ndata:testdata\nid:1\nretry:20\n:comment\n\n",
			expectErr:  nil,
		},
		{
			name:       "unhappy path: only id",
			event:      Event{ID: "123"},
			expectData: "",
			expectErr:  &InvalidEventError{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			buf := &bytes.Buffer{}
			err := tt.event.encode(buf)

			if tt.expectErr != nil {
				require.ErrorAs(t, err, &tt.expectErr)
			}

			if tt.expectErr == nil {
				require.NoError(t, err)
				require.Equal(t, tt.expectData, buf.String())
			}
		})
	}
}

func BenchmarkEventEncode(b *testing.B) {
	event := Event{
		Type: "msg",
		Data: bytes.Repeat([]byte("a"), 1024),
	}

	buf := &bytes.Buffer{}

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = event.encode(buf)
		buf.Reset()
	}

	runtime.KeepAlive(buf)
}
