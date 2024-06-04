package orders

import (
	"github.com/spf13/cobra"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/cli"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/cli/root/commands/orders/commands"
)

// ordersCmd represents the Orders command
var ordersCmd = &cobra.Command{
	Use:   "orders",
	Short: "Совершить операцию с заказом",
	Long: `Команда "orders" содержит перечень команд для обработки заказа.

Для большей информации вызовите команду:
  orders [command] --help
`,
}

func InitOrdersCmd(parentCmd *cobra.Command, c *cli.CLI) {
	for _, init := range commands.InitCommands {
		init(ordersCmd, c)
	}

	parentCmd.AddCommand(ordersCmd)
}
