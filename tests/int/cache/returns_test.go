package cache_test

import (
	"time"

	orders "gitlab.ozon.dev/yuweebix/homework-1/gen/orders/v1/proto"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
)

func (s *CacheSuite) TestGetRetuns_Success() {
	ordersSet := []*models.Order{
		{
			ID:        1,
			User:      &models.User{ID: 1},
			Expiry:    time.Date(9999, 1, 1, 0, 0, 0, 0, time.UTC),
			Status:    models.StatusReturned,
			Hash:      "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
			CreatedAt: now,
			Cost:      1,
			Weight:    1,
			Packaging: models.PackagingType(orders.PackagingType_PACKAGING_BOX.String()),
		},
		{
			ID:        2,
			User:      &models.User{ID: 1},
			Expiry:    time.Date(9999, 1, 1, 0, 0, 0, 0, time.UTC),
			Status:    models.StatusReturned,
			Hash:      "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
			CreatedAt: now,
			Cost:      1,
			Weight:    1,
			Packaging: models.PackagingType(orders.PackagingType_PACKAGING_BOX.String()),
		},
	}

	for _, o := range ordersSet {
		err := s.cache.SetOrder(s.ctx, o)
		s.Require().NoError(err)
	}

	list, err := s.cache.GetReturns(s.ctx, 2, 0)
	s.Require().NoError(err)

	for i := range list {
		s.Equal(ordersSet[i], list[i])
	}
}
