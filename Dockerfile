# -- Tahap 1: Kompilasi aplikasi (builder) --
# Gunakan image Golang yang spesifik untuk versi go.mod Anda
FROM golang:1.22.1-alpine AS builder

# Atur direktori kerja di dalam kontainer
WORKDIR /app

# Salin file go.mod dan go.sum untuk mengelola cache dependensi
COPY go.mod .
COPY go.sum .

# Unduh semua dependensi
RUN go mod download

# Salin semua kode sumber dari lokal
COPY . .

# Secara eksplisit tentukan OS dan Arsitektur target
# Ini mencegah masalah "Exec format error"
ENV GOOS=linux
ENV GOARCH=amd64

# Kompilasi aplikasi Anda menjadi file executable
RUN CGO_ENABLED=0 go build -o main -v

# -- Tahap 2: Buat image produksi yang bersih (final) --
# Gunakan image Alpine yang stabil dan aman
FROM alpine:3.18

# Atur direktori kerja di dalam image final

# Salin binary (file executable) dari tahap builder ke tahap final
COPY --from=builder main main

# Tambahkan izin eksekusi ke file binary
RUN chmod +x ./main

# Tentukan perintah untuk menjalankan aplikasi
CMD ["./main"]