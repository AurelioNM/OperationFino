run-customer-service:
	clear && cd ./apis/customer-service/ \
		&& go build -o bin/ ./cmd/customer-service \
		&& docker-compose -f ../../docker/docker-compose.yml up --build of-customer-service
	
db-size:
	docker exec of-customer-postgres psql -v ON_ERROR_STOP=1 --username "customer" --dbname "customer-service" -c \
		"SELECT COUNT(*) FROM customers;" 


db-clean:
	docker exec of-customer-postgres psql -v ON_ERROR_STOP=1 --username "customer" --dbname "customer-service" -c \
		"DELETE FROM customers;" 

