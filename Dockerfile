# Gunakan base image golang
FROM golang:1.21-alpine AS builder

# Set environment variables
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Set working directory di dalam container
WORKDIR /go/src/book-online-api

# Copy semua source code ke dalam container
COPY . .

# Download dependencies
RUN go mod download

# Build binary
RUN go build -o main .

# ------------------------------
# Stage 2: Buat image final lebih ringan
# ------------------------------
FROM alpine:latest

WORKDIR /root/

# Copy binary dari stage builder
COPY --from=builder /go/src/book-online-api/main .

# Copy folder lain yang diperlukan (opsional, kalau butuh file selain binary)
# COPY --from=builder /go/src/book-online-api/config ./config
# COPY --from=builder /go/src/book-online-api/migrations ./migrations

# Port aplikasi (ubah sesuai kebutuhan)
EXPOSE 8080

# Jalankan binary
CMD ["./main"]
