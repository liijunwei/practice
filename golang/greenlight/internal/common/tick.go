package common

import (
	"context"
	"time"
)

// TODO revise the api to return a cancel function
func Tick[T any](
	ctx context.Context,
	interval time.Duration,
	fetch func() (T, error),
) (<-chan T, <-chan error) {
	dataCh := make(chan T, 1)
	errCh := make(chan error, 1)

	// fetch immediately
	data, err := fetch()
	if err != nil {
		errCh <- err
	} else {
		dataCh <- data
	}

	go func() {
		defer close(dataCh)
		defer close(errCh)

		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				data, err := fetch()
				if err != nil {
					errCh <- err

					return
				}

				select {
				case <-ctx.Done():
					return
				case dataCh <- data:
				}
			}
		}
	}()

	return dataCh, errCh
}
