package routes

import (
	"os"
	"restapi-users-management/controllers"
	"restapi-users-management/database"
	"restapi-users-management/middleware"
	"restapi-users-management/repositories"
	"restapi-users-management/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
)

func ConnectRoutes() {
	app := fiber.New()
	app.Use(cors.New())
	app.Static("/docs", "./docs")
	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL:         "/docs/swagger.json",
		DeepLinking: true,
	}))

	authRepositories := repositories.NewAuthRepository(database.DB)
	authServices := services.NewAuthService(authRepositories)
	authControllers := controllers.NewAuthController(authServices)

	userRepositories := repositories.NewUserRepository(database.DB)
	userServices := services.NewUserService(userRepositories)
	userControllers := controllers.NewUserController(userServices)

	profileRepositories := repositories.NewProfileRepository(database.DB)
	profileServices := services.NewProfileService(profileRepositories)
	profileControllers := controllers.NewProfileController(profileServices)

	v1Api := app.Group("/api/v1")

	v1Api.Post("/auth/register", authControllers.RegisterController)
	v1Api.Post("/auth/login", authControllers.LoginController)

	// ------ Admin Route ------ //
	v1Admin := v1Api.Group("/admin", middleware.TokenMiddleware, middleware.IsAdminAccess)

	// Users (Admin)
	v1Admin.Get("/users", userControllers.FindUsersController)
	v1Admin.Get("/users/detail-user/:user_username", userControllers.FindUserController)
	v1Admin.Post("/users", userControllers.CreateUserController)
	v1Admin.Patch("/users/detail-user/:user_username", userControllers.UpdateUserController)
	v1Admin.Delete("/users/detail-user/:user_username", userControllers.DeleteUserController)

	// Profile (Admin)
	v1Admin.Get("/profile", profileControllers.FindProfileController)
	v1Admin.Patch("/profile", profileControllers.UpdateProfileController)

	// ------ User Route ------ //
	v1User := v1Api.Group("/user", middleware.TokenMiddleware, middleware.IsUserAccess)

	// Profile (User)
	v1User.Get("/profile", profileControllers.FindProfileController)
	v1User.Patch("/profile", profileControllers.UpdateProfileController)

	app.Listen(os.Getenv("APP_PORT"))

}
