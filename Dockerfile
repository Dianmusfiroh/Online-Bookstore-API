# Stage 1: Build
FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Paksa build untuk Linux AMD64, binary statis
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./main.go

# Stage 2: Run
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/main .

# Pastikan executable
RUN chmod +x /app/main

EXPOSE 3000

CMD ["/app/main"]
