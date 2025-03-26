package middleware

import (
	"restapi-users-management/database"
	"restapi-users-management/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func TokenMiddleware(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")

	if tokenString == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Token is not found / empty!",
			"status":  fiber.StatusNotFound,
		})

	}

	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": "Invalid token format!",
		})

	}

	checkToken, errCheckToken := checkSessionToken(tokenString)

	if errCheckToken != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "Token is not found!",
		})

	}

	if time.Now().After(checkToken.SessionExpiredAt) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": "Token is expired!",
		})
	}

	currentUser := map[string]string{
		"user_uuid":     checkToken.SessionUserUuid,
		"user_username": checkToken.SessionUserUsername,
		"user_role":     checkToken.SessionUserRole,
	}

	c.Locals("currentUser", currentUser)
	return c.Next()
}

func IsAdminAccess(c *fiber.Ctx) error {
	getCurrentUser := c.Locals("currentUser").(map[string]string)

	if getCurrentUser["user_role"] != "admin" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "This url can be accessed by admin",
		})
	}

	return c.Next()
}

func IsUserAccess(c *fiber.Ctx) error {
	getCurrentUser := c.Locals("currentUser").(map[string]string)

	if getCurrentUser["user_role"] != "user" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "This url can be accessed by user",
		})
	}

	return c.Next()
}

func CurrentUser(c *fiber.Ctx) map[string]string {
	getCurrentUser := c.Locals("currentUser").(map[string]string)

	return getCurrentUser
}

func checkSessionToken(tokenString string) (models.Session, error) {
	var session models.Session

	err := database.DB.Where("session_token = ?", tokenString).First(&session).Error

	return session, err
}
