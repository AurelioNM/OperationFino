FROM golang:1.21.2

WORKDIR /app

COPY ./bin/customer-service /app

COPY ./.env /app

EXPOSE 8001

CMD [ "/app/customer-service" ]
