package domain

import (
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
)

// storage интерфейс необходимых сервису функций для реализации хранилищем
type storage interface {
	// заказы
	CreateOrder(tx models.Tx, o *models.Order) error                                                            // добавить заказ в хранилище
	DeleteOrder(tx models.Tx, o *models.Order) error                                                            // удалить заказ из хранилища
	UpdateOrder(tx models.Tx, o *models.Order) error                                                            // обновить данные о заказе в хранилище
	GetOrders(tx models.Tx, userID uint64, limit uint64, offset uint64, isStored bool) ([]*models.Order, error) // выдать список заказов клиента
	GetOrdersForDelivery(tx models.Tx, orderIDs []uint64) ([]*models.Order, error)                              // получить заказы для выдачи
	GetOrder(tx models.Tx, o *models.Order) (*models.Order, error)                                              // получить заказ

	// возвраты
	GetReturns(tx models.Tx, limit uint64, offset uint64) ([]*models.Order, error) // выдать список возвратов

	// начать исполнение транзакции
	Begin(opts models.TxOptions, fn func(tx models.Tx) error) error
}

type threading interface {
	// рабочие
	AddWorkers(n int) error
	RemoveWorkers(n int) error
}

// Domain представляет слой бизнес-логики приложения
type Domain struct {
	storage   storage
	threading threading
}

// NewDomain конструктор с добавлением зависимостей
func NewDomain(s storage, m threading) *Domain {
	return &Domain{storage: s, threading: m}
}
