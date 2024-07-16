package domain_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/domain"
	mocks "gitlab.ozon.dev/yuweebix/homework-1/mocks/domain"
)

type DomainSuite struct {
	suite.Suite
}

const (
	day                 = time.Hour * 24
	returnByAllowedTime = day * 2
)

// TestDomainSuite запускает все domain unit-тесты
func TestDomainSuite(t *testing.T) {
	suite.Run(t, new(DomainSuite))
}

func (s *DomainSuite) SetupTest() (_domain *domain.Domain, _storage *mocks.MockStorage) {
	_storage = mocks.NewMockStorage(s.T())
	_domain = domain.NewDomain(_storage)
	return
}
