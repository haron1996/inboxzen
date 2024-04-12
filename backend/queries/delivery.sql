-- name: SetDeliveryTime :one
insert into deliveryTime (delivery_time, email_address)
values ($1, $2)
returning *;

-- name: DeleteDeliveryTime :exec
delete from deliveryTime where email_address = $1;

-- name: GetDeliveryTimes :many
select * from deliveryTime where email_address = $1;