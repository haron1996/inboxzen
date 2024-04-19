-- +goose Up
drop table if exists sender cascade;

create table sender (
    id text primary key,
    sender_email_address text not null
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
drop table if exists sender cascade;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
