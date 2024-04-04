-- +goose Up
drop table if exists vipDomains cascade;

create table vipDomains (
    domain_name text not null,
    date_added timestamptz not null default current_timestamp,
    email_address text not null,
    foreign key (email_address) references email (email_address) on delete cascade
);

-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
drop table if exists vipDomains cascade;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
