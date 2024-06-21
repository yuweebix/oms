package domain

import (
	"time"

	e "gitlab.ozon.dev/yuweebix/homework-1/internal/domain/errors"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
	"gitlab.ozon.dev/yuweebix/homework-1/pkg/hash"
)

const (
	returnByAllowedTime = time.Hour * 48
)

// AcceptOrder принимает заказ от курьера
func (s *Domain) AcceptOrder(o *models.Order) (_ error) {
	// срок хранения превышен
	if o.Expiry.Before(time.Now()) {
		return e.ErrOrderExpired
	}

	// проверка наличия упаковки в базе данных
	packaging, err := s.storage.GetPackaging(&o.Packaging)
	if err != nil {
		return err
	}

	// если равен нулю, то лимита нету
	if packaging.WeightLimit != 0 && o.Weight > packaging.WeightLimit {
		return e.ErrOrderTooHeavy
	}

	o.Cost += packaging.Cost
	o.Packaging = models.Packaging{
		Type:        packaging.Type,
		Cost:        packaging.Cost,
		WeightLimit: packaging.WeightLimit,
	}

	o.Status = models.StatusAccepted
	o.CreatedAt = time.Now().UTC()
	o.Hash = hash.GenerateHash() // HASH

	return s.storage.CreateOrder(o)
}

// ReturnOrder возвращает заказ курьеру
func (s *Domain) ReturnOrder(o *models.Order) (err error) {
	o, err = s.storage.GetOrder(o)
	if err != nil {
		return err
	}

	// если вернули, то не имеет значение прошёл ли срок хранения
	if o.Status != models.StatusReturned && o.Expiry.After(time.Now()) {
		return e.ErrOrderNotExpired
	}

	o.Hash = hash.GenerateHash() // HASH

	return s.storage.DeleteOrder(o)
}

// ListOrders выводит список заказов с пагинацией, сортировкой и фильтрацией
func (s *Domain) ListOrders(userID uint64, limit uint64, offset uint64, isStored bool) (list []*models.Order, err error) {
	list, err = s.storage.GetOrders(userID, limit, offset, isStored)
	if err != nil {
		return nil, err
	}

	return list, nil
}

// DeliverOrders принимает список заказов, переводит их в форму для обработки в хранилище
func (s *Domain) DeliverOrders(orderIDs []uint64) (err error) {
	if len(orderIDs) == 0 {
		return e.ErrEmpty
	}

	list, err := s.storage.GetOrdersForDelivery(orderIDs)
	if err != nil {
		return err
	}

	// когда передаются ID заказов, которых нет в базе данных
	if len(list) != len(orderIDs) {
		return e.ErrOrderNotFound
	}

	// Можно выдавать только те заказы, которые были приняты от курьера и чей срок хранения меньше текущей даты.
	// Все ID заказов должны принадлежать только одному клиенту.
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

		err = s.storage.UpdateOrder(list[i])
		if err != nil {
			return err
		}
	}

	return nil
}
