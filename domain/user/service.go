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

func (service *UserService) Create(user *User) error {
	return service.repository.Create(user)
}
