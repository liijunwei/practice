package usersapi

import (
	"errors"
	"greenlight/internal/common"
	"greenlight/internal/data"
	"greenlight/internal/validator"
	"net/http"
)

func ActivateUserHandler(models data.Models) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			TokenPlaintext string `json:"token"`
		}

		if err := common.ReadJSON(w, r, &input); err != nil {
			common.RenderBadRequest(w, r, err)
			return
		}

		v := validator.New()

		if data.ValidateTokenPlaintext(v, input.TokenPlaintext); !v.Valid() {
			common.RenderFailedValidation(w, r, v.Errors)
			return
		}

		user, err := models.Users.GetByToken(r.Context(), data.ScopeActivation, input.TokenPlaintext)
		if err != nil {
			switch {
			case errors.Is(err, data.ErrRecordNotFound):
				v.AddError("token", "invalid or expired activation token")
				common.RenderFailedValidation(w, r, v.Errors)
			default:
				common.RenderInternalServerError(w, r, err)
			}

			return
		}

		user.Status = "activated"

		if err := models.Users.Update(r.Context(), user); err != nil {
			switch {
			case errors.Is(err, data.ErrStaleObject):
				common.RenderEditStaleRecord(w, r)
			default:
				common.RenderInternalServerError(w, r, err)
			}

			return
		}

		if err := models.Tokens.DeleteAllForUser(r.Context(), data.ScopeActivation, user.ID); err != nil {
			common.RenderInternalServerError(w, r, err)
			return
		}

		if err := common.WriteResponseJSON(w, http.StatusOK, common.Envelope{"user": user}, nil); err != nil {
			common.RenderInternalServerError(w, r, err)
		}
	}
}
