package user

import "github.com/pkg/errors"

type UserService struct {
	repository UserRepositoryInterface
}

func NewUserService(repo UserRepositoryInterface) *UserService {
	return &UserService{repository: repo}
}

func (service *UserService) GetAll() ([]*User, error) {
	return service.repository.FindAll()
}

func (service *UserService) GetByID(id *UserID) (*User, error) {
	return service.repository.FindByID(id)
}

func (service *UserService) GetByName(name *Username) (*User, error) {
	return service.repository.FindByName(name)
}

type CreateUserInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (service *UserService) Register(input *CreateUserInput) error {
	user, err := NewUser(input.Name, input.Password, nil)
	if err != nil {
		return err
	}
	name := user.GetName()
	found, err := service.isNameRegistered(&name)
	if err != nil {
		return err
	}
	if found {
		return errors.WithStack(ErorrUserAlreadyRegistered)
	}
	return service.repository.Create(&CreatedUser{
		ID:       user.id,
		Name:     user.name,
		Password: user.password,
	})
}
