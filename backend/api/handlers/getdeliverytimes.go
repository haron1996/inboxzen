package handlers

import (
	"errors"
	"net/http"

	"github.com/haron1996/inboxzen/api"
	"github.com/haron1996/inboxzen/apperror"
	"github.com/haron1996/inboxzen/messages"
	"github.com/haron1996/inboxzen/mw"
	"github.com/haron1996/inboxzen/paseto"
	"github.com/haron1996/inboxzen/sqlc"
	"github.com/haron1996/inboxzen/viper"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetDeliveryTimes(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	const pLoad mw.ContextKey = "payload"

	payload := ctx.Value(pLoad).(*paseto.PayLoad)

	email := payload.Email
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

	getEmailParams := sqlc.GetEmailParams{
		EmailAddress: email,
		UserID:       userID,
	}

	emailFromDB, err := q.GetEmail(ctx, getEmailParams)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			api.ReturnResponse(w, 404, nil, true, messages.ErrNotFound)
			return &apperror.APPError{
				Message: "Email not found",
				Code:    404,
				Err:     err,
			}
		default:
			api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
			return &apperror.APPError{
				Message: "Error getting email address",
				Code:    500,
				Err:     err,
			}
		}
	}

	times, err := q.GetDeliveryTimes(ctx, emailFromDB.ID)
	if err != nil {
		api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
		return &apperror.APPError{
			Message: "Error getting delivery times",
			Code:    500,
			Err:     err,
		}
	}

	var dts []deliveryTime

	for _, tm := range times {
		hour, minute, amPm, err := parseTime(tm)
		if err != nil {
			api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
			return &apperror.APPError{
				Message: "Error parsing time",
				Code:    500,
				Err:     err,
			}
		}

		res := deliveryTime{
			ID:      tm.ID,
			Hour:    hour,
			Minutes: minute,
			AmPm:    amPm,
		}

		dts = append(dts, res)
	}

	api.ReturnResponse(w, 200, dts, false, "")

	return nil
}
