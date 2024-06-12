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

	"gitlab.ozon.dev/yuweebix/homework-1/internal/cli"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/middleware"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/service"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/storage"
)

const (
	storageFileName = "orders.json"
	logFileName     = "log.txt"
	numWorkers      = 5 // начальное количество рабочих в пуле
)

func main() {
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
		os.Exit(1)
	}

	// инициализируем хранилище, сервис и утилиту
	storageJSON, err := storage.NewStorage(storageFileName)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	service := service.NewService(storageJSON, wp)
	c := cli.NewCLI(service, logFileName)

	wp.Start()

	// инициализируем канал для считывания команд
	// и каналом для синхронизации
	ch := make(chan []string)
	inSig := make(chan struct{})
	outSig := make(chan struct{})
	var wg sync.WaitGroup

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
				os.Exit(1)
			}
			inSig <- struct{}{}

			text = strings.TrimSpace(text)
			args := strings.Fields(text)

			// выходим
			if len(args) > 0 && args[0] == "exit" {
				cancel()
				break
			}

			ch <- args
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
			wp.Stop()
			wg.Wait() // Ждем завершения всех горутин
			return
		case args := <-ch:
			wp.Enqueue(ctx, func() { c.Execute(args) }, args)
		}
	}
}
