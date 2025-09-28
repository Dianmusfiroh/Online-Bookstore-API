# Stage 1: Build
FROM --platform=linux/amd64 golang:1.22 AS builder

WORKDIR /app

# Copy dependencies dulu
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary (pastikan statically linked)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./main.go

# Stage 2: Run
FROM --platform=linux/amd64 debian:bookworm-slim

WORKDIR /app

# Copy binary dari builder
COPY --from=builder /app/main .

# Pastikan bisa dieksekusi
RUN chmod +x /app/main

EXPOSE 3000

CMD ["/app/main"]
