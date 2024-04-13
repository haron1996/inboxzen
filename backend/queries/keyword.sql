-- name: AddKeyword :one
insert into vipKeyword (id, keyword, email_address)
values ($1, $2, $3)
returning *;

-- name: GetKeywords :many
select * from vipKeyword where email_address = $1;

-- name: DeleteKeywords :exec
delete from vipKeyword where email_address = $1;