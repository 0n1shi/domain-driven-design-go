package user

import (
	"errors"

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
