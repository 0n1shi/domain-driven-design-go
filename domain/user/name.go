package user

import "github.com/pkg/errors"

type Username struct {
	Name string
}

func NewUsername(name string) (*Username, error) {
	if len(name) < 3 || len(name) > 20 {
		return nil, errors.New("name is invalid")
	}
	return &Username{Name: name}, nil
}

func (name *Username) Get() string {
	return name.Get()
}
