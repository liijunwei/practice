package moviesapi

import (
	"greenlight/internal/common"
	"greenlight/internal/data"
	"net/http"
)

func GetMovieDetailHandler(models data.Models) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		movieID, err := common.ReadIDParam(r)
		if err != nil {
			common.RenderNotFound(w, r)
			return
		}

		movie, err := models.Movies.Get(r.Context(), movieID)
		if err != nil {
			common.RenderNotFoundOrUnknownError(err, w, r)
			return
		}

		if err := common.WriteResponseJSON(w, http.StatusOK, common.Envelope{"movies": movie}, nil); err != nil {
			common.RenderInternalServerError(w, r, err)
		}
	}
}
