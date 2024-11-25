package sse

import (
	"context"
	"fmt"
	"greenlight/internal/common"
	"net/http"
	"time"

	"github.com/rs/zerolog"
)

const (
	Timeout               = 30 * time.Minute
	HeartbeatInterval     = 15 * time.Second
	RequestCanceled       = "request canceled"
	StatusRequestCanceled = 499
)

type Response[T any] struct {
	Name    string
	DataCh  <-chan T
	ErrorCh <-chan error
}

type TypedHandler[T any] func(r *http.Request) (*Response[T], error)

var sseMetrics metrics = newMetrics()

func HandlerWrapper[T any](handler TypedHandler[T]) http.HandlerFunc {
	return func(respW http.ResponseWriter, req *http.Request) {
		defer sseMetrics.record()()

		ctx := req.Context()

		ctx, cancel := context.WithTimeout(ctx, Timeout)
		defer cancel()

		req = req.WithContext(ctx)

		sseResp, err := handler(req)
		if err != nil {
			zerolog.Ctx(ctx).Err(err).Msg("failed to start SSE handler")
			common.RenderInternalServerError(respW, req, err)

			return
		}

		writer := newEventWriter[T](respW)
		writeHeader(writer.respW)

		if err := handleSSE(ctx, writer, sseResp); err != nil {
			zerolog.Ctx(ctx).Error().Str("name", sseResp.Name).Err(err).Msg("sse error")
		}
	}
}

func handleSSE[T any](
	ctx context.Context,
	writer *sseWriter[T],
	sseResp *Response[T],
) error {
	logger := zerolog.Ctx(ctx).With().Logger()

	// return means stop streaming data from server to client
	for {
		select {
		case <-ctx.Done():
			logger.Info().Str("sse_name", sseResp.Name).Msg("sse stopped: context done")
			return nil
		case err, ok := <-sseResp.ErrorCh:
			if !ok {
				logger.Info().Str("sse_name", sseResp.Name).Msg("sse error channel closed")
				return nil
			}

			if err := writeError(writer.respW, writer.buffer, err); err != nil {
				return fmt.Errorf("failed to write event error: %w", err)
			}

			return nil
		case data, ok := <-sseResp.DataCh:
			if !ok {
				logger.Info().Str("sse_name", sseResp.Name).Msg("sse data channel closed")
				return nil
			}

			if err := writeMessage(writer.respW, writer.buffer, data); err != nil {
				return fmt.Errorf("failed to write event data: %w", err)
			}

			// keep streaming
		}
	}
}
