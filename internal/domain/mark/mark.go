package mark

import (
	"strings"

	"github.com/joaolima7/-complete-api-go/internal/domain/errs"
)

type Mark struct {
	ID   uint64
	Name string
}

func NewMark(id uint64, name string) (*Mark, error) {
	trimmedName := strings.TrimSpace(name)
	if trimmedName == "" {
		return nil, errs.DomainValidation("o nome da marca não pode ser vazio", nil)
	}

	return &Mark{
		ID:   id,
		Name: trimmedName,
	}, nil
}
