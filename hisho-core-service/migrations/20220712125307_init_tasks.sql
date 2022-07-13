-- +goose Up
-- +goose StatementBegin
CREATE TABLE tasks (
    id bigint primary key,
    title varchar not null,
    term int not null default 0, 
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    done_at timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tasks;
-- +goose StatementEnd
