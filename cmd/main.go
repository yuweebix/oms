package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"gitlab.ozon.dev/yuweebix/homework-1/internal/cli"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/service"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/storage"
)

const (
	storageFileName = "orders.json"
	logFileName     = "log.txt"
)

func main() {
	// инициализируем хранилище, сервис и утилиту
	storageJSON, err := storage.NewStorage(storageFileName)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	service := service.NewService(storageJSON)
	c := cli.NewCLI(service, logFileName)

	// инициализируем канал для ввода
	ch := make(chan []string)
	// считываем команды
	go func() {
		in := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("> ")

			text, err := in.ReadString('\n')
			if err != nil {
				log.Fatalln(err)
				os.Exit(1)
			}

			text = strings.TrimSpace(text)
			args := strings.Fields(text)

			// выходим
			if len(args) > 0 && args[0] == "exit" {
				break
			}
			ch <- args
		}
	}()

	for {
		// запускаем команду
		args := <-ch
		c.Execute(args)
	}
}
