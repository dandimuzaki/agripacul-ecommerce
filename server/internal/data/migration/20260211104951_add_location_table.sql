-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS provinces (
  id serial primary key,
  code varchar(50) not null,
  name varchar(100) not null,
  raja_ongkir_id int
);

CREATE TABLE IF NOT EXISTS regencies (
  id serial primary key,
  province_id int not null references provinces(id),
  code varchar(50) not null,
  name varchar(100) not null,
  type varchar(20) not null,
  raja_ongkir_id int
);

CREATE TABLE IF NOT EXISTS districts (
  id serial primary key,
  regency_id int not null references regencies(id),
  code varchar(50) not null,
  name varchar(100) not null,
  raja_ongkir_id int
);

CREATE TABLE IF NOT EXISTS subdistricts (
  id serial primary key,
  district_id int not null references districts(id),
  code varchar(50) not null,
  name varchar(100) not null
);

CREATE TABLE IF NOT EXISTS addresses (
  id serial primary key,
  customer_id int not null references customers(id),
  recipient_name varchar(255) not null,
  label varchar(50) not null,
  province_id int not null references provinces(id),
  regency_id int not null references regencies(id),
  district_id int not null references districts(id),
  subdistrict_id int not null references subdistricts(id),
  postal_code varchar(10),
  detail_address text,
  is_default boolean default false,
);

-- one default address per customer
CREATE UNIQUE INDEX uniq_default_address_per_customer
ON addresses(customer_id)
WHERE is_default = true;

-- prevent duplicate provinces
CREATE UNIQUE INDEX IF NOT EXISTS uq_provinces_code
ON provinces (code);

-- prevent duplicate regencies
CREATE UNIQUE INDEX IF NOT EXISTS uq_regencies_province_code
ON regencies (province_id, code);

-- prevent duplicate districts
CREATE UNIQUE INDEX IF NOT EXISTS uq_districts_regency_code
ON districts (regency_id, code);

-- prevent duplicate subdistricts
CREATE UNIQUE INDEX IF NOT EXISTS uq_subdistricts_district_code
ON districts (regency_id, code);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS uq_subdistricts_district_code;
DROP INDEX IF EXISTS uq_districts_regency_code;
DROP INDEX IF EXISTS uq_regencies_province_code;
DROP INDEX IF EXISTS uq_provinces_code;
DROP TABLE IF EXISTS addresses;
DROP TABLE IF EXISTS subdistricts;
DROP TABLE IF EXISTS districts;
DROP TABLE IF EXISTS regencies;
DROP TABLE IF EXISTS provinces;
DROP INDEX IF EXISTS uniq_default_address_per_customer;
-- +goose StatementEnd
