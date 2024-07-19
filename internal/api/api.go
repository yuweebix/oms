package api

import (
	"context"

	orders "gitlab.ozon.dev/yuweebix/homework-1/gen/orders/v1/proto"
	returns "gitlab.ozon.dev/yuweebix/homework-1/gen/returns/v1/proto"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
)

// service интерфейс необходимых CLI функций для реализации сервисом
type service interface {
	// заказы
	AcceptOrder(ctx context.Context, o *models.Order) error                                                             // логика принятия заказа от курьера
	ReturnOrder(ctx context.Context, o *models.Order) error                                                             // логика возврата просроченного заказа курьеру
	ListOrders(ctx context.Context, userID uint64, limit uint64, offset uint64, isStored bool) ([]*models.Order, error) // логика вывода списка заказов
	DeliverOrders(ctx context.Context, orderIDs []uint64) error                                                         // логика выдачи заказов клиенту

	// возвраты
	AcceptReturn(ctx context.Context, o *models.Order) error                               // логика принятия возврата от клиента
	ListReturns(ctx context.Context, limit uint64, offset uint64) ([]*models.Order, error) // логика вывода возвратов
}

// logger интерфейс для производства сообщений в брокер сообщений (кафку)
type logger interface {
	Send(message any) error
}

// API представляет слой API
type API struct {
	service service
	logger  logger
	orders.UnimplementedOrdersServer
	returns.UnimplementedReturnsServer
}

// NewAPI конструктор с добавлением зависимостей
func NewAPI(s service, l logger) (api *API) {
	api = &API{
		service: s,
		logger:  l,
	}

	return api
}
