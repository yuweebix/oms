package errors

import "errors"

var (
	// storage
	// order
	ErrOrderExpired       = errors.New("order expired")
	ErrOrderNotExpired    = errors.New("order not expired")
	ErrOrderAlreadyExists = errors.New("order already exists")
	ErrOrderNotFound      = errors.New("order not found")

	// user
	ErrUserInvalid = errors.New("user invalid")

	// status
	ErrStatusInvalid = errors.New("status invalid")

	// other
	ErrEmpty = errors.New("empty")
)
