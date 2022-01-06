package user

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type UserID struct {
	val string
}

func NewUserID(newID *string) (*UserID, error) {
	if newID != nil {
		return &UserID{val: *newID}, nil
	}

	userID, err := uuid.NewUUID()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &UserID{val: userID.String()}, nil
}

func (id *UserID) Get() string {
	return id.val
}
