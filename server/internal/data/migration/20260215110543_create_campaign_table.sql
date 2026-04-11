-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS campaigns (
  id bigserial primary key,
  name varchar(255) not null,
  description text,
  type varchar(100),
  start_date timestamptz,
  end_date timestamptz,
  is_active boolean default true,
  created_at timestamptz default now(),
  updated_at timestamptz default now(),
  deleted_at timestamptz
);

CREATE TABLE IF NOT EXISTS campaign_products (
  id serial primary key,
  campaign_id int references campaigns(id),
  product_id int references products(id),
  created_at timestamptz default now(),
  updated_at timestamptz default now(),
  deleted_at timestamptz
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS campaign_products;
DROP TABLE IF EXISTS campaigns;
-- +goose StatementEnd
