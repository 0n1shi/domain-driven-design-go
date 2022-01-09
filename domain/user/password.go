package user

import (
	"unicode"

	"github.com/pkg/errors"
)

type userPassword struct {
	val string
}

type UserPassword interface {
	Get() string
}

var _ UserPassword = (*userPassword)(nil)

func NewUserPassword(password string) (UserPassword, error) {
	if !validatePassword(password) {
		return nil, errors.WithStack(ErrorInvalidUserPassword)
	}
	return &userPassword{val: password}, nil
}

func (password *userPassword) Get() string {
	return password.val
}

func validatePassword(password string) bool {
	length := false
	upper := false
	lower := false
	numeric := false

	if len(password) >= 8 {
		length = true
	}
	for _, c := range password {
		switch {
		case unicode.IsNumber(c):
			numeric = true
		case unicode.IsUpper(c):
			upper = true
		case unicode.IsLower(c):
			lower = true
		}
	}

	return length && upper && lower && numeric
}
