package seed

import (
	"log"

	"github.com/google/uuid"
	"github.com/meidhika/project-management/config"
	"github.com/meidhika/project-management/models"
	"github.com/meidhika/project-management/utils"
)

func SeedAdmin() {
	password, _ := utils.HashPassword("admin123")

	admin := models.User{
		Name:     "Super admin",
		Email:    "admin@example.com",
		Password: password,
		Role:     "admin",
		PublicID: uuid.New(),
	}
	if err := config.DB.FirstOrCreate(&admin, models.User{Email: admin.Email}).Error; err != nil {
		log.Println("Failed too seed admin", err)
	} else {
		log.Println("Admin user seeded")
	}
}