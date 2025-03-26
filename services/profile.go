package services

import (
	"restapi-users-management/inputs"
	"restapi-users-management/models"
	"restapi-users-management/repositories"
	"time"
)

type ProfileService interface {
	FindProfileService(user_username string) (models.User, error)
	UpdateProfileService(user_username string, currentUser map[string]string, updateUserInput inputs.UpdateUserInput) (models.User, error)
}

type profileService struct {
	profileRepository repositories.ProfileRepository
}

func NewProfileService(profileRepository repositories.ProfileRepository) *profileService {
	return &profileService{profileRepository}
}

func (ps *profileService) FindProfileService(user_username string) (models.User, error) {
	user, err := ps.profileRepository.FindOne(user_username)

	return user, err
}

func (ps *profileService) UpdateProfileService(user_username string, currentUser map[string]string, updateUserInput inputs.UpdateUserInput) (models.User, error) {
	checkUser, _ := ps.profileRepository.FindOne(user_username)

	checkUser.UserFirstName = updateUserInput.FirstName
	checkUser.UserLastName = updateUserInput.LastName
	checkUser.UserAddress = updateUserInput.Address
	checkUser.UserPhoneNumber = updateUserInput.PhoneNumber
	checkUser.UserRole = updateUserInput.Role
	checkUser.UserUpdatedDate = time.Now()
	checkUser.UserUpdatedUserUuid = currentUser["user_uuid"]
	checkUser.UserUpdatedUserUsername = currentUser["user_username"]

	updateProfile, err := ps.profileRepository.Update(checkUser)

	return updateProfile, err
}
