package usecase

import (
	"github.com/google/uuid"
	"github.com/joaolima7/-complete-api-go/internal/domain/product"
	"github.com/joaolima7/-complete-api-go/internal/domain/product/repository"
)

type ProductInputDTO struct {
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	MarkID string  `json:"markID"`
}

type ProductOutputDTO struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	MarkID string  `json:"markID"`
}

type CreateProductUseCase struct {
	CreateProductRepository repository.CreateProductRepository
}

func NewCreateProductUseCase(createProductRepository repository.CreateProductRepository) *CreateProductUseCase {
	return &CreateProductUseCase{CreateProductRepository: createProductRepository}
}

func (u *CreateProductUseCase) Execute(input ProductInputDTO) (*ProductOutputDTO, error) {
	product, err := product.NewProduct(uuid.NewString(), input.Name, input.Price, input.MarkID)
	if err != nil {
		return nil, err
	}

	product, err = u.CreateProductRepository.CreateProduct(product)
	if err != nil {
		return nil, err
	}

	return &ProductOutputDTO{
		ID:     product.ID,
		Name:   product.Name,
		Price:  product.Price,
		MarkID: product.MarkID,
	}, nil
}
