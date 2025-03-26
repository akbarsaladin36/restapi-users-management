package controllers

import (
	"restapi-users-management/inputs"
	"restapi-users-management/middleware"
	"restapi-users-management/responses"
	"restapi-users-management/services"

	"github.com/gofiber/fiber/v2"
)

type profileController struct {
	profileService services.ProfileService
}

func NewProfileController(profileService services.ProfileService) *profileController {
	return &profileController{profileService}
}

// FindProfileController godoc
// @Summary Get profile detail information
// @Description Get profile detail information by username
// @Tags profile
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} responses.UserResponse
// @Failure 404 {object} object{status=string,message=string}
// @Failure 400 {object} object{status=string,message=string}
// @Router /admin/profile [get]
// @Router /user/profile [get]
func (pc *profileController) FindProfileController(c *fiber.Ctx) error {
	currentUser := middleware.CurrentUser(c)

	user_username := currentUser["user_username"]

	user, err := pc.profileService.FindProfileService(user_username)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "A profile for username " + user_username + " is not found! Please try again!",
		})
	}

	userRsps := responses.GetUserResponse(user)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "A profile for username " + user_username + " is succesfully appeared!",
		"data":    userRsps,
	})
}

// UpdateProfileController godoc
// @Summary Update profile detail information
// @Description Updating profile detail information by username
// @Tags profile
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param credentials body inputs.UpdateUserInput true "update profile detail information for user"
// @Success 200 {object} responses.UserResponse
// @Failure 404 {object} object{status=string,message=string}
// @Failure 400 {object} object{status=string,message=string}
// @Router /admin/profile [patch]
// @Router /user/profile [patch]
func (pc *profileController) UpdateProfileController(c *fiber.Ctx) error {
	currentUser := middleware.CurrentUser(c)

	user_username := currentUser["user_username"]

	_, err := pc.profileService.FindProfileService(user_username)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "A profile for username " + user_username + " is not found! Please try again!",
		})
	}

	var updateUserInput inputs.UpdateUserInput

	errUpdateUserBody := c.BodyParser(&updateUserInput)

	if errUpdateUserBody != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Invalid request body",
		})
	}

	updateUser, errUpdateUser := pc.profileService.UpdateProfileService(user_username, currentUser, updateUserInput)

	if errUpdateUser != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Error when updating a profile data for username " + user_username + " ! Please try again!",
		})
	}

	userRsps := responses.GetUserResponse(updateUser)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "A profile data for username " + user_username + " are succesfully updated!",
		"data":    userRsps,
	})

}
