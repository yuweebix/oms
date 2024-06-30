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

type OrdersSuite struct {
	suite.Suite
}

func (s *OrdersSuite) SetUpTest() (_domain *domain.Domain, _storage *mocks.MockStorage, _threading *mocks.MockThreading) {
	_storage = mocks.NewMockStorage(s.T())
	_threading = mocks.NewMockThreading(s.T())
	_domain = domain.NewDomain(_storage, _threading)
	return
}

// orders accept tests

// TestOrdersSuite запускает все orders unit-тесты
func TestOrdersSuite(t *testing.T) {
	suite.Run(t, new(OrdersSuite))
}

func (s *OrdersSuite) TestAcceptOrder() {
	s.T().Parallel()

	// arrange
	order := &models.Order{
		ID:        1,
		User:      &models.User{ID: 1},
		Expiry:    time.Now().Add(day),
		Cost:      1,
		Weight:    1,
		Packaging: "box",
	}

	// act
	domain, storage, _ := s.SetUpTest()

	storage.EXPECT().CreateOrder(mock.Anything, mock.Anything).Return(nil)

	err := domain.AcceptOrder(context.Background(), order)

	// assert
	s.NoError(err)
	s.Equal(models.StatusAccepted, order.Status)
	s.NotZero(order.CreatedAt)
	s.NotEmpty(order.Hash)
}

func (s *OrdersSuite) TestAcceptOrder_Expired() {
	s.T().Parallel()

	// arrange
	order := &models.Order{
		ID:        1,
		User:      &models.User{ID: 1},
		Expiry:    time.Now().Add(-day), // срок хранения превышен
		Cost:      1,
		Weight:    1,
		Packaging: "box",
	}

	// act
	domain, _, _ := s.SetUpTest()

	err := domain.AcceptOrder(context.Background(), order)

	// assert
	s.Error(err)
	s.Equal(e.ErrOrderExpired, err)
}

func (s *OrdersSuite) TestAcceptOrder_InvalidPackaging() {
	s.T().Parallel()

	// arrange
	order := &models.Order{
		ID:        1,
		User:      &models.User{ID: 1},
		Expiry:    time.Now().Add(day),
		Cost:      1,
		Weight:    1,
		Packaging: "bucket", // https://www.youtube.com/watch?v=L8FmQoSFys0
	}

	// act
	domain, _, _ := s.SetUpTest()

	err := domain.AcceptOrder(context.Background(), order)

	// assert
	s.Error(err)
	s.Equal(e.ErrPackagingInvalid, err)
}

func (s *OrdersSuite) TestAcceptOrder_TooHeavy() {
	s.T().Parallel()

	// arrange
	order := &models.Order{
		ID:        1,
		User:      &models.User{ID: 1},
		Expiry:    time.Now().Add(day),
		Cost:      1,
		Weight:    1_000_000, // тяжело...
		Packaging: "box",
	}

	// act
	domain, _, _ := s.SetUpTest()

	err := domain.AcceptOrder(context.Background(), order)

	// assert
	s.Error(err)
	s.Equal(e.ErrOrderTooHeavy, err)
}

// orders return tests

func (s *OrdersSuite) TestReturnOrder_StatusReturned() {
	// arrange
	order := &models.Order{ID: 1}

	// act
	domain, storage, _ := s.SetUpTest()

	storage.EXPECT().RunTx(mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(ctx context.Context, opts models.TxOptions, fn func(context.Context) error) error {
		return fn(ctx)
	})
	storage.EXPECT().GetOrder(mock.Anything, mock.Anything).RunAndReturn(func(ctx context.Context, o *models.Order) (*models.Order, error) {
		o.Status = models.StatusReturned // клиент вернул -> можно вернуть курьеру
		return o, nil
	})
	storage.EXPECT().DeleteOrder(mock.Anything, mock.Anything).Return(nil)

	err := domain.ReturnOrder(context.Background(), order)

	// assert
	s.NoError(err)
}

func (s *OrdersSuite) TestReturnOrder_Expired() {
	// arrange
	order := &models.Order{ID: 1}

	// act
	domain, storage, _ := s.SetUpTest()

	storage.EXPECT().RunTx(mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(ctx context.Context, opts models.TxOptions, fn func(context.Context) error) error {
		return fn(ctx)
	})
	storage.EXPECT().GetOrder(mock.Anything, mock.Anything).RunAndReturn(func(ctx context.Context, o *models.Order) (*models.Order, error) {
		o.Expiry = time.Now().Add(-day) // просрочился -> можно вернуть курьеру
		return o, nil
	})
	storage.EXPECT().DeleteOrder(mock.Anything, mock.Anything).Return(nil)

	err := domain.ReturnOrder(context.Background(), order)

	// assert
	s.NoError(err)
}
func (s *OrdersSuite) TestReturnOrder_NotExpired() {
	// arrange
	order := &models.Order{ID: 1}

	// act
	domain, storage, _ := s.SetUpTest()

	storage.EXPECT().RunTx(mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(ctx context.Context, opts models.TxOptions, fn func(context.Context) error) error {
		return fn(ctx)
	})
	storage.EXPECT().GetOrder(mock.Anything, mock.Anything).RunAndReturn(func(ctx context.Context, o *models.Order) (*models.Order, error) {
		o.Expiry = time.Now().Add(day) // не просрочился -> нельзя вернуть курьеру
		return o, nil
	})

	err := domain.ReturnOrder(context.Background(), order)

	// assert
	s.Error(err)
	s.Equal(e.ErrOrderNotExpired, err)
}
