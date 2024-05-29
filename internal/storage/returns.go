package storage

import (
	"encoding/json"
	"errors"
	"os"
	"sort"
	"time"

	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
	"gitlab.ozon.dev/yuweebix/homework-1/pkg/hash"
)

// AddReturn добавляет заказ в хранилище
func (s *Storage) AddReturn(o *models.Order) error {
	b, errReadFile := os.ReadFile(s.fileName)
	if errReadFile != nil {
		return errReadFile
	}

	var database map[int]*models.Order
	if len(b) == 0 {
		database = make(map[int]*models.Order) // если файл пуст, инициализируем пустую мапу
	} else {
		if errUnmarshal := json.Unmarshal(b, &database); errUnmarshal != nil {
			return errUnmarshal
		}
	}

	if _, ok := database[o.ID]; !ok {
		return errors.New("order doesn't exists")
	}

	// срок хранения превышен
	if database[o.ID].Expiry.Before(time.Now()) {
		return errors.New("expired")
	}

	if database[o.ID].User.ID != o.User.ID {
		return errors.New("invalid user")
	}

	// должен быть доставлен
	if database[o.ID].Status != models.StatusDelivered {
		return errors.New("invalid status")
	}

	// помечаем заказ как принятый
	o.Status = models.StatusReturned
	o.Hash = hash.GenerateHash() // HASH
	database[o.ID] = o

	bWrite, errMarshal := json.MarshalIndent(database, "  ", "  ")
	if errMarshal != nil {
		return errMarshal
	}

	return os.WriteFile(s.fileName, bWrite, 0666)
}

// ListReturns передает список возвратов от start-ого возврата до finish-ого возврата
// почти ничем не отличается от ListOrders
// добавляется условие в цикле и немного меняется срез
func (s *Storage) ListReturns(start, finish int) ([]*models.Order, error) {
	b, errReadFile := os.ReadFile(s.fileName)
	if errReadFile != nil {
		return nil, errReadFile
	}

	var database map[int]*models.Order
	if len(b) == 0 {
		return nil, errors.New("empty")
	} else {
		if errUnmarshal := json.Unmarshal(b, &database); errUnmarshal != nil {
			return nil, errUnmarshal
		}
	}

	// записываем в список
	var list = make([]*models.Order, 0, len(database))
	for _, v := range database {
		// ЗДЕСЬ
		if v.Status == models.StatusReturned {
			list = append(list, v)
		}
		// v.Hash = hash.GenerateHash() // HASH
	}

	// сортим по order_id, чтобы список был постоянным
	sort.Slice(list, func(i, j int) bool {
		return list[i].ID < list[j].ID
	})

	// И ЗДЕСЬ
	if start > len(list) {
		return nil, nil
	}
	// минус 1, потому что отсчитываем с 1-го возврата
	return list[start-1 : min(len(list), finish)], nil
}
