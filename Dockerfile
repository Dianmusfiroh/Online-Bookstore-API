# ------------------------------
# Stage 1: Build binary
# ------------------------------
FROM golang:1.22.1-alpine AS builder

# Paksa build untuk Linux amd64
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /go/src/book-online-api

COPY . .

RUN go mod download

RUN go build -o main .

# ------------------------------
# Stage 2: Final image
# ------------------------------
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /go/src/book-online-api/main .

RUN chmod +x ./main

EXPOSE 8080

CMD ["./main"]
