package mark

import (
	"context"

	"github.com/joaolima7/-complete-api-go/internal/domain/errs"
	"github.com/joaolima7/-complete-api-go/internal/domain/mark"
	"github.com/joaolima7/-complete-api-go/internal/infra/database/db"
)

type UpdateMarkDatabaseImpl struct {
	queries *db.Queries
}

func NewUpdateMarkDatabaseImpl(queries *db.Queries) *UpdateMarkDatabaseImpl {
	return &UpdateMarkDatabaseImpl{queries: queries}
}

func (um *UpdateMarkDatabaseImpl) UpdateMark(markInput *mark.Mark) (*mark.Mark, error) {
	ctx := context.Background()

	params := db.UpdateMarkParams{
		Name: markInput.Name,
		ID:   markInput.ID,
	}

	_, err := um.queries.UpdateMark(ctx, params)
	if err != nil {
		return nil, errs.Internal("erro na atualização da marca", err)
	}

	return markInput, nil
}
