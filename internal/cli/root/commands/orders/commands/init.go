package commands

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/cli"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
	"gitlab.ozon.dev/yuweebix/homework-1/pkg/flags"
)

type initCommand func(*cobra.Command, *cli.CLI)

var InitCommands = []initCommand{InitacceptCmd, InitdeliverCmd, InitlistCmd, InitreturnCmd}

// InitacceptCmd принимает данные о закази на принятие
func InitacceptCmd(parentCmd *cobra.Command, c *cli.CLI) {
	acceptCmd.Flags().IntP("order_id", "o", -1, "ID заказа(*)")
	acceptCmd.Flags().IntP("user_id", "u", -1, "ID получателя(*)")
	acceptCmd.Flags().IntP("expiry", "e", -1, "Срок хранения в днях(*)")

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
		expiry, err := cmd.Flags().GetInt("expiry")
		if err != nil {
			return err
		}

		errAcceptOrder := c.Module.AcceptOrder(&models.Order{
			ID:     orderID,
			User:   &models.User{ID: userID},
			Expiry: time.Now().UTC().AddDate(0, 0, expiry),
		})

		if errAcceptOrder != nil {
			return errAcceptOrder
		}

		fmt.Println("Заказ принят.")
		return nil
	}

	parentCmd.AddCommand(acceptCmd)
}

// InitdeliverCmd принимает данные о закази на принятие
func InitdeliverCmd(parentCmd *cobra.Command, c *cli.CLI) {
	deliverCmd.Flags().IntSliceP("order_ids", "o", []int{}, "Список ID заказов")

	deliverCmd.RunE = func(cmd *cobra.Command, args []string) error {
		defer flags.ResetFlags(cmd)

		orderIDSlice, err := cmd.Flags().GetIntSlice("order_ids")
		if err != nil {
			return err
		}

		errDeliverOrders := c.Module.DeliverOrders(orderIDSlice)
		if errDeliverOrders != nil {
			return errDeliverOrders
		}
		fmt.Println("Заказы выданы.")
		return nil
	}

	parentCmd.AddCommand(deliverCmd)
}

// InitlistCmd принимает данные о закази на принятие
func InitlistCmd(parentCmd *cobra.Command, c *cli.CLI) {
	listCmd.Flags().IntP("limit", "l", -1, "ограничение по количеству заказов в списке")

	listCmd.RunE = func(cmd *cobra.Command, args []string) error {
		defer flags.ResetFlags(cmd)

		limit, err := cmd.Flags().GetInt("limit")
		if err != nil {
			return err
		}

		list, errListOrders := c.Module.ListOrders(limit)
		if errListOrders != nil {
			return errListOrders
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

// InitreturnCmd принимает данные о закази на принятие
func InitreturnCmd(parentCmd *cobra.Command, c *cli.CLI) {
	returnCmd.Flags().IntP("order_id", "o", -1, "ID заказа(*)")

	returnCmd.RunE = func(cmd *cobra.Command, args []string) error {
		defer flags.ResetFlags(cmd)

		orderID, err := cmd.Flags().GetInt("order_id")
		if err != nil {
			return err
		}

		errReturnOrder := c.Module.ReturnOrder(&models.Order{
			ID: orderID,
		})

		if errReturnOrder != nil {
			return errReturnOrder
		}

		fmt.Println("Заказ вернут курьеру")
		return nil
	}

	parentCmd.AddCommand(returnCmd)
}
