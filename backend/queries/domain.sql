-- name: AddDomain :one
insert into vipDomain (id, domain_name, email_id)
values ($1, $2, $3)
returning *;

-- name: GetDomains :many
select * from vipDomain where email_id = $1;

-- name: DeleteDomains :exec
delete from vipDomain where email_id = $1;