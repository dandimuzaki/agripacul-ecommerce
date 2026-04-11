-- +goose Up
-- +goose StatementBegin
ALTER TABLE IF EXISTS categories
ADD COLUMN IF NOT EXISTS icon_public_id text;

ALTER TABLE IF EXISTS banners
ADD COLUMN IF NOT EXISTS image_public_id text;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE IF EXISTS categories
DROP COLUMN IF EXISTS icon_public_id;

ALTER TABLE IF EXISTS banners
DROP COLUMN IF EXISTS image_public_id;
-- +goose StatementEnd
