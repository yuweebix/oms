package service

import "gitlab.ozon.dev/yuweebix/homework-1/internal/models"

// storage интерфейс необходимых сервису функций для реализации хранилищем
type storage interface {
	// заказы
	AddOrder(o *models.Order) error                                          // добавить заказ в хранилище
	DeleteOrder(o *models.Order) error                                       // удалить заказ из хранилища
	ListOrders(userID int) ([]*models.Order, error)                          // выдать список заказов клиента
	GetOrdersForDelivery(orderIDs map[int]struct{}) ([]*models.Order, error) // получить заказы для выдачи
	GetOrder(o *models.Order) (*models.Order, error)                         // получить заказ

	// возвраты
	ListReturns() ([]*models.Order, error) // выдать список возвратов
}

type middleware interface {
	// рабочие
	AddWorkers(n int) error
	RemoveWorkers(n int) error
}

// Service представляет слой бизнес-логики приложения
type Service struct {
	storage    storage
	middleware middleware
}

// NewService конструктор с добавлением зависимостей
func NewService(s storage, m middleware) *Service {
	return &Service{storage: s, middleware: m}
}
