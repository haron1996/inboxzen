package handlers

import (
	"net/http"
	"net/url"

	"github.com/haron1996/inboxzen/api"
	"github.com/haron1996/inboxzen/apperror"
	"github.com/haron1996/inboxzen/messages"
	"github.com/haron1996/inboxzen/utils"
	"golang.org/x/oauth2"
)

func GenerateGoogleAuthURL(w http.ResponseWriter, r *http.Request) error {

	config, err := utils.ConstructGoogleConfig()
	if err != nil {
		api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
		return &apperror.APPError{
			Message: "Error constructing config from json",
			Code:    500,
			Err:     err,
		}
	}

	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)

	u, err := url.Parse(authURL)
	if err != nil {
		api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
		return &apperror.APPError{
			Message: "Error returning oauth url",
			Code:    500,
			Err:     err,
		}
	}

	q := u.Query()

	q.Set("prompt", "select_account")

	u.RawQuery = q.Encode()

	api.ReturnResponse(w, 200, u.String(), false, "")

	return nil
}
