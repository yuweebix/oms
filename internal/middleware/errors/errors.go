package errors

import "errors"

var (
	ErrGoroutinesNumExceeded  = errors.New("too many goroutines")
	ErrGoroutinesNumSubceeded = errors.New("number of goroutines cannot be less than one")
)
