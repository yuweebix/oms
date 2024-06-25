package domain

import (
	"time"

	e "gitlab.ozon.dev/yuweebix/homework-1/internal/domain/errors"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
	"gitlab.ozon.dev/yuweebix/homework-1/pkg/hash"
)

// AcceptReturn принимает возврат от клиента
func (d *Domain) AcceptReturn(o *models.Order) (err error) {
	ro, err := d.storage.GetOrder(o) // ro - return order
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

	err = d.storage.UpdateOrder(ro)
	if err != nil {
		return err
	}
	return nil
}

// ListReturns выводит список возвратов с пагинацией
func (d *Domain) ListReturns(limit uint64, offset uint64) (list []*models.Order, err error) {
	list, err = d.storage.GetReturns(limit, offset)
	if err != nil {
		return nil, err
	}

	return list, nil
}
