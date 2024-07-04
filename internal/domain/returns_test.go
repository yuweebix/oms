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

type ReturnsSuite struct {
	suite.Suite
}

const (
	returnByAllowedTime = day * 2
)

// TestReturnsSuite запускает все returns unit-тесты
func TestReturnsSuite(t *testing.T) {
	suite.Run(t, new(ReturnsSuite))
}

func (s *ReturnsSuite) SetUpTest() (_domain *domain.Domain, _storage *mocks.MockStorage, _threading *mocks.MockThreading) {
	_storage = mocks.NewMockStorage(s.T())
	_threading = mocks.NewMockThreading(s.T())
	_domain = domain.NewDomain(_storage, _threading)
	return
}

func (s *ReturnsSuite) TestAcceptReturn_Success() {
	s.T().Parallel()

	// должны совпадать юзера, статус должен быть Delivered и ReturnBy время не должно быть превышено
	order := &models.Order{ID: 1, User: &models.User{ID: 1}}
	returnOrder := &models.Order{ID: 1, Status: models.StatusDelivered, User: &models.User{ID: 1}, ReturnBy: time.Now().Add(returnByAllowedTime)}

	domain, storage, _ := s.SetUpTest()

	storage.EXPECT().RunTx(mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(ctx context.Context, opts models.TxOptions, fn func(context.Context) error) error {
		return fn(ctx)
	})
	storage.EXPECT().GetOrder(mock.Anything, order).Return(returnOrder, nil)
	storage.EXPECT().UpdateOrder(mock.Anything, mock.Anything).Return(nil)

	err := domain.AcceptReturn(context.Background(), order)

	s.NoError(err)
	s.Equal(models.StatusReturned, returnOrder.Status)
	s.NotEmpty(returnOrder.Hash)
}

func (s *ReturnsSuite) TestAcceptReturn_StatusInvalid() {
	s.T().Parallel()

	order := &models.Order{ID: 1, User: &models.User{ID: 1}}
	returnOrder := &models.Order{ID: 1, Status: models.StatusAccepted, User: &models.User{ID: 1}, ReturnBy: time.Now().Add(24 * time.Hour)} // не доставили...

	domain, storage, _ := s.SetUpTest()

	storage.EXPECT().RunTx(mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(ctx context.Context, opts models.TxOptions, fn func(context.Context) error) error {
		return fn(ctx)
	})
	storage.EXPECT().GetOrder(mock.Anything, order).Return(returnOrder, nil)

	err := domain.AcceptReturn(context.Background(), order)

	s.EqualError(err, e.ErrStatusInvalid.Error())
}

func (s *ReturnsSuite) TestAcceptReturn_OrderExpired() {
	s.T().Parallel()

	order := &models.Order{ID: 1, User: &models.User{ID: 1}}
	returnOrder := &models.Order{ID: 1, Status: models.StatusDelivered, User: &models.User{ID: 1}, ReturnBy: time.Now().Add(-returnByAllowedTime)} // долго думал юзер, не вернул во время

	domain, storage, _ := s.SetUpTest()

	storage.EXPECT().RunTx(mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(ctx context.Context, opts models.TxOptions, fn func(context.Context) error) error {
		return fn(ctx)
	})
	storage.EXPECT().GetOrder(mock.Anything, order).Return(returnOrder, nil)

	err := domain.AcceptReturn(context.Background(), order)

	s.EqualError(err, e.ErrOrderExpired.Error())
}

func (s *ReturnsSuite) TestAcceptReturn_UserInvalid() {
	s.T().Parallel()

	order := &models.Order{ID: 1, User: &models.User{ID: 1}} // пытаемся принять возврат от первого юзера, а заказ то был от второго
	returnOrder := &models.Order{ID: 1, Status: models.StatusDelivered, User: &models.User{ID: 2}, ReturnBy: time.Now().Add(returnByAllowedTime)}
	domain, storage, _ := s.SetUpTest()

	storage.EXPECT().RunTx(mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(ctx context.Context, opts models.TxOptions, fn func(context.Context) error) error {
		return fn(ctx)
	})
	storage.EXPECT().GetOrder(mock.Anything, order).Return(returnOrder, nil)

	err := domain.AcceptReturn(context.Background(), order)

	s.EqualError(err, e.ErrUserInvalid.Error())
}

// returns list tests

// имхо, смысла тестировать не вижу, потому что только лишь можно проверить, правильно ли работает мок
// просто тестируем, написал ли прогер if err != nil
