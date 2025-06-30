package usecase

import (
	"time"

	"github.com/joaolima7/-complete-api-go/internal/domain/mark/repository"
)

type MarkInputDTO struct {
	Name string `json:"name"`
}

type MarkOutputDTO struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type CreateMarkUseCase struct {
	CreateMarkRepository repository.CreateMarkRepository
}

func NewCreateMarkUseCase(createMarkRepository repository.CreateMarkRepository) *CreateMarkUseCase {
	return &CreateMarkUseCase{CreateMarkRepository: createMarkRepository}
}

func (u *CreateMarkUseCase) Execute(input MarkInputDTO) (*MarkOutputDTO, error)
