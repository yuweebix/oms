package cli

import (
	"log"
	"os"

	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
)

// service интерфейс необходимых CLI функций для реализации сервисом
type service interface {
	// заказы
	AcceptOrder(o *models.Order) error                                        // логика принятия заказа от курьера
	ReturnOrder(o *models.Order) error                                        // логика возврата просроченного заказа курьеру
	ListOrders(userID int, limit int, isStored bool) ([]*models.Order, error) // логика вывода списка заказов
	DeliverOrders(orderIDs []int) error                                       // логика выдачи заказов клиенту

	// возвраты
	AcceptReturn(o *models.Order) error                     // логика принятия возврата от клиента
	ListReturns(start, finish int) ([]*models.Order, error) // логика вывода возвратов
}

// CLI представляет слой командной строки приложения
type CLI struct {
	service service
	Logger  *log.Logger
}

// NewCLI конструктор с добавлением зависимостей
func NewCLI(s service, logFileName string) *CLI {
	logger := createLogger(logFileName)

	return &CLI{
		service: s,
		Logger:  logger,
	}
}

// Service возвращает интерфейс service
func (c *CLI) Service() service {
	return c.service
}

func createLogger(logFileName string) *log.Logger {
	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	logger := log.New(logFile, "CLI: ", log.LstdFlags)
	return logger
}
