package repositories

import (
	"restapi-users-management/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	FindOne(user_username string) (models.User, error)
	Create(user models.User) (models.User, error)
	FindSession(user_uuid string) (models.Session, error)
	CreateSession(session models.Session) (models.Session, error)
	UpdateSession(user_uuid string, session models.Session) (models.Session, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{db}
}

func (ar *authRepository) FindOne(user_username string) (models.User, error) {
	var user models.User

	err := ar.db.Where("user_username = ?", user_username).First(&user).Error

	return user, err
}

func (ar *authRepository) FindSession(user_uuid string) (models.Session, error) {
	var session models.Session

	err := ar.db.Where("session_user_uuid = ?", user_uuid).First(&session).Error

	return session, err
}

func (ar *authRepository) Create(user models.User) (models.User, error) {
	err := ar.db.Create(&user).Error

	return user, err
}

func (ar *authRepository) CreateSession(session models.Session) (models.Session, error) {
	err := ar.db.Create(&session).Error

	return session, err
}

func (ar *authRepository) UpdateSession(user_uuid string, session models.Session) (models.Session, error) {
	err := ar.db.Where("session_user_uuid = ?", user_uuid).Save(&session).Error

	return session, err
}
