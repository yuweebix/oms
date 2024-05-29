package orders

import (
	"github.com/spf13/cobra"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/cli"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/cli/root/commands/orders/commands"
)

// ordersCmd represents the Orders command
var ordersCmd = &cobra.Command{
	Use:   "orders",
	Short: "",
	Long:  ``,
}

func InitordersCmd(parentCmd *cobra.Command, c *cli.CLI) {
	for _, init := range commands.InitCommands {
		init(ordersCmd, c)
	}

	ordersCmd.Run = func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		}
	}

	parentCmd.AddCommand(ordersCmd)
}
