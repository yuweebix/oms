package storage

import (
	"encoding/json"
	"os"

	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
)

// Storage представляет слой хранилища (потом заменить бдшкой)
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

// readJSONFileToMap читает json файл и переносит его содержимое
func readJSONFileToMap[K comparable, V any](s *Storage) (database map[K]V, err error) {
	b, err := os.ReadFile(s.fileName)
	if err != nil {
		return nil, err
	}

	// если файл пуст, инициализируем пустую мапу
	if len(b) == 0 {
		database = make(map[K]V)
		return database, nil
	}

	err = json.Unmarshal(b, &database)
	if err != nil {
		return nil, err
	}

	return database, nil
}

// writeMapToJSONFile переводит содержимое мапы в json
func writeMapToJSONFile[K comparable, V any](s *Storage, database map[K]V) (err error) {
	b, err := json.MarshalIndent(database, "  ", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.fileName, b, 0666)
}

// loadOrders хелпер для чтения из JSON файла именно в map[int]*models.Order
func (s *Storage) loadOrders() (map[int]*models.Order, error) {
	return readJSONFileToMap[int, *models.Order](s)
}

// saveOrders хелпер для записи именно map[int]*models.Orders в JSON файл
func (s *Storage) saveOrders(database map[int]*models.Order) error {
	return writeMapToJSONFile(s, database)
}
