package service

import (
	"time"

	e "gitlab.ozon.dev/yuweebix/homework-1/internal/errors"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
)

// AcceptOrder принимает заказ от курьера
func (m *Service) AcceptOrder(o *models.Order) error {
	// срок хранения превышен
	if o.Expiry.Before(time.Now()) {
		return e.ErrOrderExpired
	}

	return m.Storage.AddOrder(o)
}

// ReturnOrder возвращает заказ курьеру
func (m *Service) ReturnOrder(o *models.Order) error {
	var err error
	if o, err = m.Storage.GetOrder(o); err != nil {
		return err
	}
	if o.Expiry.After(time.Now()) {
		return e.ErrOrderNotExpired
	}
	return m.Storage.DeleteOrder(o)
}

// ListOrders выводит список заказов
func (m *Service) ListOrders(userID int, limit int) ([]*models.Order, error) {
	return m.Storage.ListOrders(userID, limit)
}

// DeliverOrders принимает список заказов, переводит их в форму для обработки в хранилище
func (m *Service) DeliverOrders(orderIDs []int) error {
	// создаем сет для быстрого поиска
	set := make(map[int]struct{}, len(orderIDs))
	for _, v := range orderIDs {
		set[v] = struct{}{}
	}

	return m.Storage.CheckOrdersForDelivery(set)
}
