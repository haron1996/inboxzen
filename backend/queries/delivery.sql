-- name: SetDeliveryTime :one
insert into deliveryTime (id, delivery_time, email_id)
values ($1, $2, $3)
returning *;

-- name: DeleteDeliveryTime :one
delete from deliveryTime where id = $1 returning *;

-- name: GetDeliveryTimes :many
select * from deliveryTime where email_id = $1;