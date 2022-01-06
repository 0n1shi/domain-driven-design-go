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

type CreateUserInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (service *UserService) Register(input *CreateUserInput) error {
	user, err := NewUser(input.Name, input.Password, nil)
	if err != nil {
		return errors.WithStack(err)
	}
	return service.repository.Create(&CreatedUser{
		ID:       user.id,
		Name:     user.name,
		Password: user.password,
	})
}
