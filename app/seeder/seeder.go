package seeder


import (
	"book-online-api/app/models"
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
)

var isSeeded bool

func Seed(db *gorm.DB) {
	if isSeeded {
		return
	}

	isSeeded = true
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	if err != nil {
			log.Fatal("Failed to hash password:", err)
	}
	// Seed users
	admin := models.User{
		Name:     "Admin",
		Email:    "admin@example.com",
		Password: string(hashedPassword),
		Role:     "admin",
	}
	db.Create(&admin)

	
}

