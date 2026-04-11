-- +goose Up
-- +goose StatementBegin
ALTER TABLE IF EXISTS promotions
RENAME COLUMN IF EXISTS is_shown TO is_public;
ALTER TABLE IF EXISTS promotions
ADD COLUMN IF NOT EXISTS used_count INT DEFAULT 0;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE IF EXISTS promotions
RENAME COLUMN IF EXISTS is_public TO is_shown;
ALTER TABLE IF EXISTS promotions
DROP COLUMN IF EXISTS used_count;
-- +goose StatementEnd
