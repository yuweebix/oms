package repository_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/repository"
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

// TestReturnsSuite запускает все return int-тесты
func TestReturnsSuite(t *testing.T) {
	suite.Run(t, new(ReturnsSuite))
}

func (s *ReturnsSuite) SetupSuite() {
	err := godotenv.Load("../../.env")
	if err != nil {
		s.FailNowf("Error loading .env file", err.Error())
	}

	connString := os.Getenv("DATABASE_TEST_URL") // при локальном тестировании закомментировать эту строчку
	// connString := os.Getenv("DATABASE_LOCAL_TEST_URL") // и разкомментировать эту
	if connString == "" {
		s.FailNow("Error reading database url from the .env")
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

func (s *ReturnsSuite) AfterTest(suiteName, testName string) {
	err := s.repository.DeleteAllOrders(s.ctx)
	if err != nil {
		s.Failf("Error truncating table orders", err.Error())
	}
}

// get returns

// т.к. тесты похожи я их все опишу здесь
// описание: создаем два возвратных заказа и один доставленный заказ, затем выполняем различные запросы с разными комбинациями лимитов и смещений (первая цифра в тесте - лимит, вторая - смещение)
// ожидаемый результат: возвращаемся ожидаемое количество возвратных заказов в зависимости от переданных параметров запроса в правильном порядке
func (s *ReturnsSuite) TestGetReturns_20() {
	err := s.repository.CreateOrder(s.ctx, orderReturned1)
	s.Require().NoError(err)

	err = s.repository.CreateOrder(s.ctx, orderReturned2)
	s.Require().NoError(err)

	err = s.repository.CreateOrder(s.ctx, orderDelivered)
	s.Require().NoError(err)

	orders, err := s.repository.GetReturns(s.ctx, 2, 0)
	s.Require().NoError(err)

	s.Len(orders, 2)
	s.Equal(orderReturned1, orders[0])
	s.Equal(orderReturned2, orders[1])
}

func (s *ReturnsSuite) TestGetReturns_10() {
	err := s.repository.CreateOrder(s.ctx, orderReturned1)
	s.Require().NoError(err)

	err = s.repository.CreateOrder(s.ctx, orderReturned2)
	s.Require().NoError(err)

	err = s.repository.CreateOrder(s.ctx, orderDelivered)
	s.Require().NoError(err)

	orders, err := s.repository.GetReturns(s.ctx, 1, 0)
	s.Require().NoError(err)

	s.Len(orders, 1)
	s.Equal(orderReturned1, orders[0])
}

func (s *ReturnsSuite) TestGetReturns_11() {
	err := s.repository.CreateOrder(s.ctx, orderReturned1)
	s.Require().NoError(err)

	err = s.repository.CreateOrder(s.ctx, orderReturned2)
	s.Require().NoError(err)

	err = s.repository.CreateOrder(s.ctx, orderDelivered)
	s.Require().NoError(err)

	orders, err := s.repository.GetReturns(s.ctx, 1, 1)
	s.Require().NoError(err)

	s.Len(orders, 1)
	s.Equal(orderReturned2, orders[0])
}

func (s *ReturnsSuite) TestGetReturns_00() {
	err := s.repository.CreateOrder(s.ctx, orderReturned1)
	s.Require().NoError(err)

	err = s.repository.CreateOrder(s.ctx, orderReturned2)
	s.Require().NoError(err)

	err = s.repository.CreateOrder(s.ctx, orderDelivered)
	s.Require().NoError(err)

	orders, err := s.repository.GetReturns(s.ctx, 0, 0)
	s.Require().NoError(err)

	s.Len(orders, 0)
}
