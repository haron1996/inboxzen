// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: email.sql

package sqlc

import (
	"context"
)

const addEmail = `-- name: AddEmail :one
insert into email (email_address, account_name, profile_picture, user_id, primaryAccount) values ($1, $2, $3, $4, $5) returning email_address, account_name, profile_picture, date_added, user_id, primaryaccount
`

type AddEmailParams struct {
	EmailAddress   string `json:"email_address"`
	AccountName    string `json:"account_name"`
	ProfilePicture string `json:"profile_picture"`
	UserID         string `json:"user_id"`
	Primaryaccount bool   `json:"primaryaccount"`
}

func (q *Queries) AddEmail(ctx context.Context, arg AddEmailParams) (Email, error) {
	row := q.db.QueryRow(ctx, addEmail,
		arg.EmailAddress,
		arg.AccountName,
		arg.ProfilePicture,
		arg.UserID,
		arg.Primaryaccount,
	)
	var i Email
	err := row.Scan(
		&i.EmailAddress,
		&i.AccountName,
		&i.ProfilePicture,
		&i.DateAdded,
		&i.UserID,
		&i.Primaryaccount,
	)
	return i, err
}

const getAllUserEmails = `-- name: GetAllUserEmails :many
select email_address, account_name, profile_picture, date_added, user_id, primaryaccount from email where user_id = $1
`

func (q *Queries) GetAllUserEmails(ctx context.Context, userID string) ([]Email, error) {
	rows, err := q.db.Query(ctx, getAllUserEmails, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Email{}
	for rows.Next() {
		var i Email
		if err := rows.Scan(
			&i.EmailAddress,
			&i.AccountName,
			&i.ProfilePicture,
			&i.DateAdded,
			&i.UserID,
			&i.Primaryaccount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getEmailByEmailAddress = `-- name: GetEmailByEmailAddress :one
select email_address, account_name, profile_picture, date_added, user_id, primaryaccount from email where email_address = $1 limit 1
`

func (q *Queries) GetEmailByEmailAddress(ctx context.Context, emailAddress string) (Email, error) {
	row := q.db.QueryRow(ctx, getEmailByEmailAddress, emailAddress)
	var i Email
	err := row.Scan(
		&i.EmailAddress,
		&i.AccountName,
		&i.ProfilePicture,
		&i.DateAdded,
		&i.UserID,
		&i.Primaryaccount,
	)
	return i, err
}
