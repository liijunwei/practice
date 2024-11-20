package moviesapi

import (
	"errors"
	"greenlight/internal/common"
	"greenlight/internal/data"
	"greenlight/internal/validator"
	"net/http"
)

func UpdateMovieDetailHandler(models data.Models) http.HandlerFunc {
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

		var input struct {
			Title   string       `json:"title"`
			Year    int32        `json:"year"`
			Runtime data.Runtime `json:"runtime"`
			Genres  []string     `json:"genres"`
		}

		if err := common.ReadJSON(w, r, &input); err != nil {
			common.RenderBadRequest(w, r, err)
			return
		}

		movie.Title = input.Title
		movie.Year = input.Year
		movie.Runtime = input.Runtime
		movie.Genres = input.Genres

		v := validator.New()
		if data.ValidateMovie(v, movie); !v.Valid() {
			common.RenderInternalServerError(w, r, err)
			return
		}

		if err := models.Movies.Update(r.Context(), movie); err != nil {
			switch {
			case errors.Is(err, data.ErrStaleObject):
				common.RenderEditStaleRecord(w, r)
			default:
				common.RenderInternalServerError(w, r, err)
			}
			return
		}

		if err := common.WriteResponseJSON(w, http.StatusOK, common.Envelope{"movie": movie}, nil); err != nil {
			common.RenderInternalServerError(w, r, err)
		}
	}
}
