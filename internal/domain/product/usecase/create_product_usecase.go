package usecase

import (
	"github.com/joaolima7/-complete-api-go/internal/domain/product"
	"github.com/joaolima7/-complete-api-go/internal/domain/product/repository"
)

type ProductInputDTO struct {
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	MarkID uint64  `json:"markID"`
}

type ProductOutputDTO struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	MarkID uint64  `json:"markID"`
}

type CreateProductUseCase struct {
	CreateProductRepository repository.CreateProductRepository
}

func NewCreateProductUseCase(createProductRepository repository.CreateProductRepository) *CreateProductUseCase {
	return &CreateProductUseCase{CreateProductRepository: createProductRepository}
}

func (u *CreateProductUseCase) Execute(input ProductInputDTO) (ProductOutputDTO, error) {
	product, err := product.NewProduct(input.Name, input.Price, input.MarkID)
	if err != nil {
		return ProductOutputDTO{}, err
	}

	product, err = u.CreateProductRepository.Execute(product)
	if err != nil {
		return ProductOutputDTO{}, err
	}

	return ProductOutputDTO{
		ID:     product.ID,
		Name:   product.Name,
		Price:  product.Price,
		MarkID: product.MarkID,
	}, nil
}
