package user

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser(uuid.NewString(), "Jhon Smith", "test@mail.com", "ATh09!k3")

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "Jhon Smith", user.Name)
	assert.Equal(t, "test@mail.com", user.Email)
	assert.NotEmpty(t, user.Password)
	assert.NotEqual(t, "ATh09!k3", user.Password)
}

func TestNewUserInvalidPassword(t *testing.T) {
	user, err := NewUser(uuid.NewString(), "Jhon Smith", "test@mail.com", "12345678")
	assert.Error(t, err)
	assert.Nil(t, user)

	user, err = NewUser(uuid.NewString(), "Jhon Smith", "test@gmail.com", "abcdefgh")
	assert.Error(t, err)
	assert.Nil(t, user)

	user, err = NewUser(uuid.NewString(), "Jhon Smith", "test@mail.com", "ABCDEFGH")
	assert.Error(t, err)
	assert.Nil(t, user)

	user, err = NewUser(uuid.NewString(), "Jhon Smith", "test@mail.com", "ATh09k3")
	assert.Error(t, err)
	assert.Nil(t, user)
}

func TestNewUserInvalidEmail(t *testing.T) {
	user, err := NewUser(uuid.NewString(), "Jhon Smith", "invalid-email", "ATh09!k3")
	assert.Error(t, err)
	assert.Nil(t, user)

	user, err = NewUser(uuid.NewString(), "Jhon Smith", "test@", "ATh09!k3")
	assert.Error(t, err)
	assert.Nil(t, user)

	user, err = NewUser(uuid.NewString(), "Jhon Smith", "@mail.com", "ATh09!k3")
	assert.Error(t, err)
	assert.Nil(t, user)

	user, err = NewUser(uuid.NewString(), "Jhon Smith", "test.mail.com", "ATh09!k3")
	assert.Error(t, err)
	assert.Nil(t, user)
}

func TestNewUserInvalidName(t *testing.T) {
	user, err := NewUser(uuid.NewString(), "Jhon123", "test@mail.com", "ATh09!k3")
	assert.Error(t, err)
	assert.Nil(t, user)

	user, err = NewUser(uuid.NewString(), "Jhon@Smith", "test@mail.com", "ATh09!k3")
	assert.Error(t, err)
	assert.Nil(t, user)

	user, err = NewUser(uuid.NewString(), "", "test@mail.com", "ATh09!k3")
	assert.Error(t, err)
	assert.Nil(t, user)

	user, err = NewUser(uuid.NewString(), "123", "test@mail.com", "ATh09!k3")
	assert.Error(t, err)
	assert.Nil(t, user)
}
