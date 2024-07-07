package models

import "time"

type Message struct {
	CreatedAt  time.Time `json:"created_at"`
	MethodName string    `json:"method"`
	RawRequest string    `json:"request"`
}

type MessageWithError struct {
	Message
	Error string `json:"error"`
}
