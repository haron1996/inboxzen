package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/haron1996/inboxzen/api"
	"github.com/haron1996/inboxzen/apperror"
	"github.com/haron1996/inboxzen/messages"
	"github.com/haron1996/inboxzen/mw"
	"github.com/haron1996/inboxzen/paseto"
	"github.com/haron1996/inboxzen/sqlc"
	"github.com/haron1996/inboxzen/utils"
	"github.com/haron1996/inboxzen/viper"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type deliveryTime struct {
	ID      string `json:"id"`
	Hour    string `json:"hour"`
	Minutes string `json:"minutes"`
	AmPm    string `json:"am_pm"`
}

func SetDeliveryTime(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	var req *deliveryTime

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
		return &apperror.APPError{
			Message: "Error returning new decoder",
			Code:    500,
			Err:     err,
		}
	}

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

	const pLoad mw.ContextKey = "payload"

	payload := ctx.Value(pLoad).(*paseto.PayLoad)

	email := payload.Email
	userID := payload.UserID

	timeString := fmt.Sprintf("%s:%s", req.Hour, req.Minutes)
	fullTimeString := fmt.Sprintf("%s %s", timeString, req.AmPm)

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

	params := sqlc.SetDeliveryTimeParams{
		ID:           utils.RandomString(),
		DeliveryTime: fullTimeString,
		EmailID:      emailFromDB.ID,
	}

	dt, err := q.SetDeliveryTime(ctx, params)
	if err != nil {
		switch {
		case err.Error() == `ERROR: duplicate key value violates unique constraint "deliverytime_delivery_time_email_address_key" (SQLSTATE 23505)`:
			api.ReturnResponse(w, 409, nil, true, messages.ErrConflict)
			return &apperror.APPError{
				Message: "Error setting delivery time",
				Code:    409,
				Err:     err,
			}

		default:
			api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
			return &apperror.APPError{
				Message: "Error setting delivery time",
				Code:    500,
				Err:     err,
			}
		}
	}

	hour, minute, amPm, err := parseTime(dt)
	if err != nil {
		api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
		return &apperror.APPError{
			Message: "Error parsing time",
			Code:    500,
			Err:     err,
		}
	}

	res := deliveryTime{
		ID:      dt.ID,
		Hour:    hour,
		Minutes: minute,
		AmPm:    amPm,
	}

	api.ReturnResponse(w, 200, res, false, "")

	return nil
}

func parseTime(dt sqlc.Deliverytime) (string, string, string, error) {
	t, err := time.Parse("03:04 pm", dt.DeliveryTime)
	if err != nil {
		return "", "", "", err
	}

	// Extract hour, minute, and am/pm
	hour := fmt.Sprintf("%02d", t.Hour())     // Ensure hour is two digits
	minute := fmt.Sprintf("%02d", t.Minute()) // Ensure minute is two digits
	amPm := strings.ToLower(t.Format("pm"))

	// Convert to 12-hour format
	if t.Hour() > 12 {
		hour = fmt.Sprintf("%02d", t.Hour()-12)
	}

	return hour, minute, amPm, nil
}
