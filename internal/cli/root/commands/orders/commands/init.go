package commands

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/cli"
	e "gitlab.ozon.dev/yuweebix/homework-1/internal/cli/errors"
	f "gitlab.ozon.dev/yuweebix/homework-1/internal/cli/flags"
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
	acceptCmd.Flags().IntP(f.OrderIDL, f.OrderIDS, f.DefaultIntValue, f.OrderIDU)
	acceptCmd.Flags().IntP(f.UserIDL, f.UserIDS, f.DefaultIntValue, f.UserIDU)
	acceptCmd.Flags().StringP(f.ExpiryL, f.ExpiryS, f.DefaultStringValue, f.ExpiryU)

	// помечаем флаги как обязательные
	acceptCmd.MarkFlagRequired(f.OrderIDL)
	acceptCmd.MarkFlagRequired(f.UserIDL)
	acceptCmd.MarkFlagRequired(f.ExpiryL)

	// функционал команды
	acceptCmd.RunE = func(cmd *cobra.Command, args []string) error {
		orderID, err = cmd.Flags().GetInt(f.OrderIDL)
		if err != nil {
			return err
		}
		userID, err = cmd.Flags().GetInt(f.UserIDL)
		if err != nil {
			return err
		}
		expiry, err = cmd.Flags().GetString(f.ExpiryL)
		if err != nil {
			return err
		}

		expiryDate, err := time.Parse(time.DateOnly, expiry)
		if err != nil {
			return e.ErrDateFormatInvalid
		}

		err = c.Service().AcceptOrder(&models.Order{
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
	deliverCmd.Flags().IntSliceP(f.OrderIDsL, f.OrderIDsS, f.DefaultIntSliceValue(), f.OrderIDsU)

	// помечаем флаги как обязательные
	deliverCmd.MarkFlagRequired(f.OrderIDsL)

	// функционал команды
	deliverCmd.RunE = func(cmd *cobra.Command, args []string) error {
		orderIDs, err = cmd.Flags().GetIntSlice(f.OrderIDsL)
		if err != nil {
			return err
		}

		err = c.Service().DeliverOrders(orderIDs)
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
	var isStored bool
	var list []*models.Order
	var err error

	// инициализируем флаги
	listCmd.Flags().IntP(f.UserIDL, f.UserIDS, f.DefaultIntValue, f.UserIDU)
	listCmd.Flags().IntP(f.LimitL, f.LimitS, f.DefaultIntValue, f.LimitU)            // опциональный флаг
	listCmd.Flags().BoolP(f.IsStoredL, f.IsStoredS, f.DefaultBoolValue, f.IsStoredU) // опциональный флаг

	// помечаем флаги как обязательные
	listCmd.MarkFlagRequired(f.UserIDL)

	// функционал команды
	listCmd.RunE = func(cmd *cobra.Command, args []string) error {
		userID, err = cmd.Flags().GetInt(f.UserIDL)
		if err != nil {
			return err
		}
		limit, err = cmd.Flags().GetInt(f.LimitL)
		if err != nil {
			return err
		}
		isStored, err = cmd.Flags().GetBool(f.IsStoredL)
		if err != nil {
			return err
		}

		// стандартное значение
		if limit < 1 {
			limit = 10
		}

		list, err = c.Service().ListOrders(userID, limit, isStored)
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
	returnCmd.Flags().IntP(f.OrderIDL, f.OrderIDS, f.DefaultIntValue, f.OrderIDU)

	// помечаем флаги как обязательные
	returnCmd.MarkFlagRequired(f.OrderIDL)

	// функционал команды
	returnCmd.RunE = func(cmd *cobra.Command, args []string) error {
		orderID, err = cmd.Flags().GetInt(f.OrderIDL)
		if err != nil {
			return err
		}

		err := c.Service().ReturnOrder(&models.Order{
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
