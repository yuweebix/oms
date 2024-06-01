package root

import (
	"github.com/spf13/cobra"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/cli"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/cli/root/commands"
)

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "Утилита для управления ПВЗ",
	Long: `Утилита содержит перечень команд, что можно производить над заказами и возвратами.

Заказы:
  - "orders accept  [flags]": Принять заказ от курьера
  - "orders return  [flags]": Вернуть заказ курьеру
  - "orders deliver [flags]": Выдать заказ клиенту
  - "orders list    [flags]": Получить список заказов
	
Возвраты:
  - "returns accept [flags]": Принять возврат от клиента
  - "returns list   [flags]": Получить список возвратов`,
}

func Execute(c *cli.CLI, args []string) error {
	rootCmd.SetArgs(args)
	err := rootCmd.Execute()
	if err != nil {
		return err
	}
	return nil
}

func InitRootCmd(c *cli.CLI) {
	for _, init := range commands.InitCommands {
		init(rootCmd, c)
	}
}
