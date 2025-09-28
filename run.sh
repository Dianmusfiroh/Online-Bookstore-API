#!/bin/sh

# Pindah ke direktori /app
cd /app

# Jalankan migrasi menggunakan program yang kita kompilasi
echo "Running database migrations..."
./migrate_runner

# Jalankan aplikasi utama
echo "Starting the main application..."
./main