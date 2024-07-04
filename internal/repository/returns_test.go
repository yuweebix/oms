package repository_test

import (
	"time"

	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
)

// get returns

// т.к. тесты похожи я их все опишу здесь
// описание: создаем два возвратных заказа и один доставленный заказ, затем выполняем различные запросы с разными комбинациями лимитов и смещений (первая цифра в тесте - лимит, вторая - смещение)
// ожидаемый результат: возвращаемся ожидаемое количество возвратных заказов в зависимости от переданных параметров запроса в правильном порядке
func (s *RepositorySuite) TestGetReturns() {
	var (
		orderReturned1 = &models.Order{
			ID:        1,
			User:      &models.User{ID: 1},
			Expiry:    time.Date(10000, 0, 0, 0, 0, 0, 0, time.UTC),
			Status:    models.StatusReturned,
			Hash:      "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
			ReturnBy:  time.Date(10001, 0, 0, 0, 0, 0, 0, time.UTC),
			CreatedAt: now,
			Cost:      1,
			Weight:    1,
			Packaging: "bag",
		}
		orderReturned2 = &models.Order{
			ID:        2,
			User:      &models.User{ID: 1},
			Expiry:    time.Date(10000, 0, 0, 0, 0, 0, 0, time.UTC),
			Status:    models.StatusReturned,
			Hash:      "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
			ReturnBy:  time.Date(10001, 0, 0, 0, 0, 0, 0, time.UTC),
			CreatedAt: now.Add(time.Second),
			Cost:      1,
			Weight:    1,
			Packaging: "bag",
		}
		orderDelivered = &models.Order{
			ID:        3,
			User:      &models.User{ID: 1},
			Expiry:    time.Date(10000, 0, 0, 0, 0, 0, 0, time.UTC),
			Status:    models.StatusDelivered,
			Hash:      "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
			ReturnBy:  time.Date(10001, 0, 0, 0, 0, 0, 0, time.UTC),
			CreatedAt: now.Add(2 * time.Second),
			Cost:      1,
			Weight:    1,
			Packaging: "bag",
		}
	)

	err := s.repository.CreateOrder(s.ctx, orderReturned1)
	s.Require().NoError(err)

	err = s.repository.CreateOrder(s.ctx, orderReturned2)
	s.Require().NoError(err)

	err = s.repository.CreateOrder(s.ctx, orderDelivered)
	s.Require().NoError(err)

	tests := []struct {
		limit          uint64
		offset         uint64
		expectedLen    int
		expectedOrders []*models.Order
	}{
		{2, 0, 2, []*models.Order{orderReturned1, orderReturned2}},
		{1, 0, 1, []*models.Order{orderReturned1}},
		{1, 1, 1, []*models.Order{orderReturned2}},
		{0, 0, 0, []*models.Order{}},
	}

	for _, tt := range tests {
		orders, err := s.repository.GetReturns(s.ctx, tt.limit, tt.offset)
		s.Require().NoError(err)

		s.Len(orders, tt.expectedLen)
		for i, expectedOrder := range tt.expectedOrders {
			s.Equal(expectedOrder, orders[i])
		}
	}
}
