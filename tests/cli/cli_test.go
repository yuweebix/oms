package cli_test

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/cli"
	e "gitlab.ozon.dev/yuweebix/homework-1/internal/cli/errors"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/cli/mocks"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
)

const (
	topic = "cli"
)

type CLISuite struct {
	suite.Suite
}

// TestCLISuite запускает все cli unit-тесты
func TestCLISuite(t *testing.T) {
	suite.Run(t, new(CLISuite))
}

func (s *CLISuite) SetUpTest() (_cli *cli.CLI, _domain *mocks.MockDomain) {
	_domain = mocks.NewMockDomain(s.T())
	_producer := mocks.NewMockProducer(s.T())

	_producer.EXPECT().Send(topic, mock.Anything).Return(nil).Maybe()

	_cli, err := cli.NewCLI(_domain, _producer, "log_text.txt")
	if err != nil {
		s.FailNowf("could not create cli", err.Error())
	}
	return
}

// no args tests

func (s *CLISuite) TestOrdersAccept_NoArgs() {
	s.T().Parallel()

	args := []string{"orders", "accept"}

	cli, _ := s.SetUpTest()

	err := cli.Execute(context.Background(), args)

	s.ErrorContains(err, "required flag(s)")
	s.ErrorContains(err, "not set")
}

func (s *CLISuite) TestOrdersDeliver_NoArgs() {
	s.T().Parallel()

	args := []string{"orders", "deliver"}

	cli, _ := s.SetUpTest()

	err := cli.Execute(context.Background(), args)

	s.ErrorContains(err, "required flag(s)")
	s.ErrorContains(err, "not set")
}

func (s *CLISuite) TestOrdersList_NoArgs() {
	s.T().Parallel()

	args := []string{"orders", "list"}

	cli, _ := s.SetUpTest()

	err := cli.Execute(context.Background(), args)

	s.ErrorContains(err, "required flag(s)")
	s.ErrorContains(err, "not set")
}

func (s *CLISuite) TestOrdersReturn_NoArgs() {
	s.T().Parallel()

	args := []string{"orders", "return"}

	cli, _ := s.SetUpTest()

	err := cli.Execute(context.Background(), args)

	s.ErrorContains(err, "required flag(s)")
	s.ErrorContains(err, "not set")
}

func (s *CLISuite) TestReturnsAccept_NoArgs() {
	s.T().Parallel()

	args := []string{"returns", "accept"}

	cli, _ := s.SetUpTest()

	err := cli.Execute(context.Background(), args)

	s.ErrorContains(err, "required flag(s)")
	s.ErrorContains(err, "not set")
}

// returns list -  особый случай, когда нету обязательных флагов

func (s *CLISuite) TestWorkers_NoArgs() {
	s.T().Parallel()

	args := []string{"workers"}

	cli, _ := s.SetUpTest()

	err := cli.Execute(context.Background(), args)

	s.ErrorContains(err, "required flag(s)")
	s.ErrorContains(err, "not set")
}

// no flag arg tests

func (s *CLISuite) TestOrdersAccept_NoFlagArg() {
	s.T().Parallel()

	args := []string{"orders", "accept", "-o"}

	cli, _ := s.SetUpTest()

	err := cli.Execute(context.Background(), args)

	s.ErrorContains(err, "flag needs an argument")
}

func (s *CLISuite) TestOrdersDeliver_NoFlagArg() {
	s.T().Parallel()

	args := []string{"orders", "deliver", "-o"}

	cli, _ := s.SetUpTest()

	err := cli.Execute(context.Background(), args)

	s.ErrorContains(err, "flag needs an argument")
}

func (s *CLISuite) TestOrdersList_NoFlagArg() {
	s.T().Parallel()

	args := []string{"orders", "list", "-u"}

	cli, _ := s.SetUpTest()

	err := cli.Execute(context.Background(), args)

	s.ErrorContains(err, "flag needs an argument")
}

func (s *CLISuite) TestOrdersReturn_NoFlagArg() {
	s.T().Parallel()

	args := []string{"orders", "return", "-o"}

	cli, _ := s.SetUpTest()

	err := cli.Execute(context.Background(), args)

	s.ErrorContains(err, "flag needs an argument")
}

func (s *CLISuite) TestReturnsAccept_NoFlagArg() {
	s.T().Parallel()

	args := []string{"returns", "accept", "-o"}

	cli, _ := s.SetUpTest()

	err := cli.Execute(context.Background(), args)

	s.ErrorContains(err, "flag needs an argument")
}

func (s *CLISuite) TestReturnsList_NoFlagArg() {
	s.T().Parallel()

	args := []string{"returns", "list", "-l"}

	cli, _ := s.SetUpTest()

	err := cli.Execute(context.Background(), args)

	s.ErrorContains(err, "flag needs an argument")
}

func (s *CLISuite) TestWorkers_NoFlagArg() {
	s.T().Parallel()

	args := []string{"workers", "-n"}

	cli, _ := s.SetUpTest()

	err := cli.Execute(context.Background(), args)

	s.ErrorContains(err, "flag needs an argument")
}

// unknow command test

func (s *CLISuite) TestUknownCommand() {
	s.T().Parallel()

	args := []string{"lmao"}

	cli, _ := s.SetUpTest()

	err := cli.Execute(context.Background(), args)

	s.ErrorContains(err, "unknown command")
}

// orders accept invalid date format test

func (s *CLISuite) TestOrdersAccept_InvalidDateFormat() {
	s.T().Parallel()

	text := "orders accept -o 1 -u 1 -e 1-1-1 -c 1 -w 1 -p bag"
	text = strings.TrimSpace(text)
	args := strings.Fields(text)

	cli, _ := s.SetUpTest()

	err := cli.Execute(context.Background(), args)

	s.EqualError(err, e.ErrDateFormatInvalid.Error())
}

// successful tests

func (s *CLISuite) TestOrdersAccept_Success() {
	s.T().Parallel()

	text := "orders accept -o 1 -u 1 -e 1111-11-11 -c 1 -w 1 -p bag"
	text = strings.TrimSpace(text)
	args := strings.Fields(text)

	cli, domain := s.SetUpTest()

	domain.EXPECT().AcceptOrder(mock.Anything, mock.Anything).Return(nil)
	err := cli.Execute(context.Background(), args)

	s.NoError(err)
	domain.AssertCalled(s.T(), "AcceptOrder", mock.Anything, mock.Anything)
}

func (s *CLISuite) TestOrdersDeliver_Success() {
	s.T().Parallel()

	text := "orders deliver -o 1,2,3"
	text = strings.TrimSpace(text)
	args := strings.Fields(text)

	cli, domain := s.SetUpTest()

	domain.EXPECT().DeliverOrders(mock.Anything, mock.Anything).Return(nil)
	err := cli.Execute(context.Background(), args)

	s.NoError(err)
	domain.AssertCalled(s.T(), "DeliverOrders", mock.Anything, mock.Anything)
}

func (s *CLISuite) TestOrdersList_Success() {
	s.T().Parallel()

	text := "orders list -u 1"
	text = strings.TrimSpace(text)
	args := strings.Fields(text)

	cli, domain := s.SetUpTest()

	domain.EXPECT().ListOrders(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return([]*models.Order{}, nil)
	err := cli.Execute(context.Background(), args)

	s.NoError(err)
	domain.AssertCalled(s.T(), "ListOrders", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
}

func (s *CLISuite) TestOrdersReturn_Success() {
	s.T().Parallel()

	text := "orders return -o 1"
	text = strings.TrimSpace(text)
	args := strings.Fields(text)

	cli, domain := s.SetUpTest()

	domain.EXPECT().ReturnOrder(mock.Anything, mock.Anything).Return(nil)
	err := cli.Execute(context.Background(), args)

	s.NoError(err)
	domain.AssertCalled(s.T(), "ReturnOrder", mock.Anything, mock.Anything)
}

func (s *CLISuite) TestReturnsAccept_Success() {
	s.T().Parallel()

	text := "returns accept -o 1 -u 1"
	text = strings.TrimSpace(text)
	args := strings.Fields(text)

	cli, domain := s.SetUpTest()

	domain.EXPECT().AcceptReturn(mock.Anything, mock.Anything).Return(nil)
	err := cli.Execute(context.Background(), args)

	s.NoError(err)
	domain.AssertCalled(s.T(), "AcceptReturn", mock.Anything, mock.Anything)
}

func (s *CLISuite) TestReturnsList_Success() {
	s.T().Parallel()

	text := "returns list"
	text = strings.TrimSpace(text)
	args := strings.Fields(text)

	cli, domain := s.SetUpTest()

	domain.EXPECT().ListReturns(mock.Anything, mock.Anything, mock.Anything).Return([]*models.Order{}, nil)
	err := cli.Execute(context.Background(), args)

	s.NoError(err)
	domain.AssertCalled(s.T(), "ListReturns", mock.Anything, mock.Anything, mock.Anything)
}

func (s *CLISuite) TestWorkers_Success() {
	s.T().Parallel()

	text := "workers -n 5"
	text = strings.TrimSpace(text)
	args := strings.Fields(text)

	cli, domain := s.SetUpTest()

	domain.EXPECT().ChangeWorkersNumber(mock.Anything, mock.Anything).Return(nil)
	err := cli.Execute(context.Background(), args)

	s.NoError(err)
	domain.AssertCalled(s.T(), "ChangeWorkersNumber", mock.Anything, mock.Anything)
}
