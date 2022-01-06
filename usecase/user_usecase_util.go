package usecase

import (
	domainUser "github.com/0n1shi/domain-driven-design/domain/user"
)

type DTOUser struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func ToDTOUser(user *domainUser.User) *DTOUser {
	return &DTOUser{
		ID:   user.ID.Get(),
		Name: user.Name.Get(),
	}
}

func ToDTOUsers(users []*domainUser.User) []*DTOUser {
	dtoUsers := []*DTOUser{}
	for _, u := range users {
		dtoUsers = append(dtoUsers, ToDTOUser(u))
	}
	return dtoUsers
}
