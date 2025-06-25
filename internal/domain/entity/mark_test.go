package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMark(t *testing.T) {
	mark, err := NewMark(1, "Brand Name")

	assert.NoError(t, err)
	assert.NotNil(t, mark)
	assert.Equal(t, uint64(1), mark.ID)
	assert.Equal(t, "Brand Name", mark.Name)
}

func TestNewMarkInvalidName(t *testing.T) {
	mark, err := NewMark(1, "")
	assert.Error(t, err)
	assert.Nil(t, mark)

	mark, err = NewMark(1, " ")
	assert.Error(t, err)
	assert.Nil(t, mark)

	mark, err = NewMark(1, "\t")
	assert.Error(t, err)
	assert.Nil(t, mark)

	mark, err = NewMark(1, "\n")
	assert.Error(t, err)
	assert.Nil(t, mark)
}
