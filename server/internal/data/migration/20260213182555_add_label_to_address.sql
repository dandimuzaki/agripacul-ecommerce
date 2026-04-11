-- +goose Up
-- +goose StatementBegin
ALTER TABLE addresses
ADD COLUMN IF NOT EXISTS recipient_name varchar(255),
ADD COLUMN IF NOT EXISTS label varchar(100),
ADD COLUMN IF NOT EXISTS is_default boolean default false;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE addresses
DROP COLUMN IF EXISTS recipient_name,
DROP COLUMN IF EXISTS label,
DROP COLUMN IF EXISTS is_default;
-- +goose StatementEnd
