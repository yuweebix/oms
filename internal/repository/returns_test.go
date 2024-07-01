package repository_test

import (
	"context"
	"errors"
	"log"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/repository"
)

var (
	ErrReturnNotFound = errors.New("could not find return")
)

type ReturnsSuite struct {
	suite.Suite
	repository *repository.Repository
	ctx        context.Context
}

var (
	orderReturned1 = &models.Order{
		ID:        1,
		User:      &models.User{ID: 1},
		Expiry:    time.Date(10000, 1, 1, 1, 1, 1, 0, time.UTC),
		Status:    models.StatusReturned,
		Hash:      "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		ReturnBy:  time.Date(10001, 1, 1, 1, 1, 1, 0, time.UTC),
		CreatedAt: now,
		Cost:      1,
		Weight:    1,
		Packaging: "bag",
	}
	orderReturned2 = &models.Order{
		ID:        2,
		User:      &models.User{ID: 1},
		Expiry:    time.Date(10000, 1, 1, 1, 1, 1, 0, time.UTC),
		Status:    models.StatusReturned,
		Hash:      "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		ReturnBy:  time.Date(10001, 1, 1, 1, 1, 1, 0, time.UTC),
		CreatedAt: now.Add(time.Second),
		Cost:      1,
		Weight:    1,
		Packaging: "bag",
	}
	orderDelivered = &models.Order{
		ID:        3,
		User:      &models.User{ID: 1},
		Expiry:    time.Date(10000, 1, 1, 1, 1, 1, 0, time.UTC),
		Status:    models.StatusDelivered,
		Hash:      "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		ReturnBy:  time.Date(10001, 1, 1, 1, 1, 1, 0, time.UTC),
		CreatedAt: now.Add(2 * time.Second),
		Cost:      1,
		Weight:    1,
		Packaging: "bag",
	}
)

// TestReturnsSuite запускает все return int-тесты
func TestReturnsSuite(t *testing.T) {
	suite.Run(t, new(ReturnsSuite))
}

func (s *ReturnsSuite) SetupSuite() {
	err := godotenv.Load("../../.env")
	if err != nil {
		s.FailNowf("Error loading .env file", err.Error())
	}

	connString := os.Getenv("DATABASE_TEST_URL")
	if connString == "" {
		log.Fatalf("Error reading environment variable DATABASE_TEST_URL")
	}

	s.ctx = context.Background()

	s.repository, err = repository.NewRepository(s.ctx, connString)
	if err != nil {
		s.FailNowf("Error connecting to the database", err.Error())
	}
}

func (s *ReturnsSuite) TearDownSuite() {
	s.repository.Close()
}

// get returns

// т.к. тесты похожи я их все опишу здесь
// описание: создаем два возвратных заказа и один доставленный заказ, затем выполняем различные запросы с разными комбинациями лимитов и смещений (первая цифра в тесте - лимит, вторая - смещение)
// ожидаемый результат: возвращаемся ожидаемое количество возвратных заказов в зависимости от переданных параметров запроса в правильном порядке
func (s *ReturnsSuite) TestGetReturns_20() {
	orders := make([]*models.Order, 0, 2)
	err := s.repository.RunTxWithRollback(s.ctx, models.TxOptions{}, func(ctxTX context.Context) error {
		err := s.repository.CreateOrder(ctxTX, orderReturned1)
		if err != nil {
			return err
		}
		err = s.repository.CreateOrder(ctxTX, orderReturned2)
		if err != nil {
			return err
		}
		err = s.repository.CreateOrder(ctxTX, orderDelivered)
		if err != nil {
			return err
		}

		orders, err = s.repository.GetReturns(ctxTX, 2, 0)
		if err != nil {
			return ErrReturnNotFound
		}

		return nil
	})

	s.NoError(err)
	s.Len(orders, 2)
	s.Equal(orderReturned1.ID, orders[0].ID)
	s.Equal(orderReturned2.ID, orders[1].ID)
}

func (s *ReturnsSuite) TestGetReturns_10() {
	orders := make([]*models.Order, 0, 2)
	err := s.repository.RunTxWithRollback(s.ctx, models.TxOptions{}, func(ctxTX context.Context) error {
		err := s.repository.CreateOrder(ctxTX, orderReturned1)
		if err != nil {
			return err
		}
		err = s.repository.CreateOrder(ctxTX, orderReturned2)
		if err != nil {
			return err
		}
		err = s.repository.CreateOrder(ctxTX, orderDelivered)
		if err != nil {
			return err
		}

		orders, err = s.repository.GetReturns(ctxTX, 1, 0)
		if err != nil {
			return ErrReturnNotFound
		}

		return nil
	})

	s.NoError(err)
	s.Len(orders, 1)
	s.Equal(orderReturned1.ID, orders[0].ID)
}

func (s *ReturnsSuite) TestGetReturns_11() {
	orders := make([]*models.Order, 0, 2)
	err := s.repository.RunTxWithRollback(s.ctx, models.TxOptions{}, func(ctxTX context.Context) error {
		err := s.repository.CreateOrder(ctxTX, orderReturned1)
		if err != nil {
			return err
		}
		err = s.repository.CreateOrder(ctxTX, orderReturned2)
		if err != nil {
			return err
		}
		err = s.repository.CreateOrder(ctxTX, orderDelivered)
		if err != nil {
			return err
		}

		orders, err = s.repository.GetReturns(ctxTX, 1, 1)
		if err != nil {
			return ErrReturnNotFound
		}

		return nil
	})

	s.NoError(err)
	s.Len(orders, 1)
	s.Equal(orderReturned2.ID, orders[0].ID)
}

func (s *ReturnsSuite) TestGetReturns_00() {
	orders := make([]*models.Order, 0, 2)
	err := s.repository.RunTxWithRollback(s.ctx, models.TxOptions{}, func(ctxTX context.Context) error {
		err := s.repository.CreateOrder(ctxTX, orderReturned1)
		if err != nil {
			return err
		}
		err = s.repository.CreateOrder(ctxTX, orderReturned2)
		if err != nil {
			return err
		}
		err = s.repository.CreateOrder(ctxTX, orderDelivered)
		if err != nil {
			return err
		}

		orders, err = s.repository.GetReturns(ctxTX, 0, 0)
		if err != nil {
			return ErrReturnNotFound
		}

		return nil
	})

	s.NoError(err)
	s.Len(orders, 0)
}
