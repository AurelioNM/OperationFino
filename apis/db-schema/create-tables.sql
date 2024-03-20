CREATE TABLE IF NOT EXISTS customer (
	customer_id varchar(26) NOT NULL,
	name varchar(30) NOT NULL,
	surname varchar(30) NOT NULL,
	email varchar(200) UNIQUE NOT NULL,
	birthdate date NOT NULL,

	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,

	CONSTRAINT customer_pk PRIMARY KEY (customer_id)
);

CREATE TABLE IF NOT EXISTS products (
	product_id varchar(26) NOT NULL,
	name varchar(30) NOT NULL,
	description varchar(255) NOT NULL,
	quantity smallint NOT NULL,

	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,

	CONSTRAINT product_pk PRIMARY KEY (product_id)
);

CREATE TYPE order_status AS ENUM(
	'STARTED',
	'PROGRESS',
	'DONE',
	'CANCELED'
);

CREATE TABLE IF NOT EXISTS orders (
	order_id varchar(26) NOT NULL,
	customer_id varchar(30) NOT NULL,
	status order_status DEFAULT 'open' NOT NULL,

	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,

	CONSTRAINT order_pk PRIMARY KEY (order_id)
);

CREATE TABLE IF NOT EXISTS order_items (
	order_item_id varchar(26) NOT NULL,
	order_id varchar(26) NOT NULL,
	product_id varchar(30) NOT NULL,
	quantity smallint NOT NULL,

	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	
	CONSTRAINT order_item_pk PRIMARY KEY (order_item_id),
	CONSTRAINT order_item_order_fk FOREIGN KEY (order_id) REFERENCES orders(order_id)
);

