package cache

import (
	"context"
	"fmt"
	"sort"
	"time"

	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
)

// CreateOrder создаёт новую запись в кеше о заказе, а также добавляет его в соответствующие списки
func (c *Cache) CreateOrder(ctx context.Context, o *models.Order) {
	c.mu.Lock()
	defer c.mu.Unlock()

	orderKey := fmt.Sprintf("order:%d", o.ID)

	// TTL будет время просрочки заказа
	duration := time.Until(o.Expiry)
	c.store.Set(orderKey, o, &duration)

	// добавим его с соответвующие списки
	listIDs := c.getOrdersList(fmt.Sprintf("orders:by_status:%s", o.Status))
	listIDs.Add(o.ID)
	listIDs = c.getOrdersList(fmt.Sprintf("orders:by_user:%d", o.User.ID))
	listIDs.Add(o.ID)
}

// GetOrder возвращает заказ из кеша; если его нет в кеше - nil
func (c *Cache) GetOrder(ctx context.Context, o *models.Order) *models.Order {
	c.mu.RLock()
	defer c.mu.RUnlock()

	orderKey := fmt.Sprintf("order:%d", o.ID)

	orderInterface, ok := c.store.Get(orderKey)
	if !ok {
		return nil
	}

	return orderInterface.(*models.Order)
}

// UpdateOrder обновляет существующую запись в кеше
func (c *Cache) UpdateOrder(ctx context.Context, o *models.Order) {
	c.mu.Lock()
	defer c.mu.Unlock()

	orderKey := fmt.Sprintf("order:%d", o.ID)

	oldInterface, ok := c.store.Get(orderKey)
	if !ok { // её нет, нечего обновлять
		return
	}
	oldOrder := oldInterface.(*models.Order)

	// удалим запись из списка фильтрации по статусу (у него больше не этот статус)
	listIDs := c.getOrdersList(fmt.Sprintf("orders:by_status:%s", oldOrder.Status))
	listIDs.Rem(o.ID)

	// если ReturnBy каким-то образом пустой, то также просто оставим Expiry
	duration := time.Until(o.Expiry)
	switch o.ReturnBy {
	case time.Time{}:
		duration = time.Until(o.ReturnBy)
	default:
	}

	c.store.Set(orderKey, o, &duration)

	// статус обновился -> нужно добавить в соответсвующий список
	listIDs = c.getOrdersList(fmt.Sprintf("orders:by_status:%s", o.Status))
	listIDs.Add(o.ID)
}

// DeleteOrder удаляет заказ и его записи в соответствующих списках
func (c *Cache) DeleteOrder(ctx context.Context, o *models.Order) {
	c.mu.Lock()
	defer c.mu.Unlock()

	orderKey := fmt.Sprintf("order:%d", o.ID)

	c.store.Del(orderKey)

	listIDs := c.getOrdersList(fmt.Sprintf("orders:by_status:%s", o.Status))
	listIDs.Rem(o.ID)
	listIDs = c.getOrdersList(fmt.Sprintf("orders:by_user:%d", o.User.ID))
	listIDs.Rem(o.ID)
}

// GetOrders возвращает список заказов заданного пользователя; если список невалидный возвращается nil
func (c *Cache) GetOrders(ctx context.Context, userID uint64, limit uint64, offset uint64, isStored bool) []*models.Order {
	c.mu.RLock()
	defer c.mu.RUnlock()

	listIDs := c.getOrdersList(fmt.Sprintf("orders:by_user:%d", userID))

	// сначала все заказы получим
	orders := make([]*models.Order, 0, listIDs.Len())
	for id := range listIDs.Get() {
		orderKey := fmt.Sprintf("order:%d", id)

		orderInterface, ok := c.store.Get(orderKey)
		if !ok {
			continue
		}

		orders = append(orders, orderInterface.(*models.Order))
	}

	// если не соответствуют заказы в кеше со списком, то возвращаем nil
	// это может произойти, когда была удалёна запись заказа из кеша из-за ttl, а список не изменился
	// чтобы это исправить, нужно удалить просроченный заказ(ы), т.е. вернуть их курьеру
	if len(orders) != int(listIDs.Len()) {
		return nil
	}

	// фильтруем по статусу
	filteredOrders := make([]*models.Order, 0, len(orders))
	switch isStored {
	case false:
		filteredOrders = orders
	case true:
		for _, order := range orders {
			if order.Status != models.StatusAccepted &&
				order.Status != models.StatusReturned {
				continue
			}
			filteredOrders = append(filteredOrders, order)
		}
	}

	// сортируем от самого нового к самому старому
	sort.Slice(filteredOrders, func(i, j int) bool {
		return filteredOrders[i].CreatedAt.After(filteredOrders[j].CreatedAt)
	})

	// пагинация
	len := uint64(len(filteredOrders))
	start := min(offset, len)
	stop := min(offset+limit, len)
	return filteredOrders[start:stop]
}

// GetOrdersForDelivery возвращает список заказов по данным идентификаторам заказов
func (c *Cache) GetOrdersForDelivery(ctx context.Context, orderIDs []uint64) []*models.Order {
	orders := make([]*models.Order, 0, len(orderIDs))
	for id := range orderIDs {
		orderKey := fmt.Sprintf("order:%d", id)

		orderInterface, ok := c.store.Get(orderKey)
		if !ok {
			continue
		}

		orders = append(orders, orderInterface.(*models.Order))
	}

	return orders
}
