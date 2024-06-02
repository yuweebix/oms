package storage

import (
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
	e "gitlab.ozon.dev/yuweebix/homework-1/internal/storage/errors"
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

	delete(database, o.ID)

	return s.saveOrders(database)
}

// ListOrders возвращает список заказов клиента
func (s *Storage) ListOrders(userID int) ([]*models.Order, error) {
	var database map[int]*models.Order
	var err error

	// запишем данные из файла в database
	if database, err = s.loadOrders(); err != nil {
		return nil, err
	}

	// записываем в список
	var list []*models.Order
	for _, v := range database {
		// совпадает ID клиента
		if v.User.ID == userID {
			list = append(list, v)
		}
	}

	return list, nil
}

// DeliverOrder помечает заказ, как переданный клиенту
// Его можно будет вернуть в течение двух дней
// на вход даются IDs заказов в форме сета
func (s *Storage) GetOrdersForDelivery(orderIDs map[int]struct{}) ([]*models.Order, error) {
	var database map[int]*models.Order
	var err error

	// запишем данные из файла в database
	if database, err = s.loadOrders(); err != nil {
		return nil, err
	}

	// создаем список всех заказов, что нам передали
	var list = make([]*models.Order, 0, len(orderIDs))
	for id := range orderIDs {
		if order, ok := database[id]; ok {
			list = append(list, order)
		}
	}

	return list, nil
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
