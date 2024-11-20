package moviesapi

import (
	"fmt"
	"greenlight/internal/common"
	"greenlight/internal/data"
	"greenlight/internal/validator"
	"net/http"
)

func CreateMovieHandler(models data.Models) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Title   string   `json:"title"`
			Year    int32    `json:"year"`
			Runtime int32    `json:"runtime"`
			Genres  []string `json:"genres"`
		}

		if err := common.ReadJSON(w, r, &input); err != nil {
			common.RenderBadRequest(w, r, err)
			return
		}

		movie := &data.Movie{
			Title:   input.Title,
			Year:    input.Year,
			Runtime: data.Runtime(input.Runtime),
			Genres:  input.Genres,
		}

		v := validator.New()
		if data.ValidateMovie(v, movie); !v.Valid() {
			common.RenderFailedValidation(w, r, v.Errors)
			return
		}

		if err := models.Movies.Create(r.Context(), movie); err != nil {
			common.RenderInternalServerError(w, r, err)
			return
		}

		headers := make(http.Header)
		headers.Set("Location", fmt.Sprintf("/v1/movies/%d", movie.ID))

		if err := common.WriteResponseJSON(w, http.StatusCreated, common.Envelope{"movie": movie}, headers); err != nil {
			common.RenderInternalServerError(w, r, err)
		}
	}
}
