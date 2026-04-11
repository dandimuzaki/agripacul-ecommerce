-- +goose Up
-- +goose StatementBegin

-- subtotal must be positive
ALTER TABLE IF EXISTS orders
ADD CONSTRAINT IF NOT EXISTS orders_subtotal_non_negative
CHECK (subtotal >= 0);

-- discount amount must be positive
ALTER TABLE IF EXISTS orders
ADD CONSTRAINT IF NOT EXISTS orders_discount_non_negative
CHECK (discount_amount >= 0);

-- total must be positive
ALTER TABLE IF EXISTS orders
ADD CONSTRAINT IF NOT EXISTS orders_total_non_negative
CHECK (total >= 0);

-- quantity must be positive
ALTER TABLE IF EXISTS order_items
ADD CONSTRAINT IF NOT EXISTS order_items_quantity_positive
CHECK (quantity > 0);

-- unit price must be positive
ALTER TABLE IF EXISTS order_items
ADD CONSTRAINT IF NOT EXISTS order_items_unit_price_non_negative
CHECK (unit_price >= 0);

-- total price must be positive
ALTER TABLE IF EXISTS order_items
ADD CONSTRAINT IF NOT EXISTS order_items_total_price_non_negative
CHECK (total_price >= 0);

-- total price must equal to unit price * quantity
ALTER TABLE IF EXISTS order_items
ADD CONSTRAINT IF NOT EXISTS order_items_total_price_valid
CHECK (total_price = unit_price * quantity);

-- prevent duplicate sku per order
ALTER TABLE IF EXISTS order_items
ADD CONSTRAINT IF NOT EXISTS order_items_order_sku_unique
UNIQUE (order_id, sku_id);

-- promo snapshot consistency
ALTER TABLE IF EXISTS orders
ADD CONSTRAINT IF NOT EXISTS orders_promo_snapshot_consistency
CHECK (
  promotion_id IS NULL OR promotion_snapshot IS NOT NULL
);

-- unique tracking number
CREATE UNIQUE INDEX IF NOT EXISTS orders_tracking_number_idx
ON order_shippings (tracking_number)
WHERE tracking_number IS NOT NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE IF EXISTS orders DROP CONSTRAINT IF EXISTS orders_subtotal_non_negative;
ALTER TABLE IF EXISTS orders DROP CONSTRAINT IF EXISTS orders_discount_non_negative;
ALTER TABLE IF EXISTS orders DROP CONSTRAINT IF EXISTS orders_total_non_negative;
ALTER TABLE IF EXISTS order_items DROP CONSTRAINT IF EXISTS order_items_quantity_positive;
ALTER TABLE IF EXISTS order_items DROP CONSTRAINT IF EXISTS order_items_unit_price_non_negative;
ALTER TABLE IF EXISTS order_items DROP CONSTRAINT IF EXISTS order_items_total_price_non_negative;
ALTER TABLE IF EXISTS order_items DROP CONSTRAINT IF EXISTS order_items_total_price_valid;
ALTER TABLE IF EXISTS order_items DROP CONSTRAINT IF EXISTS order_items_order_sku_unique;
ALTER TABLE IF EXISTS orders DROP CONSTRAINT IF EXISTS orders_promo_snapshot_consistency;
DROP INDEX IF EXISTS orders_tracking_number_idx;
-- +goose StatementEnd
