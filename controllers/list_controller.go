package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/meidhika/project-management/models"
	"github.com/meidhika/project-management/services"
	"github.com/meidhika/project-management/utils"
)

type ListController struct {
	service services.ListService
}

func NewListController(s services.ListService) *ListController {
	return &ListController{service: s}
}

func (c *ListController) CreateList(ctx *fiber.Ctx) error{
	list := new(models.List)
	if err := ctx.BodyParser(list); err != nil {
		return utils.BadRequest(ctx, "Gagal Membaca Request", err.Error())
	}

	if err := c.service.Create(list); err != nil {
		return utils.BadRequest(ctx, "Gagal Membuat List", err.Error())
	}

	return utils.Success(ctx, "List Berhasil Dibuat", list)
}

func (c *ListController) UpdateList(ctx *fiber.Ctx) error{
	publicID := ctx.Params("id")
	list := new(models.List)
	if err := ctx.BodyParser(list); err != nil {
		return utils.BadRequest(ctx, "Gagal Parsing Data", err.Error())
	}

	if _, err := uuid.Parse(publicID); err != nil { 
		return utils.BadRequest(ctx, "ID Tidak Valid", err.Error()) 
	}

	existingList, err := c.service.GetByPublicID(publicID)
	if err != nil {
		return utils.NotFound(ctx, "List Tidak Ditemukan", err.Error())
	}

	list.InternalID = existingList.InternalID
	list.PublicID = existingList.PublicID

	if err := c.service.Update(list); err != nil {
		return utils.BadRequest(ctx, "Gagal Mengupdate List", err.Error())
	}

	updatedList, err := c.service.GetByPublicID(publicID)
	if err != nil {
		return utils.NotFound(ctx, "List Tidak Ditemukan", err.Error())
	}

	return utils.Success(ctx, "List Berhasil Diupdate", updatedList)
}