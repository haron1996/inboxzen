-- +goose Up
drop table if exists vipEmailAddress cascade;

create table vipEmailAddress(
    id text not null primary key,
    vip_email_address text not null,
    date_added timestamptz not null default current_timestamp,
    email_id text not null,
    unique(vip_email_address, email_id),
    foreign key (email_id) references email(id) on delete cascade
);

-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
drop table if exists vipEmailAddress cascade;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
