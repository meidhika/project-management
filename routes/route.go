package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/meidhika/project-management/controllers"
)

func Setup(app *fiber.App, uc *controllers.UserController) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}


	app.Post("/v1/auth/register", uc.Register)
	app.Post("/v1/auth/login", uc.Login)
	app.Get("/users", uc.GetUserPagination)
	app.Get("/users/:id", uc.GetUser)
	app.Put("/users/:id", uc.UpdateUser)
	app.Delete("/users/:id", uc.DeleteUser)
}