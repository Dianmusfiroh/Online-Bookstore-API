FROM golang:1.22.1-alpine AS builder

# Paksa compile untuk Linux AMD64
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o main .

FROM alpine:latest

WORKDIR /root

COPY --from=builder /app/main .

RUN chmod +x main

EXPOSE 3000
CMD ["./main"]
