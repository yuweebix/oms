package cli

import (
	"github.com/spf13/cobra"
)

// функции для получения списков подкоманд

func getOrdersSubcommands(c *CLI) ([]*cobra.Command, error) {
	acceptCmd, err := c.ordersAcceptCmd()
	if err != nil {
		return nil, err
	}
	deliverCmd, err := c.ordersDeliverCmd()
	if err != nil {
		return nil, err
	}
	listCmd, err := c.ordersListCmd()
	if err != nil {
		return nil, err
	}
	returnCmd, err := c.ordersReturnCmd()
	if err != nil {
		return nil, err
	}
	// команду добавить сюда

	return []*cobra.Command{acceptCmd, deliverCmd, listCmd, returnCmd}, nil
}

func getReturnsSubcommands(c *CLI) ([]*cobra.Command, error) {
	acceptCmd, err := c.returnsAcceptCmd()
	if err != nil {
		return nil, err
	}
	listCmd, err := c.returnsListCmd()
	if err != nil {
		return nil, err
	}
	// команду добавить сюда

	return []*cobra.Command{acceptCmd, listCmd}, nil
}

func getWorkersSubcommands(_ *CLI) ([]*cobra.Command, error) {
	return []*cobra.Command{}, nil
}

// initRootCmd инициализирует стоковую команду
func initRootCmd(c *CLI) (*cobra.Command, error) {
	// orders
	ordersCmd, err := initCmd(c, c.ordersCmd, getOrdersSubcommands)
	if err != nil {
		return nil, err
	}

	// returns
	returnsCmd, err := initCmd(c, c.returnsCmd, getReturnsSubcommands)
	if err != nil {
		return nil, err
	}

	// workers
	workersCmd, err := initCmd(c, c.workersCmd, getWorkersSubcommands)
	if err != nil {
		return nil, err
	}

	// root
	rootCmd, err := c.rootCmd()
	if err != nil {
		return nil, err
	}
	rootCmd = initSubCmds(rootCmd, []*cobra.Command{ordersCmd, returnsCmd, workersCmd})

	return rootCmd, nil
}

// initCmd инициализирует команду и подкоманды
func initCmd(c *CLI, cmdFunc func() (*cobra.Command, error), subcmdsFunc func(c *CLI) ([]*cobra.Command, error)) (*cobra.Command, error) {
	cmd, err := cmdFunc()
	if err != nil {
		return nil, err
	}
	subcmds, err := subcmdsFunc(c)
	if err != nil {
		return nil, err
	}
	return initSubCmds(cmd, subcmds), nil
}

// initSubCmds инициализирует подкоманды
func initSubCmds(cmd *cobra.Command, subcommands []*cobra.Command) *cobra.Command {
	for _, subcmd := range subcommands {
		cmd.AddCommand(subcmd)
	}
	return cmd
}
