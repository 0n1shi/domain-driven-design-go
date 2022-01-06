package user

import "github.com/pkg/errors"

var (
	ErrorInvalidUsername        = errors.New("invalid username (length 3 ~ 20)")
	ErrorFailedToGenerateUserID = errors.New("failed to generate user id")
	ErrorInvalidUserPassword    = errors.New("at least 1 upper case, 1 lower case and numeric must be EMBEDDED somewhere in the middle of the password")
)
