package domain

import (
	"context"
	"fmt"
	"os"
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
		// происходит несколько запросов к бд, надо гарантировать, что заказ полученный при GetOrder не изменился
		// до вызова UpdateOrder другим конкурентным вызовом
		// при ReadCommitted этой проверки не произойдет - соответственно недостаточно
		IsoLevel:   models.RepeatableRead,
		AccessMode: models.ReadWrite,
	}

	ro := &models.Order{} // ro - return order

	// сначала проверим в кеши ли заказ
	cachedOrder, err := d.cache.GetOrder(ctx, o)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	if cachedOrder != nil {
		ro = cachedOrder
	}

	// начинаем транзакцию
	err = d.storage.RunTx(ctx, opts, func(ctxTX context.Context) error {
		if cachedOrder == nil {
			ro, err = d.storage.GetOrder(ctxTX, o)
			if err != nil {
				return err
			}
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
		return d.storage.UpdateOrder(ctxTX, ro)
	})
	if err != nil {
		return err
	}

	return d.cache.SetOrder(ctx, ro)
}

// ListReturns выводит список возвратов с пагинацией
func (d *Domain) ListReturns(ctx context.Context, limit uint64, offset uint64) (list []*models.Order, err error) {
	cachedList, err := d.cache.GetReturns(ctx, limit, offset)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	if len(cachedList) != 0 {
		list = cachedList
	}

	// можно обойтись и без эксплицитной транзакции, ведь мы просто читаем данные, не проверяем их и не изменяем
	if len(cachedList) == 0 {
		list, err = d.storage.GetReturns(ctx, limit, offset)
		if err != nil {
			return nil, err
		}
	}

	return list, err
}
