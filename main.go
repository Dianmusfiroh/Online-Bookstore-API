package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	// coba load .env, tapi abaikan error kalau file tidak ada
	_ = godotenv.Load()

	dsn := os.Getenv("DATABASE_URL") // Railway biasanya kasih ini
	if dsn == "" {
		// fallback manual kalau env var belum diset
		host := os.Getenv("DB_HOST")
		user := os.Getenv("DB_USER")
		pass := os.Getenv("DB_PASS")
		name := os.Getenv("DB_NAME")
		port := os.Getenv("DB_PORT")
		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			host, user, pass, name, port,
		)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect database:", err)
		return nil
	}
	return db
}
