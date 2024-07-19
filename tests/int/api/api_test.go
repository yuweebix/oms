package api_test

import (
	"log"
	"net"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	orders "gitlab.ozon.dev/yuweebix/homework-1/gen/orders/v1/proto"
	returns "gitlab.ozon.dev/yuweebix/homework-1/gen/returns/v1/proto"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/api"
	mocks "gitlab.ozon.dev/yuweebix/homework-1/mocks/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type APISuite struct {
	suite.Suite
	server *grpc.Server
	port   string
}

// TestAPISuite запускает все API unit-тесты
func TestAPISuite(t *testing.T) {
	suite.Run(t, new(APISuite))
}

func (s *APISuite) SetupSuite() {
	// читаем данные из .env
	err := godotenv.Load("../../../.env")
	if err != nil {
		log.Fatalln(err)
	}

	// сначала нужно порт получить
	s.port = os.Getenv("GRPC_TEST_PORT")
	if s.port == "" {
		log.Fatalln("Error reading GRPC_TEST_PORT from .env file")
	}

	// domain + mock expectations
	domain := mocks.NewMockService(s.T())
	domain.EXPECT().AcceptOrder(mock.Anything, mock.Anything).Return(nil).Maybe()
	domain.EXPECT().DeliverOrders(mock.Anything, mock.Anything).Return(nil).Maybe()
	domain.EXPECT().ListOrders(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, nil).Maybe()
	domain.EXPECT().ReturnOrder(mock.Anything, mock.Anything).Return(nil).Maybe()
	domain.EXPECT().AcceptReturn(mock.Anything, mock.Anything).Return(nil).Maybe()
	domain.EXPECT().ListReturns(mock.Anything, mock.Anything, mock.Anything).Return(nil, nil).Maybe()

	// api
	api := api.NewAPI(domain)

	// server
	s.server = grpc.NewServer()
	orders.RegisterOrdersServer(s.server, api)
	returns.RegisterReturnsServer(s.server, api)

	// запуск grpc сервера
	go func() {
		// слушаем
		lis, err := net.Listen("tcp", s.port)
		if err != nil {
			log.Fatalln(err)
		}

		// сёрвим
		if err := s.server.Serve(lis); err != nil {
			log.Fatalln(err)
		}
	}()

}

func (s *APISuite) TearDownSuite() {
	s.server.GracefulStop()
}

// GetClientConn вспомогательная функция для получения подключения
func (s *APISuite) GetClientConn() (conn *grpc.ClientConn) {
	conn, err := grpc.NewClient("localhost"+s.port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		s.FailNowf("could not get client connection", err.Error())
	}
	return conn
}
