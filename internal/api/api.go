package api

import (
	"context"
	"time"

	orders "gitlab.ozon.dev/yuweebix/homework-1/gen/orders/v1/proto"
	returns "gitlab.ozon.dev/yuweebix/homework-1/gen/returns/v1/proto"
	e "gitlab.ozon.dev/yuweebix/homework-1/internal/api/errors"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// domain интерфейс необходимых CLI функций для реализации сервисом
type domain interface {
	// заказы
	AcceptOrder(ctx context.Context, o *models.Order) error                                                             // логика принятия заказа от курьера
	ReturnOrder(ctx context.Context, o *models.Order) error                                                             // логика возврата просроченного заказа курьеру
	ListOrders(ctx context.Context, userID uint64, limit uint64, offset uint64, isStored bool) ([]*models.Order, error) // логика вывода списка заказов
	DeliverOrders(ctx context.Context, orderIDs []uint64) error                                                         // логика выдачи заказов клиенту

	// возвраты
	AcceptReturn(ctx context.Context, o *models.Order) error                               // логика принятия возврата от клиента
	ListReturns(ctx context.Context, limit uint64, offset uint64) ([]*models.Order, error) // логика вывода возвратов
}

// producer интерфейс для производства сообщений в брокер сообщений (кафку)
type producer interface {
	Send(message any) error
}

// API представляет слой API
type API struct {
	domain   domain
	producer producer
	orders.UnimplementedOrdersServer
	returns.UnimplementedReturnsServer
}

// NewAPI конструктор с добавлением зависимостей
func NewAPI(d domain, p producer) (api *API) {
	api = &API{
		domain:   d,
		producer: p,
	}

	return api
}

// getMessage вспомогательная функция для получения сообщения для брокера
func getMessage(ctx context.Context, message protoreflect.Message) (msg *models.Message, err error) {
	// получаем сырой запрос
	raw, err := protojson.Marshal(message.Interface())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// получаем название метода
	method, ok := grpc.Method(ctx)
	if !ok {
		return nil, status.Error(codes.Internal, e.ErrMethodNotFound.Error())
	}

	msg = &models.Message{
		CreatedAt:  time.Now().UTC(),
		MethodName: method,
		RawRequest: string(raw),
	}

	return msg, nil
}
