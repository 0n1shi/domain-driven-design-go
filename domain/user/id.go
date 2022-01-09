package user

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type userID struct {
	val string
}

type UserID interface {
	Get() string
}

var _ UserID = (*userID)(nil)

func NewUserID(newID *string) (UserID, error) {
	if newID != nil {
		return &userID{val: *newID}, nil
	}

	id, err := uuid.NewUUID()
	if err != nil {
		return nil, errors.WithStack(ErrorFailedToGenerateUserID)
	}
	return &userID{val: id.String()}, nil
}

func (id *userID) Get() string {
	return id.val
}
