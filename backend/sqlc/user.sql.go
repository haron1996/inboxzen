// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: user.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const addUser = `-- name: AddUser :one
insert into users (id) values ($1) returning id, register_date, last_login
`

func (q *Queries) AddUser(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRow(ctx, addUser, id)
	var i User
	err := row.Scan(&i.ID, &i.RegisterDate, &i.LastLogin)
	return i, err
}

const getUser = `-- name: GetUser :one
select id, register_date, last_login from users where id = $1 limit 1
`

func (q *Queries) GetUser(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRow(ctx, getUser, id)
	var i User
	err := row.Scan(&i.ID, &i.RegisterDate, &i.LastLogin)
	return i, err
}

const updateLastLogin = `-- name: UpdateLastLogin :exec
update users set last_login = $1 where id = $2
`

type UpdateLastLoginParams struct {
	LastLogin pgtype.Timestamptz `json:"last_login"`
	ID        string             `json:"id"`
}

func (q *Queries) UpdateLastLogin(ctx context.Context, arg UpdateLastLoginParams) error {
	_, err := q.db.Exec(ctx, updateLastLogin, arg.LastLogin, arg.ID)
	return err
}
