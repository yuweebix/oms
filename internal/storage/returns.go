package storage

import (
	"errors"
	"sort"
	"time"

	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
	"gitlab.ozon.dev/yuweebix/homework-1/pkg/hash"
)

// AddReturn добавляет заказ в хранилище
func (s *Storage) AddReturn(o *models.Order) error {
	var database map[int]*models.Order
	var err error

	// запишем данные из файла в database
	if database, err = s.loadOrders(); err != nil {
		return err
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

	// помечаем заказ как возвращенный
	o.Status = models.StatusReturned
	o.Hash = hash.GenerateHash() // HASH
	database[o.ID] = o

	return s.saveOrders(database)
}

// ListReturns передает список возвратов от start-ого возврата до finish-ого возврата
// почти ничем не отличается от ListOrders
// добавляется условие в цикле и немного меняется срез
func (s *Storage) ListReturns(start, finish int) ([]*models.Order, error) {
	var database map[int]*models.Order
	var err error

	// запишем данные из файла в database
	if database, err = s.loadOrders(); err != nil {
		return nil, err
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
