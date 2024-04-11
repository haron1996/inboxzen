-- +goose Up
drop table if exists vipKeyword cascade;

create table vipKeyword (
    keyword text not null,
    date_added timestamptz not null default current_timestamp,
    email_address text not null,
    unique(keyword, email_address),
    foreign key (email_address) references email(email_address) on delete cascade
);

-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
drop table if exists vipKeyword cascade;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
