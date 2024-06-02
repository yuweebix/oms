package service

import (
	"time"

	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
	e "gitlab.ozon.dev/yuweebix/homework-1/internal/service/errors"
	"gitlab.ozon.dev/yuweebix/homework-1/pkg/hash"
)

// AcceptReturn принимает возврат от клиента
func (s Service) AcceptReturn(o *models.Order) error {
	var ro *models.Order // ro - return order
	var err error

	ro, err = s.storage.GetOrder(o)
	if err != nil {
		return err
	}

	// должен быть доставлен
	if ro.Status != models.StatusDelivered {
		return e.ErrStatusInvalid
	}

	// окно возврата пройдено
	if ro.ReturnBy.Before(time.Now()) {
		return e.ErrOrderExpired
	}

	// не совпадает id клиента
	if ro.User.ID != o.User.ID {
		return e.ErrUserInvalid
	}

	// помечаем заказ как возвращенный
	ro.Status = models.StatusReturned
	ro.Hash = hash.GenerateHash() // HASH

	err = s.storage.DeleteOrder(ro)
	if err != nil {
		return err
	}
	return s.storage.AddOrder(ro)
}

// ListReturns выводит список заказов
func (s Service) ListReturns(start, finish int) ([]*models.Order, error) {
	var list []*models.Order
	var err error

	list, err = s.storage.ListReturns()
	if err != nil {
		return nil, err
	}

	// 1 <= start <= len(list)
	if start > len(list) {
		return nil, nil
	} else if start < 1 {
		start = 1
	}

	// 0 <= finish <= len(list)
	if finish > len(list) {
		finish = len(list)
	} else if finish < 0 {
		finish = 0
	}

	// start <= finish
	if start > finish {
		return nil, nil
	}

	return list[start-1 : finish], nil
}
