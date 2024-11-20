package middleware

import "net/http"

func EnableCORS(next http.Handler, trustedOrigins map[string]struct{}) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Vary", "Origin")
		w.Header().Add("Vary", "Access-Control-Request-Method")

		origin := r.Header.Get("Origin")

		_, ok := trustedOrigins[origin]
		if origin != "" && ok {
			w.Header().Set("Access-Control-Allow-Origin", origin)

			if isPreflightRequest(r) {
				w.Header().Add("Access-Control-Request-Methods", "OPTIONS, PUT, PATCH, DELETE")
				w.Header().Add("Access-Control-Allow-Headers", "Authorization, Content-Type")
				w.WriteHeader(http.StatusOK)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

func isPreflightRequest(r *http.Request) bool {
	return r.Method == http.MethodOptions && r.Header.Get("Access-Control-Request-Method") != ""
}
