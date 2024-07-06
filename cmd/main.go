package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/joho/godotenv"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/cli"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/domain"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/kafka/pub"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/kafka/sub"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/repository"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/threading"
)

const (
	logFileName = "log.txt"
	numWorkers  = 5 // начальное количество рабочих в пуле
)

var (
	brokers = []string{"broker:19092"}
	topic   = "cli"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	connString := os.Getenv("DATABASE_URL")
	if connString == "" {
		log.Fatalf("Error reading DATABASE_URL from .env file: %v", err)
	}

	outputMode := os.Getenv("OUTPUT_MODE")
	if outputMode == "" {
		log.Fatalf("Error reading OUTPUT_MODE from .env file: %v", err)
	}

	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// graceful shutdown
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)

	// может быть либо использован воркер пулом, либо будет принимать от консьюмера
	notificationChan := make(chan string, 100)

	// инициализируем пул рабочих
	wp, err := threading.NewWorkerPool(ctx, numWorkers)
	if err != nil {
		log.Fatalln(err)
	}

	if outputMode == "worker_pool" {
		wp.Notify(notificationChan)
	}

	r, err := repository.NewRepository(ctx, connString)
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Close()

	d := domain.NewDomain(r, wp)

	producer, err := pub.NewProducer(brokers)
	if err != nil {
		log.Fatalln(err)
	}
	defer producer.Close()

	// если не настроено на кафку, то не создаём консьюмера
	consumer := &sub.Consumer{}
	if outputMode == "kafka" {
		consumer, err = sub.NewConsumer(brokers, topic, notificationChan)
		if err != nil {
			log.Fatalln(err)
		}
		defer consumer.Close()
	}

	c, err := cli.NewCLI(d, producer, logFileName)
	if err != nil {
		log.Fatalln(err)
	}

	wp.Start()

	// инициализируем канал для считывания команд
	// и каналами для синхронизации
	commandChan := make(chan []string)
	inSig := make(chan struct{})
	outSig := make(chan struct{})

	// горутина для считывания команд
	wg.Add(1)
	go func() {
		defer wg.Done()
		in := bufio.NewReader(os.Stdin)
		for {
			// считываем команду
			<-outSig
			text, err := in.ReadString('\n')
			if err != nil {
				log.Fatalln(err)
			}
			inSig <- struct{}{}

			text = strings.TrimSpace(text)
			args := strings.Fields(text)

			// выходим
			select {
			case <-ctx.Done():
				return
			default:
				if len(args) == 0 { // нажимая просто Enter, выводим все уведомления на данный момент
					continue
				}
				if len(args) > 0 && args[0] == "clear" {
					fmt.Print("\033[H\033[2J") // full clear
					continue
				}
				if len(args) > 0 && args[0] == "exit" {
					cancel()
					return
				}

				commandChan <- args
			}
		}
	}()

	// горутина для обработки уведомлений
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				for notification := range notificationChan {
					fmt.Println(notification)
				}
				return
			default:
				outSig <- struct{}{}
				<-inSig
				for len(notificationChan) > 0 {
					fmt.Println(<-notificationChan)
				}
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			wp.Stop()
			if err := producer.Close(); err != nil {
				log.Println(err)
			}
			if err := consumer.Close(); err != nil {
				log.Println(err)
			}
			close(notificationChan)
			wg.Wait()
			return
		case <-sigs:
			cancel()
			fmt.Println("\nНажмите Enter, чтобы закрыть утилиту.")
		case args := <-commandChan:
			wp.Enqueue(ctx, func() error { return c.Execute(ctx, args) }, args)
		}
	}
}
