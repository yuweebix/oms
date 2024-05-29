package module

import "gitlab.ozon.dev/yuweebix/homework-1/internal/models"

type Storage interface {
	// заказы
	AddOrder(o *models.Order) error                     // добавить заказ в хранилище
	DeleteOrder(o *models.Order) error                  // удалить заказ из хранилища
	ListOrders(limit int) ([]*models.Order, error)      // выдать список заказов
	CheckOrdersForDelivery(oSet map[int]struct{}) error // отметить заказы при выдачи

	// возвраты
	AddReturn(o *models.Order) error                        // отметить доставленный заказ как вернутый
	ListReturns(start, finish int) ([]*models.Order, error) // выдать список возвратов
}

type Deps struct {
	Storage Storage
}

type Module struct {
	Deps
}

// NewModule конструктор с добавлением зависимостей
func NewModule(d Deps) *Module {
	return &Module{Deps: d}
}
