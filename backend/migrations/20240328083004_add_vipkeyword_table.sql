-- +goose Up
drop table if exists vipKeyword cascade;

create table vipKeyword (
    id text not null primary key,
    keyword text not null,
    date_added timestamptz not null default current_timestamp,
    email_id text not null,
    unique(keyword, email_id),
    foreign key (email_id) references email(id) on delete cascade
);

-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
drop table if exists vipKeyword cascade;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
