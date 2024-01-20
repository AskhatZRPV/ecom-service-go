CREATE TABLE IF NOT EXISTS cart_item (
  id SERIAL PRIMARY KEY NOT NULL,
  session_id bigint NOT NULL,
  product_id bigint NOT NULL,
  quantity bigint
);

CREATE TABLE IF NOT EXISTS category (
  id SERIAL PRIMARY KEY NOT NULL,
  title VARCHAR(50) NOT NULL,
  description text,
);

CREATE TABLE IF NOT EXISTS inventory (
  id SERIAL PRIMARY KEY NOT NULL,
  quantity int,
  description text,
);

CREATE TABLE IF NOT EXISTS user_address (
  id SERIAL PRIMARY KEY NOT NULL,
  user_id bigint NOT NULL,
  first_name VARCHAR(50) NOT NULL,
  last_name VARCHAR(50),
  address_line1 VARCHAR(255),
  city VARCHAR,
  postal_code VARCHAR(10),
  country VARCHAR(50),
  phone_number VARCHAR(10),
  created_at timestamp,
  updated_at timestamp
);

CREATE TABLE IF NOT EXISTS order_item (
  id SERIAL PRIMARY KEY NOT NULL,
  order_id bigint NOT NULL,
  product_id bigint NOT NULL,
  quantity bigint,
  price bigint
);

CREATE TABLE IF NOT EXISTS order (
  id SERIAL PRIMARY KEY NOT NULL,
  payment_id bigint,
  user_id bigint NOT NULL,
  status VARCHAR,
  created_at timestamp
);

CREATE TABLE IF NOT EXISTS payment (
  id SERIAL PRIMARY KEY NOT NULL,
  order_id bigint,
  amount bigint,
  created_at timestamp
);

CREATE TABLE IF NOT EXISTS product (
  id SERIAL PRIMARY KEY NOT NULL,
  category_id bigint NOT NULL,
  inventory_id bigint NOT NULL,
  SKU varchar,
  title VARCHAR,
  description text,
  price bigint,
);

CREATE TABLE IF NOT EXISTS shopping_session (
  id SERIAL PRIMARY KEY NOT NULL,
  user_id int,
  total int
);

CREATE TABLE IF NOT EXISTS user (
  id SERIAL PRIMARY KEY NOT NULL,
  username VARCHAR(100) UNIQUE,
  password VARCHAR(255),
  role VARCHAR(50),
  created_at timestamp,
  updated_at timestamp,
  last_login_attempt timestamp
);

CREATE TABLE IF NOT EXISTS user_account (
  id SERIAL PRIMARY KEY NOT NULL,
  user_id int,
  balance int,
);

CREATE TABLE IF NOT EXISTS session (
  id uuid,
  user_id int,
  expires_at timestamp
);

ALTER TABLE order ADD FOREIGN KEY (user_id) REFERENCES user (id);
ALTER TABLE order ADD FOREIGN KEY (payment_id) REFERENCES payment (id);
ALTER TABLE product ADD FOREIGN KEY (inventory_id) REFERENCES inventory (id);
ALTER TABLE product ADD FOREIGN KEY (category_id) REFERENCES category (id);
ALTER TABLE order_item ADD FOREIGN KEY (order_id) REFERENCES order (id);
ALTER TABLE order_item ADD FOREIGN KEY (product_id) REFERENCES product (id);
ALTER TABLE cart_item ADD FOREIGN KEY (session_id) REFERENCES shopping_session (id);
ALTER TABLE cart_item ADD FOREIGN KEY (product_id) REFERENCES product (id);
ALTER TABLE shopping_session ADD FOREIGN KEY (user_id) REFERENCES user (id);
ALTER TABLE user_account ADD FOREIGN KEY (user_id) REFERENCES user (id);
ALTER TABLE user_address ADD FOREIGN KEY (user_id) REFERENCES user (id);

