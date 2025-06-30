package repository

import "github.com/joaolima7/-complete-api-go/internal/domain/mark"

type CreateMarkRepository interface {
	CreateMark(markInput *mark.Mark) (*mark.Mark, error)
}
