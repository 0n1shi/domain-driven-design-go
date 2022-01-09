package user

type CreatedUser struct {
	ID       UserID
	Name     Username
	Password UserPassword
}

type UpdatedUser struct {
	Name     Username
	Password UserPassword
}

type UserRepositoryInterface interface {
	FindAll() ([]*User, error)
	FindByID(id UserID) (*User, error)
	FindByName(id *Username) (*User, error)
	Create(user *CreatedUser) error
	Update(user *UpdatedUser) error
	Delete(id UserID) error
}
