package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/cli"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
	"gitlab.ozon.dev/yuweebix/homework-1/pkg/flags"
)

type initCommand func(*cobra.Command, *cli.CLI)

var InitCommands = []initCommand{InitacceptCmd, InitlistCmd}

// InitacceptCmd принимает данные о закази на принятие
func InitacceptCmd(parentCmd *cobra.Command, c *cli.CLI) {
	acceptCmd.Flags().IntP("order_id", "o", -1, "ID заказа(*)")
	acceptCmd.Flags().IntP("user_id", "u", -1, "ID получателя(*)")

	acceptCmd.RunE = func(cmd *cobra.Command, args []string) error {
		defer flags.ResetFlags(cmd)
		orderID, err := cmd.Flags().GetInt("order_id")
		if err != nil {
			return err
		}
		userID, err := cmd.Flags().GetInt("user_id")
		if err != nil {
			return err
		}

		errAcceptReturn := c.Module.AcceptReturn(&models.Order{
			ID:   orderID,
			User: &models.User{ID: userID},
		})

		if errAcceptReturn != nil {
			return errAcceptReturn
		}

		fmt.Println("Заказ возвращен.")
		return nil
	}

	parentCmd.AddCommand(acceptCmd)
}

// InitlistCmd принимает данные о закази на принятие
func InitlistCmd(parentCmd *cobra.Command, c *cli.CLI) {
	listCmd.Flags().IntP("start", "s", -1, "нижняя граница по количеству заказов в списке")
	listCmd.Flags().IntP("finish", "f", -1, "верхняя граница по количеству заказов в списке")

	listCmd.RunE = func(cmd *cobra.Command, args []string) error {
		defer flags.ResetFlags(cmd)

		start, err := cmd.Flags().GetInt("start")
		if err != nil {
			return err
		}
		finish, err := cmd.Flags().GetInt("finish")
		if err != nil {
			return err
		}

		list, errListReturns := c.Module.ListReturns(start, finish)
		if errListReturns != nil {
			return errListReturns
		}

		for _, v := range list {
			fmt.Printf("Возврат: %v. Получатель: %v.\n", v.ID, v.User.ID)
		}
		return nil
	}

	parentCmd.AddCommand(listCmd)
}
