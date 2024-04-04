-- +goose Up
drop table if exists email cascade;

create table email(
    email_address text not null primary key,
    account_name text not null,
    profile_picture text not null,
    date_added timestamptz not null default current_timestamp,
    user_id text not null,
    primaryAccount boolean not null,
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
