package domain

import "context"

// рабочие
func (d *Domain) ChangeWorkersNumber(ctx context.Context, workersNum int) error {
	if workersNum > 0 {
		return d.addWorkers(ctx, workersNum)
	} else if workersNum < 0 {
		return d.removeWorkers(ctx, workersNum)
	}
	return nil
}

func (d *Domain) addWorkers(ctx context.Context, num int) error {
	return d.threading.AddWorkers(ctx, num)
}

func (d *Domain) removeWorkers(ctx context.Context, num int) error {
	return d.threading.RemoveWorkers(ctx, num)
}
