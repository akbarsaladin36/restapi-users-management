package responses

import "restapi-users-management/models"

type UserResponse struct {
	Username    string `json:"user_username"`
	Email       string `json:"user_email"`
	Password    string `json:"user_password"`
	FirstName   string `json:"user_first_name"`
	LastName    string `json:"user_last_name"`
	Address     string `json:"user_address"`
	PhoneNumber string `json:"user_phone_number"`
	Role        string `json:"user_role"`
}

func GetUserResponse(userRsps models.User) UserResponse {
	return UserResponse{
		Username:    userRsps.UserUsername,
		Email:       userRsps.UserEmail,
		Password:    userRsps.UserPassword,
		FirstName:   userRsps.UserFirstName,
		LastName:    userRsps.UserLastName,
		Address:     userRsps.UserAddress,
		PhoneNumber: userRsps.UserPhoneNumber,
		Role:        userRsps.UserRole,
	}
}
