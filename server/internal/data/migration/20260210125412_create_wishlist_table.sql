-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS wishlists (
  id bigserial primary key,
  customer_id bigint references customers(id)
);

CREATE TABLE IF NOT EXISTS wishlist_products (
  id bigserial primary key,
  wishlist_id bigint references wishlists(id),
  product_id bigint references products(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS wishlist_products;
DROP TABLE IF EXISTS wishlists;
-- +goose StatementEnd
