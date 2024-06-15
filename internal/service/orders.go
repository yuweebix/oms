package service

import (
	"errors"
	"sort"
	"time"

	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
	e "gitlab.ozon.dev/yuweebix/homework-1/internal/service/errors"
	"gitlab.ozon.dev/yuweebix/homework-1/pkg/hash"
)

// AcceptOrder принимает заказ от курьера
func (s *Service) AcceptOrder(o *models.Order) (_ error) {
	// срок хранения превышен
	if o.Expiry.Before(time.Now()) {
		return e.ErrOrderExpired
	}

	// помечаем заказ как принятый
	o.Status = models.StatusAccepted
	o.CreatedAt = time.Now().UTC()
	o.Hash = hash.GenerateHash() // HASH

	return s.storage.CreateOrder(o)
}

// ReturnOrder возвращает заказ курьеру
func (s *Service) ReturnOrder(o *models.Order) (err error) {
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

// ListOrders выводит список заказов
func (s *Service) ListOrders(userID int, limit int, isStored bool) (list []*models.Order, err error) {
	list, err = s.storage.GetOrders(userID)
	if err != nil {
		return nil, err
	}

	if isStored {
		list = filterByStoredOrders(list)
	}

	// сортим по времени получения
	sort.Slice(list, func(i, j int) bool {
		return list[i].CreatedAt.After(list[j].CreatedAt)
	})

	// 0 <= limit <= len(list)
	if limit <= 0 || limit > len(list) {
		limit = len(list)
	}

	return list[:limit], nil
}

// DeliverOrders принимает список заказов, переводит их в форму для обработки в хранилище
func (s *Service) DeliverOrders(orderIDs []int) (err error) {
	// создаем сет для быстрого поиска
	set := make(map[int]struct{}, len(orderIDs))
	for _, v := range orderIDs {
		set[v] = struct{}{}
	}

	if len(orderIDs) == 0 {
		return e.ErrEmpty
	}

	list, err := s.storage.GetOrdersForDelivery(set)
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
		list[i].ReturnBy = time.Now().UTC().AddDate(0, 0, 2)
		list[i].Hash = hash.GenerateHash() // HASH

		err = s.storage.UpdateOrder(list[i])
		if err != nil {
			if errors.Is(err, errors.ErrUnsupported) {
				err = s.storage.DeleteOrder(list[i])
				if err != nil {
					return err
				}

				err = s.storage.CreateOrder(list[i])
				if err != nil {
					return err
				}
			}
			return err
		}
	}

	return nil
}

// filterByStoredOrders фильтрует заказы, оставляя только те, что находятся в ПВЗ
func filterByStoredOrders(list []*models.Order) (newList []*models.Order) {
	for _, o := range list {
		if o.Status == models.StatusAccepted || o.Status == models.StatusReturned {
			newList = append(newList, o)
		}
	}
	return newList
}
