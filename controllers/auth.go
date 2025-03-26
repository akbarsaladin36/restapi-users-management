package controllers

import (
	"restapi-users-management/helpers"
	"restapi-users-management/inputs"
	"restapi-users-management/responses"
	"restapi-users-management/services"

	"github.com/gofiber/fiber/v2"
)

type authController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *authController {
	return &authController{authService}
}

// Register godoc
// @Summary User register
// @Description Register a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body inputs.RegisterInput true "Register credentials"
// @Success 200 {object} responses.RegisterResponse
// @Failure 302 {object} object{status=string,message=string}
// @Router /auth/register [post]
func (ac *authController) RegisterController(c *fiber.Ctx) error {
	var registerInput inputs.RegisterInput

	err := c.BodyParser(&registerInput)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "400",
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	_, errCheckUser := ac.authService.FindOneService(registerInput.Username)

	if errCheckUser == nil {
		c.Status(fiber.StatusFound).JSON(fiber.Map{
			"status":  "302",
			"message": "The username " + registerInput.Username + " is exist! Please try find a new username!",
		})
	}

	newRegisterUser, _ := ac.authService.RegisterService(registerInput)

	registerUserRsps := responses.GetRegisterResponse(newRegisterUser)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "200",
		"message": "A new user is succesfully created!",
		"data":    registerUserRsps,
	})
}

// Login godoc
// @Summary User login
// @Description authenticate user and get a token
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body inputs.LoginInput true "Login credentials"
// @Success 200 {object} responses.LoginResponse
// @Failure 302 {object} object{status=string,message=string}
// @Router /auth/login [post]
func (ac *authController) LoginController(c *fiber.Ctx) error {
	var loginInput inputs.LoginInput

	err := c.BodyParser(&loginInput)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "400",
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	_, errCheckUser := ac.authService.FindOneService(loginInput.Username)

	if errCheckUser != nil {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "404",
			"message": "The username " + loginInput.Username + " is not exist! Please register this username first!",
		})
	}

	checkLoginUser, errLoginUser := ac.authService.LoginService(loginInput)

	if errLoginUser != nil {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "404",
			"message": "Process login user are error! Please try again!",
		})
	}

	checkSession, errCheckSessionUser := ac.authService.FindSessionService(checkLoginUser.UserUuid)

	tokenString := helpers.GenerateSessionToken(300)

	if errCheckSessionUser != nil {
		ac.authService.CreateSessionService(checkLoginUser, tokenString)

		loginResponse := responses.GetLoginResponse(checkLoginUser, tokenString)

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "A login process is succesfull!",
			"data":    loginResponse,
		})
	} else {
		ac.authService.UpdateSessionService(checkLoginUser.UserUuid, tokenString, checkSession)

		loginResponse := responses.GetLoginResponse(checkLoginUser, tokenString)

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "A login process is succesfull!",
			"data":    loginResponse,
		})
	}

}
