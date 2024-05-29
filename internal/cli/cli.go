package cli

import "gitlab.ozon.dev/yuweebix/homework-1/internal/models"

type Module interface {
	// заказы
	AcceptOrder(o *models.Order) error             // логика принятия заказа от курьера
	ReturnOrder(o *models.Order) error             // логика возврата просроченного заказа курьеру
	ListOrders(limit int) ([]*models.Order, error) // логика вывода списка заказов
	DeliverOrders(oIDSlice []int) error            // логика выдачи заказов клиенту

	// возвраты
	AcceptReturn(o *models.Order) error                     // логика принятия возврата от клиента
	ListReturns(start, finish int) ([]*models.Order, error) // логика вывода возвратов
}

type Deps struct {
	Module Module
}

type CLI struct {
	Deps
}

// NewCLI конструктор с добавлением зависимостей
func NewCLI(d Deps) *CLI {
	return &CLI{Deps: d}
}
