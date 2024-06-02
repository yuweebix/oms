package storage

import (
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
)

// ListReturns передает список возвратов от start-ого возврата до finish-ого возврата
func (s *Storage) ListReturns() ([]*models.Order, error) {
	var database map[int]*models.Order
	var err error

	// запишем данные из файла в database
	if database, err = s.loadOrders(); err != nil {
		return nil, err
	}

	// записываем в список
	var list = make([]*models.Order, 0, len(database))
	for _, v := range database {
		if v.Status == models.StatusReturned {
			list = append(list, v)
		}
	}

	return list, nil
}
