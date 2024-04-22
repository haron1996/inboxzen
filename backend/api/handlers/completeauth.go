package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/charmbracelet/log"
	"github.com/haron1996/inboxzen/api"
	"github.com/haron1996/inboxzen/apperror"
	"github.com/haron1996/inboxzen/messages"
	"github.com/haron1996/inboxzen/paseto"
	"github.com/haron1996/inboxzen/sqlc"
	"github.com/haron1996/inboxzen/utils"
	"github.com/haron1996/inboxzen/viper"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/oauth2"
)

type authCode struct {
	Code string `json:"code"`
}

type userInfo struct {
	ID      string
	Name    string
	Email   string
	Picture string
}

var (
	issuedAt = time.Now().UTC()

	duration = 24 * time.Hour
)

func CompleteGoogleAuth(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	var req *authCode

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
		return &apperror.APPError{
			Message: "Error returning new decoder",
			Code:    500,
			Err:     err,
		}
	}

	config, err := utils.ConstructGoogleConfig()
	if err != nil {
		api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
		return &apperror.APPError{
			Message: "Error constructing config from json",
			Code:    500,
			Err:     err,
		}
	}

	tok, err := config.Exchange(ctx, req.Code)
	if err != nil {
		api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
		return &apperror.APPError{
			Message: "Error converting authorization code to token",
			Code:    500,
			Err:     err,
		}
	}

	b, err := json.Marshal(tok)
	if err != nil {
		api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
		return &apperror.APPError{
			Message: "Error marshalling token",
			Code:    500,
			Err:     err,
		}
	}

	client := oauth2.NewClient(ctx, oauth2.StaticTokenSource(tok))

	userInfoUrl := fmt.Sprintf("https://www.googleapis.com/oauth2/v2/userinfo?alt=json&access_token=%s", tok.AccessToken)

	response, err := client.Get(userInfoUrl)
	if err != nil {
		api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
		return &apperror.APPError{
			Message: "Error getting user info",
			Code:    500,
			Err:     err,
		}
	}

	defer response.Body.Close()

	var ui *userInfo

	err = json.NewDecoder(response.Body).Decode(&ui)
	if err != nil {
		api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
		return &apperror.APPError{
			Message: "Error decoding user info",
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

	cookie, err := r.Cookie("session")
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			email, err := q.GetEmailByEmailAddress(ctx, ui.Email)
			if err != nil {
				switch {
				case errors.Is(err, pgx.ErrNoRows):

					log.Info("register user")

					user, err := q.AddUser(ctx, ui.ID)
					if err != nil {
						api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
						return &apperror.APPError{
							Message: "Error adding user",
							Code:    500,
							Err:     err,
						}
					}

					userID := user.ID

					accountParams := sqlc.AddEmailParams{
						ID:             utils.RandomString(),
						EmailAddress:   ui.Email,
						AccountName:    ui.Name,
						ProfilePicture: ui.Picture,
						UserID:         userID,
						Primaryaccount: true,
						Oauth2Token:    b,
					}

					email, err := q.AddEmail(ctx, accountParams)
					if err != nil {
						api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
						return &apperror.APPError{
							Message: "Error adding account",
							Code:    500,
							Err:     err,
						}
					}

					// insert default keywords
					err = addDefaultKeywords(q, email.ID, ctx)
					if err != nil {
						api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
						return &apperror.APPError{
							Message: "Error adding default keyword",
							Code:    500,
							Err:     err,
						}
					}

					t, p, err := paseto.CreateToken(userID, ui.Email, issuedAt, duration)
					if err != nil {
						api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
						return &apperror.APPError{
							Message: "Error creating token",
							Code:    500,
							Err:     err,
						}
					}

					lastLoginParams := sqlc.UpdateLastLoginParams{
						LastLogin: pgtype.Timestamptz{
							Time:  issuedAt,
							Valid: true,
						},
						ID: userID,
					}

					err = q.UpdateLastLogin(ctx, lastLoginParams)
					if err != nil {
						api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
						return &apperror.APPError{
							Message: "Error updating user last login",
							Code:    500,
							Err:     err,
						}
					}

					session := http.Cookie{
						Name:     "session",
						Value:    t,
						Path:     "/",
						Secure:   true,
						SameSite: http.SameSiteStrictMode,
						HttpOnly: true,
						Expires:  p.Expiry,
					}

					http.SetCookie(w, &session)

					api.ReturnResponse(w, 200, nil, true, messages.OK)

					return nil
				default:
					api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
					return &apperror.APPError{
						Message: "Error getting account",
						Code:    500,
						Err:     err,
					}
				}
			}

			log.Info("log user in")

			userID := email.UserID

			t, p, err := paseto.CreateToken(userID, ui.Email, issuedAt, duration)
			if err != nil {
				api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
				return &apperror.APPError{
					Message: "Error creating token",
					Code:    500,
					Err:     err,
				}
			}

			lastLoginParams := sqlc.UpdateLastLoginParams{
				LastLogin: pgtype.Timestamptz{
					Time:  issuedAt,
					Valid: true,
				},
				ID: userID,
			}

			err = q.UpdateLastLogin(ctx, lastLoginParams)
			if err != nil {
				api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
				return &apperror.APPError{
					Message: "Error updating user last login",
					Code:    500,
					Err:     err,
				}
			}

			session := http.Cookie{
				Name:     "session",
				Value:    t,
				Path:     "/",
				Secure:   true,
				SameSite: http.SameSiteStrictMode,
				HttpOnly: true,
				Expires:  p.Expiry,
			}

			http.SetCookie(w, &session)

			api.ReturnResponse(w, 200, nil, true, "ok")

			return nil
		default:
			api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
			return &apperror.APPError{
				Message: "Error getting cookie",
				Code:    500,
				Err:     err,
			}
		}
	}

	log.Info("add email to user emails")

	payload, err := paseto.VerifyToken(cookie.Value)
	if err != nil {
		api.ReturnResponse(w, 500, nil, true, err.Error())
		return &apperror.APPError{
			Message: "Error verify token",
			Code:    500,
			Err:     err,
		}
	}

	userID := payload.UserID

	_, err = q.GetUser(ctx, userID)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			api.ReturnResponse(w, 500, nil, true, "user not found")
			return &apperror.APPError{
				Message: "Error getting user",
				Code:    404,
				Err:     err,
			}
		default:
			api.ReturnResponse(w, 401, nil, true, messages.ErrInternalServer)
			return &apperror.APPError{
				Message: "Error getting user",
				Code:    500,
				Err:     err,
			}
		}
	}

	emails, err := q.GetAllUserEmails(ctx, userID)
	if err != nil {
		api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
		return &apperror.APPError{
			Message: "Error getting user emails",
			Code:    500,
			Err:     err,
		}
	}

	for _, email := range emails {
		if email.EmailAddress == ui.Email {
			api.ReturnResponse(w, 409, nil, true, "email address already registered")
			return &apperror.APPError{
				Message: "Email address already registered",
				Code:    409,
				Err:     errors.New("email address already registered"),
			}
		}
	}

	addEmailParams := sqlc.AddEmailParams{
		ID:             utils.RandomString(),
		EmailAddress:   ui.Email,
		AccountName:    ui.Name,
		ProfilePicture: ui.Picture,
		UserID:         userID,
		Primaryaccount: false,
		Oauth2Token:    b,
	}

	email, err := q.AddEmail(ctx, addEmailParams)
	if err != nil {
		api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
		return &apperror.APPError{
			Message: "Error adding email",
			Code:    500,
			Err:     err,
		}
	}

	// insert default keywords
	err = addDefaultKeywords(q, email.ID, ctx)
	if err != nil {
		api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
		return &apperror.APPError{
			Message: "Error adding default keyword",
			Code:    500,
			Err:     err,
		}
	}

	t, pload, err := paseto.CreateToken(userID, email.EmailAddress, issuedAt, duration)
	if err != nil {
		api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
		return &apperror.APPError{
			Message: "Error creating token",
			Code:    500,
			Err:     err,
		}
	}

	lastLoginParams := sqlc.UpdateLastLoginParams{
		LastLogin: pgtype.Timestamptz{
			Time:  issuedAt,
			Valid: true,
		},
		ID: userID,
	}

	err = q.UpdateLastLogin(ctx, lastLoginParams)
	if err != nil {
		api.ReturnResponse(w, 500, nil, true, messages.ErrInternalServer)
		return &apperror.APPError{
			Message: "Error updating user last login",
			Code:    500,
			Err:     err,
		}
	}

	session := http.Cookie{
		Name:     "session",
		Value:    t,
		Path:     "/",
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		HttpOnly: true,
		Expires:  pload.Expiry,
	}

	http.SetCookie(w, &session)

	api.ReturnResponse(w, 200, nil, true, messages.OK)

	return nil
}

func addDefaultKeywords(q *sqlc.Queries, emailID string, ctx context.Context) error {
	defaultKeywords := []string{
		"otp",
		"code",
		"confirm",
		"confirmation",
		"verify",
		"verification",
		"reset",
		"password",
		"2fa",
		"welcome",
	}

	for _, dk := range defaultKeywords {
		params := sqlc.AddKeywordParams{
			ID:      utils.RandomString(),
			Keyword: dk,
			EmailID: emailID,
		}

		_, err := q.AddKeyword(ctx, params)
		if err != nil {
			return err
		}
	}

	return nil
}
