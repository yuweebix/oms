package repository_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/repository"
)

var (
	ErrOrderNotCreated    = errors.New("could not create order")
	ErrOrderNotFound      = errors.New("could not find order")
	ErrOrderAlreadyExists = errors.New("order already exists")
	ErrOrderNotDeleted    = errors.New("could not delete order")
	ErrOrderNotUpdated    = errors.New("could not update order")
)

type OrdersSuite struct {
	suite.Suite
	repository *repository.Repository
	ctx        context.Context
}

var (
	now = time.Now().UTC().Truncate(time.Second)
)

// TestOrdersSuite запускает все orders int-тесты
func TestOrdersSuite(t *testing.T) {
	suite.Run(t, new(OrdersSuite))
}

func (s *OrdersSuite) SetupSuite() {
	err := godotenv.Load("../../.env")
	if err != nil {
		s.FailNowf("Error loading .env file", err.Error())
	}

	connString := os.Getenv("DATABASE_TEST_URL") // при локальном тестировании закомментировать эту строчку
	// connString := os.Getenv("DATABASE_LOCAL_TEST_URL") // и разкомментировать эту
	if connString == "" {
		s.FailNow("rror reading database url from the .env")
	}

	s.ctx = context.Background()

	s.repository, err = repository.NewRepository(s.ctx, connString)
	if err != nil {
		s.FailNowf("Error connecting to the database", err.Error())
	}
}

func (s *OrdersSuite) TearDownSuite() {
	s.repository.Close()
}

// create order

// описание: создаем заказ и затем получаем его из базы данных
// ожидаемый результат: создание заказа проходит успешно, а поля заказов совпадают
func (s *OrdersSuite) TestCreateOrder_Success() {
	orderCreate := &models.Order{
		ID:        1,
		User:      &models.User{ID: 1},
		Expiry:    time.Date(10000, 0, 0, 0, 0, 0, 0, time.UTC),
		Status:    models.StatusAccepted,
		Hash:      "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		CreatedAt: now,
		Cost:      1,
		Weight:    1,
		Packaging: "bag",
	}
	orderGet := &models.Order{
		ID: 1,
	}

	err := s.repository.RunTxWithRollback(s.ctx, models.TxOptions{}, func(ctxTX context.Context) error {
		err := s.repository.CreateOrder(ctxTX, orderCreate)
		if err != nil {
			return ErrOrderNotCreated
		}

		orderGet, err = s.repository.GetOrder(ctxTX, orderGet)
		if err != nil {
			return ErrOrderNotFound
		}

		return nil
	})

	s.NoError(err)
	s.Equal(orderCreate, orderGet)
}

// описание: пытаемся создать один и тот же заказ дважды
// ожидаемый результат: первая попытка проходит успешно, а вторая - возвращает ошибку, потому что заказ уже существует
func (s *OrdersSuite) TestCreateOrder_AlreadyExists() {
	order := &models.Order{
		ID:        1,
		User:      &models.User{ID: 1},
		Expiry:    time.Date(10000, 0, 0, 0, 0, 0, 0, time.UTC),
		Status:    models.StatusAccepted,
		Hash:      "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		CreatedAt: now,
		Cost:      1,
		Weight:    1,
		Packaging: "bag",
	}

	err := s.repository.RunTxWithRollback(s.ctx, models.TxOptions{}, func(ctxTX context.Context) error {
		err := s.repository.CreateOrder(ctxTX, order)
		if err != nil {
			return ErrOrderNotCreated
		}

		err = s.repository.CreateOrder(ctxTX, order)
		if err != nil {
			return ErrOrderAlreadyExists
		}

		return nil
	})

	s.EqualError(err, ErrOrderAlreadyExists.Error())
}

// delete order

// описание: создаем заказ, удаляем его и затем пытается получить его
// ожидаемый результат: удаление заказа проходит успешно, и он больше не доступен для получения.
func (s *OrdersSuite) TestDeleteOrder_Success() {
	orderCreate := &models.Order{
		ID:        1,
		User:      &models.User{ID: 1},
		Expiry:    time.Date(1, 1, 1, 1, 1, 1, 1, time.UTC), // уже просрочен
		Status:    models.StatusAccepted,
		Hash:      "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		CreatedAt: now,
		Cost:      1,
		Weight:    1,
		Packaging: "bag",
	}
	orderGet := &models.Order{
		ID: 1,
	}

	err := s.repository.RunTxWithRollback(s.ctx, models.TxOptions{}, func(ctxTX context.Context) error {
		err := s.repository.CreateOrder(ctxTX, orderCreate)
		if err != nil {
			return ErrOrderNotCreated
		}

		err = s.repository.DeleteOrder(ctxTX, orderGet)
		if err != nil {
			return ErrOrderNotDeleted
		}

		orderGet, err = s.repository.GetOrder(ctxTX, orderGet)
		if err != nil {
			return ErrOrderNotFound
		}

		return nil
	})

	s.EqualError(err, ErrOrderNotFound.Error())
	s.Empty(orderGet)
}

// описание: пытаемся удалить заказ, которого нет в бд
// ожидаемый результат: возвращается ошибка, ведь удалять нечего
func (s *OrdersSuite) TestDeleteOrder_NotExists() {
	orderGet := &models.Order{
		ID: 1,
	}

	err := s.repository.RunTxWithRollback(s.ctx, models.TxOptions{}, func(ctxTX context.Context) error {
		err := s.repository.DeleteOrder(ctxTX, orderGet)
		if err != nil {
			return ErrOrderNotDeleted
		}

		return nil
	})

	s.EqualError(err, ErrOrderNotDeleted.Error())
}

// update order

// описание: создаем заказ, затем обновляем его и получаем обновленный заказ
// ожидаемый результат: обновление заказа проходит успешно и сохраняет все изменения
func (s *OrdersSuite) TestUpdateOrder_Success() {
	orderCreate := &models.Order{
		ID:        1,
		User:      &models.User{ID: 1},
		Expiry:    time.Date(10000, 0, 0, 0, 0, 0, 0, time.UTC),
		Status:    models.StatusAccepted,
		Hash:      "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		CreatedAt: now,
		Cost:      1,
		Weight:    1,
		Packaging: "bag",
	}
	orderUpdate := &models.Order{
		ID:        1,
		User:      &models.User{ID: 1},
		Expiry:    time.Date(10000, 0, 0, 0, 0, 0, 0, time.UTC),
		ReturnBy:  time.Date(10001, 0, 0, 0, 0, 0, 0, time.UTC),
		Status:    models.StatusDelivered,
		Hash:      "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		CreatedAt: now,
		Cost:      1,
		Weight:    1,
		Packaging: "bag",
	}
	orderGet := &models.Order{
		ID: 1,
	}

	err := s.repository.RunTxWithRollback(s.ctx, models.TxOptions{}, func(ctxTX context.Context) error {
		err := s.repository.CreateOrder(ctxTX, orderCreate)
		if err != nil {
			return ErrOrderNotCreated
		}

		err = s.repository.UpdateOrder(ctxTX, orderUpdate)
		if err != nil {
			return ErrOrderNotUpdated
		}

		orderGet, err = s.repository.GetOrder(ctxTX, orderGet)
		if err != nil {
			return ErrOrderNotFound
		}

		return nil
	})

	s.NoError(err)
	s.Equal(orderUpdate, orderGet)
}

// описание: пытаемся обновить заказ, который не существует в базе данных
// ожидаемый результат: попытка обновления возвращает ошибку, ничего не обновляется и заказа все ещё не существует после попытки изменить данные
func (s *OrdersSuite) TestUpdateOrder_NotExists() {
	orderUpdate := &models.Order{
		ID:        1,
		User:      &models.User{ID: 1},
		Expiry:    time.Date(10000, 0, 0, 0, 0, 0, 0, time.UTC),
		ReturnBy:  time.Date(10001, 1, 1, 1, 1, 1, 1, time.UTC),
		Status:    models.StatusDelivered,
		Hash:      "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		CreatedAt: now,
		Cost:      1,
		Weight:    1,
		Packaging: "bag",
	}

	err := s.repository.RunTxWithRollback(s.ctx, models.TxOptions{}, func(ctxTX context.Context) error {
		err := s.repository.UpdateOrder(ctxTX, orderUpdate)
		if err != nil {
			return ErrOrderNotUpdated
		}

		return nil
	})

	s.EqualError(err, ErrOrderNotUpdated.Error())
}

// get order

// описание: пытаемся получить заказ, которого нет в бд
// ожидаемый результат: возвращается ошибка, что заказ не найден
// P.S. успешное получение заказа через GetOrder уже описано в TestCreateOrder_Success, и имело бы идентичную логику
func (s *OrdersSuite) TestGetOrder_NotFound() {
	orderGet := &models.Order{
		ID: 1,
	}

	err := s.repository.RunTxWithRollback(s.ctx, models.TxOptions{}, func(ctxTX context.Context) error {
		_, err := s.repository.GetOrder(ctxTX, orderGet)
		if err != nil {
			return ErrOrderNotFound
		}

		return nil
	})

	s.EqualError(err, ErrOrderNotFound.Error())
}

// get orders

// описание: создаем два заказа с разными статусами и получаем их
// ожидаемый результат: заказы возвращаются в правильном порядке по дате создания (DESC)
func (s *OrdersSuite) TestGetOrders_Standard() {
	orderAccepted := &models.Order{
		ID:        1,
		User:      &models.User{ID: 1},
		Expiry:    time.Date(10000, 1, 1, 1, 1, 1, 0, time.UTC),
		Status:    models.StatusAccepted,
		Hash:      "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		CreatedAt: now,
		Cost:      1,
		Weight:    1,
		Packaging: "bag",
	}
	orderDelivered := &models.Order{
		ID:        2,
		User:      &models.User{ID: 1},
		Expiry:    time.Date(10000, 1, 1, 1, 1, 1, 0, time.UTC),
		ReturnBy:  time.Date(10001, 1, 1, 1, 1, 1, 0, time.UTC),
		Status:    models.StatusDelivered,
		Hash:      "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		CreatedAt: now.Add(time.Second), // чтобы был позже чем orderAccepted
		Cost:      1,
		Weight:    1,
		Packaging: "bag",
	}

	orders := make([]*models.Order, 0, 2)
	err := s.repository.RunTxWithRollback(s.ctx, models.TxOptions{}, func(ctxTX context.Context) error {
		err := s.repository.CreateOrder(ctxTX, orderAccepted)
		if err != nil {
			return ErrOrderNotCreated
		}
		err = s.repository.CreateOrder(ctxTX, orderDelivered)
		if err != nil {
			return ErrOrderNotCreated
		}

		orders, err = s.repository.GetOrders(ctxTX, 1, 2, 0, false)
		if err != nil {
			return ErrOrderNotFound
		}

		return nil
	})

	s.NoError(err)
	s.Len(orders, 2)
	s.Equal(orderDelivered, orders[0])
	s.Equal(orderAccepted, orders[1])
}

// описание: теперь с лимитом
// ожидаемый результат: возвращается только один заказ
func (s *OrdersSuite) TestGetOrders_Limit() {
	orderAccepted := &models.Order{
		ID:        1,
		User:      &models.User{ID: 1},
		Expiry:    time.Date(10000, 1, 1, 1, 1, 1, 0, time.UTC),
		Status:    models.StatusAccepted,
		Hash:      "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		CreatedAt: now,
		Cost:      1,
		Weight:    1,
		Packaging: "bag",
	}
	orderDelivered := &models.Order{
		ID:        2,
		User:      &models.User{ID: 1},
		Expiry:    time.Date(10000, 1, 1, 1, 1, 1, 0, time.UTC),
		ReturnBy:  time.Date(10001, 1, 1, 1, 1, 1, 0, time.UTC),
		Status:    models.StatusDelivered,
		Hash:      "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		CreatedAt: now.Add(time.Second),
		Cost:      1,
		Weight:    1,
		Packaging: "bag",
	}

	orders := make([]*models.Order, 0, 2)
	err := s.repository.RunTxWithRollback(s.ctx, models.TxOptions{}, func(ctxTX context.Context) error {
		err := s.repository.CreateOrder(ctxTX, orderAccepted)
		if err != nil {
			return ErrOrderNotCreated
		}
		err = s.repository.CreateOrder(ctxTX, orderDelivered)
		if err != nil {
			return ErrOrderNotCreated
		}

		orders, err = s.repository.GetOrders(ctxTX, 1, 1, 0, false)
		if err != nil {
			return ErrOrderNotFound
		}

		return nil
	})

	s.NoError(err)
	s.Len(orders, 1)
	s.Equal(orderDelivered, orders[0])
}

// описание: теперь смещаем
// ожидаемый результат: возвращается только второй заказ
func (s *OrdersSuite) TestGetOrders_Offset() {
	orderAccepted := &models.Order{
		ID:        1,
		User:      &models.User{ID: 1},
		Expiry:    time.Date(10000, 1, 1, 1, 1, 1, 0, time.UTC),
		Status:    models.StatusAccepted,
		Hash:      "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		CreatedAt: now,
		Cost:      1,
		Weight:    1,
		Packaging: "bag",
	}
	orderDelivered := &models.Order{
		ID:        2,
		User:      &models.User{ID: 1},
		Expiry:    time.Date(10000, 1, 1, 1, 1, 1, 0, time.UTC),
		ReturnBy:  time.Date(10001, 1, 1, 1, 1, 1, 0, time.UTC),
		Status:    models.StatusDelivered,
		Hash:      "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		CreatedAt: now.Add(time.Second),
		Cost:      1,
		Weight:    1,
		Packaging: "bag",
	}

	orders := make([]*models.Order, 0, 2)
	err := s.repository.RunTxWithRollback(s.ctx, models.TxOptions{}, func(ctxTX context.Context) error {
		err := s.repository.CreateOrder(ctxTX, orderAccepted)
		if err != nil {
			return ErrOrderNotCreated
		}
		err = s.repository.CreateOrder(ctxTX, orderDelivered)
		if err != nil {
			return ErrOrderNotCreated
		}

		orders, err = s.repository.GetOrders(ctxTX, 1, 2, 1, false)
		if err != nil {
			return ErrOrderNotFound
		}

		return nil
	})

	s.NoError(err)
	s.Len(orders, 1)
	s.Equal(orderAccepted, orders[0])
}

// описание: создаем два заказа и получаем только то, что хранится (StatusAccepted)
// ожидаемый результат: возвращается только первый заказ, ведь второй заказ доставили
func (s *OrdersSuite) TestGetOrders_IsStored() {
	orderAccepted := &models.Order{
		ID:        1,
		User:      &models.User{ID: 1},
		Expiry:    time.Date(10000, 1, 1, 1, 1, 1, 0, time.UTC),
		Status:    models.StatusAccepted,
		Hash:      "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		CreatedAt: now,
		Cost:      1,
		Weight:    1,
		Packaging: "bag",
	}
	orderDelivered := &models.Order{
		ID:        2,
		User:      &models.User{ID: 1},
		Expiry:    time.Date(10000, 1, 1, 1, 1, 1, 0, time.UTC),
		ReturnBy:  time.Date(10001, 1, 1, 1, 1, 1, 0, time.UTC),
		Status:    models.StatusDelivered,
		Hash:      "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		CreatedAt: now.Add(time.Second),
		Cost:      1,
		Weight:    1,
		Packaging: "bag",
	}

	orders := make([]*models.Order, 0, 2)
	err := s.repository.RunTxWithRollback(s.ctx, models.TxOptions{}, func(ctxTX context.Context) error {
		err := s.repository.CreateOrder(ctxTX, orderAccepted)
		if err != nil {
			return ErrOrderNotCreated
		}
		err = s.repository.CreateOrder(ctxTX, orderDelivered)
		if err != nil {
			return ErrOrderNotCreated
		}

		orders, err = s.repository.GetOrders(ctxTX, 1, 2, 0, true)
		if err != nil {
			return ErrOrderNotFound
		}

		return nil
	})

	s.NoError(err)
	s.Len(orders, 1)
	s.Equal(orderAccepted, orders[0])
}

// get orders for delivery

// описание: создаем два заказа и пытаемся получить их по их id
// ожидаемый результат: оба заказа находятся и возвращаются
func (s *OrdersSuite) TestGetOrdersForDelivery_BothMatching() {
	order1 := &models.Order{
		ID:        1,
		User:      &models.User{ID: 1},
		Expiry:    time.Date(10000, 1, 1, 1, 1, 1, 0, time.UTC),
		Status:    models.StatusAccepted,
		Hash:      "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		CreatedAt: now,
		Cost:      1,
		Weight:    1,
		Packaging: "bag",
	}
	order2 := &models.Order{
		ID:        2,
		User:      &models.User{ID: 1},
		Expiry:    time.Date(10000, 1, 1, 1, 1, 1, 0, time.UTC),
		ReturnBy:  time.Date(10001, 1, 1, 1, 1, 1, 0, time.UTC),
		Status:    models.StatusDelivered,
		Hash:      "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		CreatedAt: now.Add(time.Second),
		Cost:      1,
		Weight:    1,
		Packaging: "bag",
	}

	orders := make([]*models.Order, 0, 2)
	err := s.repository.RunTxWithRollback(s.ctx, models.TxOptions{}, func(ctxTX context.Context) error {
		err := s.repository.CreateOrder(ctxTX, order1)
		if err != nil {
			return err
		}
		err = s.repository.CreateOrder(ctxTX, order2)
		if err != nil {
			return err
		}

		orders, err = s.repository.GetOrdersForDelivery(ctxTX, []uint64{1, 2})
		if err != nil {
			return ErrOrderNotFound
		}

		return nil
	})

	s.NoError(err)
	s.Len(orders, 2)
	s.Equal(order1, orders[0])
	s.Equal(order2, orders[1])
}

// описание: теперь пытаемся получить их по одному существующему и одному несуществующему id
// ожидаемый результат: возвращается только один существующий заказ
func (s *OrdersSuite) TestGetOrdersForDelivery_OneMatching() {
	order1 := &models.Order{
		ID:        1,
		User:      &models.User{ID: 1},
		Expiry:    time.Date(10000, 1, 1, 1, 1, 1, 0, time.UTC),
		Status:    models.StatusAccepted,
		Hash:      "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		CreatedAt: now,
		Cost:      1,
		Weight:    1,
		Packaging: "bag",
	}
	order2 := &models.Order{
		ID:        2,
		User:      &models.User{ID: 1},
		Expiry:    time.Date(10000, 1, 1, 1, 1, 1, 0, time.UTC),
		ReturnBy:  time.Date(10001, 1, 1, 1, 1, 1, 0, time.UTC),
		Status:    models.StatusDelivered,
		Hash:      "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		CreatedAt: now.Add(time.Second),
		Cost:      1,
		Weight:    1,
		Packaging: "bag",
	}

	orders := make([]*models.Order, 0, 2)
	err := s.repository.RunTxWithRollback(s.ctx, models.TxOptions{}, func(ctxTX context.Context) error {
		err := s.repository.CreateOrder(ctxTX, order1)
		if err != nil {
			return err
		}
		err = s.repository.CreateOrder(ctxTX, order2)
		if err != nil {
			return err
		}

		orders, err = s.repository.GetOrdersForDelivery(ctxTX, []uint64{1, 3})
		if err != nil {
			return ErrOrderNotFound
		}

		return nil
	})

	s.NoError(err)
	s.Len(orders, 1)
	s.Equal(order1, orders[0])
}

// описание: теперь ни одного заказа нет в бд
// ожидаемый результат: ничего не находятся
func (s *OrdersSuite) TestGetOrdersForDelivery_NoneMatching() {
	order1 := &models.Order{
		ID:        1,
		User:      &models.User{ID: 1},
		Expiry:    time.Date(10000, 1, 1, 1, 1, 1, 0, time.UTC),
		Status:    models.StatusAccepted,
		Hash:      "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		CreatedAt: now,
		Cost:      1,
		Weight:    1,
		Packaging: "bag",
	}
	order2 := &models.Order{
		ID:        2,
		User:      &models.User{ID: 1},
		Expiry:    time.Date(10000, 1, 1, 1, 1, 1, 0, time.UTC),
		ReturnBy:  time.Date(10001, 1, 1, 1, 1, 1, 0, time.UTC),
		Status:    models.StatusDelivered,
		Hash:      "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		CreatedAt: now.Add(time.Second),
		Cost:      1,
		Weight:    1,
		Packaging: "bag",
	}

	orders := make([]*models.Order, 0, 2)
	err := s.repository.RunTxWithRollback(s.ctx, models.TxOptions{}, func(ctxTX context.Context) error {
		err := s.repository.CreateOrder(ctxTX, order1)
		if err != nil {
			return err
		}
		err = s.repository.CreateOrder(ctxTX, order2)
		if err != nil {
			return err
		}

		orders, err = s.repository.GetOrdersForDelivery(ctxTX, []uint64{3, 4})
		if err != nil {
			return ErrOrderNotFound
		}

		return nil
	})

	s.NoError(err)
	s.Len(orders, 0)
}
