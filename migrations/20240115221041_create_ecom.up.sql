-- BEGIN;

CREATE TABLE IF NOT EXISTS cart_items (
  id bigint PRIMARY KEY NOT NULL,
  cart_id bigint NOT NULL,
  product_id bigint NOT NULL,
  quantity bigint
);

CREATE TABLE IF NOT EXISTS carts (
  id bigint PRIMARY KEY NOT NULL,
  customer_id bigint NOT NULL,
  total_price bigint,
  created_at timestamp,
  updated_at timestamp
);

CREATE TABLE IF NOT EXISTS categories (
  id bigint PRIMARY KEY NOT NULL,
  name VARCHAR(50) NOT NULL,
  description text
);

CREATE TABLE IF NOT EXISTS customers (
  id bigint PRIMARY KEY NOT NULL,
  user_id bigint NOT NULL,
  first_name VARCHAR(50) NOT NULL,
  last_name VARCHAR(50),
  address VARCHAR(255),
  city VARCHAR,
  zip_code VARCHAR(10),
  country VARCHAR(50),
  created_at timestamp,
  updated_at timestamp
);

CREATE TABLE IF NOT EXISTS order_items (
  id bigint PRIMARY KEY NOT NULL,
  order_id bigint NOT NULL,
  product_id bigint NOT NULL,
  quantity bigint,
  price bigint
);

CREATE TABLE IF NOT EXISTS orders (
  id bigint PRIMARY KEY NOT NULL,
  customer_id bigint NOT NULL,
  status VARCHAR,
  created_at timestamp,
  updated_at timestamp
);

CREATE TABLE IF NOT EXISTS products (
  id bigint PRIMARY KEY NOT NULL,
  category_id bigint NOT NULL,
  name VARCHAR,
  description text,
  price bigint
);

CREATE TABLE IF NOT EXISTS users (
  id bigint PRIMARY KEY NOT NULL,
  username VARCHAR(100),
  password VARCHAR(255),
  role VARCHAR(50),
  created_at timestamp,
  updated_at timestamp
);


ALTER TABLE IF NOT EXISTS cart_items ADD FOREIGN KEY (cart_id) REFERENCES carts (id);

ALTER TABLE IF NOT EXISTS cart_items ADD FOREIGN KEY (product_id) REFERENCES products (id);

ALTER TABLE IF NOT EXISTS orders ADD FOREIGN KEY (customer_id) REFERENCES customers (id);

ALTER TABLE IF NOT EXISTS carts ADD FOREIGN KEY (customer_id) REFERENCES customers (id);

ALTER TABLE IF NOT EXISTS order_items ADD FOREIGN KEY (order_id) REFERENCES orders (id);

ALTER TABLE IF NOT EXISTS order_items ADD FOREIGN KEY (product_id) REFERENCES products (id);

ALTER TABLE IF NOT EXISTS products ADD FOREIGN KEY (category_id) REFERENCES categories (id);

ALTER TABLE IF NOT EXISTS customers ADD FOREIGN KEY (user_id) REFERENCES users (id);

-- COMMIT;