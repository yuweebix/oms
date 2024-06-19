package service

import (
	"time"

	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
	e "gitlab.ozon.dev/yuweebix/homework-1/internal/service/errors"
	"gitlab.ozon.dev/yuweebix/homework-1/pkg/hash"
)

// AcceptReturn принимает возврат от клиента
func (s Service) AcceptReturn(o *models.Order) (err error) {
	ro, err := s.storage.GetOrder(o) // ro - return order
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

	err = s.storage.UpdateOrder(ro)
	if err != nil {
		return err
	}
	return nil
}

// ListReturns выводит список заказов
func (s Service) ListReturns(start, finish int) (list []*models.Order, err error) {
	list, err = s.storage.GetReturns()
	if err != nil {
		return nil, err
	}

	// 1 <= start <= len(list)
	if start > len(list) {
		return nil, nil
	} else if start < 1 {
		start = 1
	}

	// 1 <= finish <= len(list)
	if finish < 1 || finish > len(list) {
		finish = len(list)
	}

	// start <= finish
	if start > finish {
		return nil, nil
	}

	return list[start-1 : finish], nil
}
