package mark

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewMark(t *testing.T) {
	mark, err := NewMark(uuid.NewString(), "Brand Name")

	assert.NoError(t, err)
	assert.NotNil(t, mark)
	assert.NotNil(t, mark.ID)
	assert.Equal(t, "Brand Name", mark.Name)
}

func TestNewMarkInvalidName(t *testing.T) {
	mark, err := NewMark(uuid.NewString(), "")
	assert.Error(t, err)
	assert.Nil(t, mark)

	mark, err = NewMark(uuid.NewString(), " ")
	assert.Error(t, err)
	assert.Nil(t, mark)

	mark, err = NewMark(uuid.NewString(), "\t")
	assert.Error(t, err)
	assert.Nil(t, mark)

	mark, err = NewMark(uuid.NewString(), "\n")
	assert.Error(t, err)
	assert.Nil(t, mark)
}
