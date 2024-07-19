package repository_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/repository"
)

type RepositorySuite struct {
	suite.Suite
	repository *repository.Repository
	ctx        context.Context
}

var (
	now = time.Now().UTC().Truncate(time.Second)
)

// TestRepositorySuite запускает все repository int-тесты
func TestRepositorySuite(t *testing.T) {
	suite.Run(t, new(RepositorySuite))
}

func (s *RepositorySuite) SetupSuite() {
	err := godotenv.Load("../../../.env")
	if err != nil {
		s.FailNowf("Error loading .env file", err.Error())
	}

	environment := os.Getenv("ENVIRONMENT")
	var connString string

	if environment == "docker" {
		connString = os.Getenv("DATABASE_TEST_URL")
	} else {
		connString = os.Getenv("DATABASE_LOCAL_TEST_URL")
	}

	if connString == "" {
		s.FailNow("Error reading database url from the .env")
	}

	s.ctx = context.Background()

	s.repository, err = repository.NewRepository(s.ctx, connString)
	if err != nil {
		s.FailNowf("Error connecting to the database", err.Error())
	}

}

func (s *RepositorySuite) TearDownSuite() {
	s.repository.Close()
}

func (s *RepositorySuite) AfterTest(suiteName, testName string) {
	err := s.repository.DeleteAllOrders(s.ctx)
	if err != nil {
		s.Failf("Error truncating table orders", err.Error())
	}
}
