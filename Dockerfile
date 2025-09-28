# -- Tahap 1: Kompilasi aplikasi (builder) --
FROM golang:1.22.1-alpine AS builder

WORKDIR /app

# Salin go.mod dan go.sum
COPY go.mod go.sum ./

# Unduh semua dependensi
RUN go mod download

# === BARIS YANG DIPERBAIKI ===
# Instal alat migrasi dengan driver pgx
# go install github.com/golang-migrate/migrate/v4/cmd/migrate@v4.18.0
RUN GO111MODULE=on go install github.com/golang-migrate/migrate/v4/cmd/migrate@v4.18.0
# ==============================

# Salin semua kode sumber dari lokal
COPY . .

# Secara eksplisit tentukan OS dan Arsitektur target
ENV GOOS=linux
ENV GOARCH=amd64

# Kompilasi aplikasi utama Anda
RUN CGO_ENABLED=0 go build -o main .

# -- Tahap 2: Buat image produksi yang bersih (final) --
FROM alpine:3.18

WORKDIR /app

# Salin binary aplikasi utama dan alat migrasi dari builder
COPY --from=builder /go/bin/migrate /usr/local/bin/migrate
COPY --from=builder /app/main .

# Salin skrip startup
COPY ./run.sh .

# Salin folder migrasi
COPY migrations migrations

# Tambahkan izin eksekusi ke file binary dan skrip
RUN chmod +x ./main ./run.sh

# Jalankan skrip startup
CMD ["./run.sh"]