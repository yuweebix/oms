package commands

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/cli"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/cli/flags"
	e "gitlab.ozon.dev/yuweebix/homework-1/internal/errors"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
)

type initCommand func(*cobra.Command, *cli.CLI)

var InitCommands = []initCommand{InitAcceptCmd, InitDeliverCmd, InitListCmd, InitReturnCmd}

// InitAcceptCmd принимает данные о заказе на принятие
func InitAcceptCmd(parentCmd *cobra.Command, c *cli.CLI) {
	var orderID, userID int
	var expiry string
	var err error

	// инициализируем флаги
	acceptCmd.Flags().IntP("order_id", "o", 0, "ID заказа(*)")
	acceptCmd.Flags().IntP("user_id", "u", 0, "ID получателя(*)")
	acceptCmd.Flags().StringP("expiry", "e", "", "Срок хранения в формате YYYY-MM-DD(*)")

	// помечаем флаги как обязательные
	acceptCmd.MarkFlagRequired("order_id")
	acceptCmd.MarkFlagRequired("user_id")
	acceptCmd.MarkFlagRequired("expiry")

	// функционал команды
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
		expiry, err = cmd.Flags().GetString("expiry")
		if err != nil {
			return err
		}

		expiryDate, err := time.Parse("2006-01-02", expiry)
		if err != nil {
			return e.ErrDateFormatInvalid
		}

		err = c.Service.AcceptOrder(&models.Order{
			ID:     orderID,
			User:   &models.User{ID: userID},
			Expiry: expiryDate,
		})
		if err != nil {
			return err
		}

		fmt.Println("Заказ принят.")
		return nil
	}

	parentCmd.AddCommand(acceptCmd)
}

// InitDeliverCmd принимает данные о заказе на принятие
func InitDeliverCmd(parentCmd *cobra.Command, c *cli.CLI) {
	var orderIDs []int
	var err error

	// инициализируем флаги
	deliverCmd.Flags().IntSliceP("order_ids", "o", flags.DefaultIntSliceValue(), "Список ID заказов")

	// помечаем флаги как обязательные
	deliverCmd.MarkFlagRequired("orders_id")

	// функционал команды
	deliverCmd.RunE = func(cmd *cobra.Command, args []string) error {
		defer flags.ResetFlags(cmd)

		orderIDs, err = cmd.Flags().GetIntSlice("order_ids")
		if err != nil {
			return err
		}

		err = c.Service.DeliverOrders(orderIDs)
		if err != nil {
			return err
		}
		fmt.Println("Заказы выданы.")
		return nil
	}

	parentCmd.AddCommand(deliverCmd)
}

// InitListCmd принимает данные о заказе на принятие
func InitListCmd(parentCmd *cobra.Command, c *cli.CLI) {
	var userID int
	var limit int
	var list []*models.Order
	var err error

	// инициализируем флаги
	listCmd.Flags().IntP("user_id", "u", flags.DefaultIntValue, "ID клиента(*)")
	listCmd.Flags().IntP("limit", "l", flags.DefaultIntValue, "ограничение по количеству заказов в списке") // опциональный флаг

	// помечаем флаги как обязательные
	listCmd.MarkFlagRequired("user_id")

	// функционал команды
	listCmd.RunE = func(cmd *cobra.Command, args []string) error {
		defer flags.ResetFlags(cmd)

		userID, err = cmd.Flags().GetInt("user_id")
		if err != nil {
			return err
		}
		limit, err = cmd.Flags().GetInt("limit")
		if err != nil {
			return err
		}

		// стандартное значение
		if limit < 1 {
			limit = 10
		}

		list, err = c.Service.ListOrders(userID, limit)
		if err != nil {
			return err
		}

		for _, v := range list {
			var status string
			if v.Status == models.StatusAccepted {
				status = "Принят"
			} else if v.Status == models.StatusDelivered {
				status = "Забран"
			} else if v.Status == models.StatusReturned {
				status = "Возвращен"
			}
			fmt.Printf("Заказ: %v. Получатель: %v. Хранится до %v. Статус: %v\n", v.ID, v.User.ID, v.Expiry, status)
		}
		return nil
	}

	parentCmd.AddCommand(listCmd)
}

// InitReturnCmd принимает данные о заказе на принятие
func InitReturnCmd(parentCmd *cobra.Command, c *cli.CLI) {
	var orderID int
	var err error

	// инициализируем флаги
	returnCmd.Flags().IntP("order_id", "o", flags.DefaultIntValue, "ID заказа(*)")

	// помечаем флаги как обязательные
	returnCmd.MarkFlagRequired("orders_id")

	// функционал команды
	returnCmd.RunE = func(cmd *cobra.Command, args []string) error {
		defer flags.ResetFlags(cmd)

		orderID, err = cmd.Flags().GetInt("order_id")
		if err != nil {
			return err
		}

		err := c.Service.ReturnOrder(&models.Order{
			ID: orderID,
		})

		if err != nil {
			return err
		}

		fmt.Println("Заказ вернут курьеру")
		return nil
	}

	parentCmd.AddCommand(returnCmd)
}
