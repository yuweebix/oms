package domain_test

import (
	"context"
	"time"

	"github.com/stretchr/testify/mock"
	e "gitlab.ozon.dev/yuweebix/homework-1/internal/domain/errors"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
)

func (s *DomainSuite) TestAcceptReturn_Success() {
	s.T().Parallel()

	// должны совпадать юзера, статус должен быть Delivered и ReturnBy время не должно быть превышено
	order := &models.Order{ID: 1, User: &models.User{ID: 1}}
	returnOrder := &models.Order{ID: 1, Status: models.StatusDelivered, User: &models.User{ID: 1}, ReturnBy: time.Now().Add(returnByAllowedTime)}

	domain, storage, cache := s.SetupTest()

	storage.EXPECT().RunTx(mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(ctx context.Context, opts models.TxOptions, fn func(context.Context) error) error {
		return fn(ctx)
	})
	storage.EXPECT().GetOrder(mock.Anything, order).Return(returnOrder, nil)
	storage.EXPECT().UpdateOrder(mock.Anything, mock.Anything).Return(nil)

	cache.EXPECT().GetOrder(mock.Anything, mock.Anything).Return(nil, nil)
	cache.EXPECT().SetOrder(mock.Anything, mock.Anything).Return(nil)

	err := domain.AcceptReturn(context.Background(), order)

	s.NoError(err)
	s.Equal(models.StatusReturned, returnOrder.Status)
	s.NotEmpty(returnOrder.Hash)
}

func (s *DomainSuite) TestAcceptReturn_StatusInvalid() {
	s.T().Parallel()

	order := &models.Order{ID: 1, User: &models.User{ID: 1}}
	returnOrder := &models.Order{ID: 1, Status: models.StatusAccepted, User: &models.User{ID: 1}, ReturnBy: time.Now().Add(24 * time.Hour)} // не доставили...

	domain, storage, cache := s.SetupTest()

	storage.EXPECT().RunTx(mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(ctx context.Context, opts models.TxOptions, fn func(context.Context) error) error {
		return fn(ctx)
	})
	storage.EXPECT().GetOrder(mock.Anything, order).Return(returnOrder, nil)

	cache.EXPECT().GetOrder(mock.Anything, mock.Anything).Return(nil, nil)

	err := domain.AcceptReturn(context.Background(), order)

	s.EqualError(err, e.ErrStatusInvalid.Error())
}

func (s *DomainSuite) TestAcceptReturn_OrderExpired() {
	s.T().Parallel()

	order := &models.Order{ID: 1, User: &models.User{ID: 1}}
	returnOrder := &models.Order{ID: 1, Status: models.StatusDelivered, User: &models.User{ID: 1}, ReturnBy: time.Now().Add(-returnByAllowedTime)} // долго думал юзер, не вернул во время

	domain, storage, cache := s.SetupTest()

	storage.EXPECT().RunTx(mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(ctx context.Context, opts models.TxOptions, fn func(context.Context) error) error {
		return fn(ctx)
	})
	storage.EXPECT().GetOrder(mock.Anything, order).Return(returnOrder, nil)

	cache.EXPECT().GetOrder(mock.Anything, mock.Anything).Return(nil, nil)

	err := domain.AcceptReturn(context.Background(), order)

	s.EqualError(err, e.ErrOrderExpired.Error())
}

func (s *DomainSuite) TestAcceptReturn_UserInvalid() {
	s.T().Parallel()

	order := &models.Order{ID: 1, User: &models.User{ID: 1}} // пытаемся принять возврат от первого юзера, а заказ то был от второго
	returnOrder := &models.Order{ID: 1, Status: models.StatusDelivered, User: &models.User{ID: 2}, ReturnBy: time.Now().Add(returnByAllowedTime)}
	domain, storage, cache := s.SetupTest()

	storage.EXPECT().RunTx(mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(ctx context.Context, opts models.TxOptions, fn func(context.Context) error) error {
		return fn(ctx)
	})
	storage.EXPECT().GetOrder(mock.Anything, order).Return(returnOrder, nil)

	cache.EXPECT().GetOrder(mock.Anything, mock.Anything).Return(nil, nil)

	err := domain.AcceptReturn(context.Background(), order)

	s.EqualError(err, e.ErrUserInvalid.Error())
}

// returns list tests

// имхо, смысла тестировать не вижу, потому что только лишь можно проверить, правильно ли работает мок
// просто тестируем, написал ли прогер if err != nil
