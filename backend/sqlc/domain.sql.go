// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: domain.sql

package sqlc

import (
	"context"
)

const addDomain = `-- name: AddDomain :one
insert into vipDomain (domain_name, email_address)
values ($1, $2)
returning domain_name, date_added, email_address
`

type AddDomainParams struct {
	DomainName   string `json:"domain_name"`
	EmailAddress string `json:"email_address"`
}

func (q *Queries) AddDomain(ctx context.Context, arg AddDomainParams) (Vipdomain, error) {
	row := q.db.QueryRow(ctx, addDomain, arg.DomainName, arg.EmailAddress)
	var i Vipdomain
	err := row.Scan(&i.DomainName, &i.DateAdded, &i.EmailAddress)
	return i, err
}

const deleteDomains = `-- name: DeleteDomains :exec
delete from vipDomain where email_address = $1
`

func (q *Queries) DeleteDomains(ctx context.Context, emailAddress string) error {
	_, err := q.db.Exec(ctx, deleteDomains, emailAddress)
	return err
}

const getDomains = `-- name: GetDomains :many
select domain_name, date_added, email_address from vipDomain where email_address = $1
`

func (q *Queries) GetDomains(ctx context.Context, emailAddress string) ([]Vipdomain, error) {
	rows, err := q.db.Query(ctx, getDomains, emailAddress)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Vipdomain{}
	for rows.Next() {
		var i Vipdomain
		if err := rows.Scan(&i.DomainName, &i.DateAdded, &i.EmailAddress); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}