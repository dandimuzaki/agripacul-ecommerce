-- +goose Up
-- +goose StatementBegin

-- Assign enum user role to users table
CREATE TYPE user_roles AS ENUM('superadmin','admin','staff','customer');
ALTER TABLE IF EXISTS users
ALTER COLUMN role TYPE user_roles
USING role::user_roles;
ALTER TABLE IF EXISTS users
ALTER COLUMN role SET DEFAULT 'customer';

-- Assign enum order status to orders table
CREATE TYPE order_status AS ENUM('created','processing','completed','cancelled');
ALTER TABLE IF EXISTS orders
ALTER COLUMN status TYPE order_status
USING status::order_status;
ALTER TABLE IF EXISTS orders
ALTER COLUMN status SET DEFAULT 'created';

-- Assign enum shipping status to order_shippings table
CREATE TYPE shipping_status AS ENUM('pending','shipped','delivered');
ALTER TABLE IF EXISTS order_shippings
ALTER COLUMN status TYPE order_status
USING status::order_status;
ALTER TABLE IF EXISTS order_shippings
ALTER COLUMN status SET DEFAULT 'pending';

-- Assign enum payment status to order_payments table
CREATE TYPE payment_status AS ENUM('pending','paid','expired','failed');
ALTER TABLE IF EXISTS order_payments
ALTER COLUMN status TYPE payment_status
USING payment_status::payment_status;
ALTER TABLE IF EXISTS order_payments
ALTER COLUMN status SET DEFAULT 'pending';

-- Assign enum promo type to promotions table
CREATE TYPE promo_types AS ENUM('direct discount','voucher code');
ALTER TABLE IF EXISTS promotions
ALTER COLUMN type TYPE promo_types
USING type::promo_types;
ALTER TABLE IF EXISTS promotions
ALTER COLUMN type SET DEFAULT 'direct discount';

-- Assign enum discount type to promotions table
CREATE TYPE discount_types AS ENUM('amount','percentage');
ALTER TABLE IF EXISTS promotions
ALTER COLUMN discount_type TYPE discount_types
USING discount_type::discount_types;
ALTER TABLE IF EXISTS promotions
ALTER COLUMN discount_type SET DEFAULT 'amount';

-- Assign banner type to banners table
CREATE TYPE banner_types AS ENUM('main','secondary');
ALTER TABLE IF EXISTS banners
ALTER COLUMN type TYPE banner_types
USING type::banner_types;
ALTER TABLE IF EXISTS banners
ALTER COLUMN type SET DEFAULT 'secondary';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- Rollback enum user role
ALTER TABLE IF EXISTS users
ALTER COLUMN role TYPE varchar(50);
ALTER TABLE IF EXISTS users
ALTER COLUMN role DROP DEFAULT;
DROP TYPE IF EXISTS user_roles;

-- Rollback enum order status
ALTER TABLE IF EXISTS orders
ALTER COLUMN status TYPE varchar(50);
ALTER TABLE IF EXISTS orders
ALTER COLUMN status DROP DEFAULT;
DROP TYPE IF EXISTS order_status;

-- Rollback enum shipping status
ALTER TABLE IF EXISTS order_shippings
ALTER COLUMN status TYPE varchar(50);
ALTER TABLE IF EXISTS order_shippings
ALTER COLUMN status DROP DEFAULT;
DROP TYPE IF EXISTS shipping_status;

-- Rollback enum payment status
ALTER TABLE IF EXISTS order_payments
ALTER COLUMN status TYPE varchar(50);
ALTER TABLE IF EXISTS order_payments
ALTER COLUMN status DROP DEFAULT;
DROP TYPE IF EXISTS payment_status;

-- Rollback enum promo type
ALTER TABLE IF EXISTS promotions
ALTER COLUMN type TYPE varchar(50);
ALTER TABLE IF EXISTS promotions
ALTER COLUMN type DROP DEFAULT;
DROP TYPE IF EXISTS promo_types;

-- Rollback enum discount type
ALTER TABLE IF EXISTS promotions
ALTER COLUMN discount_type TYPE varchar(50);
ALTER TABLE IF EXISTS promotions
ALTER COLUMN discount_type DROP DEFAULT;
DROP TYPE IF EXISTS discount_types;

-- Rollback enum banner type
ALTER TABLE IF EXISTS banners
ALTER COLUMN type TYPE varchar(50);
ALTER TABLE IF EXISTS banners
ALTER COLUMN type DROP DEFAULT;
DROP TYPE IF EXISTS banner_types;
-- +goose StatementEnd
