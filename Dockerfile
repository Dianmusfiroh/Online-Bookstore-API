# Build stage
FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build untuk linux amd64
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./main.go

# Run stage
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/main .

# Pastikan binary bisa dieksekusi
RUN chmod +x /app/main

EXPOSE 3000

CMD ["/app/main"]
