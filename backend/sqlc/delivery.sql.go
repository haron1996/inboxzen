// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: delivery.sql

package sqlc

import (
	"context"
)

const deleteDeliveryTime = `-- name: DeleteDeliveryTime :exec
delete from deliveryTime where email_address = $1
`

func (q *Queries) DeleteDeliveryTime(ctx context.Context, emailAddress string) error {
	_, err := q.db.Exec(ctx, deleteDeliveryTime, emailAddress)
	return err
}

const getDeliveryTimes = `-- name: GetDeliveryTimes :many
select delivery_time, delivery_am_pm, date_added, date_updated, email_address from deliveryTime where email_address = $1
`

func (q *Queries) GetDeliveryTimes(ctx context.Context, emailAddress string) ([]Deliverytime, error) {
	rows, err := q.db.Query(ctx, getDeliveryTimes, emailAddress)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Deliverytime{}
	for rows.Next() {
		var i Deliverytime
		if err := rows.Scan(
			&i.DeliveryTime,
			&i.DeliveryAmPm,
			&i.DateAdded,
			&i.DateUpdated,
			&i.EmailAddress,
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

const setDeliveryTime = `-- name: SetDeliveryTime :one
insert into deliveryTime (delivery_time, delivery_am_pm, email_address)
values ($1, $2, $3)
returning delivery_time, delivery_am_pm, date_added, date_updated, email_address
`

type SetDeliveryTimeParams struct {
	DeliveryTime string `json:"delivery_time"`
	DeliveryAmPm string `json:"delivery_am_pm"`
	EmailAddress string `json:"email_address"`
}

func (q *Queries) SetDeliveryTime(ctx context.Context, arg SetDeliveryTimeParams) (Deliverytime, error) {
	row := q.db.QueryRow(ctx, setDeliveryTime, arg.DeliveryTime, arg.DeliveryAmPm, arg.EmailAddress)
	var i Deliverytime
	err := row.Scan(
		&i.DeliveryTime,
		&i.DeliveryAmPm,
		&i.DateAdded,
		&i.DateUpdated,
		&i.EmailAddress,
	)
	return i, err
}
