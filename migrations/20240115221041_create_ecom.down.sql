DROP TABLE IF EXISTS cart_items;
DROP TABLE IF EXISTS carts;
DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS customers;
DROP TABLE IF EXISTS order_items;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS users;

ALTER TABLE IF EXISTS cart_items DROP CONSTRAINT cart_id;
ALTER TABLE IF EXISTS cart_items DROP CONSTRAINT product_id;
ALTER TABLE IF EXISTS orders DROP CONSTRAINT customer_id;
ALTER TABLE IF EXISTS carts DROP CONSTRAINT customer_id;
ALTER TABLE IF EXISTS order_items DROP CONSTRAINT order_id;
ALTER TABLE IF EXISTS order_items DROP CONSTRAINT product_id;
ALTER TABLE IF EXISTS products DROP CONSTRAINT category_id;
ALTER TABLE IF EXISTS customers DROP CONSTRAINT user_id;
