package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/joho/godotenv"
	"github.com/meidhika/project-management/config"
	"github.com/meidhika/project-management/controllers"
	"github.com/meidhika/project-management/utils"
)

func Setup(app *fiber.App, 
	uc *controllers.UserController,
	bc *controllers.BoardController,
	lc *controllers.ListController,
	cc *controllers.CardController) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}


	app.Post("/v1/auth/register", uc.Register)
	app.Post("/v1/auth/login", uc.Login)

	api := app.Group("/api/v1", jwtware.New(jwtware.Config{SigningKey: []byte(config.AppConfig.JWTSecret),
	ContextKey: "user",
	ErrorHandler: func (c *fiber.Ctx, err error) error{
		return utils.Unauthorized(c, "Error Unauthorized", err.Error())
		},
	}))

	// users
	userGroup := api.Group("/users")
	userGroup.Get("/page", uc.GetUserPagination) // /api/v1/users/page?page=1&limit=10&sort=-id&filter=triady
	userGroup.Get("/:id", uc.GetUser) // /api/v1/users/:id
	userGroup.Put("/:id", uc.UpdateUser)
	userGroup.Delete("/:id", uc.DeleteUser)

	// boards
	boardGroup := api.Group("/boards")
	boardGroup.Post("/", bc.CreateBoard)
	boardGroup.Put("/:id", bc.UpdateBoard)
	boardGroup.Post("/:id/members", bc.AddBoardMembers)
	boardGroup.Delete("/:id/members", bc.RemoveBoardMembers)
	boardGroup.Get("/my", bc.GetMyBoardPaginate)
	boardGroup.Get("/:board_id/lists", lc.GetListsOnBoard)

	// list
	listGroup := api.Group("/lists")
	listGroup.Post("/", lc.CreateList)
	listGroup.Put("/:id", lc.UpdateList)
	listGroup.Delete("/:id", lc.DeleteList)

	// card
	cardGroup := api.Group("/cards")
	cardGroup.Post("/", cc.CreateCard)
	cardGroup.Put("/:id", cc.UpdateCard)
	cardGroup.Delete("/:id", cc.DeleteCard)

}