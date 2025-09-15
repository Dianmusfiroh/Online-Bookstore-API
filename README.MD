# **Online Bookstore API**

Aplikasi ini adalah REST API untuk sistem toko buku online yang dikembangkan menggunakan **Go** dengan framework **GoFiber** dan database **PostgreSQL**. Proyek ini menerapkan arsitektur modular, otentikasi JWT, dan berbagai fitur e-commerce standar.

## **Collaction Postman**
file collaction ada di file postmant dan bisa juga diakses dari https://documenter.getpostman.com/view/21349149/2sB3HqHJDT#d5fae148-2b52-4688-a369-01e60a239d2f

## **Daftar Isi**

1.  [Fitur Utama](#fitur-utama)
2.  [Teknologi yang Digunakan](#teknologi-yang-digunakan)
3.  [Panduan Instalasi](#panduan-instalasi)
4.  [Database](#database)
5.  [Penggunaan API](#penggunaan-api)
    -   [Autentikasi & Otorisasi](#autentikasi--otorisasi)
    -   [CRUD Categories (Admin Only)](#crud-categories-admin-only)
    -   [CRUD Books](#crud-books)
    -   [Orders & Payment](#orders--payment)
    -   [Reporting (Admin Only)](#reporting-admin-only)
6.  [Struktur Proyek](#struktur-proyek)

---

## **Fitur Utama**

* [cite_start]**Autentikasi & Otorisasi**: Menggunakan JWT untuk mengamankan endpoint, dengan peran `user` dan `admin`[cite: 5].
* **CRUD Categories & Books**:
    * [cite_start]Admin dapat mengelola data buku dan kategori[cite: 6].
    * [cite_start]Pengguna dapat melihat daftar buku dengan fitur **pencarian**, **filter**, dan **paginasi**[cite: 37, 70].
* **Sistem Pesanan (Orders)**:
    * [cite_start]Memungkinkan pembuatan pesanan multi-item dengan validasi stok[cite: 7, 8].
    * [cite_start]Total harga dihitung secara otomatis[cite: 45].
    * [cite_start]**Integrasi Pembayaran Xendit** untuk membuat *payment link* (simulasi API eksternal)[cite: 9, 47].
* [cite_start]**Laporan Penjualan (Reporting)**: Fitur khusus admin untuk melihat laporan penjualan, buku terlaris, dan statistik harga buku[cite: 10, 50].

---

## **Teknologi yang Digunakan**

* **Bahasa Pemrograman**: Go
* **Framework**: [GoFiber](https://gofiber.io/)
* **ORM**: [GORM](https://gorm.io/)
* **Database**: PostgreSQL
* **Manajemen Paket**: Go Modules
* **Dependensi**:
    * `github.com/golang-jwt/jwt/v5`
    * `golang.org/x/crypto/bcrypt`
    * `github.com/go-playground/validator/v10`
    * `github.com/joho/godotenv`
    * `github.com/xendit/xendit-go/v7`

---

## **Panduan Instalasi**

### **1. Clone Repositori**

```bash
git clone [https://github.com/your_username/book-online-api.git](https://github.com/your_username/book-online-api.git)
cd book-online-api

2. Konfigurasi Environment
Buat file .env di root proyek dan isi dengan konfigurasi berikut:

# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_user
DB_PASSWORD=your_password
DB_NAME=book_online_db

# Autentikasi
JWT_SECRET=supersecretkey-gantidenganstringunik
# Jika Anda mengaktifkan seeder, jalankan dengan RUN_SEEDER=true
RUN_SEEDER=true 

# Xendit
XENDIT_SECRET_KEY=xnd_development_...
XENDIT_VERIFICATION_TOKEN=webhook_secret_token
APP_URL=http://localhost:3000

README.md

Markdown

# **Online Bookstore API**

Aplikasi ini adalah REST API untuk sistem toko buku online yang dikembangkan menggunakan **Go** dengan framework **GoFiber** dan database **PostgreSQL**. Proyek ini menerapkan arsitektur modular, otentikasi JWT, dan berbagai fitur e-commerce standar.

## **Daftar Isi**

1.  [Fitur Utama](#fitur-utama)
2.  [Teknologi yang Digunakan](#teknologi-yang-digunakan)
3.  [Panduan Instalasi](#panduan-instalasi)
4.  [Database](#database)
5.  [Penggunaan API](#penggunaan-api)
    -   [Autentikasi & Otorisasi](#autentikasi--otorisasi)
    -   [CRUD Categories (Admin Only)](#crud-categories-admin-only)
    -   [CRUD Books](#crud-books)
    -   [Orders & Payment](#orders--payment)
    -   [Reporting (Admin Only)](#reporting-admin-only)
6.  [Struktur Proyek](#struktur-proyek)

---

## **Fitur Utama**

* [cite_start]**Autentikasi & Otorisasi**: Menggunakan JWT untuk mengamankan endpoint, dengan peran `user` dan `admin`[cite: 5].
* **CRUD Categories & Books**:
    * [cite_start]Admin dapat mengelola data buku dan kategori[cite: 6].
    * [cite_start]Pengguna dapat melihat daftar buku dengan fitur **pencarian**, **filter**, dan **paginasi**[cite: 37, 70].
* **Sistem Pesanan (Orders)**:
    * [cite_start]Memungkinkan pembuatan pesanan multi-item dengan validasi stok[cite: 7, 8].
    * [cite_start]Total harga dihitung secara otomatis[cite: 45].
    * [cite_start]**Integrasi Pembayaran Xendit** untuk membuat *payment link* (simulasi API eksternal)[cite: 9, 47].
* [cite_start]**Laporan Penjualan (Reporting)**: Fitur khusus admin untuk melihat laporan penjualan, buku terlaris, dan statistik harga buku[cite: 10, 50].

---

## **Teknologi yang Digunakan**

* **Bahasa Pemrograman**: Go
* **Framework**: [GoFiber](https://gofiber.io/)
* **ORM**: [GORM](https://gorm.io/)
* **Database**: PostgreSQL
* **Manajemen Paket**: Go Modules
* **Dependensi**:
    * `github.com/golang-jwt/jwt/v5`
    * `golang.org/x/crypto/bcrypt`
    * `github.com/go-playground/validator/v10`
    * `github.com/joho/godotenv`
    * `github.com/xendit/xendit-go/v7`

---

## **Panduan Instalasi**

### **1. Clone Repositori**

```bash
git clone [https://github.com/your_username/book-online-api.git](https://github.com/your_username/book-online-api.git)
cd book-online-api
2. Konfigurasi Environment
Buat file .env di root proyek dan isi dengan konfigurasi berikut:

# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_user
DB_PASSWORD=your_password
DB_NAME=book_online_db

# Autentikasi
JWT_SECRET=supersecretkey-gantidenganstringunik
# Jika Anda mengaktifkan seeder, jalankan dengan RUN_SEEDER=true
RUN_SEEDER=true 

# Xendit
XENDIT_SECRET_KEY=xnd_development_...
XENDIT_VERIFICATION_TOKEN=webhook_secret_token
APP_URL=http://localhost:3000
3. Instalasi Dependensi dan Database Migration
Bash

# Instal dependensi Go
go mod tidy

# Instal tool migrasi
go install -tags 'postgres' [github.com/golang-migrate/migrate/v4/cmd/migrate@latest](https://github.com/golang-migrate/migrate/v4/cmd/migrate@latest)

# Jalankan migrasi database
migrate -path migrations -database "postgresql://<DB_USER>:<DB_PASSWORD>@<DB_HOST>:<DB_PORT>/<DB_NAME>?sslmode=disable" up

README.md

Markdown

# **Online Bookstore API**

Aplikasi ini adalah REST API untuk sistem toko buku online yang dikembangkan menggunakan **Go** dengan framework **GoFiber** dan database **PostgreSQL**. Proyek ini menerapkan arsitektur modular, otentikasi JWT, dan berbagai fitur e-commerce standar.

## **Daftar Isi**

1.  [Fitur Utama](#fitur-utama)
2.  [Teknologi yang Digunakan](#teknologi-yang-digunakan)
3.  [Panduan Instalasi](#panduan-instalasi)
4.  [Database](#database)
5.  [Penggunaan API](#penggunaan-api)
    -   [Autentikasi & Otorisasi](#autentikasi--otorisasi)
    -   [CRUD Categories (Admin Only)](#crud-categories-admin-only)
    -   [CRUD Books](#crud-books)
    -   [Orders & Payment](#orders--payment)
    -   [Reporting (Admin Only)](#reporting-admin-only)
6.  [Struktur Proyek](#struktur-proyek)

---

## **Fitur Utama**

* [cite_start]**Autentikasi & Otorisasi**: Menggunakan JWT untuk mengamankan endpoint, dengan peran `user` dan `admin`[cite: 5].
* **CRUD Categories & Books**:
    * [cite_start]Admin dapat mengelola data buku dan kategori[cite: 6].
    * [cite_start]Pengguna dapat melihat daftar buku dengan fitur **pencarian**, **filter**, dan **paginasi**[cite: 37, 70].
* **Sistem Pesanan (Orders)**:
    * [cite_start]Memungkinkan pembuatan pesanan multi-item dengan validasi stok[cite: 7, 8].
    * [cite_start]Total harga dihitung secara otomatis[cite: 45].
    * [cite_start]**Integrasi Pembayaran Xendit** untuk membuat *payment link* (simulasi API eksternal)[cite: 9, 47].
* [cite_start]**Laporan Penjualan (Reporting)**: Fitur khusus admin untuk melihat laporan penjualan, buku terlaris, dan statistik harga buku[cite: 10, 50].

---

## **Teknologi yang Digunakan**

* **Bahasa Pemrograman**: Go
* **Framework**: [GoFiber](https://gofiber.io/)
* **ORM**: [GORM](https://gorm.io/)
* **Database**: PostgreSQL
* **Manajemen Paket**: Go Modules
* **Dependensi**:
    * `github.com/golang-jwt/jwt/v5`
    * `golang.org/x/crypto/bcrypt`
    * `github.com/go-playground/validator/v10`
    * `github.com/joho/godotenv`
    * `github.com/xendit/xendit-go/v7`

---

## **Panduan Instalasi**

### **1. Clone Repositori**

```bash
git clone [https://github.com/your_username/book-online-api.git](https://github.com/your_username/book-online-api.git)
cd book-online-api
2. Konfigurasi Environment
Buat file .env di root proyek dan isi dengan konfigurasi berikut:

# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_user
DB_PASSWORD=your_password
DB_NAME=book_online_db

# Autentikasi
JWT_SECRET=supersecretkey-gantidenganstringunik
# Jika Anda mengaktifkan seeder, jalankan dengan RUN_SEEDER=true
RUN_SEEDER=true 

# Xendit
XENDIT_SECRET_KEY=xnd_development_...
XENDIT_VERIFICATION_TOKEN=webhook_secret_token
APP_URL=http://localhost:3000
3. Instalasi Dependensi dan Database Migration
Bash

# Instal dependensi Go
go mod tidy

# Instal tool migrasi
go install -tags 'postgres' [github.com/golang-migrate/migrate/v4/cmd/migrate@latest](https://github.com/golang-migrate/migrate/v4/cmd/migrate@latest)

# Jalankan migrasi database
migrate -path migrations -database "postgresql://<DB_USER>:<DB_PASSWORD>@<DB_HOST>:<DB_PORT>/<DB_NAME>?sslmode=disable" up
4. Menjalankan Aplikasi
Bash

# Jalankan seeder untuk membuat data awal (admin user, dsb.)
RUN_SEEDER=true go run main.go

# Jalankan server tanpa seeder
go run main.go

Aplikasi akan berjalan di http://localhost:3000.

Penggunaan API
Autentikasi & Otorisasi

POST /api/register: Registrasi user baru.


POST /api/login: Login dan mendapatkan JWT token.

Semua endpoint (kecuali register & login) hanya bisa diakses dengan token.

CRUD Categories (Admin Only)

POST /api/admin/categories 


GET /api/categories 


PUT /api/admin/categories/:id 


DELETE /api/admin/categories/:id 

CRUD Books

POST /api/admin/books: Buat buku baru (Admin Only), dengan gambar base64.


GET /api/books: Lihat semua buku (User & Admin), mendukung paginasi, pencarian berdasarkan judul/penulis, dan filter berdasarkan kategori.


GET /api/books/:id: Lihat detail buku.


PUT /api/admin/books/:id: Update buku (Admin Only).


DELETE /api/admin/books/:id: Hapus buku (Admin Only).

Orders & Payment
POST /api/orders: Buat pesanan multi-item dan dapatkan payment link Xendit.

POST /api/webhooks/xendit: Endpoint untuk menerima notifikasi pembayaran dari Xendit.


GET /api/orders: Lihat daftar pesanan (user hanya melihat pesanan sendiri, admin melihat semua).


GET /api/orders/:id: Lihat detail pesanan (hanya pemilik atau admin).

Reporting (Admin Only)

GET /api/admin/reports/sales: Total omzet (revenue) dan total buku terjual.


GET /api/admin/reports/bestseller: Daftar 3 buku terlaris.


GET /api/admin/reports/prices: Harga maksimum, minimum, dan rata-rata buku.

Struktur Proyek
book-online-api/
├── main.go               # Entry point
├── go.mod
├── .env                  # Environment variables
├── /internal
│   ├── config/           # Database configuration
│   ├── dto/              # Data Transfer Objects
│   ├── handlers/         # HTTP request handlers
│   ├── middleware/       # Middleware untuk otentikasi & otorisasi
│   ├── models/           # Struct/model database
│   ├── repository/       # Database access logic
│   ├── routes/           # API routes definition
│   ├── seeder/           # Seeder untuk data awal
│   └── services/         # Business logic
└── /migrations           # SQL files for database migrations