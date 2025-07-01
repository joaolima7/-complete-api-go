package mark

import (
	"context"

	"github.com/joaolima7/-complete-api-go/internal/domain/errs"
	"github.com/joaolima7/-complete-api-go/internal/domain/mark"
	"github.com/joaolima7/-complete-api-go/internal/infra/database/db"
)

type CreateMarkDatabaseImpl struct {
	queries *db.Queries
}

func NewCreateMarkDatabaseImpl(queries *db.Queries) *CreateMarkDatabaseImpl {
	return &CreateMarkDatabaseImpl{queries: queries}
}

func (cm *CreateMarkDatabaseImpl) CreateMark(markInput *mark.Mark) (*mark.Mark, error) {
	ctx := context.Background()

	params := db.CreateMarkParams{
		ID:   markInput.ID,
		Name: markInput.Name,
	}

	_, err := cm.queries.CreateMark(ctx, params)
	if err != nil {
		return nil, errs.Internal("erro na criação da marca", err)
	}

	return markInput, nil
}
