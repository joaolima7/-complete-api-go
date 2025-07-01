package usecase

import (
	"errors"

	"github.com/joaolima7/-complete-api-go/internal/domain/mark"
	"github.com/joaolima7/-complete-api-go/internal/domain/mark/repository"
)

type UpdateMarkInputDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UpdateMarkOutputDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UpdateMarkUseCase struct {
	UpdateMarkRepository   repository.UpdateMarkRepository
	FindMarkByIDRepository repository.FindMarkByIDRepository
}

func NewUpdateMarkUseCase(updateMarkRepository repository.UpdateMarkRepository, findMarkByIDRepository repository.FindMarkByIDRepository) *UpdateMarkUseCase {
	return &UpdateMarkUseCase{
		UpdateMarkRepository:   updateMarkRepository,
		FindMarkByIDRepository: findMarkByIDRepository,
	}
}

func (u *UpdateMarkUseCase) Execute(input UpdateMarkInputDTO) (*UpdateMarkOutputDTO, error) {
	if input.ID == "" {
		return nil, errors.New("ID da marca é obrigatório para atualização")
	}

	_, err := u.FindMarkByIDRepository.FindMarkByID(input.ID)
	if err != nil {
		return nil, err
	}

	mark, err := mark.NewMark(input.ID, input.Name)
	if err != nil {
		return nil, err
	}

	if mark, err = u.UpdateMarkRepository.UpdateMark(mark); err != nil {
		return nil, err
	}

	return &UpdateMarkOutputDTO{
		ID:   mark.ID,
		Name: mark.Name,
	}, nil
}
