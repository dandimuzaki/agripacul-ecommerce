-- +goose Up
-- +goose StatementBegin
-- drop column is_active
ALTER TABLE IF EXISTS skus
DROP COLUMN IF EXISTS is_active;

-- create sku status enum
CREATE TYPE sku_status AS ENUM('active','inactive','archived');

-- add column status
ALTER TABLE IF EXISTS skus
ADD COLUMN IF NOT EXISTS status sku_status DEFAULT 'inactive';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE IF EXISTS skus
DROP COLUMN IF EXISTS status;

DROP TYPE IF EXISTS sku_status;

ALTER TABLE IF EXISTS skus
ADD COLUMN IF NOT EXISTS is_active BOOLEAN DEFAULT true;
-- +goose StatementEnd
