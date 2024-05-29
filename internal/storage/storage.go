package storage

import (
	"os"
)

type Storage struct {
	fileName string
}

// NewStorage открывает json файл, что используется как хранилище, и если его нет создает его
func NewStorage(fileName string) (*Storage, error) {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0660)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return &Storage{fileName: fileName}, nil
}
