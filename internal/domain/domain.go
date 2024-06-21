package domain

import "gitlab.ozon.dev/yuweebix/homework-1/internal/models"

// storage интерфейс необходимых сервису функций для реализации хранилищем
type storage interface {
	// заказы
	CreateOrder(o *models.Order) error                                                            // добавить заказ в хранилище
	DeleteOrder(o *models.Order) error                                                            // удалить заказ из хранилища
	UpdateOrder(o *models.Order) error                                                            // обновить данные о заказе в хранилище
	GetOrders(userID uint64, limit uint64, offset uint64, isStored bool) ([]*models.Order, error) // выдать список заказов клиента
	GetOrdersForDelivery(orderIDs []uint64) ([]*models.Order, error)                              // получить заказы для выдачи
	GetOrder(o *models.Order) (*models.Order, error)                                              // получить заказ

	// возвраты
	GetReturns(limit uint64, offset uint64) ([]*models.Order, error) // выдать список возвратов

	// упаковка
	GetPackaging(p *models.Packaging) (*models.Packaging, error) // получить данные об упаковке
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
