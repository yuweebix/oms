package domain

// рабочие
func (s *Domain) ChangeWorkersNumber(workersNum int) error {
	if workersNum > 0 {
		return s.addWorkers(workersNum)
	} else if workersNum < 0 {
		return s.removeWorkers(workersNum)
	}
	return nil
}

func (s *Domain) addWorkers(num int) error {
	return s.middleware.AddWorkers(num)
}

func (s *Domain) removeWorkers(num int) error {
	return s.middleware.RemoveWorkers(num)
}
