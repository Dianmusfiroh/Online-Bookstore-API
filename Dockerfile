# ------------------------------
# Stage 1: Build binary
# ------------------------------
FROM golang:1.22.1-alpine AS builder

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /go/src/book-online-api

COPY . .

RUN go mod download

RUN go build -o main .

# Pastikan binary bisa dieksekusi
RUN chmod +x main

# ------------------------------
# Stage 2: Final image
# ------------------------------
FROM alpine:latest

WORKDIR /root/

# Copy binary dari builder
COPY --from=builder /go/src/book-online-api/main .

# Pastikan tetap eksekutabel setelah copy
RUN chmod +x ./main

EXPOSE 8080

CMD ["./main"]
