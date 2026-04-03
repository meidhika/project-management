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

	// User
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	// Board
	boardRepo := repositories.NewBoardRepository()
	boardMemberRepo := repositories.NewBoardMemberRepository()
	boardService := services.NewBoardService(boardRepo, userRepo, boardMemberRepo)
	boardController := controllers.NewBoardController(boardService)

	// list
	listRepo := repositories.NewListRepository()
	listPosRepo := repositories.NewListPositionRepository()
	listService := services.NewListService(listRepo, boardRepo, listPosRepo)
	listController := controllers.NewListController(listService)

	routes.Setup(app, userController, boardController, listController)
	port := config.AppConfig.AppPort
	log.Println("Server is runnning on port :", port)
	log.Fatal(app.Listen(":" + port))
	
}