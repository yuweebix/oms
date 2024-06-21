package errors

import "errors"

var (
	ErrOrderExpired    = errors.New("order expired")
	ErrOrderNotExpired = errors.New("order not expired")
	ErrOrderNotFound   = errors.New("order not found")
)

var (
	ErrEmpty = errors.New("empty")
)

var (
	ErrUserInvalid = errors.New("user invalid")
)

var (
	ErrStatusInvalid = errors.New("status invalid")
)

var (
	ErrOrderTooHeavy = errors.New("order is over the weight limit for the packaging")
)
