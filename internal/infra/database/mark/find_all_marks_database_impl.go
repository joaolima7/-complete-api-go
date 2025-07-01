package mark

import (
	"context"

	"github.com/joaolima7/-complete-api-go/internal/domain/mark"
	"github.com/joaolima7/-complete-api-go/internal/infra/database/db"
)

type FindAllMarksDatabaseImpl struct {
	queries *db.Queries
}

func NewFindAllMarksDatabaseImpl(queries *db.Queries) *FindAllMarksDatabaseImpl {
	return &FindAllMarksDatabaseImpl{
		queries: queries,
	}
}

func (f *FindAllMarksDatabaseImpl) FindAllMarks() ([]*mark.Mark, error) {
	ctx := context.Background()

	marks, err := f.queries.GetAllMarks(ctx)
	if err != nil {
		return nil, err
	}

	var result []*mark.Mark
	for _, m := range marks {
		result = append(result, &mark.Mark{
			ID:   m.ID,
			Name: m.Name,
		})
	}

	return result, nil
}
