package cli

import (
	"log"
	"os"
	"sync"

	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
)

// domain интерфейс необходимых CLI функций для реализации сервисом
type domain interface {
	// заказы
	AcceptOrder(o *models.Order) error                                                             // логика принятия заказа от курьера
	ReturnOrder(o *models.Order) error                                                             // логика возврата просроченного заказа курьеру
	ListOrders(userID uint64, limit uint64, offset uint64, isStored bool) ([]*models.Order, error) // логика вывода списка заказов
	DeliverOrders(orderIDs []uint64) error                                                         // логика выдачи заказов клиенту

	// возвраты
	AcceptReturn(o *models.Order) error                               // логика принятия возврата от клиента
	ListReturns(limit uint64, offset uint64) ([]*models.Order, error) // логика вывода возвратов

	// рабочие
	ChangeWorkersNumber(workersNum int) error // логика изменения количества рабочих горутин
}

// CLI представляет слой командной строки приложения
type CLI struct {
	domain domain
	logger *log.Logger
	mu     *sync.Mutex
}

// NewCLI конструктор с добавлением зависимостей
func NewCLI(s domain, logFileName string) *CLI {
	logger := createLogger(logFileName)
	c := &CLI{
		domain: s,
		logger: logger,
		mu:     &sync.Mutex{},
	}
	c.initRootCmd()
	return c
}

// Execute выполняет команду CLI
func (c *CLI) Execute(args []string) {
	c.mu.Lock()
	rootCmd.SetArgs(args)
	c.mu.Unlock()

	err := rootCmd.Execute()
	if err != nil {
		c.logger.Println(err)
	}
}

// createLogger вспомогательная функция для открытия файла и привязки к нему логгера
func createLogger(logFileName string) *log.Logger {
	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	logger := log.New(logFile, "CLI: ", log.LstdFlags)
	return logger
}
