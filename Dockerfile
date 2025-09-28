# -- Tahap 1: Kompilasi aplikasi (builder) --
# Gunakan image Golang yang spesifik untuk versi go.mod Anda
FROM golang:1.22.1-alpine AS builder

# Create and change to the app directory.
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Copy local code to the container image.
COPY . ./

# Install project dependencies
RUN go mod download

# Build the app
RUN go build -o app

# Run the service on container startup.
ENTRYPOINT ["./app"]