// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: keyword.sql

package sqlc

import (
	"context"
)

const addKeyword = `-- name: AddKeyword :one
insert into vipKeyword (id, keyword, email_id)
values ($1, $2, $3)
returning id, keyword, date_added, email_id
`

type AddKeywordParams struct {
	ID      string `json:"id"`
	Keyword string `json:"keyword"`
	EmailID string `json:"email_id"`
}

func (q *Queries) AddKeyword(ctx context.Context, arg AddKeywordParams) (Vipkeyword, error) {
	row := q.db.QueryRow(ctx, addKeyword, arg.ID, arg.Keyword, arg.EmailID)
	var i Vipkeyword
	err := row.Scan(
		&i.ID,
		&i.Keyword,
		&i.DateAdded,
		&i.EmailID,
	)
	return i, err
}

const deleteKeywords = `-- name: DeleteKeywords :exec
delete from vipKeyword where email_id = $1
`

func (q *Queries) DeleteKeywords(ctx context.Context, emailID string) error {
	_, err := q.db.Exec(ctx, deleteKeywords, emailID)
	return err
}

const getKeywords = `-- name: GetKeywords :many
select id, keyword, date_added, email_id from vipKeyword where email_id = $1
`

func (q *Queries) GetKeywords(ctx context.Context, emailID string) ([]Vipkeyword, error) {
	rows, err := q.db.Query(ctx, getKeywords, emailID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Vipkeyword{}
	for rows.Next() {
		var i Vipkeyword
		if err := rows.Scan(
			&i.ID,
			&i.Keyword,
			&i.DateAdded,
			&i.EmailID,
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
