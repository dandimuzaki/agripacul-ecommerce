-- +goose Up
-- +goose StatementBegin
CREATE TYPE inventory_log_types AS ENUM('in', 'out', 'adjustment', 'initial');

CREATE TABLE IF NOT EXISTS inventory_logs (
  id bigserial primary key,
  sku_id int references skus(id) not null,
  type inventory_log_types not null,
  quantity_change int not null,
  current_stock_after int not null,
  reference_id int,
  reference_type varchar(50),
  notes text,
  created_at timestamptz default now(),
  updated_at timestamptz default now(),
  deleted_at timestamptz
);

ALTER TABLE IF EXISTS inventory_logs
ALTER COLUMN type TYPE inventory_log_types USING type::inventory_log_types;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS inventory_logs;
DROP TYPE IF EXISTS inventory_log_types;
-- +goose StatementEnd
