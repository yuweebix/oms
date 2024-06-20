package hash

import (
	"github.com/google/uuid"
)

// GenerateHash возвращает случайный "хэш"
func GenerateHash() string {
	return uuid.New().String()
}
