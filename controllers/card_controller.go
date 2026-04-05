package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/meidhika/project-management/models"
	"github.com/meidhika/project-management/services"
	"github.com/meidhika/project-management/utils"
)

type CardController struct {
	service services.CardService
}

func NewCardController(s services.CardService) *CardController {
	return &CardController{service: s}
}

func (c *CardController) CreateCard(ctx *fiber.Ctx) error {
	type CreateCardRequest struct {
		ListPublicID string `json:"list_id"`
		Title string `json:"title"`
		Description string `json:"description"`
		DueDate time.Time `json:"due_date"`
		Position int `json:"position"`	
	}

	var req CreateCardRequest
	if err := ctx.BodyParser(&req); err != nil {
		return utils.BadRequest(ctx, "Gagal Mengambil Data", err.Error())
	}

	card := &models.Card{
		Title : req.Title,
		Description : req.Description,
		DueDate : &req.DueDate,
		Position : req.Position,
	}

	if err := c.service.Create(card, req.ListPublicID); err != nil {
		return utils.InternalServerError(ctx, "Gagal Membuat Card", err.Error())
	}

	return utils.Success(ctx, "Card berhasil dibuat", card)
}

func (c *CardController) UpdateCard(ctx *fiber.Ctx) error{
	publicID := ctx.Params("id")
	type updateCardRequest struct {
		ListPublicID string `json:"list_id"`
		Title string `json:"title"`
		Description string `json:"description"`
		DueDate *time.Time `json:"due_date"`
		Position int `json:"position"`	
	}

	var req updateCardRequest
	if err := ctx.BodyParser(&req); err != nil {
		return utils.BadRequest(ctx, "Gagal Parsing Data", err.Error())
	}

	if _, err := uuid.Parse(publicID); err != nil {
		return utils.BadRequest(ctx, "ID Tidak Valid", err.Error())
	}

	card := &models.Card{
		Title : req.Title,
		Description : req.Description,
		DueDate : req.DueDate,
		Position : req.Position,
		PublicID: uuid.MustParse(publicID),
	}

	if err := c.service.Update(card, req.ListPublicID); err != nil {
		return utils.InternalServerError(ctx, "Gagal update data", err.Error())
	}

	return utils.Success(ctx, "Card Berhasil diperbaharui", card)
}


func (c *CardController) DeleteCard(ctx *fiber.Ctx) error {
	publicID := ctx.Params("id")
	if _, err := uuid.Parse(publicID); err != nil {
		return utils.BadRequest(ctx, "ID Tidak Valid", err.Error())
	}

	card, err := c.service.GetByPublicID(publicID)
	if err != nil {
		return utils.NotFound(ctx, "Card Tidak Ditemukan", err.Error())
	}

	if err := c.service.Delete(uint(card.InternalID)); err != nil {
		return utils.BadRequest(ctx, "Gagal Menghapus Card", err.Error())
	}

	return utils.Success(ctx, "Card Berhasil Dihapus", publicID)
}

func (c *CardController) GetCardDetails(ctx *fiber.Ctx) error {
	cardPublicID := ctx.Params("id")

	card, err := c.service.GetByPublicID(cardPublicID)
	if err != nil {
		return utils.InternalServerError(ctx, "Error Saat Mengambil Data", err.Error())
	}

	if card == nil{
		return utils.NotFound(ctx, "Card Tidak Ditemukan", err.Error())
	}
	return utils.Success(ctx, "Data Berhasil diambil", card)
}