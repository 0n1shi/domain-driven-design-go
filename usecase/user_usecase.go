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

func (usecase *UserUsecase) FindByID(id string) (*DTOUser, error) {
	userID, err := domainUser.NewUserID(&id)
	if err != nil {
		return nil, err
	}
	user, err := usecase.service.FindByID(userID)
	if err != nil {
		return nil, err
	}
	return ToDTOUser(user), nil
}

type CreateUserInput struct {
	Name     string
	Password string
}

func (usecase *UserUsecase) Create(input *CreateUserInput) error {
	return usecase.service.Create(&domainUser.CreateUserInput{
		Name:     input.Name,
		Password: input.Password,
	})
}
