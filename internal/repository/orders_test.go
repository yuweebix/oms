package repository_test

import (
	"time"

	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
)

// create order

// описание: создаем заказ и затем получаем его из базы данных
// ожидаемый результат: создание заказа проходит успешно, а поля заказов совпадают
func (s *RepositorySuite) TestCreateOrder_Success() {
	orderCreate := &models.Order{
		ID:        1,
		User:      &models.User{ID: 1},
		Expiry:    time.Date(10000, 1, 1, 0, 0, 0, 0, time.UTC),
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

	err := s.repository.CreateOrder(s.ctx, orderCreate)
	s.Require().NoError(err)

	orderGet, err = s.repository.GetOrder(s.ctx, orderGet)
	s.Require().NoError(err)

	s.Equal(orderCreate, orderGet)
}

// описание: пытаемся создать один и тот же заказ дважды
// ожидаемый результат: первая попытка проходит успешно, а вторая - возвращает ошибку, потому что заказ уже существует
func (s *RepositorySuite) TestCreateOrder_AlreadyExists() {
	order := &models.Order{
		ID:        1,
		User:      &models.User{ID: 1},
		Expiry:    time.Date(10000, 1, 1, 0, 0, 0, 0, time.UTC),
		Status:    models.StatusAccepted,
		Hash:      "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		CreatedAt: now,
		Cost:      1,
		Weight:    1,
		Packaging: "bag",
	}

	err := s.repository.CreateOrder(s.ctx, order)
	s.Require().NoError(err)

	err = s.repository.CreateOrder(s.ctx, order)
	s.Error(err)
}

// delete order

// описание: создаем заказ, удаляем его и затем пытается получить его
// ожидаемый результат: удаление заказа проходит успешно, и он больше не доступен для получения.
func (s *RepositorySuite) TestDeleteOrder_Success() {
	orderCreate := &models.Order{
		ID:        1,
		User:      &models.User{ID: 1},
		Expiry:    time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC), // уже просрочен
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

	err := s.repository.CreateOrder(s.ctx, orderCreate)
	s.Require().NoError(err)

	err = s.repository.DeleteOrder(s.ctx, orderGet)
	s.Require().NoError(err)

	orderGet, err = s.repository.GetOrder(s.ctx, orderGet)
	s.Error(err)
	s.Empty(orderGet)
}

// описание: пытаемся удалить заказ, которого нет в бд
// ожидаемый результат: возвращается ошибка, ведь удалять нечего
func (s *RepositorySuite) TestDeleteOrder_NotExists() {
	orderGet := &models.Order{
		ID: 1,
	}

	err := s.repository.DeleteOrder(s.ctx, orderGet)
	s.Error(err)
}

// update order

// описание: создаем заказ, затем обновляем его и получаем обновленный заказ
// ожидаемый результат: обновление заказа проходит успешно и сохраняет все изменения
func (s *RepositorySuite) TestUpdateOrder_Success() {
	orderCreate := &models.Order{
		ID:        1,
		User:      &models.User{ID: 1},
		Expiry:    time.Date(10000, 1, 1, 0, 0, 0, 0, time.UTC),
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
		Expiry:    time.Date(10000, 1, 1, 0, 0, 0, 0, time.UTC),
		ReturnBy:  time.Date(10001, 1, 1, 0, 0, 0, 0, time.UTC),
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

	err := s.repository.CreateOrder(s.ctx, orderCreate)
	s.Require().NoError(err)

	err = s.repository.UpdateOrder(s.ctx, orderUpdate)
	s.Require().NoError(err)

	orderGet, err = s.repository.GetOrder(s.ctx, orderGet)
	s.Require().NoError(err)

	s.Equal(orderUpdate, orderGet)
}

// описание: пытаемся обновить заказ, который не существует в базе данных
// ожидаемый результат: попытка обновления возвращает ошибку, ничего не обновляется и заказа все ещё не существует после попытки изменить данные
func (s *RepositorySuite) TestUpdateOrder_NotExists() {
	orderUpdate := &models.Order{
		ID:        1,
		User:      &models.User{ID: 1},
		Expiry:    time.Date(10000, 1, 1, 0, 0, 0, 0, time.UTC),
		ReturnBy:  time.Date(10001, 1, 1, 0, 0, 0, 0, time.UTC),
		Status:    models.StatusDelivered,
		Hash:      "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		CreatedAt: now,
		Cost:      1,
		Weight:    1,
		Packaging: "bag",
	}

	err := s.repository.UpdateOrder(s.ctx, orderUpdate)
	s.Error(err)
}

// get order

// описание: пытаемся получить заказ, которого нет в бд
// ожидаемый результат: возвращается ошибка, что заказ не найден
// P.S. успешное получение заказа через GetOrder уже описано в TestCreateOrder_Success, и имело бы идентичную логику
func (s *RepositorySuite) TestGetOrder_NotFound() {
	orderGet := &models.Order{
		ID: 1,
	}

	_, err := s.repository.GetOrder(s.ctx, orderGet)
	s.Error(err)
}

// get orders

// описание: создаем два заказа с разными статусами и получаем их
// ожидаемый результат: заказы возвращаются в правильном порядке по дате создания (DESC)
func (s *RepositorySuite) TestGetOrders_Standard() {
	orderAccepted := &models.Order{
		ID:        1,
		User:      &models.User{ID: 1},
		Expiry:    time.Date(10000, 1, 1, 0, 0, 0, 0, time.UTC),
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
		Expiry:    time.Date(10000, 1, 1, 0, 0, 0, 0, time.UTC),
		ReturnBy:  time.Date(10001, 1, 1, 0, 0, 0, 0, time.UTC),
		Status:    models.StatusDelivered,
		Hash:      "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		CreatedAt: now.Add(time.Second), // чтобы был позже чем orderAccepted
		Cost:      1,
		Weight:    1,
		Packaging: "bag",
	}

	err := s.repository.CreateOrder(s.ctx, orderAccepted)
	s.Require().NoError(err)

	err = s.repository.CreateOrder(s.ctx, orderDelivered)
	s.Require().NoError(err)

	orders, err := s.repository.GetOrders(s.ctx, 1, 2, 0, false)
	s.Require().NoError(err)

	s.Len(orders, 2)
	s.Equal(orderDelivered, orders[0])
	s.Equal(orderAccepted, orders[1])
}

// описание: теперь с лимитом
// ожидаемый результат: возвращается только один заказ
func (s *RepositorySuite) TestGetOrders_Limit() {
	orderAccepted := &models.Order{
		ID:        1,
		User:      &models.User{ID: 1},
		Expiry:    time.Date(10000, 1, 1, 0, 0, 0, 0, time.UTC),
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
		Expiry:    time.Date(10000, 1, 1, 0, 0, 0, 0, time.UTC),
		ReturnBy:  time.Date(10001, 1, 1, 0, 0, 0, 0, time.UTC),
		Status:    models.StatusDelivered,
		Hash:      "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		CreatedAt: now.Add(time.Second),
		Cost:      1,
		Weight:    1,
		Packaging: "bag",
	}

	err := s.repository.CreateOrder(s.ctx, orderAccepted)
	s.Require().NoError(err)

	err = s.repository.CreateOrder(s.ctx, orderDelivered)
	s.Require().NoError(err)

	orders, err := s.repository.GetOrders(s.ctx, 1, 1, 0, false)
	s.Require().NoError(err)

	s.Len(orders, 1)
	s.Equal(orderDelivered, orders[0])
}

// описание: теперь смещаем
// ожидаемый результат: возвращается только второй заказ
func (s *RepositorySuite) TestGetOrders_Offset() {
	orderAccepted := &models.Order{
		ID:        1,
		User:      &models.User{ID: 1},
		Expiry:    time.Date(10000, 1, 1, 0, 0, 0, 0, time.UTC),
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
		Expiry:    time.Date(10000, 1, 1, 0, 0, 0, 0, time.UTC),
		ReturnBy:  time.Date(10001, 1, 1, 0, 0, 0, 0, time.UTC),
		Status:    models.StatusDelivered,
		Hash:      "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		CreatedAt: now.Add(time.Second),
		Cost:      1,
		Weight:    1,
		Packaging: "bag",
	}

	err := s.repository.CreateOrder(s.ctx, orderAccepted)
	s.Require().NoError(err)

	err = s.repository.CreateOrder(s.ctx, orderDelivered)
	s.Require().NoError(err)

	orders, err := s.repository.GetOrders(s.ctx, 1, 2, 1, false)
	s.Require().NoError(err)

	s.Len(orders, 1)
	s.Equal(orderAccepted, orders[0])
}

// описание: создаем два заказа и получаем только то, что хранится (StatusAccepted)
// ожидаемый результат: возвращается только первый заказ, ведь второй заказ доставили
func (s *RepositorySuite) TestGetOrders_IsStored() {
	orderAccepted := &models.Order{
		ID:        1,
		User:      &models.User{ID: 1},
		Expiry:    time.Date(10000, 1, 1, 0, 0, 0, 0, time.UTC),
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
		Expiry:    time.Date(10000, 1, 1, 0, 0, 0, 0, time.UTC),
		ReturnBy:  time.Date(10001, 1, 1, 0, 0, 0, 0, time.UTC),
		Status:    models.StatusDelivered,
		Hash:      "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		CreatedAt: now.Add(time.Second),
		Cost:      1,
		Weight:    1,
		Packaging: "bag",
	}

	err := s.repository.CreateOrder(s.ctx, orderAccepted)
	s.Require().NoError(err)

	err = s.repository.CreateOrder(s.ctx, orderDelivered)
	s.Require().NoError(err)

	orders, err := s.repository.GetOrders(s.ctx, 1, 2, 0, true)
	s.Require().NoError(err)

	s.Len(orders, 1)
	s.Equal(orderAccepted, orders[0])
}

// get orders for delivery

// описание: создаем два заказа и пытаемся получить их по их id
// ожидаемый результат: оба заказа находятся и возвращаются
func (s *RepositorySuite) TestGetOrdersForDelivery_BothMatching() {
	order1 := &models.Order{
		ID:        1,
		User:      &models.User{ID: 1},
		Expiry:    time.Date(10000, 1, 1, 0, 0, 0, 0, time.UTC),
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
		Expiry:    time.Date(10000, 1, 1, 0, 0, 0, 0, time.UTC),
		ReturnBy:  time.Date(10001, 1, 1, 0, 0, 0, 0, time.UTC),
		Status:    models.StatusDelivered,
		Hash:      "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		CreatedAt: now.Add(time.Second),
		Cost:      1,
		Weight:    1,
		Packaging: "bag",
	}

	err := s.repository.CreateOrder(s.ctx, order1)
	s.Require().NoError(err)

	err = s.repository.CreateOrder(s.ctx, order2)
	s.Require().NoError(err)

	orders, err := s.repository.GetOrdersForDelivery(s.ctx, []uint64{1, 2})
	s.Require().NoError(err)

	s.Len(orders, 2)
	s.Equal(order1, orders[0])
	s.Equal(order2, orders[1])
}

// описание: теперь пытаемся получить их по одному существующему и одному несуществующему id
// ожидаемый результат: возвращается только один существующий заказ
func (s *RepositorySuite) TestGetOrdersForDelivery_OneMatching() {
	order1 := &models.Order{
		ID:        1,
		User:      &models.User{ID: 1},
		Expiry:    time.Date(10000, 1, 1, 0, 0, 0, 0, time.UTC),
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
		Expiry:    time.Date(10000, 1, 1, 0, 0, 0, 0, time.UTC),
		ReturnBy:  time.Date(10001, 1, 1, 0, 0, 0, 0, time.UTC),
		Status:    models.StatusDelivered,
		Hash:      "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		CreatedAt: now.Add(time.Second),
		Cost:      1,
		Weight:    1,
		Packaging: "bag",
	}

	err := s.repository.CreateOrder(s.ctx, order1)
	s.Require().NoError(err)

	err = s.repository.CreateOrder(s.ctx, order2)
	s.Require().NoError(err)

	orders, err := s.repository.GetOrdersForDelivery(s.ctx, []uint64{1, 3})
	s.Require().NoError(err)

	s.Len(orders, 1)
	s.Equal(order1, orders[0])
}

// описание: теперь ни одного заказа нет в бд
// ожидаемый результат: ничего не находятся
func (s *RepositorySuite) TestGetOrdersForDelivery_NoneMatching() {
	order1 := &models.Order{
		ID:        1,
		User:      &models.User{ID: 1},
		Expiry:    time.Date(10000, 1, 1, 0, 0, 0, 0, time.UTC),
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
		Expiry:    time.Date(10000, 1, 1, 0, 0, 0, 0, time.UTC),
		ReturnBy:  time.Date(10001, 1, 1, 0, 0, 0, 0, time.UTC),
		Status:    models.StatusDelivered,
		Hash:      "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		CreatedAt: now.Add(time.Second),
		Cost:      1,
		Weight:    1,
		Packaging: "bag",
	}

	err := s.repository.CreateOrder(s.ctx, order1)
	s.Require().NoError(err)

	err = s.repository.CreateOrder(s.ctx, order2)
	s.Require().NoError(err)

	orders, err := s.repository.GetOrdersForDelivery(s.ctx, []uint64{3, 4})
	s.Require().NoError(err)

	s.Len(orders, 0)
}
