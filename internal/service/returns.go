package service

import (
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
)

// AcceptReturn принимает заказ от курьера
func (m Service) AcceptReturn(o *models.Order) error {
	return m.Storage.AddReturn(o)
}

// ListReturns выводит список заказов
func (m Service) ListReturns(start, finish int) ([]*models.Order, error) {
	return m.Storage.ListReturns(start, finish)
}
