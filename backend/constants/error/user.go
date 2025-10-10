package error

import "errors"

var (
	ErrUserNotFound         = errors.New("User not found")
	ErrPasswordInCorrect    = errors.New("Password incorrect")
	ErrUsernameExists       = errors.New("Username exists")
	ErrEmailExists          = errors.New("Email exists")
	ErrPasswordDoesNotMatch = errors.New("Password doesn't match")
)

var UserErrors = []error{
	ErrUserNotFound,
	ErrPasswordDoesNotMatch,
	ErrPasswordInCorrect,
	ErrUsernameExists,
}
