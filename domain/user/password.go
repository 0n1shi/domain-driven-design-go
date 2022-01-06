package user

import (
	"unicode"

	"github.com/pkg/errors"
)

type UserPassword struct {
	val string
}

func NewUserPassword(password string) (*UserPassword, error) {
	if !validatePassword(password) {
		return nil, errors.WithStack(ErrorInvalidUserPassword)
	}
	return &UserPassword{val: password}, nil
}

func (password *UserPassword) Get() string {
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
