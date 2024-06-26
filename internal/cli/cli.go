package cli

import (
	"context"
	"log"
	"os"
	"sync"

	"gitlab.ozon.dev/yuweebix/homework-1/internal/cli/flags"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
	"gitlab.ozon.dev/yuweebix/homework-1/pkg/utils"
)

// domain интерфейс необходимых CLI функций для реализации сервисом
type domain interface {
	// заказы
	AcceptOrder(ctx context.Context, o *models.Order) error                                                             // логика принятия заказа от курьера
	ReturnOrder(ctx context.Context, o *models.Order) error                                                             // логика возврата просроченного заказа курьеру
	ListOrders(ctx context.Context, userID uint64, limit uint64, offset uint64, isStored bool) ([]*models.Order, error) // логика вывода списка заказов
	DeliverOrders(ctx context.Context, orderIDs []uint64) error                                                         // логика выдачи заказов клиенту

	// возвраты
	AcceptReturn(ctx context.Context, o *models.Order) error                               // логика принятия возврата от клиента
	ListReturns(ctx context.Context, limit uint64, offset uint64) ([]*models.Order, error) // логика вывода возвратов

	// рабочие
	ChangeWorkersNumber(ctx context.Context, workersNum int) error // логика изменения количества рабочих горутин
}

// CLI представляет слой командной строки приложения
type CLI struct {
	domain domain
	logger *log.Logger
	mu     *sync.Mutex
}

// NewCLI конструктор с добавлением зависимостей
func NewCLI(d domain, logFileName string) *CLI {
	logger := createLogger(logFileName)
	c := &CLI{
		domain: d,
		logger: logger,
		mu:     &sync.Mutex{},
	}
	c.initRootCmd()
	return c
}

// Execute выполняет команду CLI
func (c *CLI) Execute(ctx context.Context, args []string) {
	c.mu.Lock()
	rootCmd.SetArgs(args)
	c.mu.Unlock()

	err := rootCmd.ExecuteContext(ctx)
	if err != nil {
		c.logger.Println(err)
	}

	// позле вызова help, нужно отдельно сбрасывать, потому что она за границами run-функций
	c.mu.Lock()
	if utils.ContainsHelpFlag(args) {
		flags.ResetAllHelpFlags(rootCmd)
	}
	c.mu.Unlock()
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
