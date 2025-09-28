# Step 1: Build Go binary
FROM golang:1.22 AS builder

WORKDIR /app

# Copy go mod & sum
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary statically untuk linux/amd64
RUN GOOS=linux GOARCH=amd64 go build -o out .

# Step 2: Run stage
FROM debian:bullseye-slim

WORKDIR /app
COPY --from=builder /app/out .

RUN chmod +x ./out

EXPOSE 3000
CMD ["./out"]
