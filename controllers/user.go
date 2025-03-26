package controllers

import (
	"restapi-users-management/inputs"
	"restapi-users-management/middleware"
	"restapi-users-management/responses"
	"restapi-users-management/services"

	"github.com/gofiber/fiber/v2"
)

type userController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *userController {
	return &userController{userService}
}

// FindUsersController godoc
// @Summary Get all users
// @Description Get all users for admin role
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} responses.UserResponse
// @Failure 404 {object} object{status=string,message=string}
// @Failure 400 {object} object{status=string,message=string}
// @Router /admin/users [get]
func (uc *userController) FindUsersController(c *fiber.Ctx) error {
	users, err := uc.userService.FindUsersService()

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "All users data is succesfully appeared!",
		})
	}

	var userResponse []responses.UserResponse

	for _, user := range users {
		usersRsps := responses.GetUserResponse(user)

		userResponse = append(userResponse, usersRsps)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "All users data is succesfully appeared!",
		"data":    userResponse,
	})
}

// FindUserController godoc
// @Summary Get user detail information
// @Description Get user detail information by username
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user_username path string true "Username"
// @Success 200 {object} responses.UserResponse
// @Failure 404 {object} object{status=string,message=string}
// @Failure 400 {object} object{status=string,message=string}
// @Router /admin/users/detail-user/{user_username} [get]
func (uc *userController) FindUserController(c *fiber.Ctx) error {
	user_username := c.Params("user_username")

	user, err := uc.userService.FindUserService(user_username)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "A user data " + user_username + "are not found!",
		})
	}

	userRsps := responses.GetUserResponse(user)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "A user data " + user_username + " are succesfully appeared!",
		"data":    userRsps,
	})
}

// CreateUserController godoc
// @Summary Create a new user
// @Description registering a new user
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param credentials body inputs.CreateUserInput true "Create a new user"
// @Success 200 {object} responses.UserResponse
// @Failure 404 {object} object{status=string,message=string}
// @Failure 400 {object} object{status=string,message=string}
// @Router /admin/users [post]
func (uc *userController) CreateUserController(c *fiber.Ctx) error {
	var createUserInput inputs.CreateUserInput

	err := c.BodyParser(&createUserInput)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "400",
			"message": "Invalid request body",
		})
	}

	_, errCheckUser := uc.userService.FindUserService(createUserInput.Username)

	if errCheckUser == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "400",
			"message": "Your username " + createUserInput.Username + " is registered! Please try find another username!",
		})
	}

	currentUser := middleware.CurrentUser(c)

	newCreateUser, errCreateUser := uc.userService.CreateUserService(createUserInput, currentUser)

	if errCreateUser != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "400",
			"message": "Process creating a new user is failed! Please try again!",
		})
	}

	newUserRsps := responses.GetUserResponse(newCreateUser)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "A new user are succesfully created!",
		"data":    newUserRsps,
	})
}

// UpdateUserController godoc
// @Summary Update existing user detail information
// @Description Update user detail information from username as parameter
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user_username path string true "Username"
// @Param credentials body inputs.UpdateUserInput true "update existing user data"
// @Success 200 {object} responses.UserResponse
// @Failure 404 {object} object{status=string,message=string}
// @Failure 400 {object} object{status=string,message=string}
// @Router /admin/users/detail-user/{user_username} [patch]
func (uc *userController) UpdateUserController(c *fiber.Ctx) error {
	user_username := c.Params("user_username")

	_, errUser := uc.userService.FindUserService(user_username)

	if errUser != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "A user data " + user_username + "are not found!",
		})
	}

	var updateUserInput inputs.UpdateUserInput

	err := c.BodyParser(&updateUserInput)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "400",
			"message": "Invalid request body",
		})
	}

	currentUser := middleware.CurrentUser(c)

	updateUser, errUpdateUser := uc.userService.UpdateUserService(user_username, updateUserInput, currentUser)

	if errUpdateUser != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Error when updating a user data for username " + user_username + " ! Please try again!",
		})
	}

	userRsps := responses.GetUserResponse(updateUser)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "A user data " + user_username + " are succesfully updated!",
		"data":    userRsps,
	})
}

// DeleteUserController godoc
// @Summary Delete user detail information
// @Description Delete user detail information by username
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user_username path string true "Username"
// @Success 200 {object} object{status=string,message=string}
// @Failure 404 {object} object{status=string,message=string}
// @Failure 400 {object} object{status=string,message=string}
// @Router /admin/users/detail-user/{user_username} [delete]
func (uc *userController) DeleteUserController(c *fiber.Ctx) error {
	user_username := c.Params("user_username")

	_, errUser := uc.userService.FindUserService(user_username)

	if errUser != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "A user data " + user_username + "are not found!",
		})
	}

	_, errDeleteUser := uc.userService.DeleteUserService(user_username)

	if errDeleteUser != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Error when deleting a user data for username " + user_username + " ! Please try again!",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "A user data " + user_username + " are succesfully deleted!",
	})
}
