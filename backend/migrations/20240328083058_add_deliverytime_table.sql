-- +goose Up
drop table if exists deliveryTime cascade;

create table deliveryTime (
    delivery_time text not null,
    date_added timestamptz not null default current_timestamp,
    email_address text not null,
    unique(delivery_time, email_address),
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
