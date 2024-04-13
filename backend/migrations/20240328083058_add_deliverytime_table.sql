-- +goose Up
drop table if exists deliveryTime cascade;

create table deliveryTime (
    id text not null primary key,
    delivery_time text not null,
    date_added timestamptz not null default current_timestamp,
    email_address text not null,
    unique(delivery_time, email_address),
    foreign key (email_address) references email(email_address) on delete cascade
);

-- +goose StatementBegin
SELECT 'up SQL query';
create or replace function delivery_time_to_lowercase()
returns trigger as $$
begin
    new.delivery_time = lower(new.delivery_time);
    return new;
end;
$$ language plpgsql;

create trigger delivery_time_lowercase_trigger
before insert or update on deliveryTime
for each row execute function delivery_time_to_lowercase();
-- +goose StatementEnd

-- +goose Down
drop table if exists deliveryTime cascade;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
