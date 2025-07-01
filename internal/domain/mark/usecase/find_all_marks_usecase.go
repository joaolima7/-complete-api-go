package usecase

import "github.com/joaolima7/-complete-api-go/internal/domain/mark/repository"

type FindAllMarksOutputDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type FindAllMarksUseCase struct {
	FindAllMarksRepository repository.FindAllMarksRepository
}

func NewFindAllMarksUseCase(findAllMarksRepository repository.FindAllMarksRepository) *FindAllMarksUseCase {
	return &FindAllMarksUseCase{FindAllMarksRepository: findAllMarksRepository}
}

func (u *FindAllMarksUseCase) Execute() ([]FindAllMarksOutputDTO, error) {
	marks, err := u.FindAllMarksRepository.FindAllMarks()
	if err != nil {
		return nil, err
	}

	var output []FindAllMarksOutputDTO
	for _, mark := range marks {
		output = append(output, FindAllMarksOutputDTO{
			ID:   mark.ID,
			Name: mark.Name,
		})
	}

	return output, nil
}
