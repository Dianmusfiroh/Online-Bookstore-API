# -- Stage 1: Build the application --
FROM golang:1.20-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum to cache dependencies
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy all source code
COPY . .

# Build the Go binary for the final stage
# Set the target OS and Architecture to prevent 'exec format error'
ENV GOOS=linux
ENV GOARCH=amd64
RUN CGO_ENABLED=0 go build -o main -v

# -- Stage 2: Create a minimal production image --
FROM alpine:latest

# Set working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Add execute permissions to the binary
RUN chmod +x ./main

# Set the command to run the application
CMD ["./main"]