package repository

import "github.com/joaolima7/-complete-api-go/internal/domain/product"

type CreateProductRepository interface {
	Execute(product *product.Product) (*product.Product, error)
}
