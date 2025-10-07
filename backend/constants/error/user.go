package error

import "errors"

var (
	ErrUserNotFound         = errors.New("User not found")
	ErrPasswordInCorrect    = errors.New("Password incorrect")
	ErrUsernameExists       = errors.New("Username exists")
	ErrPasswordDoesNotMatch = errors.New("Password doesn't match")
)

var UserErrors = []error{
	ErrUserNotFound,
	ErrPasswordDoesNotMatch,
	ErrPasswordInCorrect,
	ErrUsernameExists,
}
