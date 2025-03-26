package responses

import (
	"restapi-users-management/models"
	"time"
)

type LoginResponse struct {
	UserUuid     string `json:"user_uuid"`
	UserUsername string `json:"user_username"`
	UserRole     string `json:"user_role"`
	UserStatusCd string `json:"user_status_cd"`
	UserToken    string `json:"user_token"`
}

type RegisterResponse struct {
	UserUsername            string    `json:"user_username"`
	UserEmail               string    `json:"user_email"`
	UserStatusCd            string    `json:"user_status_cd"`
	UserRole                string    `json:"user_role"`
	UserCreatedDate         time.Time `json:"user_created_date"`
	UserCreatedUserUuid     string    `json:"user_created_user_uuid"`
	UserCreatedUserUsername string    `json:"user_created_user_username"`
}

func GetLoginResponse(userRsps models.User, tokenString string) LoginResponse {
	return LoginResponse{
		UserUuid:     userRsps.UserUuid,
		UserUsername: userRsps.UserUsername,
		UserRole:     userRsps.UserRole,
		UserStatusCd: userRsps.UserStatusCd,
		UserToken:    tokenString,
	}
}

func GetRegisterResponse(userRsps models.User) RegisterResponse {
	return RegisterResponse{
		UserUsername:            userRsps.UserUsername,
		UserEmail:               userRsps.UserEmail,
		UserStatusCd:            userRsps.UserStatusCd,
		UserRole:                userRsps.UserRole,
		UserCreatedDate:         userRsps.UserCreatedDate,
		UserCreatedUserUuid:     userRsps.UserCreatedUserUuid,
		UserCreatedUserUsername: userRsps.UserCreatedUserUsername,
	}
}
