package storage

import (
	"errors"
	"sort"
	"time"

	e "gitlab.ozon.dev/yuweebix/homework-1/internal/errors"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
	"gitlab.ozon.dev/yuweebix/homework-1/pkg/hash"
)

// AddOrder добавляет заказ в хранилище
func (s *Storage) AddOrder(o *models.Order) error {
	var database map[int]*models.Order
	var err error

	// запишем данные из файла в database
	if database, err = s.loadOrders(); err != nil {
		return err
	}

	// не можем создать заказ с таким же id
	if _, ok := database[o.ID]; ok {
		return e.ErrOrderAlreadyExists
	}

	// помечаем заказ как принятый
	o.Status = models.StatusAccepted
	o.Hash = hash.GenerateHash() // HASH
	database[o.ID] = o

	return s.saveOrders(database)
}

// DeleteOrder удаляет заказ из хранилища
func (s *Storage) DeleteOrder(o *models.Order) error {
	var database map[int]*models.Order
	var err error

	// запишем данные из файла в database
	if database, err = s.loadOrders(); err != nil {
		return err
	}

	if _, ok := database[o.ID]; !ok {
		return e.ErrOrderNotFound
	}

	// o.Hash = hash.GenerateHash() // HASH

	delete(database, o.ID)

	return s.saveOrders(database)
}

// ListOrders передает список заказов с указанным максимумом
func (s *Storage) ListOrders(limit int) ([]*models.Order, error) {
	var database map[int]*models.Order
	var err error

	// запишем данные из файла в database
	if database, err = s.loadOrders(); err != nil {
		return nil, err
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
func (s *Storage) CheckOrdersForDelivery(orderIDs map[int]struct{}) error {
	if len(orderIDs) == 0 {
		return errors.New("empty")
	}

	var database map[int]*models.Order
	var err error

	// запишем данные из файла в database
	if database, err = s.loadOrders(); err != nil {
		return err
	}

	// создаем список всех заказов, что нам передали
	var list = make([]*models.Order, 0, len(orderIDs))
	for _, v := range database {
		if _, ok := orderIDs[v.ID]; ok {
			list = append(list, v)
		}
	}

	// когда передаются ID заказов, которых нет в базе данных
	if len(list) != len(orderIDs) {
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

	return s.saveOrders(database)
}

// GetOrder пересылает полный объект заказа
func (s *Storage) GetOrder(o *models.Order) (*models.Order, error) {
	var database map[int]*models.Order
	var err error
	var ok bool

	if database, err = readJSONFileToMap[int, *models.Order](s); err != nil {
		return nil, err
	}

	if o, ok = database[o.ID]; !ok {
		return nil, e.ErrOrderNotFound
	}

	return o, nil
}
