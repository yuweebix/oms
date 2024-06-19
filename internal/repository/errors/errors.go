package errors

import "errors"

var (
	ErrNoRowsAffected      = errors.New("no rows affected")
	ErrTooManyRowsAffected = errors.New("too many rows affected")
)
