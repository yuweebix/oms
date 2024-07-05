package cli

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"github.com/spf13/cobra"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/cli/flags"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
	"gitlab.ozon.dev/yuweebix/homework-1/pkg/utils"
)

const (
	topic = "cli"
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

type producer interface {
	Send(topic string, message any) error
}

type message struct {
	CreatedAt  time.Time `json:"created_at"`
	MethodName string    `json:"method"`
	RawRequest string    `json:"request"`
}

// CLI представляет слой командной строки приложения
type CLI struct {
	domain   domain
	producer producer
	logger   *log.Logger
	mu       *sync.Mutex
	cmd      *cobra.Command
}

// NewCLI конструктор с добавлением зависимостей
func NewCLI(d domain, p producer, logFileName string) (c *CLI, err error) {
	logger := createLogger(logFileName)

	c = &CLI{
		domain:   d,
		producer: p,
		logger:   logger,
		mu:       &sync.Mutex{},
	}

	c.cmd, err = initRootCmd(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

// Execute выполняет команду CLI
func (c *CLI) Execute(ctx context.Context, args []string) error {
	c.mu.Lock()
	c.cmd.SetArgs(args)
	c.mu.Unlock()

	err := c.cmd.ExecuteContext(ctx)
	if err != nil {
		return err
	}

	// позле вызова help, нужно отдельно сбрасывать, потому что она за границами run-функций
	c.mu.Lock()
	if utils.ContainsHelpFlag(args) {
		flags.ResetAllHelpFlags(c.cmd)
	}
	c.mu.Unlock()

	return nil
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
