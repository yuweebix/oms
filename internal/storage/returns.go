package storage

import (
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
)

// GetReturns передает список возвратов от start-ого возврата до finish-ого возврата
func (s *Storage) GetReturns() (list []*models.Order, err error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// запишем данные из файла в database
	database, err := s.loadOrders()
	if err != nil {
		return nil, err
	}

	// записываем в список
	for _, v := range database {
		if v.Status == models.StatusReturned {
			list = append(list, v)
		}
	}

	return list, nil
}
