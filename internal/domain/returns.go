package domain

import (
	"time"

	e "gitlab.ozon.dev/yuweebix/homework-1/internal/domain/errors"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
	"gitlab.ozon.dev/yuweebix/homework-1/pkg/hash"
)

// AcceptReturn принимает возврат от клиента
func (d *Domain) AcceptReturn(o *models.Order) (err error) {
	opts := models.TxOptions{
		// данные обновляются, поэтому следует использовать ReapeatableRead
		IsoLevel:   models.RepeatableRead,
		AccessMode: models.ReadWrite,
	}

	// начинаем транзакцию
	err = d.storage.Begin(opts, func(tx models.Tx) error {
		ro, err := d.storage.GetOrder(tx, o) // ro - return order
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
		ro.Hash = hash.GenerateHash() // Генерация HASH

		// обновляем заказ в хранилище
		err = d.storage.UpdateOrder(tx, ro)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

// ListReturns выводит список возвратов с пагинацией
func (d *Domain) ListReturns(limit uint64, offset uint64) (list []*models.Order, err error) {
	opts := models.TxOptions{
		// также как и при ListOrders, просто читаются данные и не важно, чтобы они точно совпадали, ведь ничего не проверяется и не изменяется
		IsoLevel:   models.ReadUncommitted,
		AccessMode: models.ReadOnly,
	}

	// начинаем транзакцию
	err = d.storage.Begin(opts, func(tx models.Tx) error {
		list, err = d.storage.GetReturns(tx, limit, offset)
		if err != nil {
			return err
		}
		return nil
	})

	return list, err
}
