package mysql

import (
	"time"

	domainUser "github.com/0n1shi/domain-driven-design/domain/user"
	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"size:36"`
	Name      string `gorm:"not null,unique"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
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
