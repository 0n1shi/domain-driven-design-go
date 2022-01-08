package redis

import (
	"time"

	domainUser "github.com/0n1shi/domain-driven-design/domain/user"
)

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func ToDomainUser(u User) (*domainUser.User, error) {
	return domainUser.NewUser(u.Name, u.Password, &u.ID)
}

func ToDomainUsers(users []User) ([]*domainUser.User, error) {
	domainUsers := []*domainUser.User{}
	for _, u := range users {
		domainUser, err := ToDomainUser(u)
		if err != nil {
			return nil, err
		}
		domainUsers = append(domainUsers, domainUser)
	}
	return domainUsers, nil
}
