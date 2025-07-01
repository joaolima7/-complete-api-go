package usecase

import (
	"github.com/google/uuid"
	"github.com/joaolima7/-complete-api-go/internal/domain/mark"
	"github.com/joaolima7/-complete-api-go/internal/domain/mark/repository"
)

type MarkInputDTO struct {
	Name string `json:"name"`
}

type MarkOutputDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CreateMarkUseCase struct {
	CreateMarkRepository repository.CreateMarkRepository
}

func NewCreateMarkUseCase(createMarkRepository repository.CreateMarkRepository) *CreateMarkUseCase {
	return &CreateMarkUseCase{CreateMarkRepository: createMarkRepository}
}

func (u *CreateMarkUseCase) Execute(input MarkInputDTO) (*MarkOutputDTO, error) {
	mark, err := mark.NewMark(uuid.NewString(), input.Name)
	if err != nil {
		return nil, err
	}
	if _, err := u.CreateMarkRepository.CreateMark(mark); err != nil {
		return nil, err
	}
	return &MarkOutputDTO{
		ID:   mark.ID,
		Name: mark.Name,
	}, nil
}
