-- +goose Up
-- +goose StatementBegin
ALTER TABLE IF EXISTS cart_items
ADD COLUMN IF NOT EXISTS is_selected boolean default false;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE IF EXISTS cart_items
DROP COLUMN IF EXISTS is_selected;
-- +goose StatementEnd
