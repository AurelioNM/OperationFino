CREATE TABLE IF NOT EXISTS customers (
	customer_id varchar(26) NOT NULL,
	name varchar(30) NOT NULL,
	surname varchar(30) NOT NULL,
	email varchar(200) UNIQUE NOT NULL,
	birthdate varchar(10) NOT NULL,

	created_at timestamp NOT NULL,
	updated_at timestamp,

	CONSTRAINT customer_pk PRIMARY KEY (customer_id)
);

CREATE INDEX IF NOT EXISTS customers_email_idx ON customers USING btree (email);
