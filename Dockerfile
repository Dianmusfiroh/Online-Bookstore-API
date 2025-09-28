FROM golang:1.22.1 as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o main -v

# -- Stage 2: Final --
FROM alpine:latest
COPY --from=builder /app/main .
# Add execute permissions to the binary
RUN chmod +x ./main

CMD ["./main"]