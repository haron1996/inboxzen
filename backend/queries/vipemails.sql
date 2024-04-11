-- name: AddVipEmail :one
insert into vipEmailAddress (vip_email_address, email_address)
values ($1, $2)
returning *;

-- name: GetVipEmails :many
select * from vipEmailAddress where email_address = $1;

-- name: DeleteVipEmails :exec
delete from vipEmailAddress where email_address = $1;