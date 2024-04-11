package handlers

import (
	"net/http"

	"github.com/haron1996/inboxzen/api"
	"github.com/haron1996/inboxzen/messages"
	"github.com/haron1996/inboxzen/mw"
	"github.com/haron1996/inboxzen/paseto"
)

func CheckUserLoginStatus(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	const pLoad mw.ContextKey = "payload"

	payload := ctx.Value(pLoad).(*paseto.PayLoad)

	api.ReturnResponse(w, 200, payload, false, messages.OK)

	return nil
}
