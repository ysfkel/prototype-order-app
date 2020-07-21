FROM golang:1.13.4-stretch as builder
ENV GO111MODULE=on
WORKDIR /app
COPY . . 
RUN CGO_ENABLED=0 go build .

FROM alpine:3.10
WORKDIR /app
COPY --from=builder /app/order-app .

COPY --from=builder /app/test_data/postgres-orders.csv ./test_data/postgres-orders.csv
COPY --from=builder /app/test_data/postgres-order_items.csv ./test_data/postgres-order_items.csv
COPY --from=builder /app/test_data/postgres-deliveries.csv ./test_data/postgres-deliveries.csv
COPY --from=builder /app/test_data/mongo-customers.csv ./test_data/mongo-customers.csv
COPY --from=builder /app/test_data/mongo-customer_companies.csv ./test_data/mongo-customer_companies.csv

ENTRYPOINT ["./order-app"]
