package commands

import (
	"fmt"
	"math"

	"github.com/spf13/cobra"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/cli"
	f "gitlab.ozon.dev/yuweebix/homework-1/internal/cli/flags"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
)

type initCommand func(*cobra.Command, *cli.CLI)

var InitCommands = []initCommand{InitAcceptCmd, InitListCmd}

// InitAcceptCmd принимает данные о заказе на принятие
func InitAcceptCmd(parentCmd *cobra.Command, c *cli.CLI) {
	var orderID, userID int
	var err error

	// инициализируем флаги
	acceptCmd.Flags().IntP(f.OrderIDL, f.OrderIDS, f.DefaultIntValue, f.OrderIDU)
	acceptCmd.Flags().IntP(f.UserIDL, f.UserIDS, f.DefaultIntValue, f.UserIDU)

	// помечаем флаги как обязательные
	acceptCmd.MarkFlagRequired(f.OrderIDL)
	acceptCmd.MarkFlagRequired(f.UserIDL)

	acceptCmd.RunE = func(cmd *cobra.Command, args []string) error {
		orderID, err = cmd.Flags().GetInt(f.OrderIDL)
		if err != nil {
			return err
		}
		userID, err = cmd.Flags().GetInt(f.UserIDL)
		if err != nil {
			return err
		}

		err := c.Service().AcceptReturn(&models.Order{
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
	listCmd.Flags().IntP(f.StartL, f.StartS, f.DefaultIntValue, f.StartU)    // опциональный флаг
	listCmd.Flags().IntP(f.FinishL, f.FinishS, f.DefaultIntValue, f.FinishU) // опциональный флаг

	listCmd.RunE = func(cmd *cobra.Command, args []string) error {
		start, err = cmd.Flags().GetInt(f.StartL)
		if err != nil {
			return err
		}
		finish, err = cmd.Flags().GetInt(f.FinishL)
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

		list, err = c.Service().ListReturns(start, finish)
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
