package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
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
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type setDeliveryTimesRequest struct {
	Times []string `json:"times"`
}

type setDeliveryTimesResponse struct {
	ID           string `json:"id"`
	DeliveryTime string `json:"delivery_time"`
}

func SetDeliveryTime(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	var req *setDeliveryTimesRequest

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

	emailID := emailFromDB.ID

	err = q.DeleteDeliveryTimes(ctx, emailID)
	if err != nil {
		api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
		return &apperror.APPError{
			Message: "Error deleting delivery times",
			Code:    500,
			Err:     err,
		}
	}

	var deleveryTimes []setDeliveryTimesResponse

	for _, timeStr := range req.Times {
		parsedTime, err := time.Parse("15:04", timeStr)
		if err != nil {
			api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
			return &apperror.APPError{
				Message: "Error formatting string time",
				Code:    500,
				Err:     err,
			}
		}

		// Get midnight of the same day
		midnight := time.Date(parsedTime.Year(), parsedTime.Month(), parsedTime.Day(), 0, 0, 0, 0, parsedTime.Location())

		// Calculate difference in microseconds
		microseconds := parsedTime.Sub(midnight).Microseconds()

		params := sqlc.SetDeliveryTimeParams{
			ID: utils.RandomString(),
			DeliveryTime: pgtype.Time{
				Microseconds: microseconds,
				Valid:        true,
			},
			EmailID: emailID,
		}

		deliveryTime, err := q.SetDeliveryTime(ctx, params)
		if err != nil {
			api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
			return &apperror.APPError{
				Message: "Error setting delivery time",
				Code:    500,
				Err:     err,
			}
		}

		microsecondsSinceMidnight := deliveryTime.DeliveryTime.Microseconds

		// Create new time by adding microseconds to midnight
		newTime := midnight.Add(time.Duration(microsecondsSinceMidnight) * time.Microsecond)

		// Format the time in 24-hour format
		timeStr := newTime.Format("15:04")

		dt := setDeliveryTimesResponse{
			ID:           deliveryTime.ID,
			DeliveryTime: timeStr,
		}

		deleveryTimes = append(deleveryTimes, dt)
	}

	api.ReturnResponse(w, 200, deleveryTimes, false, "")

	return nil
}
