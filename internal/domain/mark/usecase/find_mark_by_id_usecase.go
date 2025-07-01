package usecase

import (
	"github.com/joaolima7/-complete-api-go/internal/domain/errs"
	"github.com/joaolima7/-complete-api-go/internal/domain/mark/repository"
)

type FindMarkByIDInputDTO struct {
	ID string `json:"id"`
}

type FindMarkByIDOutputDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type FindMarkByIDUseCase struct {
	FindMarkByIDRepository repository.FindMarkByIDRepository
}

func NewFindMarkByIDUseCase(findMarkByIDRepository repository.FindMarkByIDRepository) *FindMarkByIDUseCase {
	return &FindMarkByIDUseCase{FindMarkByIDRepository: findMarkByIDRepository}
}

func (u *FindMarkByIDUseCase) Execute(input FindMarkByIDInputDTO) (*FindMarkByIDOutputDTO, error) {
	if input.ID == "" {
		return nil, errs.NotFound("ID da marca é obrigatório para a busca", nil)
	}

	mark, err := u.FindMarkByIDRepository.FindMarkByID(input.ID)
	if err != nil {
		return nil, err
	}

	output := &FindMarkByIDOutputDTO{
		ID:   mark.ID,
		Name: mark.Name,
	}

	return output, nil
}
