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
	"gitlab.ozon.dev/yuweebix/homework-1/internal/middleware"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/repository"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/service"
)

const (
	// storageFileName = "orders.json"
	// logFileName     = "log.txt"
	numWorkers = 5 // начальное количество рабочих в пуле
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)

	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// graceful shutdown
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)

	// инициализируем пул рабочих
	notificationChan := make(chan string, 100)
	wp, err := middleware.NewWorkerPool(ctx, numWorkers, notificationChan)
	if err != nil {
		log.Fatalln(err)
	}

	// инициализируем хранилище, сервис и утилиту
	// storageJSON, err := storage.NewStorage(storageFileName)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// service := service.NewService(storageJSON, wp)
	// c := cli.NewCLI(service, logFileName)

	repository, err := repository.NewRepository(ctx, connString)
	if err != nil {
		log.Fatalln(err)
	}
	service := service.NewService(repository, wp)
	c := cli.NewCLI(service, connString)

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
			wg.Wait() // Ждем завершения всех горутин
			return
		case <-sigs:
			cancel()
			fmt.Println("\nНажмите Enter, чтобы завершить программу.")
			wp.Stop()
			wg.Wait() // Ждем завершения всех горутин
			return
		case args := <-commandChan:
			wp.Enqueue(ctx, func() { c.Execute(args) }, args)
		}
	}
}
