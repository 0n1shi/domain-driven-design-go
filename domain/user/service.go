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

func (service *UserService) GetByID(id UserID) (*User, error) {
	return service.repository.FindByID(id)
}

func (service *UserService) GetByName(name Username) (*User, error) {
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
	found := service.isNameRegistered(name)
	if found {
		return errors.WithStack(ErorrUserAlreadyRegistered)
	}
	hashedPassword, err := hashPassword(user.password.Get())
	if err != nil {
		return err
	}
	return service.repository.Create(&CreatedUser{
		ID:       user.id,
		Name:     user.name,
		Password: hashedPassword,
	})
}

func (service *UserService) isNameRegistered(name Username) bool {
	_, err := service.repository.FindByName(name)
	if err != nil {
		return false
	}
	return true
}
