package errors

import "errors"

var (
	ErrNoOrdersAffected      = errors.New("no orders affected")
	ErrTooManyOrdersAffected = errors.New("too many orders affected")
)
