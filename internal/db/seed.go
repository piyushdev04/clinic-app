package db

import (
	"clinic-app/internal/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedUsers(database *gorm.DB) {
	passwordReceptionist, _ := bcrypt.GenerateFromPassword([]byte("receptionist123"), bcrypt.DefaultCost)
	passwordDoctor, _ := bcrypt.GenerateFromPassword([]byte("doctor123"), bcrypt.DefaultCost)

	database.FirstOrCreate(&models.User{}, models.User{
		Username:     "receptionist",
		PasswordHash: string(passwordReceptionist),
		Role:         "receptionist",
	})
	database.FirstOrCreate(&models.User{}, models.User{
		Username:     "doctor",
		PasswordHash: string(passwordDoctor),
		Role:         "doctor",
	})
}
