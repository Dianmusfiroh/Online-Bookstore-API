# Build stage
FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build untuk linux amd64
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./main.go

# Run stage
FROM scratch

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 3000

ENTRYPOINT ["/app/main"]
