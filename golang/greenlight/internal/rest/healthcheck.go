package rest

import (
	"greenlight/internal/common"
	"net/http"
)

func HealthcheckHandler(environment, version string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		env := common.Envelope{
			"status": "available",
			"system_info": map[string]any{
				"environment": environment,
				"version":     version,
			},
		}

		if err := common.WriteResponseJSON(w, http.StatusOK, env, nil); err != nil {
			common.RenderInternalServerError(w, r, err)
		}
	}
}
