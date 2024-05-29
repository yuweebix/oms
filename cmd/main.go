package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"gitlab.ozon.dev/yuweebix/homework-1/internal/cli"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/cli/root"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/module"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/storage"
)

const (
	fileName = "orders.json"
)

func main() {
	// инициализируем хранилище
	storageJSON, err := storage.NewStorage(fileName)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	// ининциализируем модуль
	service := module.NewModule(module.Deps{
		Storage: storageJSON,
	})
	// иницилизирум CLI
	c := cli.NewCLI(cli.Deps{
		Module: service,
	})
	// ининцилизируем главную команду
	root.InitRootCmd(c)

	// считываем команды
	in := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")

		text, err := in.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		text = strings.TrimSuffix(text, "\n")
		args := strings.Split(text, " ")

		// выходим
		if args[0] == "exit" {
			break
		}

		// запускаем команду
		if err := root.Execute(c, args); err != nil {
			fmt.Println(err.Error())
		}
	}
}
