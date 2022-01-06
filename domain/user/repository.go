package user

type UserRepositoryInterface interface {
	FindAll() ([]*User, error)
	FindByID(id *UserID) (*User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(id *UserID) error
}
