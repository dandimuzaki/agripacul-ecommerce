-- +goose Up
-- +goose StatementBegin
-- change table name
ALTER TABLE IF EXISTS images
ALTER COLUMN sku_id DROP NOT NULL;

ALTER TABLE IF EXISTS images
ADD COLUMN IF NOT EXISTS product_id int NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE IF EXISTS images
DROP COLUMN IF EXISTS product_id;

ALTER TABLE images
ALTER COLUMN sku_id SET NOT NULL;

-- +goose StatementEnd
