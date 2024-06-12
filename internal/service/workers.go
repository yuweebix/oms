package service

// рабочие
func (s *Service) ChangeWorkersNumber(workersNum int) error {
	if workersNum > 0 {
		return s.addWorkers(workersNum)
	} else if workersNum < 0 {
		return s.removeWorkers(workersNum)
	}
	return nil
}

func (s *Service) addWorkers(num int) error {
	return s.middleware.AddWorkers(num)
}

func (s *Service) removeWorkers(num int) error {
	return s.middleware.RemoveWorkers(num)
}
