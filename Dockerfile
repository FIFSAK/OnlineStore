FROM golang:1.21.0 as builder

WORKDIR /app

COPY . .


RUN go mod download

EXPOSE 8080

CMD ["go", "run", "/app/api-gateway/main.go", "."]

