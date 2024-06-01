package commands

import (
	"fmt"
	"math"

	"github.com/spf13/cobra"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/cli"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/cli/flags"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
)

type initCommand func(*cobra.Command, *cli.CLI)

var InitCommands = []initCommand{InitAcceptCmd, InitListCmd}

// InitAcceptCmd принимает данные о заказе на принятие
func InitAcceptCmd(parentCmd *cobra.Command, c *cli.CLI) {
	var orderID, userID int
	var err error

	// инициализируем флаги
	acceptCmd.Flags().IntP("order_id", "o", flags.DefaultIntValue, "ID заказа(*)")
	acceptCmd.Flags().IntP("user_id", "u", flags.DefaultIntValue, "ID получателя(*)")

	// помечаем флаги как обязательные
	acceptCmd.MarkFlagRequired("order_id")
	acceptCmd.MarkFlagRequired("user_id")

	acceptCmd.RunE = func(cmd *cobra.Command, args []string) error {
		defer flags.ResetFlags(cmd)

		orderID, err = cmd.Flags().GetInt("order_id")
		if err != nil {
			return err
		}
		userID, err = cmd.Flags().GetInt("user_id")
		if err != nil {
			return err
		}

		err := c.Service.AcceptReturn(&models.Order{
			ID:   orderID,
			User: &models.User{ID: userID},
		})
		if err != nil {
			return err
		}

		fmt.Println("Заказ возвращен.")
		return nil
	}

	parentCmd.AddCommand(acceptCmd)
}

// InitListCmd принимает данные о заказе на принятие
func InitListCmd(parentCmd *cobra.Command, c *cli.CLI) {
	var start, finish int
	var list []*models.Order
	var err error

	// инициализируем флаги
	listCmd.Flags().IntP("start", "s", flags.DefaultIntValue, "нижняя граница по количеству заказов в списке")   // опциональный флаг
	listCmd.Flags().IntP("finish", "f", flags.DefaultIntValue, "верхняя граница по количеству заказов в списке") // опциональный флаг

	listCmd.RunE = func(cmd *cobra.Command, args []string) error {
		defer flags.ResetFlags(cmd)

		start, err = cmd.Flags().GetInt("start")
		if err != nil {
			return err
		}
		finish, err = cmd.Flags().GetInt("finish")
		if err != nil {
			return err
		}

		// стандартное значение
		if start < 0 {
			start = 1
		}
		if finish < 0 {
			finish = math.MaxInt
		}

		list, err = c.Service.ListReturns(start, finish)
		if err != nil {
			return err
		}

		for _, v := range list {
			fmt.Printf("Возврат: %v. Получатель: %v.\n", v.ID, v.User.ID)
		}
		return nil
	}

	parentCmd.AddCommand(listCmd)
}
