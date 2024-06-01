package returns

import (
	"github.com/spf13/cobra"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/cli"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/cli/root/commands/returns/commands"
)

// returnsCmd represents the Returns command
var returnsCmd = &cobra.Command{
	Use:   "returns",
	Short: "Совершить операцию с возвратом",
	Long: `Команда "returns" содержит перечень команд для обработки возврата.

Для большей информации вызовите команду:
  returns [command] --help`,
}

func InitReturnCmd(parentCmd *cobra.Command, c *cli.CLI) {
	for _, init := range commands.InitCommands {
		init(returnsCmd, c)
	}
	returnsCmd.Run = func(cmd *cobra.Command, args []string) {
		cmd.Help()
	}
	parentCmd.AddCommand(returnsCmd)
}
