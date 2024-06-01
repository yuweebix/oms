package models

type Status int

const (
	StatusPending = iota
	StatusAccepted
	StatusDelivered
	StatusReturned
)
