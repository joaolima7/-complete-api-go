package product

import (
	"strings"

	"github.com/joaolima7/-complete-api-go/internal/domain/errs"
)

type Product struct {
	ID     string
	Name   string
	Price  float64
	MarkID string
}

func NewProduct(id string, name string, price float64, markID string) (*Product, error) {
	trimmedName := strings.TrimSpace(name)
	if trimmedName == "" {
		return nil, errs.DomainValidation("o nome do produto não pode ser vazio", nil)
	}
	if price <= 0 {
		return nil, errs.DomainValidation("o preço do produto deve ser maior que zero", nil)
	}

	product := &Product{
		ID:     id,
		Name:   trimmedName,
		Price:  price,
		MarkID: markID,
	}

	return product, nil
}
