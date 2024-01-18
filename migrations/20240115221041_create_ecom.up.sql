CREATE TABLE IF NOT EXISTS cart_items (
  id SERIAL PRIMARY KEY NOT NULL,
  cart_id bigint NOT NULL,
  product_id bigint NOT NULL,
  quantity bigint
);

CREATE TABLE IF NOT EXISTS carts (
  id SERIAL PRIMARY KEY NOT NULL,
  customer_id bigint NOT NULL,
  total_price bigint,
  created_at timestamp,
  updated_at timestamp
);

CREATE TABLE IF NOT EXISTS categories (
  id SERIAL PRIMARY KEY NOT NULL,
  name VARCHAR(50) NOT NULL,
  description text
);

CREATE TABLE IF NOT EXISTS customers (
  id SERIAL PRIMARY KEY NOT NULL,
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
  id SERIAL PRIMARY KEY NOT NULL,
  order_id bigint NOT NULL,
  product_id bigint NOT NULL,
  quantity bigint,
  price bigint
);

CREATE TABLE IF NOT EXISTS orders (
  id SERIAL PRIMARY KEY NOT NULL,
  customer_id bigint NOT NULL,
  status VARCHAR,
  created_at timestamp,
  updated_at timestamp
);

CREATE TABLE IF NOT EXISTS products (
  id SERIAL PRIMARY KEY NOT NULL,
  category_id bigint NOT NULL,
  name VARCHAR,
  description text,
  price bigint
);

CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY NOT NULL,
  username VARCHAR(100) UNIQUE,
  password VARCHAR(255),
  role VARCHAR(50),
  created_at timestamp,
  updated_at timestamp,
  last_login_attempt timestamp
);

ALTER TABLE cart_items ADD FOREIGN KEY (cart_id) REFERENCES carts (id);
ALTER TABLE cart_items ADD FOREIGN KEY (product_id) REFERENCES products (id);
ALTER TABLE orders ADD FOREIGN KEY (customer_id) REFERENCES customers (id);
ALTER TABLE carts ADD FOREIGN KEY (customer_id) REFERENCES customers (id);
ALTER TABLE order_items ADD FOREIGN KEY (order_id) REFERENCES orders (id);
ALTER TABLE order_items ADD FOREIGN KEY (product_id) REFERENCES products (id);
ALTER TABLE products ADD FOREIGN KEY (category_id) REFERENCES categories (id);
ALTER TABLE customers ADD FOREIGN KEY (user_id) REFERENCES users (id);
