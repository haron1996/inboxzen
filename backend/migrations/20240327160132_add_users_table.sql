-- +goose Up
drop table if exists users cascade;

create table users (
    id text not null primary key,
    register_date timestamptz not null default current_timestamp,
    last_login timestamptz
);

-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
drop table if exists users cascade;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
