package entity

import (
	"regexp"

	"github.com/joaolima7/-complete-api-go/internal/domain/errs"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uint64
	Name     string
	Email    string
	Password string
}

func NewUser(id uint64, name string, email string, password string) (*User, error) {
	user := User{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: password,
	}

	if err := user.ValidatePassword(); err != nil {
		return nil, err
	}

	if err := user.ValidateEmail(); err != nil {
		return nil, err
	}

	if err := user.ValidateName(); err != nil {
		return nil, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errs.Internal("falha na criptografia de senha", err)
	}

	user.Password = string(hash)

	return &user, nil
}

func (u *User) ValidatePassword() error {
	patterns := []struct {
		regex   string
		message string
	}{
		{`^.{8,}$`, "a senha deve conter no mínimo 8 caracteres"},
		{`[[:lower:]]`, "a senha deve conter pelo menos uma letra minúscula"},
		{`[[:upper:]]`, "a senha deve conter pelo menos uma letra maiúscula"},
		{`[0-9]`, "a senha deve conter pelo menos um número"},
		{`[!@#\$%\^&\*\(\)_\+\-=\[\]{};':"\\|,.<>\/?]`, "a senha deve conter pelo menos um caractere especial"},
	}

	for _, p := range patterns {
		ok, err := regexp.MatchString(p.regex, u.Password)
		if err != nil {
			return errs.Internal("erro ao validar expressão regular", err)
		}
		if !ok {
			return errs.DomainValidation(p.message, nil)
		}
	}

	return nil
}

func (u *User) ValidateEmail() error {
	if ok, err := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, u.Email); !ok {
		return errs.DomainValidation("email inválido", err)
	}
	return nil
}

func (u *User) ValidateName() error {
	if ok, err := regexp.MatchString(`^[a-zA-Z\s]+$`, u.Name); !ok {
		return errs.DomainValidation("nome inválido, deve conter apenas letras e espaços", err)
	}
	return nil
}
