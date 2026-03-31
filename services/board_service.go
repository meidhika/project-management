package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/meidhika/project-management/models"
	"github.com/meidhika/project-management/repositories"
)

type BoardService interface {
	Create(board *models.Board) error
	Update(board *models.Board) error
	GetByPublicID(publicID string) (*models.Board, error)
	AddMember(boardPublicID string, userPublicIDS []string) error
	RemoveMember(boardPublicID string, userPublicIDs []string) error
	GetAllByUserPaginate(userID, filter, sort string, limit, offset int) ([]models.Board, int64, error)
}

type boardService struct {
	boardRepo repositories.BoardRepository
	userRepo repositories.UserRepository
	boardMemberRepo repositories.BoardMemberRepository
}

func NewBoardService(
	boardRepo repositories.BoardRepository, 
	userRepo repositories.UserRepository,
	boardMemberRepo repositories.BoardMemberRepository,
) BoardService {
	return &boardService{boardRepo, userRepo, boardMemberRepo}
}

func (s *boardService) Create(board *models.Board) error {
	user, err := s.userRepo.FindByPublicID(board.OwnerPublicID.String())

	if err != nil {
		return errors.New("owner not found")
	}
	board.PublicID = uuid.New()
	board.OwnerID = user.InternalID

	return s.boardRepo.Create(board)
}


func (s *boardService) Update(board *models.Board) error {
	return s.boardRepo.Update(board)
}

func (s *boardService) GetByPublicID(publicID string) (*models.Board, error) {
	return s.boardRepo.FindByPublicID(publicID)
}

func (s *boardService) AddMember(boardPublicID string, userPublicIDS []string) error {
	board, err := s.boardRepo.FindByPublicID(boardPublicID)
	if err != nil {
		return errors.New("Board Not Found")
	}
	var userInternalIDs []uint
	for _, userPublicID := range userPublicIDS {
		user, err := s.userRepo.FindByPublicID(userPublicID)

		if err != nil {
			return errors.New("user not found: " + userPublicID)
		}
		userInternalIDs = append(userInternalIDs, uint(user.InternalID))
	}
	// check keanggotaan user di board
	existingMembers, err := s.boardMemberRepo.GetMembers(string(board.PublicID.String()))
	if err != nil {
		return err
	}

	// buat map untuk mengecek anggota yang sudah ada
	memberMap := make(map[uint]bool)
	for _, member := range existingMembers {
		memberMap[uint(member.InternalID)] = true
	}

	var newMemberIDs []uint
	for _, userID := range userInternalIDs {
		if !memberMap[userID] {
			newMemberIDs = append(newMemberIDs, userID)
		}
	}

	if len(newMemberIDs) == 0 {
		return nil
	}

	return s.boardRepo.AddMember(uint(board.InternalID), newMemberIDs)
}

func (s *boardService) RemoveMember(boardPublicID string, userPublicIDs []string) error {
	board, err := s.boardRepo.FindByPublicID(boardPublicID)
	if err != nil {
		return errors.New("Board Not Found")
	}

	// validasi userPublicIDs dan konversi ke userInternalIDs
	var userInternalIDs []uint
	for _, userPublicID := range userPublicIDs {
		user, err := s.userRepo.FindByPublicID(userPublicID)
		if err != nil {
			return errors.New("User Not Found: " + userPublicID)
		}
		userInternalIDs = append(userInternalIDs, uint(user.InternalID))
	}

	// validasi keanggotaan user di board
	existingMembers, err := s.boardMemberRepo.GetMembers(string(board.PublicID.String()))
	if err != nil {
		return err
	}

	// buat map untuk mengecek anggota yang sudah ada
	memberMap := make(map[uint]bool)
	for _, member := range existingMembers {
		memberMap[uint(member.InternalID)] = true
	}

	var removedMemberIDs []uint
	for _, userID := range userInternalIDs {
		if memberMap[userID] {
			removedMemberIDs = append(removedMemberIDs, userID)
		}
	}

	return s.boardRepo.RemoveMember(uint(board.InternalID), removedMemberIDs)
}	

func (s *boardService) GetAllByUserPaginate(userID, filter, sort string, limit, offset int) ([]models.Board, int64, error) {
	return s.boardRepo.FindAllByUserPaginate(userID, filter, sort, limit, offset)
}