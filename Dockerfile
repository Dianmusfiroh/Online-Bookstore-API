# -- Tahap 1: Kompilasi aplikasi (builder) --
FROM golang:1.22.1-alpine AS builder

WORKDIR /app

# Salin go.mod dan go.sum untuk mengunduh dependensi secara efisien
COPY go.mod go.sum ./

# Unduh semua dependensi
RUN go mod download

# Salin semua kode sumber dari lokal ke direktori kerja
COPY . .

# Secara eksplisit tentukan OS dan Arsitektur target
# Ini mencegah masalah "Exec format error"
ENV GOOS=linux
ENV GOARCH=amd64

# Kompilasi aplikasi Anda menjadi file executable
# Asumsi: main.go berada di root direktori proyek Anda
RUN CGO_ENABLED=0 go build -o main .

# -- Tahap 2: Buat image produksi yang bersih (final) --
FROM alpine:3.18

WORKDIR /app

# Salin binary (file executable) dari tahap builder ke tahap final
COPY --from=builder /app/main .

# Tambahkan izin eksekusi ke file binary
RUN chmod +x ./main

# Tentukan perintah untuk menjalankan aplikasi
CMD ["./main"]