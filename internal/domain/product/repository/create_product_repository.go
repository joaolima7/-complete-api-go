package repository

import "github.com/joaolima7/-complete-api-go/internal/domain/product"

type CreateProductRepository interface {
	Execute(productInput *product.Product) (*product.Product, error)
}
