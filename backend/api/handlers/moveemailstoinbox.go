package handlers

import (
	"errors"
	"fmt"
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
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/api/gmail/v1"
)

func MoveEmailsToInbox(w http.ResponseWriter, r *http.Request) error {
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

	for _, t := range times {
		microsecondsSinceMidnight := t.DeliveryTime.Microseconds // Example microseconds since midnight

		// Get midnight of the current day
		now := time.Now()

		midnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

		// Create new time by adding microseconds to midnight
		timeFromDB := midnight.Add(time.Duration(microsecondsSinceMidnight) * time.Microsecond)

		moveEmailToInbox := time.Now().Unix() == timeFromDB.Unix() || time.Now().After(timeFromDB) && time.Since(timeFromDB) <= 1*time.Minute

		if moveEmailToInbox {
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

			listLabelsResponse, err := srv.Users.Labels.List(user).Do()
			if err != nil {
				return err
			}

			customLabelName := c.HoldLabel

			var customLabelID string

			for _, label := range listLabelsResponse.Labels {
				if label.Name == customLabelName {
					customLabelID = label.Id
					break
				}
			}

			query := fmt.Sprintf("label:%s", customLabelName)

			var allMessages []*gmail.Message

			listMessagesResponse, err := srv.Users.Messages.List(user).Q(query).Do()
			if err != nil {
				api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
				return &apperror.APPError{
					Message: "Error getting messages",
					Code:    500,
					Err:     err,
				}
			}

			allMessages = append(allMessages, listMessagesResponse.Messages...)

			for listMessagesResponse.NextPageToken != "" {
				listMessagesResponse, err = srv.Users.Messages.List(user).PageToken(listMessagesResponse.NextPageToken).Q(query).Do()
				if err != nil {
					api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
					return &apperror.APPError{
						Message: "Error getting messages",
						Code:    500,
						Err:     err,
					}
				}

				allMessages = append(allMessages, listMessagesResponse.Messages...)
			}

			for _, message := range allMessages {
				modifyMessageRequest := &gmail.ModifyMessageRequest{
					RemoveLabelIds: []string{customLabelID},
					AddLabelIds:    []string{"INBOX"},
				}

				_, err = srv.Users.Messages.Modify(user, message.Id, modifyMessageRequest).Do()
				if err != nil {
					api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
					return &apperror.APPError{
						Message: "Error moving message to inbox",
						Code:    500,
						Err:     err,
					}
				}
			}

			api.ReturnResponse(w, 200, nil, true, messages.OK)
		}
	}

	api.ReturnResponse(w, 200, nil, true, messages.OK)

	return nil
}
