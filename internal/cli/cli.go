package cli

import "gitlab.ozon.dev/yuweebix/homework-1/internal/models"

// service интерфейс необходимых CLI функций для реализации сервисом
type service interface {
	// заказы
	AcceptOrder(o *models.Order) error                         // логика принятия заказа от курьера
	ReturnOrder(o *models.Order) error                         // логика возврата просроченного заказа курьеру
	ListOrders(userID int, limit int) ([]*models.Order, error) // логика вывода списка заказов
	DeliverOrders(orderIDs []int) error                        // логика выдачи заказов клиенту

	// возвраты
	AcceptReturn(o *models.Order) error                     // логика принятия возврата от клиента
	ListReturns(start, finish int) ([]*models.Order, error) // логика вывода возвратов
}

// Deps содержит интерфейсы зависимостей
type Deps struct {
	Service service
	// inject dependency here
}

// CLI представляет слой командной строки приложения
type CLI struct {
	Deps
}

// NewCLI конструктор с добавлением зависимостей
func NewCLI(d Deps) *CLI {
	return &CLI{Deps: d}
}
