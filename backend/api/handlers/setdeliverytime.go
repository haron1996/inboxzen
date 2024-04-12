package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/charmbracelet/log"
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

	params := sqlc.SetDeliveryTimeParams{
		DeliveryTime: strings.Join([]string{req.Hour, req.Minutes}, ":"),
		DeliveryAmPm: req.AmPm,
		EmailAddress: email,
	}

	time, err := q.SetDeliveryTime(ctx, params)
	if err != nil {
		api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
		return &apperror.APPError{
			Message: "Error setting delivery time",
			Code:    500,
			Err:     err,
		}
	}

	log.Info(time)

	api.ReturnResponse(w, 200, time, false, messages.OK)

	return nil
}
