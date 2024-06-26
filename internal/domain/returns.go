package domain

import (
	"context"
	"time"

	e "gitlab.ozon.dev/yuweebix/homework-1/internal/domain/errors"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
	"gitlab.ozon.dev/yuweebix/homework-1/pkg/hash"
)

// AcceptReturn принимает возврат от клиента
func (d *Domain) AcceptReturn(ctx context.Context, o *models.Order) (err error) {
	// вынесем генерацию хэша за транзакцию
	hash := hash.GenerateHash() // HASH

	opts := models.TxOptions{
		// данные обновляются, поэтому следует использовать ReapeatableRead
		IsoLevel:   models.RepeatableRead,
		AccessMode: models.ReadWrite,
	}

	// начинаем транзакцию
	err = d.storage.RunTx(ctx, opts, func(ctxTX context.Context) error {
		ro, err := d.storage.GetOrder(ctxTX, o) // ro - return order
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
		ro.Hash = hash

		// обновляем заказ в хранилище
		err = d.storage.UpdateOrder(ctxTX, ro)
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
func (d *Domain) ListReturns(ctx context.Context, limit uint64, offset uint64) (list []*models.Order, err error) {
	// можно обойтись и без эксплисивной транзакции
	list, err = d.storage.GetReturns(ctx, limit, offset)

	if err != nil {
		return nil, err
	}

	return list, err
}
