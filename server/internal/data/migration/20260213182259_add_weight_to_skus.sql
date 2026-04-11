-- +goose Up
-- +goose StatementBegin
ALTER TABLE IF EXISTS skus
ADD COLUMN IF NOT EXISTS weight_gram decimal default 1000,
ADD COLUMN IF NOT EXISTS sale_price decimal,
ADD COLUMN IF NOT EXISTS min_stock int default 0;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE IF EXISTS skus
DROP COLUMN IF EXISTS weight_gram,
DROP COLUMN IF EXISTS sale_price,
DROP COLUMN IF EXISTS min_stock;
-- +goose StatementEnd
