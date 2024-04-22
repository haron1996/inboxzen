package handlers

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

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
	"google.golang.org/api/gmail/v1"
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
			Message: "Error constructing gmail service",
			Code:    500,
			Err:     err,
		}
	}

	user := "me"

	var allMessages []*gmail.Message

	// Initial request to fetch the first page of messages
	listMessagesResponse, err := srv.Users.Messages.List(user).Do()
	if err != nil {
		// Handle error
		api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
		return &apperror.APPError{
			Message: "Error getting messages",
			Code:    500,
			Err:     err,
		}
	}

	// Add messages from the first page to the list
	allMessages = append(allMessages, listMessagesResponse.Messages...)

	// Check if there are more pages
	for listMessagesResponse.NextPageToken != "" {
		// Fetch the next page of messages
		listMessagesResponse, err = srv.Users.Messages.List(user).PageToken(listMessagesResponse.NextPageToken).Do()
		if err != nil {
			// Handle error
			api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
			return &apperror.APPError{
				Message: "Error getting messages",
				Code:    500,
				Err:     err,
			}
		}

		// Add messages from the current page to the list
		allMessages = append(allMessages, listMessagesResponse.Messages...)
	}

	// Iterate through each message
	for _, msg := range allMessages {
		// Get the full message details
		fullMessage, err := srv.Users.Messages.Get(user, msg.Id).Do()
		if err != nil {
			// Handle error
			api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
			return &apperror.APPError{
				Message: "Error getting message details",
				Code:    500,
				Err:     err,
			}
		}

		// Extract sender name, email, subject line, and date
		var senderName, senderEmail string
		var subject, date string
		for _, header := range fullMessage.Payload.Headers {
			switch header.Name {
			case "From":
				senderName, senderEmail, _ = parseAddress(header.Value)
			case "Subject":
				subject = header.Value
			case "Date":
				date = header.Value
			}
		}

		// Log sender details, subject, and date
		log.Infof("SENDER NAME: %s", senderName)
		log.Infof("SENDER EMAIL ADDRESS: %s", senderEmail)
		log.Infof("SUBJECT: %s", subject)
		log.Infof("DATE: %s", date)

		// Access the message body
		var body string
		for _, part := range fullMessage.Payload.Parts {
			if part.MimeType == "text/plain" {
				// If the message is in plain text format
				body = part.Body.Data
			} else if part.MimeType == "text/html" {
				// If the message is in HTML format
				decodedBody, err := base64.URLEncoding.DecodeString(part.Body.Data)
				if err != nil {
					// Handle error
					api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
					return &apperror.APPError{
						Message: "Error decoding HTML body",
						Code:    500,
						Err:     err,
					}
				}
				body = string(decodedBody)
			}
		}

		// Now 'body' contains the message body
		fmt.Println(body)
	}

	log.Infof("ALL MESSAGES: %d", len(allMessages))

	return nil
}

// Function to parse sender address into name and email parts
func parseAddress(address string) (name, email string, err error) {
	parts := strings.Split(address, " <")
	if len(parts) == 2 {
		name = parts[0]
		email = strings.TrimRight(parts[1], ">")
	} else {
		email = address
	}
	return name, email, nil
}
