-- +goose Up
drop table if exists sender_email cascade;

create table sender_email (
    sender_id text,
    email_id text,
    primary key (sender_id, email_id),
    foreign key (sender_id) references sender(id) on delete cascade,
    foreign key (email_id) references email(id) on delete cascade
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
drop table if exists sender_email cascade;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
