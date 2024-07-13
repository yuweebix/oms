package domain

import (
	"context"

	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
)

// storage интерфейс необходимых сервису функций для реализации хранилищем
type storage interface {
	// заказы
	CreateOrder(ctx context.Context, o *models.Order) error                                                            // добавить заказ в хранилище
	DeleteOrder(ctx context.Context, o *models.Order) error                                                            // удалить заказ из хранилища
	UpdateOrder(ctx context.Context, o *models.Order) error                                                            // обновить данные о заказе в хранилище
	GetOrders(ctx context.Context, userID uint64, limit uint64, offset uint64, isStored bool) ([]*models.Order, error) // выдать список заказов клиента
	GetOrdersForDelivery(ctx context.Context, orderIDs []uint64) ([]*models.Order, error)                              // получить заказы для выдачи
	GetOrder(ctx context.Context, o *models.Order) (*models.Order, error)                                              // получить заказ

	// возвраты
	GetReturns(ctx context.Context, limit uint64, offset uint64) ([]*models.Order, error) // выдать список возвратов

	// начать исполнение транзакции
	RunTx(ctx context.Context, opts models.TxOptions, fn func(ctxTX context.Context) error) error
}

// Domain представляет слой бизнес-логики приложения
type Domain struct {
	storage storage
}

// NewDomain конструктор с добавлением зависимостей
func NewDomain(s storage) *Domain {
	return &Domain{storage: s}
}
