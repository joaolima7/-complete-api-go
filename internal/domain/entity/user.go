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

func New(id uint64, name string, email string, password string) (*User, *errs.AppError) {
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

func (u *User) ValidatePassword() *errs.AppError {
	if ok, err := regexp.MatchString(`^.{8,}$`, u.Password); !ok {
		return errs.DomainValidation("a senha deve conter no minimo 8 caracteres", err)
	}
	if ok, err := regexp.MatchString(`[[:lower:]]`, u.Password); !ok {
		return errs.DomainValidation("a senha deve conter pelo menos uma letra minúscula", err)
	}
	if ok, err := regexp.MatchString(`[[:upper:]]`, u.Password); !ok {
		return errs.DomainValidation("a senha deve conter pelo menos uma letra maiúscula", err)
	}
	if ok, err := regexp.MatchString(`[0-9]`, u.Password); !ok {
		return errs.DomainValidation("a senha deve conter pelo menos um número", err)
	}
	if ok, err := regexp.MatchString(`[!@#\$%\^&\*\(\)_\+\-=\[\]{};':"\\|,.<>\/?]`, u.Password); !ok {
		return errs.DomainValidation("a senha deve conter pelo menos um caractere especial", err)
	}

	return nil
}

func (u *User) ValidateEmail() *errs.AppError {
	if ok, err := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, u.Email); !ok {
		return errs.DomainValidation("email inválido", err)
	}
	return nil
}

func (u *User) ValidateName() *errs.AppError {
	if ok, err := regexp.MatchString(`^[a-zA-Z\s]+$`, u.Name); !ok {
		return errs.DomainValidation("nome inválido, deve conter apenas letras e espaços", err)
	}
	return nil
}
