-- +goose Up
drop table if exists vipDomain cascade;

create table vipDomain (
    id text not null primary key,
    domain_name text not null,
    date_added timestamptz not null default current_timestamp,
    email_address text not null,
    unique(domain_name, email_address),
    foreign key (email_address) references email (email_address) on delete cascade
);

-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
drop table if exists vipDomain cascade;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
