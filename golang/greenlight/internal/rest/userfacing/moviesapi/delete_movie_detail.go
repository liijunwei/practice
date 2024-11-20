package moviesapi

import (
	"greenlight/internal/common"
	"greenlight/internal/data"
	"net/http"
)

func DeleteMovieHandler(models data.Models) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		movieID, err := common.ReadIDParam(r)
		if err != nil {
			common.RenderNotFound(w, r)
			return
		}

		if err := models.Movies.Delete(r.Context(), movieID); err != nil {
			common.RenderNotFoundOrUnknownError(err, w, r)
			return
		}

		if err := common.WriteResponseJSON(w, http.StatusOK, common.Envelope{"movie": "movie deleted"}, nil); err != nil {
			common.RenderInternalServerError(w, r, err)
		}
	}
}
