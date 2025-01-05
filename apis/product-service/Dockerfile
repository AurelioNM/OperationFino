FROM golang:1.21.2

WORKDIR /app

COPY ./bin/product-service /app

COPY ./.env /app

COPY ./swagger.yml /app

EXPOSE 8002

CMD [ "/app/product-service" ]
