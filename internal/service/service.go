package service

import "gitlab.ozon.dev/yuweebix/homework-1/internal/models"

// storage интерфейс необходимых сервису функций для реализации хранилищем
type storage interface {
	// заказы
	AddOrder(o *models.Order) error                         // добавить заказ в хранилище
	DeleteOrder(o *models.Order) error                      // удалить заказ из хранилища
	ListOrders(limit int) ([]*models.Order, error)          // выдать список заказов
	CheckOrdersForDelivery(orderIDs map[int]struct{}) error // отметить заказы при выдачи
	GetOrder(o *models.Order) (*models.Order, error)        // получить заказ

	// возвраты
	AddReturn(o *models.Order) error                        // отметить доставленный заказ как вернутый
	ListReturns(start, finish int) ([]*models.Order, error) // выдать список возвратов

}

// Deps содержит интерфейсы зависимостей
type Deps struct {
	Storage storage
	// inject dependency here
}

// Service представляет слой бизнес-логики приложения
type Service struct {
	Deps
}

// NewService конструктор с добавлением зависимостей
func NewService(d Deps) *Service {
	return &Service{Deps: d}
}
