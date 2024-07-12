package domain_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/domain"
	mocks "gitlab.ozon.dev/yuweebix/homework-1/mocks/domain"
)

type WorkersSuite struct {
	suite.Suite
}

// TestWorkersSuite запускает все workers unit-тесты
func TestWorkersSuite(t *testing.T) {
	suite.Run(t, new(WorkersSuite))
}

func (s *WorkersSuite) SetUpTest() (_domain *domain.Domain, _storage *mocks.MockStorage, _threading *mocks.MockThreading) {
	_storage = mocks.NewMockStorage(s.T())
	_threading = mocks.NewMockThreading(s.T())
	_domain = domain.NewDomain(_storage, _threading)
	return
}

func (s *WorkersSuite) TestChangeWorkersNumber_AddWorkers() {
	s.T().Parallel()

	numWorkers := 5
	domain, _, threading := s.SetUpTest()

	threading.EXPECT().AddWorkers(mock.Anything, numWorkers).Return(nil)

	err := domain.ChangeWorkersNumber(context.Background(), numWorkers)

	s.NoError(err)
	threading.AssertCalled(s.T(), "AddWorkers", mock.Anything, numWorkers) // тест ради теста, но я хотел чекнуть AssertCalled
}

func (s *WorkersSuite) TestChangeWorkersNumber_RemoveWorkers() {
	s.T().Parallel()

	// arrange
	numWorkers := -3
	domain, _, threading := s.SetUpTest()

	threading.EXPECT().RemoveWorkers(mock.Anything, numWorkers).Return(nil)

	// act
	err := domain.ChangeWorkersNumber(context.Background(), numWorkers)

	// assert
	s.NoError(err)
	threading.AssertCalled(s.T(), "RemoveWorkers", mock.Anything, numWorkers) // тоже тест ради теста
}

func (s *WorkersSuite) TestChangeWorkersNumber_NoChange() {
	s.T().Parallel()

	// arrange
	numWorkers := 0
	domain, _, threading := s.SetUpTest()

	// act
	err := domain.ChangeWorkersNumber(context.Background(), numWorkers)

	// assert
	s.NoError(err)
	threading.AssertNotCalled(s.T(), "AddWorkers", mock.Anything, mock.Anything) // тут я чекал AssertNotCalled
	threading.AssertNotCalled(s.T(), "RemoveWorkers", mock.Anything, mock.Anything)
}
