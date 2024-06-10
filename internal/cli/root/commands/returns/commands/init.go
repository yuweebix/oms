package commands

import (
	"github.com/spf13/cobra"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/cli"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
)

type initCommand func(*cobra.Command, *cli.CLI)

var InitCommands = []initCommand{InitAcceptCmd, InitListCmd}

// InitAcceptCmd принимает данные о заказе на принятие
func InitAcceptCmd(parentCmd *cobra.Command, c *cli.CLI) {
	// инициализируем флаги
	acceptCmd.Flags().IntP(flagOrderID.Unzip())
	acceptCmd.Flags().IntP(flagUserID.Unzip())

	// помечаем флаги как обязательные
	acceptCmd.MarkFlagRequired(flagOrderID.Name)
	acceptCmd.MarkFlagRequired(flagUserID.Name)

	acceptCmd.RunE = func(cmd *cobra.Command, args []string) (err error) {
		orderID, err := cmd.Flags().GetInt(flagOrderID.Name)
		if err != nil {
			return err
		}
		flagUserID.Value, err = cmd.Flags().GetInt(flagUserID.Name)
		if err != nil {
			return err
		}

		err = c.Service().AcceptReturn(&models.Order{
			ID:   orderID,
			User: &models.User{ID: flagUserID.Value},
		})
		if err != nil {
			return err
		}

		c.Logger.Println("Заказ возвращен.")
		return nil
	}

	parentCmd.AddCommand(acceptCmd)
}

// InitListCmd принимает данные о заказе на принятие
func InitListCmd(parentCmd *cobra.Command, c *cli.CLI) {
	// инициализируем флаги
	listCmd.Flags().IntP(flagStart.Unzip())  // опциональный флаг
	listCmd.Flags().IntP(flagFinish.Unzip()) // опциональный флаг

	listCmd.RunE = func(cmd *cobra.Command, args []string) (err error) {
		start, err := cmd.Flags().GetInt(flagStart.Name)
		if err != nil {
			return err
		}
		finish, err := cmd.Flags().GetInt(flagFinish.Name)
		if err != nil {
			return err
		}

		list, err := c.Service().ListReturns(start, finish)
		if err != nil {
			return err
		}

		for _, v := range list {
			c.Logger.Printf("Возврат: %v. Получатель: %v.\n", v.ID, v.User.ID)
		}
		return nil
	}

	parentCmd.AddCommand(listCmd)
}
