package usecase

import (
	domainUser "github.com/0n1shi/domain-driven-design/domain/user"
	"github.com/pkg/errors"
)

type UserUsecase struct {
	service *domainUser.UserService
}

func NewUserUsecase(service *domainUser.UserService) *UserUsecase {
	return &UserUsecase{service: service}
}

func (usecase *UserUsecase) FindAll() ([]*DTOUser, error) {
	users, err := usecase.service.FindAll()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ToDTOUsers(users), nil
}
