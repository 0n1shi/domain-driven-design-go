package redis

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"

	"github.com/0n1shi/domain-driven-design/domain/user"
)

type UserRepository struct {
	client *redis.Client
}

var _ user.UserRepositoryInterface = (*UserRepository)(nil)

func NewUserRepository(client *redis.Client) *UserRepository {
	return &UserRepository{client: client}
}

func (repo *UserRepository) FindAll() ([]*user.User, error) {
	keys, err := repo.client.Keys("*").Result()
	if err != nil {
		return nil, err
	}
	users := []*user.User{}
	for _, key := range keys {
		strData, err := repo.client.Get(key).Result()
		if err != nil {
			return nil, err
		}
		var user User
		if err = json.Unmarshal([]byte(strData), &user); err != nil {
			return nil, err
		}
		u, err := ToDomainUser(user)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (repo *UserRepository) FindByID(id user.UserID) (*user.User, error) {
	strData, err := repo.client.Get(id.Get()).Result()
	if err != nil {
		return nil, err
	}
	var user User
	if err = json.Unmarshal([]byte(strData), &user); err != nil {
		return nil, err
	}
	u, err := ToDomainUser(user)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (repo *UserRepository) FindByName(name user.Username) (*user.User, error) {
	keys, err := repo.client.Keys("*").Result()
	if err != nil {
		return nil, err
	}
	for _, key := range keys {
		strData, err := repo.client.Get(key).Result()
		if err != nil {
			return nil, err
		}
		var user User
		if err = json.Unmarshal([]byte(strData), &user); err != nil {
			return nil, err
		}
		if user.Name != name.Get() {
			continue
		}
		u, err := ToDomainUser(user)
		if err != nil {
			return nil, err
		}
		return u, nil
	}
	return nil, errors.WithStack(fmt.Errorf("user found by name \"%s\"", name.Get()))
}

func (repo *UserRepository) Create(user *user.CreatedUser) error {
	data, err := json.Marshal(&User{
		ID:       user.ID.Get(),
		Name:     user.Name.Get(),
		Password: user.Password.Get(),
	})
	if err != nil {
		return errors.WithStack(err)
	}
	return repo.client.Set(user.ID.Get(), string(data), 0).Err()
}

func (repo *UserRepository) Update(user *user.UpdatedUser) error {
	return nil
}

func (repo *UserRepository) Delete(id user.UserID) error {
	return nil
}
