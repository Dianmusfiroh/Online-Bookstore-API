# Stage 1: Build
FROM golang:1.22 AS builder

# Set working directory di container
WORKDIR /app

# Copy go.mod dan go.sum dulu (biar dependency bisa cache)
COPY go.mod go.sum ./
RUN go mod download

# Copy semua source code
COPY . .

# Build binary dari main.go
RUN go build -o main ./main.go

# Stage 2: Run
FROM debian:bookworm-slim

WORKDIR /app

# Copy binary hasil build dari stage 1
COPY --from=builder /app/main .

# Pastikan binary bisa dieksekusi
RUN chmod +x /app/main

# Expose port (ubah kalau pakai selain 3000)
EXPOSE 3000

# Jalankan binary
CMD ["/app/main"]
