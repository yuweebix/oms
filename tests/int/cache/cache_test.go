package cache_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/cache"
)

type CacheSuite struct {
	suite.Suite
	cache *cache.Cache
	ctx   context.Context
}

var (
	now = time.Now().UTC().Truncate(time.Second)
)

// TestCacheSuite запускает все cache int-тесты
func TestCacheSuite(t *testing.T) {
	suite.Run(t, new(CacheSuite))
}

func (s *CacheSuite) SetupSuite() {
	err := godotenv.Load("../../../.env")
	if err != nil {
		s.FailNowf("Error loading .env file", err.Error())
	}

	addr := os.Getenv("REDIS_LOCAL_TEST_ADDR")

	if addr == "" {
		s.FailNow("Error reading redis addr from the .env")
	}

	s.ctx = context.Background()

	s.cache, err = cache.NewCache(addr, "", 0)
	if err != nil {
		s.FailNowf("Error connecting to the redis", err.Error())
	}

}

func (s *CacheSuite) TearDownSuite() {
	s.cache.Close()
}

func (s *CacheSuite) AfterTest(suiteName, testName string) {
	s.cache.FlushAll(s.ctx)
}
