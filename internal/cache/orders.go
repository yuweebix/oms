package cache

import (
	"context"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/cache/schemas"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
)

// SetOrder задаёт значение данного заказа
func (c *Cache) SetOrder(ctx context.Context, o *models.Order) (err error) {
	orderKey := fmt.Sprintf("order:%d", o.ID)

	// старый статус нужно удалить (если он имеется)
	// если его нету, то это не должно быть проблемой
	oldStatus, err := c.client.HGet(ctx, orderKey, "status").Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return err
	}
	err = c.client.ZRem(ctx, fmt.Sprintf("orders:by_status:%s", oldStatus), orderKey).Err()
	if err != nil {
		return err
	}

	// будет храниться в форме хеша
	err = c.client.HSet(ctx, orderKey, schemas.FromModelsOrder(o)).Err()
	if err != nil {
		return err
	}

	// для фильтрации по статусу (см. GetReturns)
	err = c.client.ZAdd(ctx, fmt.Sprintf("orders:by_status:%s", o.Status), redis.Z{
		Score:  float64(o.CreatedAt.Unix()),
		Member: orderKey,
	}).Err()
	if err != nil {
		return err
	}

	// для фильтрации по id пользователя (см. GetOrders)
	err = c.client.ZAdd(ctx, fmt.Sprintf("user:%d:orders", o.User.ID), redis.Z{
		Score:  float64(o.CreatedAt.Unix()),
		Member: orderKey,
	}).Err()
	if err != nil {
		return err
	}

	return nil
}

// GetOrder используется для получения заказа из кеша
func (c *Cache) GetOrder(ctx context.Context, o *models.Order) (result *models.Order, err error) {
	orderKey := fmt.Sprintf("order:%d", o.ID)

	cmd := c.client.HGetAll(ctx, orderKey)
	res, err := cmd.Result()
	if err != nil {
		return nil, err
	}
	// если ничего не получили возвращаем nil
	if len(res) == 0 {
		return nil, nil
	}

	order := &schemas.Order{}
	err = cmd.Scan(order)
	if err != nil {
		return nil, err
	}

	return schemas.ToModelsOrder(order), nil
}

// DeleteOrder удаляет из кеша заказ и связанные с ним элементы множеств
func (c *Cache) DeleteOrder(ctx context.Context, o *models.Order) (err error) {
	orderKey := fmt.Sprintf("order:%d", o.ID)

	// сначала надо получить
	order := &schemas.Order{}
	err = c.client.HGetAll(ctx, orderKey).Scan(order)
	if err != nil {
		return err
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
	err = c.client.ZRem(ctx, fmt.Sprintf("user:%d:orders", order.UserID), orderKey).Err()
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

		// находится в пвз, только если заказ принят или же возвращен
		if isStored {
			if order.Status != models.StatusAccepted &&
				order.Status != models.StatusReturned {
				continue
			}
		}

		list = append(list, order)
	}

	if len(list) == 0 {
		return nil, nil
	}

	return list, nil
}

// GetOrders
func (c *Cache) GetOrdersForDelivery(ctx context.Context, orderIDs []uint64) (list []*models.Order, err error) {
	for _, orderID := range orderIDs {
		order, err := c.GetOrder(ctx, &models.Order{ID: orderID})
		if err != nil {
			return nil, err
		}

		if order == nil {
			continue
		}

		list = append(list, order)
	}

	return list, nil
}
