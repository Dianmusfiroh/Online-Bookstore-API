#!/bin/sh

# Pindah ke direktori /app
cd /app

# Jalankan migrasi terlebih dahulu
echo "Running database migrations..."
migrate -database "pgx://postgres:xidOShwWcHtxASEEHcptmJTfFsEDZHVX@$PGHOST:$PGPORT/$PGDATABASE?sslmode=disable" -path migrations up
cd migrations
./main
cd ..
# Jalankan aplikasi utama
echo "Starting the main application..."
./main