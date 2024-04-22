-- name: SetDeliveryTime :one
insert into deliveryTime (id, delivery_time, email_id)
values ($1, $2, $3)
returning *;

-- name: DeleteDeliveryTimes :exec
delete from deliveryTime where email_id = $1 returning *;

-- name: GetDeliveryTimes :many
select * from deliveryTime where email_id = $1;