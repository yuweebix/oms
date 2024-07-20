package domain

import (
	"context"
	"fmt"
	"os"
	"time"

	e "gitlab.ozon.dev/yuweebix/homework-1/internal/domain/errors"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
	"gitlab.ozon.dev/yuweebix/homework-1/pkg/hash"
	"gitlab.ozon.dev/yuweebix/homework-1/pkg/utils"
)

const (
	returnByAllowedTime = time.Hour * 48
)

// AcceptOrder принимает заказ от курьера
func (d *Domain) AcceptOrder(ctx context.Context, o *models.Order) (err error) {
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
	o.Cost += utils.ConvertToMicrocurrency(float64(packaging.Cost))

	o.Status = models.StatusAccepted
	o.CreatedAt = time.Now().UTC()
	o.Hash = hash.GenerateHash() // HASH

	// можно обойтись и без эксплицитной транзакции, потому что постгресс за нас создаст транзакцию и перечитывать данные при создании нету смысла
	err = d.storage.CreateOrder(ctx, o)
	if err != nil {
		return err
	}

	cacheErr := d.cache.SetOrder(ctx, o)
	if cacheErr != nil {
		fmt.Fprintln(os.Stderr, cacheErr)
	}

	return nil
}

// ReturnOrder возвращает заказ курьеру
func (d *Domain) ReturnOrder(ctx context.Context, o *models.Order) (err error) {
	// вынесем генерацию хэша за транзакцию
	hash := hash.GenerateHash() // HASH

	opts := models.TxOptions{
		// допустим, мы прочитали запись до того, как в ходе конкурентной транзакции поменяли статус заказа
		// в таком случае, ReadCommitted это не учтёт и удалит
		// в случае же RepeatableRead такого не произойдет
		IsoLevel:   models.RepeatableRead,
		AccessMode: models.ReadWrite,
	}

	// сначала проверим в кеши ли заказ
	cachedOrder, cacheErr := d.cache.GetOrder(ctx, o)
	if cacheErr != nil {
		fmt.Fprintln(os.Stderr, cacheErr)
	}

	// начинаем транзакцию
	err = d.storage.RunTx(ctx, opts, func(ctxTX context.Context) error {
		// если нет, то пытаемся достать из бд
		switch cachedOrder {
		case nil: // не закеширован
			o, err = d.storage.GetOrder(ctxTX, o)
			if err != nil {
				return err
			}
		default: // закеширован
			o = cachedOrder
		}

		// если заказ вернули в пвз, то не имеет значение прошёл ли срок хранения
		if o.Status != models.StatusReturned && o.Expiry.After(time.Now()) {
			return e.ErrOrderNotExpired
		}

		o.Hash = hash

		return d.storage.DeleteOrder(ctxTX, o)
	})
	if err != nil {
		return err
	}

	// не забываем удалить из кеша
	cacheErr = d.cache.DeleteOrder(ctx, o)
	if cacheErr != nil {
		fmt.Fprintln(os.Stderr, cacheErr)
	}

	return nil
}

// ListOrders выводит список заказов с пагинацией, сортировкой и фильтрацией
func (d *Domain) ListOrders(ctx context.Context, userID uint64, limit uint64, offset uint64, isStored bool) (list []*models.Order, err error) {
	cachedList, cacheErr := d.cache.GetOrders(ctx, userID, limit, offset, isStored)
	if cacheErr != nil {
		fmt.Fprintln(os.Stderr, cacheErr)
	}

	switch cachedList {
	case nil:
		// можно обойтись и без эксплицитной транзакции, ведь мы просто читаем данные, не проверяем их и не изменяем
		list, err = d.storage.GetOrders(ctx, userID, limit, offset, isStored)
		if err != nil {
			return nil, err
		}
	default:
		list = cachedList
	}

	return list, nil
}

// DeliverOrders принимает список заказов, переводит их в форму для обработки в хранилище
func (d *Domain) DeliverOrders(ctx context.Context, orderIDs []uint64) (err error) {
	if len(orderIDs) == 0 {
		return e.ErrEmpty
	}

	// вынесем сложную операцию за транзакцию
	// насчёт размера можно не беспокоиться, потому что мы ассёртим, что размеры совпадают
	hashes := make([]string, len(orderIDs))
	for i := range hashes {
		hashes[i] = hash.GenerateHash() // HASH
	}

	opts := models.TxOptions{
		// более критично, чтобы данные были более синхронизированы при их изменении
		// подобная ситуация, что и в ReturnOrder только ещё важнее в связи с работой с множеством записей
		IsoLevel:   models.RepeatableRead,
		AccessMode: models.ReadWrite,
	}

	list := make([]*models.Order, 0, len(orderIDs))

	// сначала проверим в кеше ли заказы, но нужно, чтобы совпали lenы
	cachedList, cacheErr := d.cache.GetOrdersForDelivery(ctx, orderIDs)
	if cacheErr != nil {
		fmt.Fprintln(os.Stderr, cacheErr)
	}

	// начинаем транзакцию
	err = d.storage.RunTx(ctx, opts, func(ctxTX context.Context) error {
		switch {
		case len(cachedList) == len(orderIDs): // из кеша вернулось нужное количество заказов
			list = cachedList
		default:
			list, err = d.storage.GetOrdersForDelivery(ctxTX, orderIDs)
			if err != nil {
				return err
			}
		}

		// если длины всё ещё не равны, значит что-то не так с введёнными данными
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
			list[i].Hash = hashes[i]

			err = d.storage.UpdateOrder(ctxTX, list[i])
			if err != nil {
				return err
			}
			cacheErr = d.cache.SetOrder(ctx, list[i])
			if cacheErr != nil {
				fmt.Fprintln(os.Stderr, cacheErr)
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
