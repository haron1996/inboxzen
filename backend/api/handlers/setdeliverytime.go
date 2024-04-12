package handlers

import (
	"encoding/json"
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
	"github.com/haron1996/inboxzen/viper"
	"github.com/jackc/pgx/v5/pgxpool"
)

type deliveryTime struct {
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

	timeString := fmt.Sprintf("%s:%s", req.Hour, req.Minutes)
	fullTimeString := fmt.Sprintf("%s %s", timeString, req.AmPm)

	params := sqlc.SetDeliveryTimeParams{
		DeliveryTime: fullTimeString,
		EmailAddress: email,
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

	// Parse the time string
	t, err := time.Parse("03:04 pm", dt.DeliveryTime)
	if err != nil {
		api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
		return &apperror.APPError{
			Message: "Error parsing time",
			Code:    500,
			Err:     err,
		}
	}

	// Extract hour, minute, and am/pm
	hour := fmt.Sprintf("%02d", t.Hour())     // Ensure hour is two digits
	minute := fmt.Sprintf("%02d", t.Minute()) // Ensure minute is two digits
	amPm := strings.ToLower(t.Format("pm"))

	// Convert to 12-hour format
	if t.Hour() > 12 {
		hour = fmt.Sprintf("%02d", t.Hour()-12)
	}

	res := deliveryTime{
		Hour:    hour,
		Minutes: minute,
		AmPm:    amPm,
	}

	api.ReturnResponse(w, 200, res, false, messages.OK)

	return nil
}
