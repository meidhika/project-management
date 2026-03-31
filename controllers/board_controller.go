package controllers

import (
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/meidhika/project-management/models"
	"github.com/meidhika/project-management/services"
	"github.com/meidhika/project-management/utils"
)

type BoardController struct {
	service services.BoardService
}

func NewBoardController(s services.BoardService) *BoardController {
	return &BoardController{service: s,}
}

func (c *BoardController) CreateBoard(ctx *fiber.Ctx) error {
	var userID uuid.UUID
	var err error
	board := new(models.Board)
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)


	if err := ctx.BodyParser(board); err != nil {
		return utils.BadRequest(ctx, "Gagal membaca request", err.Error())
	}

	userID, err = uuid.Parse(claims["pub_id"].(string))
	if err != nil {
		return utils.BadRequest(ctx, "Gagal memparsing user ID", err.Error())
	}

	board.OwnerPublicID = userID

	if err:= c.service.Create(board); err != nil {
		return utils.BadRequest(ctx, "Gagal menyimpan data", err.Error())
	}

	return utils.Success(ctx, "Board berhasil dibuat", board)
}

func (c *BoardController) UpdateBoard(ctx *fiber.Ctx) error {
	publicID := ctx.Params("id")
	board := new(models.Board)

	if err := ctx.BodyParser(board); err != nil {
		return utils.BadRequest(ctx, "Gagal Parsing Data", err.Error())
	}

	if _, err := uuid.Parse(publicID); err != nil {
		return utils.BadRequest(ctx, "ID Tidak Valid", err.Error())
	}

	existingBoard, err := c.service.GetByPublicID(publicID)
	if err != nil {
		return utils.NotFound(ctx, "Board Tidak Ditemukan", err.Error())
	}

	board.InternalID = existingBoard.InternalID
	board.PublicID = existingBoard.PublicID
	board.OwnerID = existingBoard.OwnerID
	board.OwnerPublicID = existingBoard.OwnerPublicID
	board.CreatedAt = existingBoard.CreatedAt

	if err := c.service.Update(board); err != nil {
		return utils.BadRequest(ctx, "Gagal update board", err.Error())
	}

	return utils.Success(ctx, "Board berhasil diperbarui", board)
}

func (c *BoardController) AddBoardMembers(ctx *fiber.Ctx) error {
	publicID := ctx.Params("id")
	var userIDs []string

	if err := ctx.BodyParser(&userIDs); err != nil {
		return utils.BadRequest(ctx, "Gagal Parsing Data", err.Error())
	}

	if err := c.service.AddMember(publicID, userIDs); err != nil {
		return utils.BadRequest(ctx, "Gagal Menambahkan Anggota", err.Error())
	}

	return utils.Success(ctx, "Anggota berhasil ditambahkan", nil)
}

func (c *BoardController) RemoveBoardMembers(ctx *fiber.Ctx) error {
	publicID := ctx.Params("id")
	var userIDs []string

	if err := ctx.BodyParser(&userIDs); err != nil {
		return utils.BadRequest(ctx, "Gagal Parsing Data", err.Error())
	}

	if err := c.service.RemoveMember(publicID, userIDs); err != nil {
		return utils.BadRequest(ctx, "Gagal Menghapus Anggota", err.Error())
	}

	return utils.Success(ctx, "Anggota berhasil dihapus", nil)
}

func (c *BoardController) GetMyBoardPaginate(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["pub_id"].(string)

	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Query("limit", "10"))
	offset := (page - 1) * limit
	filter := ctx.Query("filter", "")
	sort := ctx.Query("sort", "")

	boards, total, err := c.service.GetAllByUserPaginate(userID, filter, sort, limit, offset)
	if err != nil {
		return utils.InternalServerError(ctx, "Gagal mengambil data board", err.Error())
	}

	meta := utils.PaginationMeta{
		Page: page,
		Limit: limit,
		Total: int(total),
		TotalPage: int(math.Ceil(float64(total) / float64(limit))),
		Filter: filter,
		Sort: sort,
	}

	return utils.SuccessPagination(ctx, "Berhasil mengambil data board", boards, meta)
}