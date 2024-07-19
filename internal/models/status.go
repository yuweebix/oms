package models

type Status string

const (
	StatusPending   Status = "STATUS_PENDING"
	StatusAccepted  Status = "STATUS_ACCEPTED"
	StatusDelivered Status = "STATUS_DELIVERED"
	StatusReturned  Status = "STATUS_RETURNED"
)
