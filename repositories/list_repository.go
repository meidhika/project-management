package repositories

import (
	"github.com/meidhika/project-management/config"
	"github.com/meidhika/project-management/models"
)

type ListRepository interface {
}

type listRepository struct{}

func NewListRepository() ListRepository {
	return &listRepository{}
}

func (r *listRepository) Create(list *models.List) error {
	return config.DB.Create(&list).Error
}