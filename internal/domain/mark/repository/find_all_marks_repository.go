package repository

import "github.com/joaolima7/-complete-api-go/internal/domain/mark"

type FindAllMarksRepository interface {
	FindAllMarks() ([]*mark.Mark, error)
}
