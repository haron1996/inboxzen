-- name: SetDeliveryTime :one
insert into deliveryTime (id, delivery_time, email_address)
values ($1, $2, $3)
returning *;

-- name: DeleteDeliveryTime :one
delete from deliveryTime where id = $1 returning *;

-- name: GetDeliveryTimes :many
select * from deliveryTime where email_address = $1;