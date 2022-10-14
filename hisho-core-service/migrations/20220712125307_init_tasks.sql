-- +goose Up
-- +goose StatementBegin
CREATE TABLE tasks (
    id bigint primary key,
    title varchar not null,
    term int not null default 0, 
    is_green bool not null default false,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    closed_at timestamp,
    closed_reason text,
    is_archived bool not null default false
);

CREATE TABLE users (
    id bigint primary key,
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tasks;
-- +goose StatementEnd
