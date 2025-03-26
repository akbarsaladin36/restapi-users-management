package repositories

import (
	"restapi-users-management/models"

	"gorm.io/gorm"
)

type ProfileRepository interface {
	FindOne(user_username string) (models.User, error)
	Update(user models.User) (models.User, error)
}

type profileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) *profileRepository {
	return &profileRepository{db}
}

func (pr *profileRepository) FindOne(user_username string) (models.User, error) {
	var user models.User

	err := pr.db.Where("user_username = ?", user_username).First(&user).Error

	return user, err
}

func (pr *profileRepository) Update(user models.User) (models.User, error) {
	err := pr.db.Save(&user).Error

	return user, err
}
