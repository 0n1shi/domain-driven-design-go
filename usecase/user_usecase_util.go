package usecase

import (
	domainUser "github.com/0n1shi/domain-driven-design/domain/user"
	"github.com/pkg/errors"
)

type DTOUser struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func ToDTOUser(user *domainUser.User) *DTOUser {
	id := user.GetID()
	name := user.GetName()
	password := user.GetPassword()
	return &DTOUser{
		ID:       id.Get(),
		Name:     name.Get(),
		Password: password.Get(),
	}
}

func ToDTOUsers(users []*domainUser.User) []*DTOUser {
	dtoUsers := []*DTOUser{}
	for _, u := range users {
		dtoUsers = append(dtoUsers, ToDTOUser(u))
	}
	return dtoUsers
}

var publicErrors = []error{
	domainUser.ErrorInvalidUsername,
	domainUser.ErrorFailedToGenerateUserID,
	domainUser.ErrorInvalidUserPassword,
	domainUser.ErorrUserAlreadyRegistered,
}

func IsPublicErorr(err error) bool {
	for _, e := range publicErrors {
		if errors.Is(err, e) {
			return true
		}
	}
	return false
}
