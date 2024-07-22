package cache

import (
	"context"
	"fmt"
	"sort"

	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
)

// GetReturns возвращает список заказов, что вернули в пвз; если список невалидный, возвращается nil
func (c *Cache) GetReturns(ctx context.Context, limit uint64, offset uint64) []*models.Order {
	c.mu.RLock()
	defer c.mu.RUnlock()

	listIDs := c.getOrdersList(fmt.Sprintf("orders:by_status:%s", models.StatusReturned))

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

	// та же ситуация, что и в GetOrders
	if len(orders) != int(listIDs.Len()) {
		return nil
	}

	// сортим
	sort.Slice(orders, func(i, j int) bool {
		return orders[i].CreatedAt.Before(orders[j].CreatedAt)
	})

	// пагинируем
	len := uint64(len(orders))
	start := min(offset, len)
	stop := min(offset+limit, len)
	return orders[start:stop]
}
