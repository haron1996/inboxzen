package handlers

import (
	"encoding/base64"
	"fmt"
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
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetFirstTimeSenders(w http.ResponseWriter, r *http.Request) error {
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

	srv, err := utils.ConstructGmailService(ctx, q, userID, email)
	if err != nil {
		api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
		return &apperror.APPError{
			Message: "Error costructing gmail service",
			Code:    500,
			Err:     err,
		}
	}

	user := "me"

	//Q("is:inbox is:unread")

	listMessagesResponse, err := srv.Users.Messages.List(user).Do()
	if err != nil {
		api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
		return &apperror.APPError{
			Message: "Error getting messages",
			Code:    500,
			Err:     err,
		}
	}

	for _, message := range listMessagesResponse.Messages {
		// Extract message body
		var body string
		if message.Payload != nil {
			if len(message.Payload.Parts) > 0 {
				// If the message has parts, concatenate all parts
				for _, part := range message.Payload.Parts {
					if part.Body != nil && part.Body.Data != "" {
						body += part.Body.Data
					}
				}
			} else if message.Payload.Body != nil && message.Payload.Body.Data != "" {
				// If the message has no parts, use the message body directly
				body = message.Payload.Body.Data
			}
		}

		// Decode base64 encoded body
		decodedBody, err := base64.URLEncoding.DecodeString(body)
		if err != nil {
			log.Fatalf("Unable to decode message body: %v", err)
		}

		// Print the decoded message body
		fmt.Println(string(decodedBody))
	}

	return nil
}
