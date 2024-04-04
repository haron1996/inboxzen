-- +goose Up
drop table if exists vipEmailAddress cascade;

create table vipEmailAddress(
    vip_email_address text not null,
    date_added timestamptz not null default current_timestamp,
    email_address text not null,
    foreign key (email_address) references email(email_address) on delete cascade
);

-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
drop table if exists vipEmailAddress cascade;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
