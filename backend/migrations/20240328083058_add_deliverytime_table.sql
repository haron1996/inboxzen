-- +goose Up
drop table if exists deliveryTime cascade;

create table deliveryTime (
    id text primary key,
    delivery_time time not null,
    date_added timestamptz not null default current_timestamp,
    email_id text not null,
    unique(delivery_time, email_id),
    foreign key (email_id) references email(id) on delete cascade
);

-- +goose StatementBegin
-- +goose StatementEnd

-- +goose Down
drop table if exists deliveryTime cascade;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
