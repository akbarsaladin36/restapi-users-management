package services

import (
	"restapi-users-management/helpers"
	"restapi-users-management/inputs"
	"restapi-users-management/models"
	"restapi-users-management/repositories"
	"strings"
	"time"
)

type UserService interface {
	FindUsersService() ([]models.User, error)
	FindUserService(user_username string) (models.User, error)
	CreateUserService(createUserInput inputs.CreateUserInput, currentUser map[string]string) (models.User, error)
	UpdateUserService(user_username string, updateUserInput inputs.UpdateUserInput, currentUser map[string]string) (models.User, error)
	DeleteUserService(user_username string) (models.User, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *userService {
	return &userService{userRepository}
}

func (us *userService) FindUsersService() ([]models.User, error) {
	users, err := us.userRepository.FindAll()

	return users, err
}

func (us *userService) FindUserService(user_username string) (models.User, error) {
	user, err := us.userRepository.FindOne(user_username)

	return user, err
}

func (us *userService) CreateUserService(createUserInput inputs.CreateUserInput, currentUser map[string]string) (models.User, error) {
	hashedPassword, _ := helpers.HashPassword(createUserInput.Password)

	userUUID := strings.ReplaceAll(helpers.GenerateUUID(createUserInput.Username), "-", "")

	user := models.User{
		UserUuid:                userUUID,
		UserUsername:            createUserInput.Username,
		UserEmail:               createUserInput.Email,
		UserPassword:            hashedPassword,
		UserStatusCd:            "active",
		UserRole:                "user",
		UserCreatedDate:         time.Now(),
		UserCreatedUserUuid:     currentUser["user_uuid"],
		UserCreatedUserUsername: currentUser["user_username"],
	}

	newCreateUser, err := us.userRepository.Create(user)

	return newCreateUser, err
}

func (us *userService) UpdateUserService(user_username string, updateUserInput inputs.UpdateUserInput, currentUser map[string]string) (models.User, error) {
	checkUser, _ := us.userRepository.FindOne(user_username)

	checkUser.UserFirstName = updateUserInput.FirstName
	checkUser.UserLastName = updateUserInput.LastName
	checkUser.UserAddress = updateUserInput.Address
	checkUser.UserPhoneNumber = updateUserInput.PhoneNumber
	checkUser.UserRole = updateUserInput.Role
	checkUser.UserUpdatedDate = time.Now()
	checkUser.UserUpdatedUserUuid = currentUser["user_uuid"]
	checkUser.UserUpdatedUserUsername = currentUser["user_username"]

	updateUser, err := us.userRepository.Update(checkUser)

	return updateUser, err
}

func (us *userService) DeleteUserService(user_username string) (models.User, error) {
	checkUser, _ := us.userRepository.FindOne(user_username)

	deleteUser, err := us.userRepository.Delete(checkUser)

	return deleteUser, err
}
