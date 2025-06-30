package repository

import "github.com/joaolima7/-complete-api-go/internal/domain/product"

type CreateProductRepository interface {
	CreateProduct(productInput *product.Product) (*product.Product, error)
}
