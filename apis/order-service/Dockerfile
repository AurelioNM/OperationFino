FROM golang:1.21.2

WORKDIR /app

COPY ./bin/order-service /app

COPY ./.env /app

COPY ./swagger.yml /app

EXPOSE 8002

CMD [ "/app/order-service" ]
