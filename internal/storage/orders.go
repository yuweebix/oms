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

// AddOrder добавляет заказ в хранилище
func (s *Storage) AddOrder(o *models.Order) error {
	b, errReadFile := os.ReadFile(s.fileName)
	if errReadFile != nil {
		return errReadFile
	}

	// срок хранения превышен
	if o.Expiry.Before(time.Now()) {
		return errors.New("expired")
	}

	var database map[int]*models.Order
	if len(b) == 0 {
		database = make(map[int]*models.Order) // если файл пуст, инициализируем пустую мапу
	} else {
		if errUnmarshal := json.Unmarshal(b, &database); errUnmarshal != nil {
			return errUnmarshal
		}
	}

	if _, ok := database[o.ID]; ok {
		return errors.New("order already exists")
	}

	// помечаем заказ как принятый
	o.Status = models.StatusAccepted
	o.Hash = hash.GenerateHash() // HASH
	database[o.ID] = o

	bWrite, errMarshal := json.MarshalIndent(database, "  ", "  ")
	if errMarshal != nil {
		return errMarshal
	}

	return os.WriteFile(s.fileName, bWrite, 0666)
}

// DeleteOrder удаляет заказ из хранилище
func (s *Storage) DeleteOrder(o *models.Order) error {
	b, errReadFile := os.ReadFile(s.fileName)
	if errReadFile != nil {
		return errReadFile
	}

	var database map[int]*models.Order
	if len(b) == 0 {
		return errors.New("order doesn't exist")
	} else {
		if errUnmarshal := json.Unmarshal(b, &database); errUnmarshal != nil {
			return errUnmarshal
		}
	}

	if _, ok := database[o.ID]; !ok {
		return errors.New("order doesn't exist")
	}

	// срок хранения не превышен
	if database[o.ID].Expiry.After(time.Now()) {
		return errors.New("not expired")
	}

	// o.Hash = hash.GenerateHash() // HASH

	delete(database, o.ID)

	bWrite, errMarshal := json.MarshalIndent(database, "  ", "  ")
	if errMarshal != nil {
		return errMarshal
	}

	return os.WriteFile(s.fileName, bWrite, 0666)
}

// ListOrders передает список заказов с указанным максимум
func (s *Storage) ListOrders(limit int) ([]*models.Order, error) {
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
		list = append(list, v)
		// v.Hash = hash.GenerateHash() // HASH
	}

	// сортим по order_id, чтобы список был постоянным
	sort.Slice(list, func(i, j int) bool {
		return list[i].ID < list[j].ID
	})

	// чтобы не выходить за границы, берем минимум из длины листа и лимита
	return list[:min(len(list), limit)], nil
}

// DeliverOrder помечает заказ, как переданный клиенту
// Его можно будет вернуть в течение двух дней
// на вход даются IDs заказов в форме сета
func (s *Storage) CheckOrdersForDelivery(oSet map[int]struct{}) error {
	if len(oSet) == 0 {
		return errors.New("empty")
	}

	b, errReadFile := os.ReadFile(s.fileName)
	if errReadFile != nil {
		return errReadFile
	}

	var database map[int]*models.Order
	if len(b) == 0 {
		return errors.New("empty")
	} else {
		if errUnmarshal := json.Unmarshal(b, &database); errUnmarshal != nil {
			return errUnmarshal
		}
	}

	// создаем список всех заказов, что нам передали
	var list = make([]*models.Order, 0, len(oSet))
	for _, v := range database {
		if _, ok := oSet[v.ID]; ok {
			list = append(list, v)
		}
	}

	// когда передаются ID заказов, которых нет в базе данных
	if len(list) != len(oSet) {
		return errors.New("invalid order IDs")
	}

	// Можно выдавать только те заказы, которые были приняты от курьера и чей срок хранения меньше текущей даты.
	// Все ID заказов должны принадлежать только одному клиенту.
	user_id := list[0].User.ID
	for _, v := range list {
		if v.Status != models.StatusAccepted {
			return errors.New("invalid status")
		}
		if v.User.ID != user_id {
			return errors.New("invalid user ID")
		}
		if v.Expiry.Before(time.Now()) {
			return errors.New("expired")
		}
	}

	// помечаем как переданные клиенту и оставляем два дня на возврат
	for i := range list {
		list[i].Status = models.StatusDelivered
		list[i].Expiry = time.Now().UTC().AddDate(0, 0, 2)
		list[i].Hash = hash.GenerateHash() // HASH
	}

	bWrite, errMarshal := json.MarshalIndent(database, "  ", "  ")
	if errMarshal != nil {
		return errMarshal
	}

	return os.WriteFile(s.fileName, bWrite, 0666)
}
