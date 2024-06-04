package models

import "time"

type Order struct {
	ID        int       `json:"order_id"`
	User      *User     `json:"user_id"`
	Expiry    time.Time `json:"stored_until"`
	ReturnBy  time.Time `json:"return_by"`
	Status    Status    `json:"status"`
	Hash      string    `json:"hash"`
	CreatedAt time.Time `json:"created_at"`
}

type User struct {
	ID int `json:"user_id"`
}
