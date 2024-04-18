-- name: AddEmail :one
insert into email (id, email_address, account_name, profile_picture, user_id, primaryAccount, oauth2_token) values ($1, $2, $3, $4, $5, $6, $7) returning *;

-- name: GetEmailByEmailAddress :one
select * from email where email_address = $1 limit 1;

-- name: GetAllUserEmails :many
select * from email where user_id = $1;

-- name: GetEmail :one
select * from email where email_address = $1 and user_id = $2 limit 1;

-- name: Activate :exec
update email set running = true where email_address = $1 and user_id = $2;

-- name: UpdateHoldFilterID :exec
update email set hold_filter_id = $1 where email_address = $2 and user_id = $3;