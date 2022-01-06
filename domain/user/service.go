package user

type UserService struct {
	repository UserRepositoryInterface
}

func NewUserService(repo UserRepositoryInterface) *UserService {
	return &UserService{repository: repo}
}

func (service *UserService) FindAll() ([]*User, error) {
	return service.repository.FindAll()
}

func (service *UserService) FindByID(id *UserID) (*User, error) {
	return service.repository.FindByID(id)
}

type CreateUserInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (service *UserService) Create(input *CreateUserInput) error {
	user, err := NewUser(input.Name, input.Password, nil)
	if err != nil {
		return err
	}
	return service.repository.Create(&CreatedUser{
		ID:       user.id,
		Name:     user.name,
		Password: user.password,
	})
}
