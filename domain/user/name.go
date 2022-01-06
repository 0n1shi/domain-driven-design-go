package user

import "github.com/pkg/errors"

type Username struct {
	val string
}

func NewUsername(name string) (*Username, error) {
	if len(name) < 3 || len(name) > 20 {
		return nil, errors.WithStack(ErrorInvalidUsername)
	}
	return &Username{val: name}, nil
}

func (name *Username) Get() string {
	return name.val
}
