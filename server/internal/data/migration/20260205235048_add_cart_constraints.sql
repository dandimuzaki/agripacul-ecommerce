-- +goose Up
-- +goose StatementBegin
-- unique cart per customer
ALTER TABLE IF EXISTS carts
ADD CONSTRAINT IF NOT EXISTS carts_customer_unique
UNIQUE (customer_id);

-- prevent duplicate SKU in same cart
ALTER TABLE IF EXISTS cart_items
ADD CONSTRAINT IF NOT EXISTS cart_items_cart_sku_unique
UNIQUE (cart_id, sku_id);

-- quantity must be positive
ALTER TABLE IF EXISTS cart_items
ADD CONSTRAINT IF NOT EXISTS cart_items_quantity_positive
CHECK (quantity > 0);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE carts DROP CONSTRAINT IF EXISTS carts_customer_unique;
ALTER TABLE cart_items DROP CONSTRAINT IF EXISTS cart_items_cart_sku_unique;
ALTER TABLE cart_items DROP CONSTRAINT IF EXISTS cart_items_quantity_positive;
-- +goose StatementEnd
