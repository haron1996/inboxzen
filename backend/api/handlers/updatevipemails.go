package handlers

import (
	"encoding/json"
	"net/http"

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

type emails struct {
	Emails []string `json:"emails"`
}

func UpdateVipEmails(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	var req *emails

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

	err = q.DeleteVipEmails(ctx, email)
	if err != nil {
		api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
		return &apperror.APPError{
			Message: "Error deleting vip emails",
			Code:    500,
			Err:     err,
		}
	}

	for _, e := range req.Emails {
		params := sqlc.AddVipEmailParams{
			VipEmailAddress: e,
			EmailAddress:    email,
		}

		_, err := q.AddVipEmail(ctx, params)
		if err != nil {
			switch {
			case err.Error() == `ERROR: duplicate key value violates unique constraint "vipemailaddress_vip_email_address_email_address_key" (SQLSTATE 23505)`:
				log.Warnf("%s skipped", params.VipEmailAddress)
				continue
			default:
				api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
				return &apperror.APPError{
					Message: "Error adding vip email",
					Code:    500,
					Err:     err,
				}
			}
		}

	}

	emails, err := q.GetVipEmails(ctx, email)
	if err != nil {
		api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
		return &apperror.APPError{
			Message: "Error getting vip emails",
			Code:    500,
			Err:     err,
		}
	}

	api.ReturnResponse(w, 200, emails, false, messages.OK)

	return nil
}
