package repository

import "github.com/joaolima7/-complete-api-go/internal/domain/mark"

type FindMarkByIDRepository interface {
	FindMarkByID(markID string) (*mark.Mark, error)
}
