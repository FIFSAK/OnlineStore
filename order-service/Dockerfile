FROM golang:1.21.0 as builder

WORKDIR /app

COPY ../go.mod /app/
COPY ../go.sum /app/
COPY . /app/

RUN go mod download

EXPOSE 10003

CMD ["go", "run", "/app/order-service/main.go", "."]
