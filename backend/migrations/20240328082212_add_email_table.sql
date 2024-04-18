-- +goose Up
drop table if exists email cascade;

create table email(
    id text not null primary key,
    email_address text not null unique,
    account_name text not null,
    profile_picture text not null,
    date_added timestamptz not null default current_timestamp,
    user_id text not null,
    primaryAccount boolean not null,
    running boolean not null default false,
    oauth2_token bytea not null,
    hold_filter_id text,
    block_filter_id text,
    unique(user_id, email_address),
    foreign key (user_id) references users(id) on delete cascade
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
drop table if exists email cascade;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
