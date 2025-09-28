# -- Tahap 1: Kompilasi aplikasi (builder) --
FROM golang:1.22.1-alpine AS builder

WORKDIR /app

# Salin go.mod dan go.sum
COPY go.mod go.sum ./

# Unduh dependensi proyek
RUN go mod download

# Instal alat migrasi
# GO111MODULE=on memastikan go install bekerja dengan benar
RUN GO111MODULE=on go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Salin semua kode sumber dari lokal
COPY . .

# Kompilasi aplikasi utama Anda
ENV GOOS=linux
ENV GOARCH=amd64
RUN CGO_ENABLED=0 go build -o main .

# -- Tahap 2: Buat image produksi yang bersih (final) --
FROM alpine:3.18

WORKDIR /app

# Salin binary aplikasi utama dan alat migrasi dari builder
COPY --from=builder /go/bin/migrate /usr/local/bin/migrate
COPY --from=builder /app/main .

# Salin skrip startup
COPY ./run.sh .

# Tambahkan izin eksekusi ke file binary dan skrip
RUN chmod +x ./main ./run.sh

# Jalankan skrip startup
CMD ["./run.sh"]