package root

import (
	"github.com/spf13/cobra"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/cli"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/cli/root/commands"
)

var rootCmd = &cobra.Command{
	Use:   "*",
	Short: "",
	Long:  ``,
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
