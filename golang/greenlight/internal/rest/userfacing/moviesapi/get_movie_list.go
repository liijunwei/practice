package moviesapi

import (
	"greenlight/internal/common"
	"greenlight/internal/data"
	"greenlight/internal/validator"
	"net/http"
)

func GetMovieList(models data.Models) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Title   string
			Genres  []string
			Filters data.Filters
		}

		v := validator.New()
		qs := r.URL.Query()

		input.Title = common.ReadString(qs, "title", "")
		input.Genres = common.ReadCSV(qs, "genres", []string{})
		input.Filters.Page = common.ReadInt(qs, "page", 1, v)
		input.Filters.PageSize = common.ReadInt(qs, "page_size", 20, v)
		input.Filters.Sort = common.ReadString(qs, "sort", "id")
		input.Filters.SortSafelist = []string{"id", "title", "year", "runtime", "-id", "-title", "-year", "-runtime"}

		if data.ValidateFilters(v, input.Filters); !v.Valid() {
			common.RenderFailedValidation(w, r, v.Errors)
			return
		}

		if !v.Valid() {
			common.RenderFailedValidation(w, r, v.Errors)
			return
		}

		movies, metadata, err := models.Movies.GetAll(r.Context(), input.Title, input.Genres, input.Filters)
		if err != nil {
			common.RenderInternalServerError(w, r, err)
			return
		}

		if err := common.WriteResponseJSON(w, http.StatusOK, common.Envelope{"movies": movies, "metadata": metadata}, nil); err != nil {
			common.RenderInternalServerError(w, r, err)
		}
	}
}
