package handlers

import (
	"net/http"

	"github.com/haron1996/inboxzen/api"
	"github.com/haron1996/inboxzen/apperror"
	"github.com/haron1996/inboxzen/messages"
	"github.com/haron1996/inboxzen/mw"
	"github.com/haron1996/inboxzen/paseto"
	"github.com/haron1996/inboxzen/sqlc"
	"github.com/haron1996/inboxzen/utils"
	"github.com/haron1996/inboxzen/viper"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Activate(w http.ResponseWriter, r *http.Request) error {
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

	domains, err := q.GetDomains(ctx, email)
	if err != nil {
		api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
		return &apperror.APPError{
			Message: "Error getting vip domains",
			Code:    500,
			Err:     err,
		}
	}

	emails, err := q.GetVipEmails(ctx, email)
	if err != nil {
		api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
		return &apperror.APPError{
			Message: "Error getting user emails",
			Code:    500,
			Err:     err,
		}
	}

	keywords, err := q.GetKeywords(ctx, email)
	if err != nil {
		api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
		return &apperror.APPError{
			Message: "Error getting vip keywords",
			Code:    500,
			Err:     err,
		}
	}

	filterParams := utils.NewFilterParams(q, ctx, userID, email, domains, emails, keywords)

	err = filterParams.CreateCustomLabels()
	if err != nil {
		api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
		return &apperror.APPError{
			Message: "Error creating label",
			Code:    500,
			Err:     err,
		}
	}

	err = filterParams.CreateHoldFilter()
	if err != nil {
		api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
		return &apperror.APPError{
			Message: "Error creating hold filter",
			Code:    500,
			Err:     err,
		}
	}

	params := sqlc.ActivateParams{
		EmailAddress: email,
		UserID:       userID,
	}

	err = q.Activate(ctx, params)
	if err != nil {
		api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
		return &apperror.APPError{
			Message: "Error activating email",
			Code:    500,
			Err:     err,
		}
	}

	api.ReturnResponse(w, 200, nil, true, messages.OK)

	return nil
}