package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/meidhika/project-management/config"
	"github.com/meidhika/project-management/controllers"
	"github.com/meidhika/project-management/database/seed"
	"github.com/meidhika/project-management/repositories"
	"github.com/meidhika/project-management/routes"
	"github.com/meidhika/project-management/services"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()
	seed.SeedAdmin()

	app := fiber.New()
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	routes.Setup(app, userController)
	port := config.AppConfig.AppPort
	log.Println("Server is runnning on port :", port)
	log.Fatal(app.Listen(":" + port))
	
}