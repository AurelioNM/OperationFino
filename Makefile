run-customer-service:
	clear && cd ./apis/customer-service \
		&& go build -o ./bin ./cmd/customer-service \
		&& ./bin/customer-service
