package user

import "github.com/pkg/errors"

type username struct {
	val string // must be unique value
}

type Username interface {
	Get() string
}

var _ Username = (*username)(nil)

func NewUsername(name string) (*username, error) {
	if len(name) < 3 || len(name) > 20 {
		return nil, errors.WithStack(ErrorInvalidUsername)
	}
	return &username{val: name}, nil
}

func (name *username) Get() string {
	return name.val
}
