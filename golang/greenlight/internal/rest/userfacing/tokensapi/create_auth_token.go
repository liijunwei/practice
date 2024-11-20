package tokensapi

import (
	"greenlight/internal/common"
	"greenlight/internal/data"
	"greenlight/internal/validator"
	"net/http"
	"time"
)

func CreateAuthenticationTokenHandler(models data.Models) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := common.ReadJSON(w, r, &input); err != nil {
			common.RenderBadRequest(w, r, err)
			return
		}

		v := validator.New()
		data.ValidateEmail(v, input.Email)
		data.ValidatePasswordPlaintext(v, input.Password)

		if !v.Valid() {
			common.RenderFailedValidation(w, r, v.Errors)
			return
		}

		ctx := r.Context()
		user, err := models.Users.GetByEmail(ctx, input.Email)
		if err != nil {
			common.RenderNotFoundOrUnknownError(err, w, r)
			return
		}

		match, err := user.Password.Matches(input.Password)
		if err != nil {
			common.RenderInternalServerError(w, r, err)
			return
		}

		if !match {
			common.RenderInvlidCredentials(w, r)
			return
		}

		token, err := models.Tokens.New(r.Context(), user.ID, 24*time.Hour, data.ScopeAuthentication)
		if err != nil {
			common.RenderInternalServerError(w, r, err)
			return
		}

		if err := common.WriteResponseJSON(w, http.StatusCreated, common.Envelope{"authentication_token": token}, nil); err != nil {
			common.RenderInternalServerError(w, r, err)
		}
	}
}
