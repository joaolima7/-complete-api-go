package product

import (
	"context"
	"database/sql"

	"github.com/joaolima7/-complete-api-go/internal/domain/errs"
	"github.com/joaolima7/-complete-api-go/internal/domain/product"
	"github.com/joaolima7/-complete-api-go/internal/infra/database/db"
)

type CreateProductDatabaseImpl struct {
	queries *db.Queries
}

func NewCreateProductDatabaseImpl(database *sql.DB) *CreateProductDatabaseImpl {
	return &CreateProductDatabaseImpl{queries: db.New(database)}
}

func (cp *CreateProductDatabaseImpl) CreateProduct(productInput *product.Product) (*product.Product, error) {
	ctx := context.Background()

	params := db.CreateProductParams{
		ID:     productInput.ID,
		Name:   productInput.Name,
		Price:  productInput.Price,
		MarkID: sql.NullString{String: productInput.MarkID, Valid: productInput.MarkID != ""},
	}

	_, err := cp.queries.CreateProduct(ctx, params)
	if err != nil {
		return nil, errs.Internal("erro ao criar produto", err)
	}

	return productInput, nil
}
