run-customer-service:
	clear && cd ./apis/customer-service/ \
		&& go build -o bin/ ./cmd/customer-service \
		&& docker-compose -f ../../docker/docker-compose.yml up --build of-customer-service
