package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/cache/schemas"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
)

// SetOrder задаёт значение данного заказа
func (c *Cache) SetOrder(ctx context.Context, o *models.Order) (err error) {
	orderKey := fmt.Sprintf("order:%d", o.ID)

	// юзер никогда не меняется у заказа, поэтому менять его не нужно
	c.zRemOrdersByStatus(ctx, orderKey)

	// будет храниться в форме хеша
	err = c.client.HSet(ctx, orderKey, schemas.FromModelsOrder(o)).Err()
	if err != nil {
		return err
	}

	// z - пара поле и его значение, по которому сортируется множество
	z := redis.Z{Score: float64(o.CreatedAt.Unix()), Member: orderKey}

	// для фильтрации по статусу (см. GetReturns)
	err = c.client.ZAdd(ctx, fmt.Sprintf("orders:by_status:%s", o.Status), z).Err()
	if err != nil {
		return err
	}

	// для фильтрации по id пользователя (см. GetOrders)
	err = c.client.ZAddNX(ctx, fmt.Sprintf("user:%d:orders", o.User.ID), z).Err()
	if err != nil {
		return err
	}

	// TTL
	switch o.ReturnBy {
	case time.Time{}: // ReturnBy не был задан
		err = c.client.ExpireAt(ctx, orderKey, o.Expiry).Err()
		if err != nil {
			return err
		}
	default: // ReturnBy задан, используем его
		err = c.client.ExpireAt(ctx, orderKey, o.ReturnBy).Err()
		if err != nil {
			return err
		}
	}

	return nil
}

// GetOrder используется для получения заказа из кеша
func (c *Cache) GetOrder(ctx context.Context, o *models.Order) (result *models.Order, err error) {
	orderKey := fmt.Sprintf("order:%d", o.ID)

	order := &schemas.Order{}
	err = c.client.HGetAll(ctx, orderKey).Scan(order)
	if err != nil {
		return nil, err
	}

	// если ничего не получили возвращаем nil
	empty := &schemas.Order{}
	if *order != *empty {
		return schemas.ToModelsOrder(order), nil
	}

	// очистим множества
	err = c.zRemOrdersByStatus(ctx, orderKey)
	if err != nil {
		return nil, err
	}
	err = c.zRemUserOrders(ctx, orderKey)
	return nil, err
}

// DeleteOrder удаляет из кеша заказ и связанные с ним элементы множеств
func (c *Cache) DeleteOrder(ctx context.Context, o *models.Order) (err error) {
	orderKey := fmt.Sprintf("order:%d", o.ID)

	// сначала надо получить
	order, err := c.GetOrder(ctx, o)
	if err != nil {
		return err
	}

	// проверить, не пустой ли
	if order == nil {
		return nil
	}

	// а теперь удалить
	err = c.client.Del(ctx, orderKey).Err()
	if err != nil {
		return err
	}

	err = c.client.ZRem(ctx, "orders:by_created_at", orderKey).Err()
	if err != nil {
		return err
	}
	err = c.client.ZRem(ctx, fmt.Sprintf("orders:by_status:%s", order.Status), orderKey).Err()
	if err != nil {
		return err
	}
	err = c.client.ZRem(ctx, fmt.Sprintf("user:%d:orders", order.User.ID), orderKey).Err()
	if err != nil {
		return err
	}

	return nil
}

// GetOrders возвращает список заказов клиента из кеша
func (c *Cache) GetOrders(ctx context.Context, userID uint64, limit uint64, offset uint64, isStored bool) (list []*models.Order, err error) {
	userKey := fmt.Sprintf("user:%d:orders", userID)

	orderKeys, err := c.client.ZRangeArgs(ctx, redis.ZRangeArgs{
		Key:   userKey,
		Start: 0,
		Stop:  -1,
		Rev:   true,
	}).Result()
	if err != nil {
		return nil, err
	}

	// начинаем с момента оффсета
	// не можем превысить длину среза или же ограничения по запросу
	size := min(uint64(len(orderKeys)), offset+limit)
	for i := offset; i < size; i++ {
		key := orderKeys[i]

		// нужен лишь id заказа для получения всего заказа
		orderID := extractOrderID(key)
		order, err := c.GetOrder(ctx, &models.Order{ID: orderID})
		if err != nil {
			return nil, err
		}

		// проверить, не пустой ли
		if order == nil {
			continue
		}

		// находится в пвз, только если заказ принят или же возвращен
		if isStored {
			if order.Status != models.StatusAccepted &&
				order.Status != models.StatusReturned {
				continue
			}
		}

		list = append(list, order)
	}

	// вроде бы, если ничего не зааппендили в лист
	// `len(list) == 0 && cap(list) == 0`
	// то и так будет `list == nil` работать, но на всякий оставил
	if len(list) == 0 {
		return nil, nil
	}

	return list, nil
}

// GetOrdersForDelivery возващает заказы из кеша по их IDs
func (c *Cache) GetOrdersForDelivery(ctx context.Context, orderIDs []uint64) (list []*models.Order, err error) {
	for _, orderID := range orderIDs {
		order, err := c.GetOrder(ctx, &models.Order{ID: orderID})
		if err != nil {
			return nil, err
		}

		// проверить, не пустой ли
		if order == nil {
			continue
		}

		list = append(list, order)
	}

	return list, nil
}
