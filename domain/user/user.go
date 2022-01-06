package user

import "github.com/pkg/errors"

type User struct {
	id       UserID
	name     Username
	password UserPassword
}

func NewUser(name string, password string, id *string) (*User, error) {
	userID, err := NewUserID(id)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	username, err := NewUsername(name)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	userPassword, err := NewUserPassword(password)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &User{
		id:       *userID,
		name:     *username,
		password: *userPassword,
	}, nil
}

func (user *User) GetID() UserID {
	return user.id
}

func (user *User) GetName() Username {
	return user.name
}

func (user *User) GetPassword() UserPassword {
	return user.password
}
