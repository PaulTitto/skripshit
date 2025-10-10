package error

import "errors"

var (
	ErrInternalServerError = errors.New("Internal Server Error")
	ErrSQLError            = errors.New("SQL Error and Database Server Failed to execute query")
	ErrTooManyRequests     = errors.New("Too many requests")
	ErrUnauthorized        = errors.New("Unauthorized")
	ErrForbidden           = errors.New("Forbidden")
	ErrInvalidToken        = errors.New("Invalid Token")
)

var GeneralErrors = []error{
	ErrInternalServerError,
	ErrSQLError,
	ErrTooManyRequests,
	ErrUnauthorized,
	ErrForbidden,
	ErrInvalidToken,
}
