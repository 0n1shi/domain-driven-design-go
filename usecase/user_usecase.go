package usecase

import (
	domainUser "github.com/0n1shi/domain-driven-design/domain/user"
)

type UserUsecase struct {
	service *domainUser.UserService
}

func NewUserUsecase(service *domainUser.UserService) *UserUsecase {
	return &UserUsecase{service: service}
}

func (usecase *UserUsecase) FindAll() ([]*DTOUser, error) {
	users, err := usecase.service.GetAll()
	if err != nil {
		return nil, err
	}
	return ToDTOUsers(users), nil
}

func (usecase *UserUsecase) FindByID(id string) (*DTOUser, error) {
	userID, err := domainUser.NewUserID(&id)
	if err != nil {
		return nil, err
	}
	user, err := usecase.service.GetByID(userID)
	if err != nil {
		return nil, err
	}
	return ToDTOUser(user), nil
}

type CreateUserInput struct {
	Name     string `json:"name" validate:"required,min=3,max=20"`
	Password string `json:"password" validate:"required"`
}

func (usecase *UserUsecase) Create(input *CreateUserInput) (*DTOUser, error) {
	if err := usecase.service.Register(&domainUser.CreateUserInput{
		Name:     input.Name,
		Password: input.Password,
	}); err != nil {
		return nil, err
	}

	name, err := domainUser.NewUsername(input.Name)
	if err != nil {
		return nil, err
	}

	user, err := usecase.service.GetByName(name)
	if err != nil {
		return nil, err
	}

	return ToDTOUser(user), nil
}
