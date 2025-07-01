package mark

import (
	"context"
	"database/sql"

	"github.com/joaolima7/-complete-api-go/internal/domain/errs"
	"github.com/joaolima7/-complete-api-go/internal/domain/mark"
	"github.com/joaolima7/-complete-api-go/internal/infra/database/db"
)

type FindMarkByIDDatabaseImpl struct {
	queries *db.Queries
}

func NewFindMarkByIDDatabaseImpl(queries *db.Queries) *FindMarkByIDDatabaseImpl {
	return &FindMarkByIDDatabaseImpl{
		queries: queries,
	}
}

func (f *FindMarkByIDDatabaseImpl) FindMarkByID(markID string) (*mark.Mark, error) {
	ctx := context.Background()

	markData, err := f.queries.GetMarkById(ctx, markID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NotFound("erro na busca da marca", err)

		}
		return nil, errs.Internal("erro ao buscar marca no banco de dados", err)
	}

	return &mark.Mark{
		ID:   markData.ID,
		Name: markData.Name,
	}, nil
}
