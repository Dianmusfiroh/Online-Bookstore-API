#!/bin/sh

# Jalankan migrasi terlebih dahulu
echo "Running database migrations..."
migrate -database "pgx://$PGUSER:$PGPASSWORD@$PGHOST:$PGPORT/$PGDATABASE?sslmode=disable" -path migrations up

# Jalankan aplikasi utama
echo "Starting the main application..."
./main