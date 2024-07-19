package cache_test

import (
	"time"

	orders "gitlab.ozon.dev/yuweebix/homework-1/gen/orders/v1/proto"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
)

func (s *CacheSuite) TestSetOrder_Success() {
	orderSet := &models.Order{
		ID:        1,
		User:      &models.User{ID: 1},
		Expiry:    time.Date(9999, 1, 1, 0, 0, 0, 0, time.UTC),
		Status:    models.StatusAccepted,
		Hash:      "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		CreatedAt: now,
		Cost:      1,
		Weight:    1,
		Packaging: models.PackagingType(orders.PackagingType_PACKAGING_BOX.String()),
	}
	orderGet := &models.Order{
		ID: 1,
	}

	err := s.cache.SetOrder(s.ctx, orderSet)
	s.Require().NoError(err)
	orderGet, err = s.cache.GetOrder(s.ctx, orderGet)
	s.Require().NoError(err)

	s.Equal(orderSet, orderGet)
}

func (s *CacheSuite) TestSetOrders_SetTwice() {
	orderSet := &models.Order{
		ID:        1,
		User:      &models.User{ID: 1},
		Expiry:    time.Date(9999, 1, 1, 0, 0, 0, 0, time.UTC),
		Status:    models.StatusAccepted,
		Hash:      "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		CreatedAt: now,
		Cost:      1,
		Weight:    1,
		Packaging: models.PackagingType(orders.PackagingType_PACKAGING_BOX.String()),
	}
	orderGet := &models.Order{
		ID: 1,
	}

	err := s.cache.SetOrder(s.ctx, orderSet)
	s.Require().NoError(err)
	err = s.cache.SetOrder(s.ctx, orderSet)
	s.Require().NoError(err)
	orderGet, err = s.cache.GetOrder(s.ctx, orderGet)
	s.Require().NoError(err)

	s.Equal(orderSet, orderGet)
}

func (s *CacheSuite) TestDeleteOrder_Success() {
	orderSet := &models.Order{
		ID:        1,
		User:      &models.User{ID: 1},
		Expiry:    time.Date(9999, 1, 1, 0, 0, 0, 0, time.UTC),
		Status:    models.StatusAccepted,
		Hash:      "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		CreatedAt: now,
		Cost:      1,
		Weight:    1,
		Packaging: models.PackagingType(orders.PackagingType_PACKAGING_BOX.String()),
	}
	orderGet := &models.Order{
		ID: 1,
	}

	err := s.cache.SetOrder(s.ctx, orderSet)
	s.Require().NoError(err)
	err = s.cache.DeleteOrder(s.ctx, orderGet)
	s.Require().NoError(err)
	orderGet, err = s.cache.GetOrder(s.ctx, orderGet)
	s.Require().NoError(err)

	s.Nil(orderGet)
}

func (s *CacheSuite) TestDeleteOrder_NotExists() {
	orderGet := &models.Order{
		ID: 1,
	}

	err := s.cache.DeleteOrder(s.ctx, orderGet)
	s.Require().NoError(err)
	orderGet, err = s.cache.GetOrder(s.ctx, orderGet)
	s.Require().NoError(err)

	s.Nil(orderGet)
}

func (s *CacheSuite) TestGetOrders_Empty() {
	list, err := s.cache.GetOrders(s.ctx, 1, 1, 0, false)
	s.Require().NoError(err)
	s.Empty(list)
}

func (s *CacheSuite) TestGetOrders_Success() {
	ordersSet := []*models.Order{
		{
			ID:        1,
			User:      &models.User{ID: 1},
			Expiry:    time.Date(9999, 1, 1, 0, 0, 0, 0, time.UTC),
			Status:    models.StatusAccepted,
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
			Status:    models.StatusAccepted,
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

	list, err := s.cache.GetOrders(s.ctx, 1, 2, 0, false)
	s.Require().NoError(err)

	for i := range list {
		s.Equal(ordersSet[i], list[len(list)-i-1])
	}
}

func (s *CacheSuite) TestGetOrdersForDelivery_Empty() {
	list, err := s.cache.GetOrdersForDelivery(s.ctx, []uint64{})
	s.Require().NoError(err)

	s.Empty(list)
}

func (s *CacheSuite) TestGetOrdersForDelivery_Success() {
	ordersSet := []*models.Order{
		{
			ID:        1,
			User:      &models.User{ID: 1},
			Expiry:    time.Date(9999, 1, 1, 0, 0, 0, 0, time.UTC),
			Status:    models.StatusAccepted,
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
			Status:    models.StatusAccepted,
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

	list, err := s.cache.GetOrdersForDelivery(s.ctx, []uint64{1, 2})
	s.Require().NoError(err)

	for i := range list {
		s.Equal(ordersSet[i], list[i])
	}
}
