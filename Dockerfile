# Stage 1: Build
FROM golang:1.22 AS builder

# Set working directory
WORKDIR /app

# Copy go.mod dan go.sum dulu biar dependency bisa cache
COPY go.mod go.sum ./
RUN go mod download

# Copy semua source code
COPY . .

# Build binary dari main.go
RUN go build -o main .

# Stage 2: Run
FROM debian:bookworm-slim

WORKDIR /app

# Copy binary hasil build
COPY --from=builder /app/main .

# Expose port default (ubah kalau aplikasimu pakai port lain)
EXPOSE 8080

# Jalankan binary
CMD ["./main"]
