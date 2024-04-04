-- +goose Up
drop table if exists deliveryTime cascade;

create table deliveryTime (
    delivery_time time not null,
    delivery_am_pm text check (delivery_am_pm in ('AM', 'PM')),
    date_added timestamptz not null default current_timestamp,
    date_updated timestamptz,
    email_address text not null,
    foreign key (email_address) references email(email_address) on delete cascade
);

-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
drop table if exists deliveryTime cascade;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
