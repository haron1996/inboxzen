-- name: AddVipEmail :one
insert into vipEmailAddress (id, vip_email_address, email_id)
values ($1, $2, $3)
returning *;

-- name: GetVipEmails :many
select * from vipEmailAddress where email_id = $1;

-- name: DeleteVipEmails :exec
delete from vipEmailAddress where email_id = $1;