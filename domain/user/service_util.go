package user

import (
	"golang.org/x/crypto/bcrypt"
)

func (service *UserService) isNameRegistered(name *Username) bool {
	_, err := service.repository.FindByName(name)
	if err != nil {
		return false
	}
	return true
}

func hashPassword(password string) (*UserPassword, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return nil, err
	}
	hashedPassword, err := NewUserPassword(string(bytes))
	if err != nil {
		return nil, err
	}
	return hashedPassword, nil
}
