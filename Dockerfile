# -- Tahap 1: Kompilasi aplikasi (builder) --
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Salin go.mod dan go.sum
COPY go.mod go.sum ./

# Unduh dependensi proyek
RUN go mod download

# Salin semua kode sumber dari lokal
COPY . .

# Kompilasi aplikasi utama Anda
ENV GOOS=linux
ENV GOARCH=amd64
RUN CGO_ENABLED=0 go build -o main .

# === Kompilasi program migrasi terpisah ===
# Catatan: kita tidak perlu lagi menginstal alat migrate secara eksternal
RUN CGO_ENABLED=0 go build -o migrate_runner ./migrations

# -- Tahap 2: Buat image produksi yang bersih (final) --
FROM alpine:3.18

WORKDIR /app

# Salin binary aplikasi utama dan program migrasi dari builder
COPY --from=builder /app/main .
COPY --from=builder /app/migrate_runner .

# Salin skrip startup
COPY ./run.sh .

# Salin folder migrasi
COPY migrations migrations

# Tambahkan izin eksekusi ke semua file
RUN chmod +x ./main ./migrate_runner ./run.sh

# Jalankan skrip startup
CMD ["./run.sh"]