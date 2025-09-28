FROM golang:1.22.1 as builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o main .

FROM debian:bookworm-slim
WORKDIR /root
COPY --from=builder /app/main .

# pastikan binary bisa dieksekusi
RUN chmod +x ./main

CMD ["./main"]
