package logger

import (
	"errors"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/haron1996/inboxzen/apperror"
)

type apiHandler func(w http.ResponseWriter, r *http.Request) error

func MakeHandler(h apiHandler, logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			var apiErr *apperror.APPError
			switch {
			case errors.As(err, &apiErr):
				logger.Errorf("Error message: %s, Error code: %d, Error: %v", apiErr.Message, apiErr.Code, apiErr.Err)
				return
			default:
				logger.Errorf("Unknown error: %v", err)
				return
			}
		}
	}
}
