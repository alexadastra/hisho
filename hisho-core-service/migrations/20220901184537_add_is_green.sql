-- +goose Up
-- +goose StatementBegin
ALTER TABLE tasks ADD COLUMN is_green BOOLEAN NOT NULL DEFAULT false;
ALTER TABLE tasks RENAME done_at TO closed_at;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE tasks DROP COLUMN is_green;
ALTER TABLE tasks RENAME closed_at TO done_at;
-- +goose StatementEnd
