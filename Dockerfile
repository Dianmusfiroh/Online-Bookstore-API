# Step 1: Build Go binary
FROM golang:1.22 AS builder

# Set working directory
WORKDIR /app

# Copy go mod & sum
COPY go.mod go.sum ./
RUN go mod download

# Copy all source code
COPY . .

# Build binary
RUN go build -o out .

# Step 2: Run stage (image lebih kecil)
FROM debian:bullseye-slim

WORKDIR /app

# Copy binary dari builder stage
COPY --from=builder /app/out .

# Set permission supaya bisa dieksekusi
RUN chmod +x ./out

# Railway inject variable PORT, jadi expose default 3000
EXPOSE 3000

# Start app
CMD ["./out"]
