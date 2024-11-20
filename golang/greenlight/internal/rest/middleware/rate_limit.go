package middleware

import (
	"greenlight/internal/common"
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

// TODO move to config
type limiterConfig struct {
	RPS     float64 `json:"rps"`
	Burst   int     `json:"burst"`
	Enabled bool    `json:"enabled"`
}

func RateLimit(next http.Handler, rps float64, burst int) http.Handler {
	type client struct {
		limiter  *rate.Limiter
		lastSeen time.Time
	}

	var mu sync.Mutex
	var clients = make(map[string]*client) // works for single machine

	go func() {
		ticker := time.NewTicker(1 * time.Minute)

		for range ticker.C {
			mu.Lock()

			for ip, client := range clients {
				if time.Since(client.lastSeen) > 3*time.Minute {
					delete(clients, ip)
				}
			}

			mu.Unlock()
		}
	}()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			common.ServerErrorResponse(w, r, err, true)
			return
		}

		mu.Lock()

		if _, ok := clients[ip]; !ok {
			clients[ip] = &client{limiter: rate.NewLimiter(rate.Limit(rps), burst)}
		}

		clients[ip].lastSeen = time.Now()

		if !clients[ip].limiter.Allow() {
			mu.Unlock()

			// app.logger.Info().
			// 	Str("request_method", r.Method).
			// 	Str("request_url", r.URL.String()).
			// 	Str("user_agent", r.UserAgent()).
			// 	Msg("rate_limit triggered")

			rateLimitExceededResponse(w, r)

			return
		}

		mu.Unlock()

		next.ServeHTTP(w, r)
	})
}

func rateLimitExceededResponse(w http.ResponseWriter, r *http.Request) {
	message := "rate limit exceeded"
	common.ErrorResponse(w, r, http.StatusTooManyRequests, message)
}
