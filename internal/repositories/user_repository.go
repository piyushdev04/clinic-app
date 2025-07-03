package repositories

import (
	"clinic-app/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetByUsername(username string) (*models.User, error)
}

type GormUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &GormUserRepository{db}
}

func (r *GormUserRepository) GetByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
