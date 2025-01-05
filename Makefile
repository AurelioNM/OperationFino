run-customer-service:
	clear && cd ./apis/customer-service/ \
		&& go build -o bin/ ./cmd/customer-service \
		&& docker-compose -f ../../docker/docker-compose.yml up --build of-customer-service

run-product-service:
	clear && cd ./apis/product-service/ \
		&& go build -o bin/ ./cmd/product-service \
		&& docker-compose -f ../../docker/docker-compose.yml up --build of-product-service

test-customer-service:
	cd ./apis/customer-service \
		&& go test -v ./... && cd -

db-size:
	docker exec of-customer-postgres psql -v ON_ERROR_STOP=1 --username "customer" --dbname "customer-service" -c \
		"SELECT COUNT(*) AS customers_size FROM customers;" 

db-get:
	docker exec of-customer-postgres psql -v ON_ERROR_STOP=1 --username "customer" --dbname "customer-service" -c \
		"SELECT * FROM customers;" 

db-clean:
	docker exec of-customer-postgres psql -v ON_ERROR_STOP=1 --username "customer" --dbname "customer-service" -c \
		"DELETE FROM customers;"

