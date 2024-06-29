package domain_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/domain"
	e "gitlab.ozon.dev/yuweebix/homework-1/internal/domain/errors"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/domain/mocks"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
)

const (
	day                 = time.Hour * 24
	returnByAllowedTime = day * 2
)

// OrdersSuite - абстрактный Test Suite, что будут эмбедить все тесты
type OrdersSuite struct {
	suite.Suite
	domain    *domain.Domain
	storage   *mocks.MockStorage
	threading *mocks.MockThreading
}

// BeforeTest рефрешит моки для каждого теста, изолируя их
func (s *OrdersSuite) BeforeTest(suiteName, testName string) {
	s.storage = mocks.NewMockStorage(s.T())
	s.threading = mocks.NewMockThreading(s.T())
	s.domain = domain.NewDomain(s.storage, s.threading)

	switch suiteName {
	case "OrdersReturnSuite":
		s.storage.EXPECT().RunTx(mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(ctx context.Context, opts models.TxOptions, fn func(context.Context) error) error {
			return fn(ctx)
		})

		switch testName {
		case "TestReturnOrder_StatusReturned":
			s.storage.EXPECT().GetOrder(mock.Anything, mock.Anything).RunAndReturn(func(ctx context.Context, o *models.Order) (*models.Order, error) {
				o.Status = models.StatusReturned
				return o, nil
			})
			s.storage.EXPECT().DeleteOrder(mock.Anything, mock.Anything).Return(nil)

		case "TestReturnOrder_Expired":
			s.storage.EXPECT().GetOrder(mock.Anything, mock.Anything).RunAndReturn(func(ctx context.Context, o *models.Order) (*models.Order, error) {
				o.Expiry = time.Now().Add(-day)
				return o, nil
			})
			s.storage.EXPECT().DeleteOrder(mock.Anything, mock.Anything).Return(nil)
		case "TestReturnOrder_NotExpired":
			s.storage.EXPECT().GetOrder(mock.Anything, mock.Anything).RunAndReturn(func(ctx context.Context, o *models.Order) (*models.Order, error) {
				o.Expiry = time.Now().Add(day)
				return o, nil
			})
		}
	}
}

// orders accept tests

// OrdersAcceptSuite - Test Suite метода AcceptOrder
type OrdersAcceptSuite struct {
	OrdersSuite
}

// TestOrdersAcceptSuite запускает все тесты метода AcceptOrder
func TestOrdersAcceptSuite(t *testing.T) {
	suite.Run(t, new(OrdersAcceptSuite))
}

func (s *OrdersAcceptSuite) TestAcceptOrder() {
	s.T().Parallel()

	// arrange
	order := &models.Order{
		ID:        1001,
		User:      &models.User{ID: 123},
		Expiry:    time.Now().Add(day),
		Cost:      322,
		Weight:    1,
		Packaging: "box",
	}

	// act
	s.storage.EXPECT().CreateOrder(mock.Anything, order).Return(nil)
	err := s.domain.AcceptOrder(context.Background(), order)

	// assert
	s.NoError(err)
	s.Equal(models.StatusAccepted, order.Status)
	s.NotZero(order.CreatedAt)
	s.NotEmpty(order.Hash)
}

func (s *OrdersAcceptSuite) TestAcceptOrder_Expired() {
	s.T().Parallel()

	// arrange
	order := &models.Order{
		ID:        1002,
		User:      &models.User{ID: 123},
		Expiry:    time.Now().Add(-day), // срок хранения превышен
		Cost:      322,
		Weight:    1,
		Packaging: "box",
	}

	// act
	err := s.domain.AcceptOrder(context.Background(), order)

	// assert
	s.Error(err)
	s.Equal(e.ErrOrderExpired, err)
}

func (s *OrdersAcceptSuite) TestAcceptOrder_InvalidPackaging() {
	s.T().Parallel()

	// arrange
	order := &models.Order{
		ID:        1003,
		User:      &models.User{ID: 123},
		Expiry:    time.Now().Add(day),
		Cost:      322,
		Weight:    1,
		Packaging: "bucket",
	}

	// act
	err := s.domain.AcceptOrder(context.Background(), order)

	// assert
	s.Error(err)
	s.Equal(e.ErrPackagingInvalid, err)
}

func (s *OrdersAcceptSuite) TestAcceptOrder_TooHeavy() {
	s.T().Parallel()

	// arrange
	order := &models.Order{
		ID:        1004,
		User:      &models.User{ID: 123},
		Expiry:    time.Now().Add(day),
		Cost:      322,
		Weight:    1_000_000, // тяжело...
		Packaging: "box",
	}

	// act
	err := s.domain.AcceptOrder(context.Background(), order)

	// assert
	s.Error(err)
	s.Equal(e.ErrOrderTooHeavy, err)
}

// orders return tests

// OrdersReturnSuite - Test Suite метода AcceptOrder
type OrdersReturnSuite struct {
	OrdersSuite
}

// TestOrdersReturnSuite запускает все тесты метода AcceptOrder
func TestOrdersReturnSuite(t *testing.T) {
	suite.Run(t, new(OrdersReturnSuite))
}

func (s *OrdersReturnSuite) TestReturnOrder_StatusReturned() {
	// arrange
	order := &models.Order{ID: 1005}

	// act
	err := s.domain.ReturnOrder(context.Background(), order)

	// assert
	s.NoError(err)
}

func (s *OrdersReturnSuite) TestReturnOrder_Expired() {
	// arrange
	order := &models.Order{ID: 1006}

	// act
	err := s.domain.ReturnOrder(context.Background(), order)

	// assert
	s.NoError(err)
}
func (s *OrdersReturnSuite) TestReturnOrder_NotExpired() {
	// arrange
	order := &models.Order{ID: 1007}

	// act
	err := s.domain.ReturnOrder(context.Background(), order)

	// assert
	s.Error(err)
	s.Equal(e.ErrOrderNotExpired, err)
}
