package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"gitlab.ozon.dev/yuweebix/homework-1/internal/cli"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/cli/root"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/service"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/storage"
)

const (
	storageFileName = "orders.json"
	logFileName     = "log.txt"
)

func main() {
	// инициализируем хранилище
	storageJSON, err := storage.NewStorage(storageFileName)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	// инициализируем модуль
	service := service.NewService(storageJSON)
	// инициализируем CLI
	c := cli.NewCLI(service, logFileName)
	// инициализируем главную команду
	root.InitRootCmd(c)
	// считываем команды
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

		// запускаем команду
		err = root.Execute(c, args)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
