package commands

import (
	"github.com/spf13/cobra"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/cli"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/cli/root/commands/orders"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/cli/root/commands/returns"
)

type initCommand func(*cobra.Command, *cli.CLI)

var InitCommands = []initCommand{orders.InitOrdersCmd, returns.InitReturnCmd}
