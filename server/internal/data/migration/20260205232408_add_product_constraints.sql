-- +goose Up
-- +goose StatementBegin

-- unique category name
CREATE UNIQUE INDEX categories_name_lower_idx
ON categories (LOWER(name));

-- category name should be not an empty string
ALTER TABLE categories
ADD CONSTRAINT categories_name_not_empty
CHECK (length(trim(name)) > 0);

-- unique product name
CREATE UNIQUE INDEX products_category_name_idx
ON products (category_id, LOWER(name));

-- prevent category delete
ALTER TABLE products
ADD CONSTRAINT products_category_fk
FOREIGN KEY (category_id)
REFERENCES categories(id)
ON DELETE RESTRICT;

-- unique variant type per product
CREATE UNIQUE INDEX variant_types_product_name_idx
ON variant_types (product_id, LOWER(name));

-- variant type should be not an empty string
ALTER TABLE variant_types
ADD CONSTRAINT variant_types_name_not_empty
CHECK (length(trim(name)) > 0);

-- unique variant value per type
CREATE UNIQUE INDEX variant_values_type_value_idx
ON variant_values (variant_type_id, LOWER(value));

-- variant value should be not an empty string
ALTER TABLE variant_values
ADD CONSTRAINT variant_values_value_not_empty
CHECK (length(trim(value)) > 0);

-- unique sku code
CREATE UNIQUE INDEX skus_sku_code_lower_idx
ON skus (LOWER(sku_code));

-- price should be positive
ALTER TABLE skus
ADD CONSTRAINT skus_price_positive
CHECK (price >= 0);

-- stock should be positive
ALTER TABLE skus
ADD CONSTRAINT skus_stock_non_negative
CHECK (stock >= 0);

-- unique sku_id + variant_value_id
ALTER TABLE sku_variant_values
ADD CONSTRAINT sku_variant_unique
UNIQUE (sku_id, variant_value_id);

-- prevent empty string on image url
ALTER TABLE images
ADD CONSTRAINT image_url_not_empty
CHECK (length(trim(image_url)) > 0);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS categories_name_lower_idx;
ALTER TABLE categories DROP CONSTRAINT IF EXISTS categories_name_not_empty;
DROP INDEX IF EXISTS products_category_name_idx;
ALTER TABLE products DROP CONSTRAINT IF EXISTS products_category_fk;
DROP INDEX IF EXISTS variant_types_product_name_idx;
ALTER TABLE variant_types DROP CONSTRAINT IF EXISTS variant_types_name_not_empty;
DROP INDEX IF EXISTS variant_values_type_value_idx;
ALTER TABLE variant_values DROP CONSTRAINT IF EXISTS variant_values_value_not_empty;
DROP INDEX IF EXISTS skus_sku_code_lower_idx;
ALTER TABLE skus DROP CONSTRAINT IF EXISTS skus_price_positive;
ALTER TABLE skus DROP CONSTRAINT IF EXISTS skus_stock_non_negative;
ALTER TABLE sku_variant_values DROP CONSTRAINT IF EXISTS sku_variant_unique;
ALTER TABLE sku_images DROP CONSTRAINT IF EXISTS image_url_not_empty;
-- +goose StatementEnd
