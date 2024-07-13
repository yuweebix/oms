package domain_test

import (
	"context"
	"time"

	"github.com/stretchr/testify/mock"
	e "gitlab.ozon.dev/yuweebix/homework-1/internal/domain/errors"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
)

// orders accept tests

func (s *DomainSuite) TestAcceptOrder_Success() {
	s.T().Parallel()

	order := &models.Order{
		ID:        1,
		User:      &models.User{ID: 1},
		Expiry:    time.Now().Add(day),
		Cost:      1,
		Weight:    1,
		Packaging: "box",
	}

	domain, storage := s.SetUpTest()

	storage.EXPECT().CreateOrder(mock.Anything, mock.Anything).Return(nil)

	err := domain.AcceptOrder(context.Background(), order)

	s.NoError(err)
	s.Equal(models.StatusAccepted, order.Status)
	s.NotEmpty(order.Hash)
}

func (s *DomainSuite) TestAcceptOrder_Expired() {
	s.T().Parallel()

	order := &models.Order{
		ID:        1,
		User:      &models.User{ID: 1},
		Expiry:    time.Now().Add(-day), // срок хранения превышен
		Cost:      1,
		Weight:    1,
		Packaging: "box",
	}

	domain, _ := s.SetUpTest()

	err := domain.AcceptOrder(context.Background(), order)

	s.EqualError(err, e.ErrOrderExpired.Error())
}

func (s *DomainSuite) TestAcceptOrder_InvalidPackaging() {
	s.T().Parallel()

	order := &models.Order{
		ID:        1,
		User:      &models.User{ID: 1},
		Expiry:    time.Now().Add(day),
		Cost:      1,
		Weight:    1,
		Packaging: "bucket", // https://www.youtube.com/watch?v=L8FmQoSFys0
	}

	domain, _ := s.SetUpTest()

	err := domain.AcceptOrder(context.Background(), order)

	s.EqualError(err, e.ErrPackagingInvalid.Error())
}

func (s *DomainSuite) TestAcceptOrder_TooHeavy() {
	s.T().Parallel()

	order := &models.Order{
		ID:        1,
		User:      &models.User{ID: 1},
		Expiry:    time.Now().Add(day),
		Cost:      1,
		Weight:    1_000_000, // тяжело...
		Packaging: "box",
	}

	domain, _ := s.SetUpTest()

	err := domain.AcceptOrder(context.Background(), order)

	s.EqualError(err, e.ErrOrderTooHeavy.Error())
}

// orders return tests

func (s *DomainSuite) TestReturnOrder_Success_StatusReturned() {

	order := &models.Order{ID: 1}

	domain, storage := s.SetUpTest()

	storage.EXPECT().RunTx(mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(ctx context.Context, opts models.TxOptions, fn func(context.Context) error) error {
		return fn(ctx)
	})
	storage.EXPECT().GetOrder(mock.Anything, mock.Anything).RunAndReturn(func(ctx context.Context, o *models.Order) (*models.Order, error) {
		o.Status = models.StatusReturned // клиент вернул -> можно вернуть курьеру
		return o, nil
	})
	storage.EXPECT().DeleteOrder(mock.Anything, mock.Anything).Return(nil)

	err := domain.ReturnOrder(context.Background(), order)

	s.NoError(err)
}

func (s *DomainSuite) TestReturnOrder_Expired() {

	order := &models.Order{ID: 1}

	domain, storage := s.SetUpTest()

	storage.EXPECT().RunTx(mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(ctx context.Context, opts models.TxOptions, fn func(context.Context) error) error {
		return fn(ctx)
	})
	storage.EXPECT().GetOrder(mock.Anything, mock.Anything).RunAndReturn(func(ctx context.Context, o *models.Order) (*models.Order, error) {
		o.Expiry = time.Now().Add(-day) // просрочился -> можно вернуть курьеру
		return o, nil
	})
	storage.EXPECT().DeleteOrder(mock.Anything, mock.Anything).Return(nil)

	err := domain.ReturnOrder(context.Background(), order)

	s.NoError(err)
}
func (s *DomainSuite) TestReturnOrder_NotExpired() {

	order := &models.Order{ID: 1}

	domain, storage := s.SetUpTest()

	storage.EXPECT().RunTx(mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(ctx context.Context, opts models.TxOptions, fn func(context.Context) error) error {
		return fn(ctx)
	})
	storage.EXPECT().GetOrder(mock.Anything, mock.Anything).RunAndReturn(func(ctx context.Context, o *models.Order) (*models.Order, error) {
		o.Expiry = time.Now().Add(day) // не просрочился -> нельзя вернуть курьеру
		return o, nil
	})

	err := domain.ReturnOrder(context.Background(), order)

	s.EqualError(err, e.ErrOrderNotExpired.Error())
}

// orders list tests

// имхо, смысла тестировать не вижу, потому что только лишь можно проверить, правильно ли работает мок
// просто тестируем, написал ли прогер if err != nil

// orders deliver tests

func (s *DomainSuite) TestDeliverOrders_Success() {
	s.T().Parallel()

	orderIDs := []uint64{1, 2, 3}
	orders := []*models.Order{
		{ID: 1, Status: models.StatusAccepted, User: &models.User{ID: 1}, Expiry: time.Now().Add(day)}, // количество заказов совпадает
		{ID: 2, Status: models.StatusAccepted, User: &models.User{ID: 1}, Expiry: time.Now().Add(day)}, // ернулись заказы одного и того же юзера
		{ID: 3, Status: models.StatusAccepted, User: &models.User{ID: 1}, Expiry: time.Now().Add(day)}, // срок хранения не превышен
	}

	domain, storage := s.SetUpTest()

	storage.EXPECT().RunTx(mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(ctx context.Context, opts models.TxOptions, fn func(context.Context) error) error {
		return fn(ctx)
	})
	storage.EXPECT().GetOrdersForDelivery(mock.Anything, orderIDs).Return(orders, nil)
	storage.EXPECT().UpdateOrder(mock.Anything, mock.Anything).Return(nil).Times(len(orderIDs))

	err := domain.DeliverOrders(context.Background(), orderIDs)

	s.NoError(err)
	for _, order := range orders {
		s.Equal(models.StatusDelivered, order.Status)
		s.NotZero(order.ReturnBy)
		s.NotEmpty(order.Hash)
	}
}

func (s *DomainSuite) TestDeliverOrders_EmptyOrderIDs() {
	s.T().Parallel()

	domain, _ := s.SetUpTest()

	err := domain.DeliverOrders(context.Background(), []uint64{}) // пусто...

	s.Error(err)
	s.Equal(e.ErrEmpty, err)
}

func (s *DomainSuite) TestDeliverOrders_OrderNotFound() {
	s.T().Parallel()

	orderIDs := []uint64{1, 2, 3}

	domain, storage := s.SetUpTest()

	storage.EXPECT().RunTx(mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(ctx context.Context, opts models.TxOptions, fn func(context.Context) error) error {
		return fn(ctx)
	})
	storage.EXPECT().GetOrdersForDelivery(mock.Anything, orderIDs).Return([]*models.Order{}, nil) // опять пусто... только теперь ничего не вернули

	err := domain.DeliverOrders(context.Background(), orderIDs)

	s.EqualError(err, e.ErrOrderNotFound.Error())
}

func (s *DomainSuite) TestDeliverOrders_StatusInvalid() {
	s.T().Parallel()

	orderIDs := []uint64{1}
	orders := []*models.Order{
		{ID: 1, Status: models.StatusReturned, User: &models.User{ID: 1}, Expiry: time.Now().Add(day)}, // статус не тот
	}

	domain, storage := s.SetUpTest()

	storage.EXPECT().RunTx(mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(ctx context.Context, opts models.TxOptions, fn func(context.Context) error) error {
		return fn(ctx)
	})
	storage.EXPECT().GetOrdersForDelivery(mock.Anything, orderIDs).Return(orders, nil)

	err := domain.DeliverOrders(context.Background(), orderIDs)

	s.EqualError(err, e.ErrStatusInvalid.Error())
}

func (s *DomainSuite) TestDeliverOrders_UserInvalid() {
	s.T().Parallel()

	orderIDs := []uint64{1, 2}
	orders := []*models.Order{
		{ID: 1, Status: models.StatusAccepted, User: &models.User{ID: 1}, Expiry: time.Now().Add(day)}, // два различных юзера
		{ID: 2, Status: models.StatusAccepted, User: &models.User{ID: 2}, Expiry: time.Now().Add(day)}, // а мы можем только заказы одному передать
	}

	domain, storage := s.SetUpTest()

	storage.EXPECT().RunTx(mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(ctx context.Context, opts models.TxOptions, fn func(context.Context) error) error {
		return fn(ctx)
	})
	storage.EXPECT().GetOrdersForDelivery(mock.Anything, orderIDs).Return(orders, nil)

	err := domain.DeliverOrders(context.Background(), orderIDs)

	s.EqualError(err, e.ErrUserInvalid.Error())
}

func (s *DomainSuite) TestDeliverOrders_OrderExpired() {
	s.T().Parallel()

	orderIDs := []uint64{1}
	orders := []*models.Order{
		{ID: 1, Status: models.StatusAccepted, User: &models.User{ID: 1}, Expiry: time.Now().Add(-day)}, // срок превышен
	}

	domain, storage := s.SetUpTest()

	storage.EXPECT().RunTx(mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(ctx context.Context, opts models.TxOptions, fn func(context.Context) error) error {
		return fn(ctx)
	})
	storage.EXPECT().GetOrdersForDelivery(mock.Anything, orderIDs).Return(orders, nil)

	err := domain.DeliverOrders(context.Background(), orderIDs)

	s.EqualError(err, e.ErrOrderExpired.Error())
}
