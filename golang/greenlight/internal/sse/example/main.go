package main

import (
	"cmp"
	"errors"
	"expvar"
	"net/http"
	"os"
	"time"

	"greenlight/internal/common"
	"greenlight/internal/sse"

	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.New(os.Stdout)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/single", sse.Wrap(singleEvent()))
	mux.HandleFunc("GET /api/echo", sse.Wrap(echoStream()))
	mux.Handle("GET /api/debug", expvar.Handler())

	addr := "127.0.0.1:3000"
	logger.Info().
		Str("address", addr).
		Str("example1", "curl 'http://localhost:3000/api/single'").
		Str("example2", "curl 'http://localhost:3000/api/echo'").
		Str("example3", "curl 'http://localhost:3000/api/echo?msg=hello'").
		Str("example4", "curl 'http://localhost:3000/api/debug'").
		Msg("http server started")

	if err := http.ListenAndServe(addr, mux); err != nil {
		logger.Fatal().Err(err).Msg("http server error")
	}
}

func singleEvent() sse.TypedHandler[string] {
	return func(*http.Request) (*sse.Response[string], error) {
		dataCh := make(chan string, 1)
		dataCh <- "Hello, world!"
		close(dataCh)

		return &sse.Response[string]{
			Name:   "single",
			DataCh: dataCh,
		}, nil
	}
}

func echoStream() sse.TypedHandler[string] {
	return func(req *http.Request) (*sse.Response[string], error) {
		ctx := req.Context()

		msg := req.URL.Query().Get("msg")
		msg = cmp.Or(msg, "hello, SSE!")

		if msg == "boom1" {
			return nil, errors.New("boom in handler!")
		}

		dataCh, errorCh := common.Tick(ctx, time.Second, func() (string, error) {
			if msg == "boom2" {
				return "", errors.New("boom in loop!")
			}

			return msg, nil
		})

		return &sse.Response[string]{
			Name:    "echo",
			DataCh:  dataCh,
			ErrorCh: errorCh,
		}, nil
	}
}
