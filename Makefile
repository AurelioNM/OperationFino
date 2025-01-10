run-customer:
	clear && cd ./apis/customer-service/ \
		&& go build -o bin/ ./cmd/customer-service \
		&& docker-compose -f ../../docker/docker-compose.yml up --build of-customer-service

run-product:
	clear && cd ./apis/product-service/ \
		&& go build -o bin/ ./cmd/product-service \
		&& docker-compose -f ../../docker/docker-compose.yml up --build of-product-service

run-order:
	clear && cd ./apis/order-service/ \
		&& go build -o bin/ ./cmd/order-service \
		&& docker-compose -f ../../docker/docker-compose.yml up --build of-order-service

test-customer:
	cd ./apis/customer-service \
		&& go test -v ./... && cd -

db-size:
	docker exec of-customer-postgres psql -v ON_ERROR_STOP=1 --username "customer" --dbname "customer-service" -c \
		"SELECT COUNT(*) AS customers_size FROM customers;"  \
	&& docker exec of-product-postgres psql -v ON_ERROR_STOP=1 --username "product" --dbname "product-service" -c \
		"SELECT COUNT(*) AS products_size FROM products;"  \

db-get:
	docker exec of-customer-postgres psql -v ON_ERROR_STOP=1 --username "customer" --dbname "customer-service" -c \
		"SELECT * FROM customers LIMIT 10;" \
	&& docker exec of-product-postgres psql -v ON_ERROR_STOP=1 --username "product" --dbname "product-service" -c \
		"SELECT * FROM products LIMIT 10;"

db-test-data:
	docker exec of-customer-postgres psql -v ON_ERROR_STOP=1 --username "customer" --dbname "customer-service" -c \
		"SELECT email as customer_email FROM customers LIMIT 20;" \
	&& docker exec of-product-postgres psql -v ON_ERROR_STOP=1 --username "product" --dbname "product-service" -c \
		"SELECT name as product_name FROM products LIMIT 20;"
db-clean:
	docker exec of-customer-postgres psql -v ON_ERROR_STOP=1 --username "customer" --dbname "customer-service" -c \
		"DELETE FROM customers;" \
	&& docker exec of-product-postgres psql -v ON_ERROR_STOP=1 --username "product" --dbname "product-service" -c \
		"DELETE FROM products;"

