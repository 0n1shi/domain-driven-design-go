package db

import (
	"github.com/0n1shi/domain-driven-design/domain/user"
)

type DBUserRepository struct{}

var _ user.UserRepositoryInterface = (*DBUserRepository)(nil)

func NewDBUserRepository() *DBUserRepository {
	return &DBUserRepository{}
}

func (repo *DBUserRepository) FindAll() ([]*user.User, error) {
	return []*user.User{}, nil
}
func (repo *DBUserRepository) FindByID() (*user.User, error) {
	return &user.User{}, nil
}
func (repo *DBUserRepository) Create(user *user.User) error {
	return nil
}
func (repo *DBUserRepository) Update(user *user.User) error {
	return nil
}
func (repo *DBUserRepository) Delete(id *user.UserID) error {
	return nil
}
