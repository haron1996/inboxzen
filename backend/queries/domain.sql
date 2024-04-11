-- name: AddDomain :one
insert into vipDomain (domain_name, email_address)
values ($1, $2)
returning *;

-- name: GetDomains :many
select * from vipDomain where email_address = $1;

-- name: DeleteDomains :exec
delete from vipDomain where email_address = $1;