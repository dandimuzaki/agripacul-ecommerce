-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS companies (
  id serial primary key,
  name varchar(255) not null,
  description text,
  logo_url text,
  phone_number varchar(50),
  instagram_url text,
  twitter_url text,
  whatsapp_url text,
  contact_email varchar(255),
  created_at timestamptz default now(),
  updated_at timestamptz default now(),
  deleted_at timestamptz
);

CREATE TABLE IF NOT EXISTS company_addresses (
  id serial primary key,
  company_id int not null references companies(id),
  label varchar(50),
  province_id int not null references provinces(id),
  regency_id int not null references regencies(id),
  district_id int not null references districts(id),
  subdistrict_id int not null references subdistricts(id),
  postal_code varchar(10),
  detail_address text,
  is_shipping_origin boolean default false,
  created_at timestamptz default now(),
  updated_at timestamptz default now(),
  deleted_at timestamptz
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS company_addresses;
DROP TABLE IF EXISTS companies;
-- +goose StatementEnd
