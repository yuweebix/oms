package models

import "time"

type Order struct {
	ID        uint64    `json:"id"`
	User      *User     `json:"user"`
	Expiry    time.Time `json:"stored_until"`
	ReturnBy  time.Time `json:"return_by"`
	Status    Status    `json:"status"`
	Hash      string    `json:"hash"`
	CreatedAt time.Time `json:"created_at"`
	Cost      float64   `json:"cost"`
	Weight    float64   `json:"weight"`
}

type User struct {
	ID uint64 `json:"id"`
}

// Unzip метод для распаковки значений в сеттер
func (o *Order) Unzip() (uint64, uint64, time.Time, time.Time, Status, string, time.Time, float64, float64) {
	return o.ID, o.User.ID, o.Expiry, o.ReturnBy, o.Status, o.Hash, o.CreatedAt, o.Cost, o.Weight
}
