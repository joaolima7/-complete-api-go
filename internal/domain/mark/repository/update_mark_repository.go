package repository

import "github.com/joaolima7/-complete-api-go/internal/domain/mark"

type UpdateMarkRepository interface {
	UpdateMark(markInput *mark.Mark) (*mark.Mark, error)
}
