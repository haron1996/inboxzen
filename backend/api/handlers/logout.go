package handlers

import (
	"net/http"
	"time"

	"github.com/haron1996/inboxzen/api"
	"github.com/haron1996/inboxzen/apperror"
	"github.com/haron1996/inboxzen/messages"
	"github.com/haron1996/inboxzen/mw"
	"github.com/haron1996/inboxzen/paseto"
)

func LogOut(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	const pLoad mw.ContextKey = "payload"

	payload := ctx.Value(pLoad).(*paseto.PayLoad)

	userID := payload.UserID

	email := payload.Email

	duration := 1 * time.Microsecond

	t, _, err := paseto.CreateToken(userID, email, issuedAt, duration)
	if err != nil {
		api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
		return &apperror.APPError{
			Message: "Error creating access token",
			Code:    500,
			Err:     err,
		}
	}

	session := http.Cookie{
		Name:     "session",
		Value:    t,
		Path:     "/",
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		HttpOnly: true,
		Expires:  time.Now().UTC().Add(-24 * time.Hour),
	}

	http.SetCookie(w, &session)

	api.ReturnResponse(w, 200, nil, true, messages.OK)

	return nil
}
