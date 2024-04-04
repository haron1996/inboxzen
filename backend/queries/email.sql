-- name: AddEmail :one
insert into email (email_address, account_name, profile_picture, user_id, primaryAccount) values ($1, $2, $3, $4, $5) returning *;

-- name: GetEmailByEmailAddress :one
select * from email where email_address = $1 limit 1;

-- name: GetAllUserEmails :many
select * from email where user_id = $1;