-- +goose Up
-- +goose StatementBegin
ALTER TABLE IF EXISTS payment_methods
ADD COLUMN IF NOT EXISTS icon_url TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE IF EXISTS payment_methods
DROP COLUMN IF EXISTS icon_url;
-- +goose StatementEnd
