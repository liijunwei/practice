package main

import (
	"bytes"
	"cmp"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

// go clean -testcache && CONCURRENCY=100  go test -v .
// go clean -testcache && CONCURRENCY=1000 go test -v .
func TestConcurrentTransferAtoB(t *testing.T) {
	var wg sync.WaitGroup

	concurrencyStr := cmp.Or(os.Getenv("CONCURRENCY"), "100")
	concurrency, err := strconv.Atoi(concurrencyStr)
	boom(err)

	wg.Add(concurrency)

	for i := range concurrency {
		go doTransfer(t, i, &wg)
	}

	wg.Wait()
	t.Log("All transfers completed")
}

func doTransfer(t *testing.T, goroutineNum int, wg *sync.WaitGroup) {
	t.Helper()

	defer wg.Done()

	info := make(map[string]any)
	info["from_account_id"] = "79726a51-56bb-4bca-b1f0-2cbded8cac8a"
	info["to_account_id"] = "83b93359-f715-4fb2-9c1e-f3e5645793e5"
	info["amount"] = 0.1

	transferBody, err := json.Marshal(info)
	require.NoError(t, err)

	ctx := context.Background()
	transferAPI := "http://localhost:8080/transfer"

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, transferAPI, bytes.NewReader(transferBody))
	require.NoError(t, err)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if resp == nil {
			t.Errorf("failed to transfer without error response: %v", err)
			return
		}

		require.NotNil(t, resp)
		if b, errR := io.ReadAll(resp.Body); errR == nil {
			t.Errorf("failed to transfer with error response: %v, response: %s", err, string(b))
			return
		}

		t.Errorf("failed to transfer: %v", err)
		return
	}

	require.NoError(t, err)
	defer resp.Body.Close()

	if b, errR := io.ReadAll(resp.Body); errR == nil {
		t.Logf("api response: %s", string(b))
		return
	}

	require.Equal(t, http.StatusOK, resp.StatusCode)

	t.Logf("Goroutine %d completed transfer", goroutineNum)
}
