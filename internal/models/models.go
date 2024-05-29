package models

import "time"

type Order struct {
	ID     int       `json:"order_id"`
	User   *User     `json:"user_id"`
	Expiry time.Time `json:"stored_until"`
	Status Status    `json:"status"`
	Hash   string    `json:"hash"`
}

type User struct {
	ID int `json:"user_id"`
}
