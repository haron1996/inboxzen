-- name: AddUser :one
insert into users (id) values ($1) returning *;

-- name: GetUser :one
select * from users where id = $1 limit 1;

-- name: UpdateLastLogin :exec
update users set last_login = $1 where id = $2;