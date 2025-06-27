package product

import (
	"database/sql"

	"github.com/joaolima7/-complete-api-go/internal/domain/product"
)

type CreateProductDatabaseImpl struct {
	DB *sql.DB
}

func NewCreateProductDatabaseImpl(db *sql.DB) *CreateProductDatabaseImpl {
	return &CreateProductDatabaseImpl{DB: db}
}

func (cp *CreateProductDatabaseImpl) Execute(product *product.Product) (*product.Product, error)
