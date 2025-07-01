package repository

type DeleteMarkByIDRepository interface {
	DeleteMark(markID string) error
}
