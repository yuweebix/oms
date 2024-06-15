package storage

import (
	"errors"

	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
	e "gitlab.ozon.dev/yuweebix/homework-1/internal/storage/errors"
)

// CreateOrder добавляет заказ в хранилище
func (s *Storage) CreateOrder(o *models.Order) (err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// запишем данные из файла в database
	database, err := s.loadOrders()
	if err != nil {
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
func (s *Storage) DeleteOrder(o *models.Order) (err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// запишем данные из файла в database
	database, err := s.loadOrders()
	if err != nil {
		return err
	}

	if _, ok := database[o.ID]; !ok {
		return e.ErrOrderNotFound
	}

	delete(database, o.ID)

	return s.saveOrders(database)
}

// GetOrders возвращает список заказов клиента
func (s *Storage) GetOrders(userID int) (list []*models.Order, err error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// запишем данные из файла в database
	database, err := s.loadOrders()
	if err != nil {
		return nil, err
	}

	// записываем в список
	for _, v := range database {
		// совпадает ID клиента
		if v.User.ID == userID {
			list = append(list, v)
		}
	}

	return list, nil
}

func (s *Storage) UpdateOrder(o *models.Order) (err error) {
	return errors.ErrUnsupported
}

// GetOrdersForDelivery возваращает список заказов клиенту на выдачу
func (s *Storage) GetOrdersForDelivery(orderIDs map[int]struct{}) (list []*models.Order, err error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// запишем данные из файла в database
	database, err := s.loadOrders()
	if err != nil {
		return nil, err
	}

	// создаем список всех заказов, что нам передали
	for id := range orderIDs {
		if order, ok := database[id]; ok {
			list = append(list, order)
		}
	}

	return list, nil
}

// GetOrder пересылает полный объект заказа
func (s *Storage) GetOrder(o *models.Order) (result *models.Order, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// запишем данные из файла в database
	database, err := s.loadOrders()
	if err != nil {
		return nil, err
	}

	result, ok := database[o.ID]
	if !ok {
		return nil, e.ErrOrderNotFound
	}

	return result, nil
}
