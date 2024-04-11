package mw

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/haron1996/inboxzen/api"
	"github.com/haron1996/inboxzen/messages"
	"github.com/haron1996/inboxzen/paseto"
	"github.com/haron1996/inboxzen/sqlc"
	"github.com/haron1996/inboxzen/viper"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ContextKey string

const pLoad ContextKey = "payload"

func AuthenticateRequest() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			payload, err := verifyToken(r)
			if err != nil {
				log.Errorf("Error verifying token: %v", err)
				api.ReturnResponse(w, 401, nil, true, messages.ErrInvalidToken)
				return
			}

			if payload != nil {
				account, err := getUser(ctx, payload.UserID)
				if err != nil {
					log.Errorf("Errror getting user: %v", err)
					api.ReturnResponse(w, 401, nil, true, messages.ErrInvalidToken)
					return
				}

				if payload.IssuedAt.Unix() != account.LastLogin.Time.Unix() {
					log.Error("Invalid access token")
					api.ReturnResponse(w, 401, nil, true, messages.ErrInvalidToken)
					return
				}

				ctx := context.WithValue(r.Context(), pLoad, payload)

				next.ServeHTTP(w, r.WithContext(ctx))
			}
		}

		return http.HandlerFunc(fn)
	}
}

func verifyToken(r *http.Request) (*paseto.PayLoad, error) {
	c, err := r.Cookie("session")
	if err != nil {
		return nil, err
	}

	payload, err := paseto.VerifyToken(c.Value)
	if err != nil {
		return nil, err
	}

	return payload, nil
}

func getUser(ctx context.Context, userID string) (*sqlc.User, error) {
	config, err := viper.LoadConfig(".")
	if err != nil {
		return nil, fmt.Errorf("something went wrong: %v", err)
	}

	pool, err := pgxpool.New(ctx, config.DBString)
	if err != nil {
		return nil, fmt.Errorf("something went wrong: %v", err)
	}

	defer pool.Close()

	q := sqlc.New(pool)

	user, err := q.GetUser(ctx, userID)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return nil, errors.New("user not found")
		default:
			return nil, fmt.Errorf("error getting user: %v", err)
		}
	}

	return &user, nil
}
