-- +goose Up
-- +goose StatementBegin
-- promotion date window
ALTER TABLE IF EXISTS promotions
ADD CONSTRAINT IF NOT EXISTS promotions_date_valid
CHECK (start_date < end_date);

-- discount must be positive
ALTER TABLE IF EXISTS promotions
ADD CONSTRAINT IF NOT EXISTS promotions_discount_positive
CHECK (discount_value > 0);

-- minimum order >= 0
ALTER TABLE IF EXISTS promotions
ADD CONSTRAINT IF NOT EXISTS promotions_min_order_non_negative
CHECK (minimum_order_value >= 0);

-- maximum discount >= 0
ALTER TABLE IF EXISTS promotions
ADD CONSTRAINT IF NOT EXISTS promotions_max_discount_non_negative
CHECK (maximum_discount IS NULL OR maximum_discount >= 0);

-- percentage discount ≤ 100
ALTER TABLE IF EXISTS promotions
ADD CONSTRAINT IF NOT EXISTS promotions_percentage_discount_valid
CHECK (
  discount_type <> 'percentage'
  OR discount_value <= 100
);

-- usage limit ≥ 0
ALTER TABLE IF EXISTS promotions
ADD CONSTRAINT IF NOT EXISTS promotions_usage_limit_valid
CHECK (usage_limit >= 0);

-- unique voucher code
CREATE UNIQUE INDEX IF NOT EXISTS promotions_voucher_code_lower_idx
ON promotions (LOWER(voucher_code))
WHERE voucher_code IS NOT NULL;

-- prevent duplicate product per promotion
ALTER TABLE IF EXISTS promo_products
ADD CONSTRAINT IF NOT EXISTS promo_products_unique
UNIQUE (promotion_id, product_id);

-- one promotion per order
ALTER TABLE IF EXISTS promo_usages
ADD CONSTRAINT IF NOT EXISTS promo_usages_unique_order
UNIQUE (order_id);

-- prevent same customer using same promo twice (optional rule)
ALTER TABLE IF EXISTS promo_usages
ADD CONSTRAINT IF NOT EXISTS promo_usages_customer_promo_unique
UNIQUE (promotion_id, customer_id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE promotions DROP CONSTRAINT IF EXISTS promotions_date_valid;
ALTER TABLE promotions DROP CONSTRAINT IF EXISTS promotions_discount_positive;
ALTER TABLE promotions DROP CONSTRAINT IF EXISTS promotions_min_order_non_negative;
ALTER TABLE promotions DROP CONSTRAINT IF EXISTS promotions_max_discount_non_negative;
ALTER TABLE promotions DROP CONSTRAINT IF EXISTS promotions_percentage_discount_valid;
ALTER TABLE promotions DROP CONSTRAINT IF EXISTS promotions_usage_limit_valid;
DROP INDEX IF EXISTS promotions_voucher_code_lower_idx;
ALTER TABLE promo_products DROP CONSTRAINT IF EXISTS promo_products_unique;
ALTER TABLE promo_usages DROP CONSTRAINT IF EXISTS promo_usages_unique_order;
ALTER TABLE promo_usages DROP CONSTRAINT IF EXISTS promo_usages_customer_promo_unique;
-- +goose StatementEnd
