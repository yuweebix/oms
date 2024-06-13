package cli

import (
	"time"

	"github.com/spf13/cobra"
	e "gitlab.ozon.dev/yuweebix/homework-1/internal/cli/errors"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/cli/flags"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
)

type initCommand func(*cobra.Command)

// Функции для получения списков команд и подкоманд
func getInitCommands(c *CLI) []initCommand {
	return []initCommand{c.initOrdersCmd, c.initReturnsCmd, c.initWorkersCmd}
}

func getInitOrdersSubcommands(c *CLI) []initCommand {
	return []initCommand{c.initOrdersAcceptCmd, c.initOrdersDeliverCmd, c.initOrdersListCmd, c.initOrdersReturnCmd}
}

func getInitReturnsSubcommands(c *CLI) []initCommand {
	return []initCommand{c.initReturnsAcceptCmd, c.initReturnsListCmd}
}

// initRootCmd инициализация стоковой команды
func (c *CLI) initRootCmd() {
	for _, init := range getInitCommands(c) {
		init(rootCmd)
	}
}

// initOrdersCmd инициализация перечня команд orders
func (c *CLI) initOrdersCmd(parentCmd *cobra.Command) {
	for _, init := range getInitOrdersSubcommands(c) {
		init(ordersCmd)
	}

	parentCmd.AddCommand(ordersCmd)
}

// initReturnCmd инициализация перечня команд returns
func (c *CLI) initReturnsCmd(parentCmd *cobra.Command) {
	for _, init := range getInitReturnsSubcommands(c) {
		init(returnsCmd)
	}

	parentCmd.AddCommand(returnsCmd)
}

// orders инициализация подкоманд

// initOrdersAcceptCmd принимает данные о заказе на принятие
func (c *CLI) initOrdersAcceptCmd(parentCmd *cobra.Command) {
	// инициализируем флаги
	ordersAcceptCmd.Flags().IntP(flagOrderID.Unzip())
	ordersAcceptCmd.Flags().IntP(flagUserID.Unzip())
	ordersAcceptCmd.Flags().StringP(flagExpiry.Unzip())

	// помечаем флаги как обязательные
	ordersAcceptCmd.MarkFlagRequired(flagOrderID.Name)
	ordersAcceptCmd.MarkFlagRequired(flagUserID.Name)
	ordersAcceptCmd.MarkFlagRequired(flagExpiry.Name)

	// функционал команды
	ordersAcceptCmd.RunE = func(cmd *cobra.Command, args []string) (err error) {
		orderID, userID, expiry, err := func() (orderID, userID int, expiry string, err error) {
			c.mu.Lock()
			defer c.mu.Unlock()
			defer flags.ResetFlags(ordersAcceptCmd)

			orderID, err = cmd.Flags().GetInt(flagOrderID.Name)
			if err != nil {
				return
			}
			userID, err = cmd.Flags().GetInt(flagUserID.Name)
			if err != nil {
				return
			}
			expiry, err = cmd.Flags().GetString(flagExpiry.Name)
			if err != nil {
				return
			}

			return
		}()
		if err != nil {
			return err
		}

		flagExpiryDate, err := time.Parse(time.DateOnly, expiry)
		if err != nil {
			return e.ErrDateFormatInvalid
		}

		err = c.service.AcceptOrder(&models.Order{
			ID:     orderID,
			User:   &models.User{ID: userID},
			Expiry: flagExpiryDate,
		})
		if err != nil {
			return err
		}

		c.logger.Println("Заказ принят.")
		return nil
	}

	parentCmd.AddCommand(ordersAcceptCmd)
}

// initOrdersDeliverCmd принимает данные о заказе на принятие
func (c *CLI) initOrdersDeliverCmd(parentCmd *cobra.Command) {
	// инициализируем флаги
	ordersDeliverCmd.Flags().IntSliceP(flagOrderIDs.Unzip())

	// помечаем флаги как обязательные
	ordersDeliverCmd.MarkFlagRequired(flagOrderIDs.Name)

	// функционал команды
	ordersDeliverCmd.RunE = func(cmd *cobra.Command, args []string) (err error) {
		orderIDs, err := func() (orderIDs []int, err error) {
			c.mu.Lock()
			defer c.mu.Unlock()
			defer flags.ResetFlags(ordersDeliverCmd)

			orderIDs, err = cmd.Flags().GetIntSlice(flagOrderIDs.Name)
			if err != nil {
				return
			}

			return
		}()
		if err != nil {
			return err
		}

		err = c.service.DeliverOrders(orderIDs)
		if err != nil {
			return err
		}

		c.logger.Println("Заказы выданы.")
		return nil
	}

	parentCmd.AddCommand(ordersDeliverCmd)
}

// initOrdersListCmd принимает данные о заказе на принятие
func (c *CLI) initOrdersListCmd(parentCmd *cobra.Command) {
	// инициализируем флаги
	ordersListCmd.Flags().IntP(flagUserID.Unzip())
	ordersListCmd.Flags().IntP(flagLimit.Unzip())     // опциональный флаг
	ordersListCmd.Flags().BoolP(flagIsStored.Unzip()) // опциональный флаг

	// помечаем флаги как обязательные
	ordersListCmd.MarkFlagRequired(flagUserID.Name)

	// функционал команды
	ordersListCmd.RunE = func(cmd *cobra.Command, args []string) (err error) {
		userID, limit, isStored, err := func() (userID, limit int, isStored bool, err error) {
			c.mu.Lock()
			defer c.mu.Unlock()
			defer flags.ResetFlags(ordersListCmd)

			userID, err = cmd.Flags().GetInt(flagUserID.Name)
			if err != nil {
				return
			}
			limit, err = cmd.Flags().GetInt(flagLimit.Name)
			if err != nil {
				return
			}
			isStored, err = cmd.Flags().GetBool(flagIsStored.Name)
			if err != nil {
				return
			}

			return
		}()
		if err != nil {
			return err
		}

		list, err := c.service.ListOrders(userID, limit, isStored)
		if err != nil {
			return err
		}

		for _, v := range list {
			c.logger.Printf("Заказ: %v. Получатель: %v. Хранится до %v. Статус: %v\n", v.ID, v.User.ID, v.Expiry, getStatusMessage(v))
		}
		return nil
	}

	parentCmd.AddCommand(ordersListCmd)
}

// initOrdersReturnCmd принимает данные о заказе на принятие
func (c *CLI) initOrdersReturnCmd(parentCmd *cobra.Command) {
	// инициализируем флаги
	ordersReturnCmd.Flags().IntP(flagOrderID.Unzip())

	// помечаем флаги как обязательные
	ordersReturnCmd.MarkFlagRequired(flagOrderID.Name)

	// функционал команды
	ordersReturnCmd.RunE = func(cmd *cobra.Command, args []string) (err error) {
		orderID, err := func() (orderID int, err error) {
			c.mu.Lock()
			defer c.mu.Unlock()
			defer flags.ResetFlags(ordersReturnCmd)

			orderID, err = cmd.Flags().GetInt(flagOrderID.Name)
			if err != nil {
				return
			}

			return
		}()
		if err != nil {
			return err
		}

		err = c.service.ReturnOrder(&models.Order{
			ID: orderID,
		})

		if err != nil {
			return err
		}

		c.logger.Println("Заказ вернут курьеру.")
		return nil
	}

	parentCmd.AddCommand(ordersReturnCmd)
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

// returns инициализация подкоманд

// initReturnsAcceptCmd принимает данные о заказе на принятие
func (c *CLI) initReturnsAcceptCmd(parentCmd *cobra.Command) {
	// инициализируем флаги
	returnsAcceptCmd.Flags().IntP(flagOrderID.Unzip())
	returnsAcceptCmd.Flags().IntP(flagUserID.Unzip())

	// помечаем флаги как обязательные
	returnsAcceptCmd.MarkFlagRequired(flagOrderID.Name)
	returnsAcceptCmd.MarkFlagRequired(flagUserID.Name)

	returnsAcceptCmd.RunE = func(cmd *cobra.Command, args []string) (err error) {
		orderID, userID, err := func() (orderID, userID int, err error) {
			c.mu.Lock()
			defer c.mu.Unlock()
			defer flags.ResetFlags(returnsAcceptCmd)

			orderID, err = cmd.Flags().GetInt(flagOrderID.Name)
			if err != nil {
				return
			}
			userID, err = cmd.Flags().GetInt(flagUserID.Name)
			if err != nil {
				return
			}

			return
		}()
		if err != nil {
			return err
		}

		err = c.service.AcceptReturn(&models.Order{
			ID:   orderID,
			User: &models.User{ID: userID},
		})
		if err != nil {
			return err
		}

		c.logger.Println("Заказ возвращен.")
		return nil
	}

	parentCmd.AddCommand(returnsAcceptCmd)
}

// initReturnsListCmd принимает данные о заказе на принятие
func (c *CLI) initReturnsListCmd(parentCmd *cobra.Command) {
	// инициализируем флаги
	returnsListCmd.Flags().IntP(flagStart.Unzip())  // опциональный флаг
	returnsListCmd.Flags().IntP(flagFinish.Unzip()) // опциональный флаг

	returnsListCmd.RunE = func(cmd *cobra.Command, args []string) (err error) {
		start, finish, err := func() (start, finish int, err error) {
			c.mu.Lock()
			defer c.mu.Unlock()
			defer flags.ResetFlags(returnsListCmd)

			start, err = cmd.Flags().GetInt(flagStart.Name)
			if err != nil {
				return
			}
			finish, err = cmd.Flags().GetInt(flagFinish.Name)
			if err != nil {
				return
			}

			return
		}()
		if err != nil {
			return err
		}

		list, err := c.service.ListReturns(start, finish)
		if err != nil {
			return err
		}

		for _, v := range list {
			c.logger.Printf("Возврат: %v. Получатель: %v.\n", v.ID, v.User.ID)
		}
		return nil
	}

	parentCmd.AddCommand(returnsListCmd)
}

// initWorkersCmd принимает количество рабочих горутин
func (c *CLI) initWorkersCmd(parentCmd *cobra.Command) {
	// инициализируем флаги
	workersCmd.Flags().IntP(flagWorkersNum.Unzip())

	workersCmd.MarkFlagRequired(flagWorkersNum.Name)

	workersCmd.RunE = func(cmd *cobra.Command, args []string) (err error) {
		num, err := func() (num int, err error) {
			c.mu.Lock()
			defer c.mu.Unlock()
			defer flags.ResetFlags(workersCmd)

			num, err = cmd.Flags().GetInt(flagWorkersNum.Name)
			if err != nil {
				return
			}

			return
		}()
		if err != nil {
			return err
		}

		err = c.service.ChangeWorkersNumber(num)
		if err != nil {
			return err
		}

		c.logger.Println("Количество рабочих горутин было изменено.")
		return nil
	}

	parentCmd.AddCommand(workersCmd)
}
