package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/charmbracelet/log"
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

type domainNames struct {
	DomainNames []string `json:"domain_names"`
}

func AddDomains(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	var req *domainNames

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

	err = q.DeleteDomains(ctx, emailFromDB.ID)
	if err != nil {
		api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
		return &apperror.APPError{
			Message: "Error deleting domain",
			Code:    500,
			Err:     err,
		}
	}

	for _, domainName := range req.DomainNames {
		params := sqlc.AddDomainParams{
			ID:         utils.RandomString(),
			DomainName: domainName,
			EmailID:    emailFromDB.ID,
		}

		_, err := q.AddDomain(ctx, params)
		if err != nil {
			switch {
			case err.Error() == `ERROR: duplicate key value violates unique constraint "vipdomain_domain_name_email_address_key" (SQLSTATE 23505)`:
				log.Warnf("%s skipped", params.DomainName)
				continue
			default:
				api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
				return &apperror.APPError{
					Message: "Error adding domain",
					Code:    500,
					Err:     err,
				}
			}
		}

	}

	domains, err := q.GetDomains(ctx, emailFromDB.ID)
	if err != nil {
		api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
		return &apperror.APPError{
			Message: "Error getting domains",
			Code:    500,
			Err:     err,
		}
	}

	emails, err := q.GetVipEmails(ctx, emailFromDB.ID)
	if err != nil {
		api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
		return &apperror.APPError{
			Message: "Error getting user emails",
			Code:    500,
			Err:     err,
		}
	}

	keywords, err := q.GetKeywords(ctx, emailFromDB.ID)
	if err != nil {
		api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
		return &apperror.APPError{
			Message: "Error getting vip keywords",
			Code:    500,
			Err:     err,
		}
	}

	filterParams := utils.NewFilterParams(q, ctx, userID, email, domains, emails, keywords, c)

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

	api.ReturnResponse(w, 200, domains, false, "")

	return nil
}
