package commands

import (
	"time"

	"github.com/spf13/cobra"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/cli"
	e "gitlab.ozon.dev/yuweebix/homework-1/internal/cli/errors"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
)

type initCommand func(*cobra.Command, *cli.CLI)

var InitCommands = []initCommand{InitAcceptCmd, InitDeliverCmd, InitListCmd, InitReturnCmd}

// InitAcceptCmd принимает данные о заказе на принятие
func InitAcceptCmd(parentCmd *cobra.Command, c *cli.CLI) {
	// инициализируем флаги
	acceptCmd.Flags().IntP(flagOrderID.Unzip())
	acceptCmd.Flags().IntP(flagUserID.Unzip())
	acceptCmd.Flags().StringP(flagExpiry.Unzip())

	// помечаем флаги как обязательные
	acceptCmd.MarkFlagRequired(flagOrderID.Name)
	acceptCmd.MarkFlagRequired(flagUserID.Name)
	acceptCmd.MarkFlagRequired(flagExpiry.Name)

	// функционал команды
	acceptCmd.RunE = func(cmd *cobra.Command, args []string) (err error) {
		orderID, err := cmd.Flags().GetInt(flagOrderID.Name)
		if err != nil {
			return err
		}
		userID, err := cmd.Flags().GetInt(flagUserID.Name)
		if err != nil {
			return err
		}
		expiry, err := cmd.Flags().GetString(flagExpiry.Name)
		if err != nil {
			return err
		}

		flagExpiryDate, err := time.Parse(time.DateOnly, expiry)
		if err != nil {
			return e.ErrDateFormatInvalid
		}

		err = c.Service().AcceptOrder(&models.Order{
			ID:     orderID,
			User:   &models.User{ID: userID},
			Expiry: flagExpiryDate,
		})
		if err != nil {
			return err
		}

		c.Logger.Println("Заказ принят.")
		return nil
	}

	parentCmd.AddCommand(acceptCmd)
}

// InitDeliverCmd принимает данные о заказе на принятие
func InitDeliverCmd(parentCmd *cobra.Command, c *cli.CLI) {
	// инициализируем флаги
	deliverCmd.Flags().IntSliceP(flagOrderIDs.Unzip())

	// помечаем флаги как обязательные
	deliverCmd.MarkFlagRequired(flagOrderIDs.Name)

	// функционал команды
	deliverCmd.RunE = func(cmd *cobra.Command, args []string) (err error) {
		orderIDs, err := cmd.Flags().GetIntSlice(flagOrderIDs.Name)
		if err != nil {
			return
		}

		err = c.Service().DeliverOrders(orderIDs)
		if err != nil {
			return err
		}

		c.Logger.Println("Заказы выданы.")
		return nil
	}

	parentCmd.AddCommand(deliverCmd)
}

// InitListCmd принимает данные о заказе на принятие
func InitListCmd(parentCmd *cobra.Command, c *cli.CLI) {
	// инициализируем флаги
	listCmd.Flags().IntP(flagUserID.Unzip())
	listCmd.Flags().IntP(flagLimit.Unzip())     // опциональный флаг
	listCmd.Flags().BoolP(flagIsStored.Unzip()) // опциональный флаг

	// помечаем флаги как обязательные
	listCmd.MarkFlagRequired(flagUserID.Name)

	// функционал команды
	listCmd.RunE = func(cmd *cobra.Command, args []string) (err error) {
		userID, err := cmd.Flags().GetInt(flagUserID.Name)
		if err != nil {
			return err
		}
		limit, err := cmd.Flags().GetInt(flagLimit.Name)
		if err != nil {
			return err
		}
		isStored, err := cmd.Flags().GetBool(flagIsStored.Name)
		if err != nil {
			return err
		}

		list, err := c.Service().ListOrders(userID, limit, isStored)
		if err != nil {
			return err
		}

		for _, v := range list {
			c.Logger.Printf("Заказ: %v. Получатель: %v. Хранится до %v. Статус: %v\n", v.ID, v.User.ID, v.Expiry, getStatusMessage(v))
		}
		return nil
	}

	parentCmd.AddCommand(listCmd)
}

// InitReturnCmd принимает данные о заказе на принятие
func InitReturnCmd(parentCmd *cobra.Command, c *cli.CLI) {
	// инициализируем флаги
	returnCmd.Flags().IntP(flagOrderID.Unzip())

	// помечаем флаги как обязательные
	returnCmd.MarkFlagRequired(flagOrderID.Name)

	// функционал команды
	returnCmd.RunE = func(cmd *cobra.Command, args []string) (err error) {
		orderID, err := cmd.Flags().GetInt(flagOrderID.Name)
		if err != nil {
			return err
		}

		err = c.Service().ReturnOrder(&models.Order{
			ID: orderID,
		})

		if err != nil {
			return err
		}

		c.Logger.Println("Заказ вернут курьеру.")
		return nil
	}

	parentCmd.AddCommand(returnCmd)
}

func getStatusMessage(o *models.Order) (status string) {
	if o.Status == models.StatusAccepted {
		status = "Принят"
	} else if o.Status == models.StatusDelivered {
		status = "Забран"
	} else if o.Status == models.StatusReturned {
		status = "Возвращен"
	}
	return status
}
