package handlers

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/haron1996/inboxzen/api"
	"github.com/haron1996/inboxzen/apperror"
	"github.com/haron1996/inboxzen/messages"
	"github.com/haron1996/inboxzen/mw"
	"github.com/haron1996/inboxzen/paseto"
	"github.com/haron1996/inboxzen/sqlc"
	"github.com/haron1996/inboxzen/viper"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SwitchAccount(w http.ResponseWriter, r *http.Request) error {
	emailAddress := chi.URLParam(r, "email")

	ctx := r.Context()

	const pLoad mw.ContextKey = "payload"

	payload := ctx.Value(pLoad).(*paseto.PayLoad)

	userID := payload.UserID

	c, err := viper.LoadConfig(".")
	if err != nil {
		api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
		return &apperror.APPError{
			Message: "Error loading config",
			Code:    500,
			Err:     err,
		}
	}

	p, err := pgxpool.New(ctx, c.DBString)
	if err != nil {
		api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
		return &apperror.APPError{
			Message: "Error creating pool",
			Code:    500,
			Err:     err,
		}
	}

	defer p.Close()

	q := sqlc.New(p)

	email, err := q.GetEmailByEmailAddress(ctx, emailAddress)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			api.ReturnResponse(w, 401, nil, true, messages.ErrUnauthorized)
			return &apperror.APPError{
				Message: "Error getting email account",
				Code:    401,
				Err:     err,
			}
		default:
			api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
			return &apperror.APPError{
				Message: "Error getting email account",
				Code:    500,
				Err:     err,
			}
		}
	}

	emails, err := q.GetAllUserEmails(ctx, userID)
	if err != nil {
		api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
		return &apperror.APPError{
			Message: "Error getting user email accounts",
			Code:    500,
			Err:     err,
		}
	}

	var filteredEmails []sqlc.Email

	for _, e := range emails {
		if e.EmailAddress == emailAddress {
			continue
		}

		filteredEmails = append(filteredEmails, e)
	}

	userAccount := userAccount{
		Email:  email,
		Emails: filteredEmails,
	}

	t, pload, err := paseto.CreateToken(userID, emailAddress, issuedAt, duration)
	if err != nil {
		api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
		return &apperror.APPError{
			Message: "Error creating token",
			Code:    500,
			Err:     err,
		}
	}

	lastLoginParams := sqlc.UpdateLastLoginParams{
		LastLogin: pgtype.Timestamptz{
			Time:  issuedAt,
			Valid: true,
		},
		ID: userID,
	}

	err = q.UpdateLastLogin(ctx, lastLoginParams)
	if err != nil {
		api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
		return &apperror.APPError{
			Message: "Error updating user last login",
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
		Expires:  pload.Expiry,
	}

	http.SetCookie(w, &session)

	api.ReturnResponse(w, 200, userAccount, false, messages.OK)

	return nil
}
