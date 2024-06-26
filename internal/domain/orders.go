package domain

import (
	"context"
	"time"

	e "gitlab.ozon.dev/yuweebix/homework-1/internal/domain/errors"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
	"gitlab.ozon.dev/yuweebix/homework-1/pkg/hash"
)

const (
	returnByAllowedTime = time.Hour * 48
)

// AcceptOrder принимает заказ от курьера
func (d *Domain) AcceptOrder(o *models.Order) (_ error) {
	// срок хранения превышен
	if o.Expiry.Before(time.Now()) {
		return e.ErrOrderExpired
	}

	packaging, ok := models.GetPackaging(o.Packaging)
	// нету такой упаковки
	if !ok {
		return e.ErrPackagingInvalid
	}
	// если равен нулю, то лимита нету
	if packaging.WeightLimit != 0 && o.Weight > packaging.WeightLimit {
		return e.ErrOrderTooHeavy
	}
	o.Cost += packaging.Cost

	o.Status = models.StatusAccepted
	o.CreatedAt = time.Now().UTC()
	o.Hash = hash.GenerateHash() // HASH

	opts := models.TxOptions{
		// RepeatableRead излишне: мы просто вставляем новый заказ, а не выполняем сложные операции
		// ReadUncommitted недостаточно: из-за возможного грязного чтения
		IsoLevel:   models.ReadCommitted,
		AccessMode: models.ReadWrite,
	}

	return d.storage.RunTx(context.Background(), opts, func(ctxTX context.Context) error {
		return d.storage.CreateOrder(ctxTX, o)
	})
}

// ReturnOrder возвращает заказ курьеру
func (d *Domain) ReturnOrder(o *models.Order) (_ error) {
	opts := models.TxOptions{
		// RepeatableRead излишне: потому что читаем только одну запись и удаляем её в рамках одной транзакции
		// ReadUncommitted недостаточно: из-за возможного грязного чтения
		IsoLevel:   models.ReadCommitted,
		AccessMode: models.ReadWrite,
	}

	// начинаем транзакцию
	return d.storage.RunTx(context.Background(), opts, func(ctxTX context.Context) error {
		o, err := d.storage.GetOrder(ctxTX, o)
		if err != nil {
			return err
		}

		// если заказ вернули в пвз, то не имеет значение прошёл ли срок хранения
		if o.Status != models.StatusReturned && o.Expiry.After(time.Now()) {
			return e.ErrOrderNotExpired
		}

		o.Hash = hash.GenerateHash() // HASH

		return d.storage.DeleteOrder(ctxTX, o)
	})
}

// ListOrders выводит список заказов с пагинацией, сортировкой и фильтрацией
func (d *Domain) ListOrders(userID uint64, limit uint64, offset uint64, isStored bool) (list []*models.Order, err error) {
	opts := models.TxOptions{
		// вообще можно и без транзакции, ведь мы просто читаем данные,
		// и нам даже не критично, чтобы они были актуальными, потому что ничего не изменяется и не проверяется
		// но для некоторой "гибкости", и для того, чтобы соответствовали все функции некому "стандарту", что первый аргумент - ctxTX, оставил
		// ReadCommitted излишне: всё сделаем по минимуму
		IsoLevel:   models.ReadUncommitted,
		AccessMode: models.ReadOnly,
	}

	// начинаем транзакцию
	err = d.storage.RunTx(context.Background(), opts, func(ctxTX context.Context) error {
		list, err = d.storage.GetOrders(ctxTX, userID, limit, offset, isStored)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return list, nil
}

// DeliverOrders принимает список заказов, переводит их в форму для обработки в хранилище
func (d *Domain) DeliverOrders(orderIDs []uint64) (err error) {
	if len(orderIDs) == 0 {
		return e.ErrEmpty
	}

	opts := models.TxOptions{
		// кмк, более критично, чтобы данные были более синхронизированы при изменении данных, а не добавления
		IsoLevel:   models.RepeatableRead,
		AccessMode: models.ReadWrite,
	}

	// начинаем транзакцию
	err = d.storage.RunTx(context.Background(), opts, func(ctxTX context.Context) error {
		list, err := d.storage.GetOrdersForDelivery(ctxTX, orderIDs)
		if err != nil {
			return err
		}

		// когда передаются ID заказов, которых нет в базе данных
		if len(list) != len(orderIDs) {
			return e.ErrOrderNotFound
		}

		// можно выдавать только те заказы, которые были приняты от курьера и чей срок хранения меньше текущей даты
		// все ID заказов должны принадлежать только одному клиенту
		user_id := list[0].User.ID
		for _, v := range list {
			if v.Status != models.StatusAccepted {
				return e.ErrStatusInvalid
			}
			if v.User.ID != user_id {
				return e.ErrUserInvalid
			}
			if v.Expiry.Before(time.Now()) {
				return e.ErrOrderExpired
			}
		}

		// помечаем как переданные клиенту и оставляем два дня на возврат
		for i := range list {
			list[i].Status = models.StatusDelivered
			list[i].ReturnBy = time.Now().UTC().Add(returnByAllowedTime)
			list[i].Hash = hash.GenerateHash() // HASH

			err = d.storage.UpdateOrder(ctxTX, list[i])
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
