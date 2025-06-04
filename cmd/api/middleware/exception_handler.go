package middleware

import (
	"errors"
	httputils "github.com/knands42/lorecrafter/cmd/api/utils"
	"github.com/knands42/lorecrafter/internal/utils"
	"net/http"
)

type HandlerFuncWithError func(http.ResponseWriter, *http.Request) error

func ErrorHandlerMiddleware(next HandlerFuncWithError) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := next(w, r)
		if err == nil {
			return
		}

		var valErr *utils.ValidationError
		switch {
		case errors.As(err, &valErr):
			httputils.WriteJSONValidationError(w, valErr.Errors)
			return

		default:
			httputils.WriteJSONError(w, http.StatusInternalServerError, "Failed to register user")
			return
		}
	}
}
