package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (service *UserService) isNameRegistered(name *Username) (bool, error) {
	_, err := service.repository.FindByName(name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
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
