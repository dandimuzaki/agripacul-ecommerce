-- +goose Up
-- +goose StatementBegin
-- add new column
ALTER TABLE IF EXISTS products
ADD COLUMN IF NOT EXISTS slug TEXT NOT NULL,
ADD COLUMN IF NOT EXISTS main_image_url TEXT,
ADD COLUMN IF NOT EXISTS main_image_public_id TEXT,
ADD COLUMN IF NOT EXISTS average_rating DECIMAL DEFAULT 0,
ADD COLUMN IF NOT EXISTS review_count BIGINT DEFAULT 0,
ADD COLUMN IF NOT EXISTS sold_count BIGINT DEFAULT 0,
ADD COLUMN IF NOT EXISTS min_price NUMERIC(12,2) DEFAULT 0,
ADD COLUMN IF NOT EXISTS max_price NUMERIC(12,2) DEFAULT 0;

-- Uniqueness
CREATE UNIQUE INDEX IF NOT EXISTS products_slug_lower_idx
ON products (LOWER(slug));

-- Data integrity
ALTER TABLE IF EXISTS products
ADD CONSTRAINT IF NOT EXISTS avg_rating_range
CHECK (average_rating >= 0 AND average_rating <= 5),

ADD CONSTRAINT IF NOT EXISTS review_count_non_negative
CHECK (review_count >= 0),

ADD CONSTRAINT IF NOT EXISTS sold_count_non_negative
CHECK (sold_count >= 0),

ADD CONSTRAINT pIF NOT EXISTS rice_non_negative
CHECK (min_price >= 0 AND max_price >= 0),

ADD CONSTRAINT IF NOT EXISTS min_price_le_max_price
CHECK (min_price <= max_price);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE IF EXISTS products
DROP COLUMN IF EXISTS slug,
DROP COLUMN IF EXISTS main_image_url,
DROP COLUMN IF EXISTS average_rating,
DROP COLUMN IF EXISTS review_count,
DROP COLUMN IF EXISTS sold_count,
DROP COLUMN IF EXISTS min_price,
DROP COLUMN IF EXISTS max_price;

DROP INDEX IF EXISTS products_slug_lower_idx;
ALTER TABLE IF EXISTS products
DROP CONSTRAINT IF EXISTS avg_rating_range,
DROP CONSTRAINT IF EXISTS review_count_non_negative,
DROP CONSTRAINT IF EXISTS sold_count_non_negative,
DROP CONSTRAINT IF EXISTS price_non_negative,
DROP CONSTRAINT IF EXISTS min_price_le_max_price;
-- +goose StatementEnd
