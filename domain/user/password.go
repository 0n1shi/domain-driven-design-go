package user

import (
	"errors"
)

type UserPassword struct {
	val string
}

func NewUserPassword(password string) (*UserPassword, error) {
	if !ValidatePassword(password) {
		return nil, errors.New("at least 1 upper case, 1 lower case and numeric must be EMBEDDED somewhere in the middle of the password.")
	}
	return &UserPassword{val: password}, nil
}

func (password *UserPassword) Get() string {
	return password.val
}
