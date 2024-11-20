package usersapi

import (
	"errors"
	"greenlight/internal/common"
	"greenlight/internal/data"
	"greenlight/internal/mailer"
	"greenlight/internal/validator"
	"net/http"
	"sync"
	"time"
)

func CreateUserHandler(models data.Models, mailer mailer.Mailer, sender string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := common.ReadJSON(w, r, &input); err != nil {
			common.RenderBadRequest(w, r, err)
			return
		}

		user := &data.User{
			Name:   input.Name,
			Email:  input.Email,
			Status: "pending",
		}

		if err := user.Password.Set(input.Password); err != nil {
			common.RenderInternalServerError(w, r, err)
			return
		}

		v := validator.New()

		if data.ValidateUser(v, user); !v.Valid() {
			common.RenderFailedValidation(w, r, v.Errors)
			return
		}

		if err := models.Users.Create(r.Context(), user); err != nil {
			switch {
			case errors.Is(err, data.ErrDuplicatedEmail):
				v.AddError("email", "email already taken") // FIXME: take care of "preventing enumeration attack" when necessary
				common.RenderFailedValidation(w, r, v.Errors)
			default:
				common.RenderInternalServerError(w, r, err)
			}

			return
		}

		if err := models.Permissions.AddForUser(r.Context(), user.ID, "movies:read"); err != nil {
			common.RenderInternalServerError(w, r, err)
			return
		}

		token, err := models.Tokens.New(r.Context(), user.ID, 3*24*time.Hour, data.ScopeActivation)
		if err != nil {
			common.RenderInternalServerError(w, r, err)
			return
		}

		var wg sync.WaitGroup

		ctx := r.Context()

		common.RunInBackground(ctx, &wg, func() {
			common.SendEmail(ctx, user, token, mailer, sender)
		})

		if err := common.WriteResponseJSON(w, http.StatusAccepted, common.Envelope{"user": user}, nil); err != nil {
			common.RenderInternalServerError(w, r, err)
		}
	}
}
