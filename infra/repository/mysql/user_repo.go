package mysql

import (
	"github.com/0n1shi/domain-driven-design/domain/user"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

var _ user.UserRepositoryInterface = (*UserRepository)(nil)

func NewUserRepository(db *gorm.DB) *UserRepository {

	return &UserRepository{db: db}
}

func (repo *UserRepository) FindAll() ([]*user.User, error) {
	users := []User{}
	result := repo.db.Find(&users)
	if result.Error != nil {
		return nil, errors.WithStack(result.Error)
	}
	domainUsers, err := ToDomainUsers(users)
	if err != nil {
		return nil, err
	}
	return domainUsers, nil
}

func (repo *UserRepository) FindByID(id *user.UserID) (*user.User, error) {
	user := User{}
	result := repo.db.Find(&user)
	if result.Error != nil {
		return nil, errors.WithStack(result.Error)
	}
	domainUser, err := ToDomainUser(user)
	if err != nil {
		return nil, err
	}
	return domainUser, nil
}

func (repo *UserRepository) Create(user *user.CreatedUser) error {
	id := user.ID
	name := user.Name
	password := user.Password
	result := repo.db.Create(&User{
		ID:       id.Get(),
		Name:     name.Get(),
		Password: password.Get(),
	})
	return errors.WithStack(result.Error)
}

func (repo *UserRepository) Update(user *user.UpdatedUser) error {
	return nil
}

func (repo *UserRepository) Delete(id *user.UserID) error {
	return nil
}
