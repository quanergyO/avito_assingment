FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY . .


RUN go mod download
RUN go build -o avito_cmd ./cmd/main.go

CMD ["./avito_cmd"]