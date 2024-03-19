CREATE TABLE clients (
	client_id varchar(26) NOT NULL,
	name varchar(30) NOT NULL,
	surname varchar(30) NOT NULL,
	email varchar(200) UNIQUE NOT NULL,
	birthdate date NOT NULL,

	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,

	CONSTRAINT clients_pk PRIMARY KEY (client_id)
);

CREATE TABLE products (
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

CREATE TABLE orders (
	order_id varchar(26) NOT NULL,
	client_id varchar(30) NOT NULL,
	status order_status DEFAULT 'open' NOT NULL,

	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,

	CONSTRAINT order_pk PRIMARY KEY (order_id),
	CONSTRAINT order_client_fk FOREIGN KEY (client_id) REFERENCES clients(client_id)
);

CREATE TABLE order_items (
	order_item_id varchar(26) NOT NULL,
	order_id varchar(26) NOT NULL,
	product_id varchar(30) NOT NULL,
	quantity smallint NOT NULL,

	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	
	CONSTRAINT order_item_pk PRIMARY KEY (order_item_id),
	CONSTRAINT order_item_order_fk FOREIGN KEY (order_id) REFERENCES clients(order_id),
	CONSTRAINT order_item_product_fk FOREIGN KEY (product_id) REFERENCES clients(product_id)
);

