-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS order_payments (
  id serial primary key,
  order_id int not null references orders(id),
  payment_method_id int not null references payment_methods(id),
  amount numeric(12,2) not null,
  status varchar(50) not null,
  transaction_ref varchar(100),
  paid_at timestamptz,
  failed_at timestamptz,
  expired_at timestamptz,
  created_at timestamptz default now(),
  updated_at timestamptz default now(),
  deleted_at timestamptz
);

CREATE TABLE IF NOT EXISTS order_shippings (
  id serial primary key,
  order_id int not null references orders(id),
  recipient_name varchar(100) not null,
  label varchar(50),
  phone_number varchar(30) not null,
  province varchar(100) not null,
  regency varchar(100) not null,
  district varchar(100) not null,
  subdistrict varchar(100) not null,
  postal_code varchar(10),
  detail_address text,
  courier_name varchar(100) not null,
  courier_code varchar(20) not null,
  courier_service varchar(50),
  cost numeric(12,2) not null,
  etd int,
  status varchar(50) not null,
  tracking_number varchar(100),
  shipped_at timestamptz,
  delivered_at timestamptz,
  created_at timestamptz default now(),
  updated_at timestamptz default now(),
  deleted_at timestamptz
);

ALTER TABLE order_items 
ADD COLUMN IF NOT EXISTS created_at timestamptz default now(),
ADD COLUMN IF NOT EXISTS updated_at timestamptz default now(),
ADD COLUMN IF NOT EXISTS deleted_at timestamptz;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS order_payments;
DROP TABLE IF EXISTS order_shippings;
-- +goose StatementEnd
