package models

type Status string

const (
	StatusPending   Status = "pending"
	StatusAccepted  Status = "accepted"
	StatusDelivered Status = "delivered"
	StatusReturned  Status = "returned"
)
