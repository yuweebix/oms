package models

type Status int

const (
	StatusPending = iota + 1
	StatusAccepted
	StatusDelivered
	StatusReturned
)
