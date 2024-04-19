package utils

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/haron1996/inboxzen/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/oauth2"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

type FilterParams struct {
	Q           *sqlc.Queries
	Ctx         context.Context
	UserID      string
	Email       string
	VipDomains  []sqlc.Vipdomain
	VipEmails   []sqlc.Vipemailaddress
	VipKeywords []sqlc.Vipkeyword
}

func NewFilterParams(q *sqlc.Queries, ctx context.Context, userID, email string, domains []sqlc.Vipdomain, emails []sqlc.Vipemailaddress, keywords []sqlc.Vipkeyword) *FilterParams {
	return &FilterParams{
		Q:           q,
		Ctx:         ctx,
		UserID:      userID,
		Email:       email,
		VipDomains:  domains,
		VipEmails:   emails,
		VipKeywords: keywords,
	}
}

func ConstructGmailService(ctx context.Context, q *sqlc.Queries, userID, email string) (*gmail.Service, error) {
	getOauthParams := sqlc.GetEmailParams{
		EmailAddress: email,
		UserID:       userID,
	}

	e, err := q.GetEmail(ctx, getOauthParams)
	if err != nil {
		return nil, err
	}

	tok := &oauth2.Token{}

	err = json.Unmarshal(e.Oauth2Token, tok)
	if err != nil {
		return nil, err
	}

	config, err := ConstructGoogleConfig()
	if err != nil {
		return nil, err
	}

	client := config.Client(ctx, tok)

	srv, err := gmail.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, err
	}

	return srv, nil
}

func (p *FilterParams) CreateCustomLabels() error {
	srv, err := ConstructGmailService(p.Ctx, p.Q, p.UserID, p.Email)
	if err != nil {
		return err
	}

	user := "me"

	newLabelNames := []string{"GIZ", "Blocked By GIZ"}

	existingLabels, err := srv.Users.Labels.List(user).Do()
	if err != nil {
		return err
	}

	existingLabelNames := make(map[string]bool)

	for _, l := range existingLabels.Labels {
		existingLabelNames[l.Name] = true
	}

	for _, ln := range newLabelNames {
		if !existingLabelNames[ln] {
			newLabel := &gmail.Label{
				Name: ln,
			}
			_, err := srv.Users.Labels.Create(user, newLabel).Do()
			if err != nil {
				return err
			}

		}
	}

	return nil
}

func (p *FilterParams) CreateHoldFilter() error {
	srv, err := ConstructGmailService(p.Ctx, p.Q, p.UserID, p.Email)
	if err != nil {
		return err
	}

	user := "me"

	listLabelsResponse, err := srv.Users.Labels.List(user).Do()
	if err != nil {
		return err
	}

	customLabelName := "GIZ"

	var customLabelID string

	for _, label := range listLabelsResponse.Labels {
		if label.Name == customLabelName {
			customLabelID = label.Id
			break
		}
	}

	filterCriteria := &gmail.FilterCriteria{
		// Query: "in:inbox AND is:unread AND -label:" + customLabelName,
		Query: "in:inbox AND is:unread",
	}

	for _, email := range p.VipEmails {
		filterCriteria.Query += " -from:" + email.VipEmailAddress
	}

	for _, domain := range p.VipDomains {
		//filterCriteria.Query += " -from:@" + domain.DomainName
		filterCriteria.Query += " -from:*@" + domain.DomainName
	}

	for _, keyword := range p.VipKeywords {
		filterCriteria.Query += " -subject:\"" + keyword.Keyword + "\""
		// filterCriteria.Query += " -\"" + keyword + "\""
	}

	filterCriteria.Query = strings.TrimSpace(filterCriteria.Query)

	getEmailParams := sqlc.GetEmailParams{
		EmailAddress: p.Email,
		UserID:       p.UserID,
	}

	userEmailAccount, err := p.Q.GetEmail(p.Ctx, getEmailParams)
	if err != nil {
		return err
	}

	holdFilterID := userEmailAccount.HoldFilterID

	existingFilter, err := srv.Users.Settings.Filters.Get(user, holdFilterID.String).Do()
	if err != nil {
		return err
	}

	if existingFilter.Id != "" {
		err = srv.Users.Settings.Filters.Delete(user, existingFilter.Id).Do()
		if err != nil {
			return err
		}
	}

	f := &gmail.Filter{
		Criteria: filterCriteria,
		Action: &gmail.FilterAction{
			RemoveLabelIds: []string{"INBOX"},
			AddLabelIds:    []string{customLabelID},
		},
	}

	filter, err := srv.Users.Settings.Filters.Create(user, f).Do()
	if err != nil {
		return err
	}

	updateHoldFilterParams := sqlc.UpdateHoldFilterIDParams{
		HoldFilterID: pgtype.Text{
			String: filter.Id,
			Valid:  true,
		},
		EmailAddress: p.Email,
		UserID:       p.UserID,
	}

	err = p.Q.UpdateHoldFilterID(p.Ctx, updateHoldFilterParams)
	if err != nil {
		return err
	}

	return nil
}
