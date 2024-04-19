-- +goose Up
drop table if exists vipDomain cascade;

create table vipDomain (
    id text primary key,
    domain_name text not null,
    date_added timestamptz not null default current_timestamp,
    email_id text not null,
    unique(domain_name, email_id),
    foreign key (email_id) references email (id) on delete cascade
);

-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
drop table if exists vipDomain cascade;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
