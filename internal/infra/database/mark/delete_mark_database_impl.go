package mark

import (
	"context"

	"github.com/joaolima7/-complete-api-go/internal/domain/errs"
	"github.com/joaolima7/-complete-api-go/internal/infra/database/db"
)

type DeleteMarkDatabaseImpl struct {
	queries *db.Queries
}

func NewDeleteMarkDatabaseImpl(queries *db.Queries) *DeleteMarkDatabaseImpl {
	return &DeleteMarkDatabaseImpl{
		queries: queries,
	}
}

func (d *DeleteMarkDatabaseImpl) DeleteMark(markID string) error {
	ctx := context.Background()

	err := d.queries.DeleteMark(ctx, markID)
	if err != nil {
		return errs.Internal("erro ao deletar marca no banco de dados", err)
	}

	return nil
}
