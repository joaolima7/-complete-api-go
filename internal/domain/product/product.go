package product

import (
	"strings"

	"github.com/google/uuid"
	"github.com/joaolima7/-complete-api-go/internal/domain/errs"
)

type Product struct {
	ID     string
	Name   string
	Price  float64
	MarkID uint64
}

func NewProduct(name string, price float64, markID uint64) (*Product, error) {
	trimmedName := strings.TrimSpace(name)
	if trimmedName == "" {
		return nil, errs.DomainValidation("o nome do produto não pode ser vazio", nil)
	}
	if price <= 0 {
		return nil, errs.DomainValidation("o preço do produto deve ser maior que zero", nil)
	}
	if markID == 0 {
		return nil, errs.DomainValidation("o ID da marca não pode ser zero", nil)
	}

	product := &Product{
		ID:     uuid.NewString(),
		Name:   trimmedName,
		Price:  price,
		MarkID: markID,
	}

	return product, nil
}
