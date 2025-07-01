package usecase

import (
	"errors"

	"github.com/joaolima7/-complete-api-go/internal/domain/mark/repository"
)

type DeleteMarkInputDTO struct {
	ID string `json:"id"`
}

type DeleteMarkUseCase struct {
	DeleteMarkRepository   repository.DeleteMarkByIDRepository
	FindMarkByIDRepository repository.FindMarkByIDRepository
}

func NewDeleteMarkUseCase(deleteMarkRepository repository.DeleteMarkByIDRepository, findMarkByIDRepository repository.FindMarkByIDRepository) *DeleteMarkUseCase {
	return &DeleteMarkUseCase{DeleteMarkRepository: deleteMarkRepository, FindMarkByIDRepository: findMarkByIDRepository}
}

func (u *DeleteMarkUseCase) Execute(input DeleteMarkInputDTO) error {
	if input.ID == "" {
		return errors.New("ID da marca é obrigatório para exclusão")
	}

	if _, err := u.FindMarkByIDRepository.FindMarkByID(input.ID); err != nil {
		return err
	}

	if err := u.DeleteMarkRepository.DeleteMark(input.ID); err != nil {
		return err
	}
	return nil
}
