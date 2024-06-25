package domain

// рабочие
func (d *Domain) ChangeWorkersNumber(workersNum int) error {
	if workersNum > 0 {
		return d.addWorkers(workersNum)
	} else if workersNum < 0 {
		return d.removeWorkers(workersNum)
	}
	return nil
}

func (d *Domain) addWorkers(num int) error {
	return d.threading.AddWorkers(num)
}

func (d *Domain) removeWorkers(num int) error {
	return d.threading.RemoveWorkers(num)
}
