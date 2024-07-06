package cli

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	e "gitlab.ozon.dev/yuweebix/homework-1/internal/cli/errors"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/cli/flags"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
	"gitlab.ozon.dev/yuweebix/homework-1/pkg/utils"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// orders функционал

func (c *CLI) ordersAcceptCmdRunE(cmd *cobra.Command, args []string) (err error) {
	msg := message{
		CreatedAt:  time.Now().UTC(),
		MethodName: getMethodName(cmd),
		RawRequest: getRawRequest(cmd),
	}

	orderID, userID, expiry, cost, weight, packaging, err := c.getOrdersAcceptCmdFlagValues(cmd)
	if err != nil {
		if err := c.producer.Send(topic, messageWithError{msg, err.Error()}); err != nil {
			return err
		}
		return err
	}

	flagExpiryDate, err := time.Parse(time.DateOnly, expiry)
	if err != nil {
		if err := c.producer.Send(topic, messageWithError{msg, e.ErrDateFormatInvalid.Error()}); err != nil {
			return err
		}
		return e.ErrDateFormatInvalid
	}

	ctx := cmd.Context()
	err = c.domain.AcceptOrder(ctx, &models.Order{
		ID:        orderID,
		User:      &models.User{ID: userID},
		Expiry:    flagExpiryDate,
		Cost:      utils.ConvertToMicrocurrency(cost),
		Weight:    weight,
		Packaging: models.PackagingType(packaging),
	})
	if err != nil {
		if err := c.producer.Send(topic, messageWithError{msg, err.Error()}); err != nil {
			return err
		}
		return err
	}

	c.logger.Println("Заказ принят.")

	err = c.producer.Send(topic, msg)
	if err != nil {
		return err
	}

	return nil
}

func (c *CLI) ordersDeliverCmdRunE(cmd *cobra.Command, _ []string) (err error) {
	msg := message{
		CreatedAt:  time.Now().UTC(),
		MethodName: getMethodName(cmd),
		RawRequest: getRawRequest(cmd),
	}

	orderIDs, err := c.getOrdersDeliverCmdFlagValues(cmd)
	if err != nil {
		if err := c.producer.Send(topic, messageWithError{msg, err.Error()}); err != nil {
			return err
		}
		return err
	}

	ctx := cmd.Context()
	err = c.domain.DeliverOrders(ctx, orderIDs)
	if err != nil {
		if err := c.producer.Send(topic, messageWithError{msg, err.Error()}); err != nil {
			return err
		}
		return err
	}

	c.logger.Println("Заказы выданы.")

	err = c.producer.Send(topic, msg)
	if err != nil {
		return err
	}

	return nil
}

func (c *CLI) ordersListCmdRunE(cmd *cobra.Command, _ []string) (err error) {
	msg := message{
		CreatedAt:  time.Now().UTC(),
		MethodName: getMethodName(cmd),
		RawRequest: getRawRequest(cmd),
	}

	userID, limit, offset, isStored, err := c.getOrdersListCmdFlagValues(cmd)
	if err != nil {
		if err := c.producer.Send(topic, messageWithError{msg, err.Error()}); err != nil {
			return err
		}
		return err
	}

	ctx := cmd.Context()
	list, err := c.domain.ListOrders(ctx, userID, limit, offset, isStored)
	if err != nil {
		if err := c.producer.Send(topic, messageWithError{msg, err.Error()}); err != nil {
			return err
		}
		return err
	}

	for _, v := range list {
		c.logger.Printf("Заказ: %v. Получатель: %v. Хранится до %v. Статус: %v\n", v.ID, v.User.ID, v.Expiry, getStatusMessage(v))
	}

	err = c.producer.Send(topic, msg)
	if err != nil {
		return err
	}

	return nil
}

func (c *CLI) ordersReturnCmdRunE(cmd *cobra.Command, _ []string) (err error) {
	msg := message{
		CreatedAt:  time.Now().UTC(),
		MethodName: getMethodName(cmd),
		RawRequest: getRawRequest(cmd),
	}

	orderID, err := c.getOrdersReturnCmdFlagValues(cmd)
	if err != nil {
		if err := c.producer.Send(topic, messageWithError{msg, err.Error()}); err != nil {
			return err
		}
		return err
	}

	ctx := cmd.Context()
	err = c.domain.ReturnOrder(ctx, &models.Order{
		ID: orderID,
	})

	if err != nil {
		if err := c.producer.Send(topic, messageWithError{msg, err.Error()}); err != nil {
			return err
		}
		return err
	}

	c.logger.Println("Заказ вернут курьеру.")

	err = c.producer.Send(topic, msg)
	if err != nil {
		return err
	}

	return nil
}

// returns функционал

func (c *CLI) returnsAcceptCmdRunE(cmd *cobra.Command, _ []string) (err error) {
	msg := message{
		CreatedAt:  time.Now().UTC(),
		MethodName: getMethodName(cmd),
		RawRequest: getRawRequest(cmd),
	}

	orderID, userID, err := c.getReturnsAcceptCmdFlagValues(cmd)
	if err != nil {
		if err := c.producer.Send(topic, messageWithError{msg, err.Error()}); err != nil {
			return err
		}
		return err
	}

	ctx := cmd.Context()
	err = c.domain.AcceptReturn(ctx, &models.Order{
		ID:   orderID,
		User: &models.User{ID: userID},
	})
	if err != nil {
		if err := c.producer.Send(topic, messageWithError{msg, err.Error()}); err != nil {
			return err
		}
		return err
	}

	c.logger.Println("Заказ возвращен.")

	err = c.producer.Send(topic, msg)
	if err != nil {
		return err
	}

	return nil
}

func (c *CLI) returnsListCmdRunE(cmd *cobra.Command, _ []string) (err error) {
	msg := message{
		CreatedAt:  time.Now().UTC(),
		MethodName: getMethodName(cmd),
		RawRequest: getRawRequest(cmd),
	}

	limit, offset, err := c.getReturnsListCmdFlagValues(cmd)
	if err != nil {
		if err := c.producer.Send(topic, messageWithError{msg, err.Error()}); err != nil {
			return err
		}
		return err
	}

	ctx := cmd.Context()
	list, err := c.domain.ListReturns(ctx, limit, offset)
	if err != nil {
		if err := c.producer.Send(topic, messageWithError{msg, err.Error()}); err != nil {
			return err
		}
		return err
	}

	for _, v := range list {
		c.logger.Printf("Возврат: %v. Получатель: %v.\n", v.ID, v.User.ID)
	}

	err = c.producer.Send(topic, msg)
	if err != nil {
		return err
	}

	return nil
}

// worker функционал

func (c *CLI) workersCmdRunE(cmd *cobra.Command, _ []string) (err error) {
	msg := message{
		CreatedAt:  time.Now().UTC(),
		MethodName: getMethodName(cmd),
		RawRequest: getRawRequest(cmd),
	}

	num, err := c.getWorkersCmdFlagValues(cmd)
	if err != nil {
		if err := c.producer.Send(topic, messageWithError{msg, err.Error()}); err != nil {
			return err
		}
		return err
	}

	ctx := cmd.Context()
	err = c.domain.ChangeWorkersNumber(ctx, num)
	if err != nil {
		if err := c.producer.Send(topic, messageWithError{msg, err.Error()}); err != nil {
			return err
		}
		return err
	}

	c.logger.Println("Количество рабочих горутин было изменено.")

	err = c.producer.Send(topic, msg)
	if err != nil {
		return err
	}

	return nil
}

// вспомогательные методы

// getStatusMessage вспомогательная функция для orders list
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

// stringToUint64Slice конвертирует строку вида "1,2,3" в слайс uint64
func stringToUint64Slice(s string) ([]uint64, error) {
	idsStrSlice := strings.Split(s, ",")
	orderIDs := make([]uint64, len(idsStrSlice))
	for i, idStr := range idsStrSlice {
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			return nil, err
		}
		orderIDs[i] = id
	}
	return orderIDs, nil
}

// методы получения значений флагов

func (c *CLI) getOrdersAcceptCmdFlagValues(cmd *cobra.Command) (orderID, userID uint64, expiry string, cost, weight float64, packaging string, err error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	defer flags.ResetFlags(cmd)

	orderID, err = cmd.Flags().GetUint64(flagOrderID.Name)
	if err != nil {
		return
	}
	userID, err = cmd.Flags().GetUint64(flagUserID.Name)
	if err != nil {
		return
	}
	expiry, err = cmd.Flags().GetString(flagExpiry.Name)
	if err != nil {
		return
	}
	cost, err = cmd.Flags().GetFloat64(flagCost.Name)
	if err != nil {
		return
	}
	weight, err = cmd.Flags().GetFloat64(flagWeight.Name)
	if err != nil {
		return
	}
	packaging, err = cmd.Flags().GetString(flagPackaging.Name)
	if err != nil {
		return
	}

	return
}

func (c *CLI) getOrdersDeliverCmdFlagValues(cmd *cobra.Command) (orderIDs []uint64, err error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	defer flags.ResetFlags(cmd)

	idsStr, err := cmd.Flags().GetString(flagOrderIDs.Name)
	if err != nil {
		return nil, err
	}

	orderIDs, err = stringToUint64Slice(idsStr)
	if err != nil {
		return nil, err
	}

	return orderIDs, nil
}

func (c *CLI) getOrdersListCmdFlagValues(cmd *cobra.Command) (userID, limit, offset uint64, isStored bool, err error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	defer flags.ResetFlags(cmd)

	userID, err = cmd.Flags().GetUint64(flagUserID.Name)
	if err != nil {
		return
	}
	limit, err = cmd.Flags().GetUint64(flagLimit.Name)
	if err != nil {
		return
	}
	offset, err = cmd.Flags().GetUint64(flagOffset.Name)
	if err != nil {
		return
	}
	isStored, err = cmd.Flags().GetBool(flagIsStored.Name)
	if err != nil {
		return
	}

	return
}

func (c *CLI) getOrdersReturnCmdFlagValues(cmd *cobra.Command) (orderID uint64, err error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	defer flags.ResetFlags(cmd)

	orderID, err = cmd.Flags().GetUint64(flagOrderID.Name)
	if err != nil {
		return
	}

	return
}

func (c *CLI) getReturnsAcceptCmdFlagValues(cmd *cobra.Command) (orderID, userID uint64, err error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	defer flags.ResetFlags(cmd)

	orderID, err = cmd.Flags().GetUint64(flagOrderID.Name)
	if err != nil {
		return
	}
	userID, err = cmd.Flags().GetUint64(flagUserID.Name)
	if err != nil {
		return
	}

	return
}

func (c *CLI) getReturnsListCmdFlagValues(cmd *cobra.Command) (limit, offset uint64, err error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	defer flags.ResetFlags(cmd)

	limit, err = cmd.Flags().GetUint64(flagLimit.Name)
	if err != nil {
		return
	}
	offset, err = cmd.Flags().GetUint64(flagOffset.Name)
	if err != nil {
		return
	}

	return
}

func (c *CLI) getWorkersCmdFlagValues(cmd *cobra.Command) (num int, err error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	defer flags.ResetFlags(cmd)

	num, err = cmd.Flags().GetInt(flagWorkersNum.Name)
	if err != nil {
		return
	}

	return
}

func getRawRequest(cmd *cobra.Command) string {
	var rawRequest strings.Builder
	rawRequest.WriteString(cmd.CommandPath())

	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		if flag.Changed {
			rawRequest.WriteString(fmt.Sprintf(" --%s %s", flag.Name, flag.Value))
		}
	})

	return rawRequest.String()
}

func getMethodName(cmd *cobra.Command) string {
	commandPath := cmd.CommandPath()
	components := strings.Split(commandPath, " ")
	titleCaser := cases.Title(language.Und)
	for i := range components {
		components[i] = titleCaser.String(components[i])
	}
	return strings.Join(components, "")
}
