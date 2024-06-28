package cli

import (
	"github.com/spf13/cobra"
	f "gitlab.ozon.dev/yuweebix/homework-1/internal/cli/flags"
)

// инициализируем дефолтные флаги

// orders флаги
var (
	flagOrderID = f.Flag[uint64]{
		Name:      "order_id",
		Shorthand: "o",
		Usage:     "ID заказа(*)",
		Value:     0,
	}
	flagOrderIDs = f.Flag[[]uint64]{
		Name:      "order_ids",
		Shorthand: "o",
		Usage:     "IDs заказов(*)",
		Value:     []uint64{},
	}
	flagUserID = f.Flag[uint64]{
		Name:      "user_id",
		Shorthand: "u",
		Usage:     "ID клиента(*)",
		Value:     0,
	}
	flagExpiry = f.Flag[string]{
		Name:      "expiry",
		Shorthand: "e",
		Usage:     "Срок хранения в формате YYYY-MM-DD(*)",
		Value:     "",
	}
	flagCost = f.Flag[float64]{
		Name:      "cost",
		Shorthand: "c",
		Usage:     "Стоимость заказа(*)",
		Value:     0,
	}
	flagWeight = f.Flag[float64]{
		Name:      "weight",
		Shorthand: "w",
		Usage:     "Вес заказа в кг(*)",
		Value:     0,
	}
	flagPackaging = f.Flag[string]{
		Name:      "packaging",
		Shorthand: "p",
		Usage:     "Тип упаковки (bag, box, wrap)(*)",
		Value:     "",
	}
	flagLimit = f.Flag[uint64]{
		Name:      "limit",
		Shorthand: "l",
		Usage:     "Ограничение по количеству заказов в списке",
		Value:     0,
	}
	flagOffset = f.Flag[uint64]{
		Name:      "offset",
		Shorthand: "s",
		Usage:     "Смещение по количеству заказов в списке",
		Value:     0,
	}
	flagIsStored = f.Flag[bool]{
		Name:      "is_stored",
		Shorthand: "i",
		Usage:     "Показать заказы клиента, находящиеся в нашем ПВЗ",
		Value:     false,
	}
)

// returns флаги
var (
// flagOrderID
// flagUserID
// flagLimit
// flagOffset
)

// workers флаги
var (
	flagWorkersNum = f.Flag[int]{
		Name:      "number",
		Shorthand: "n",
		Usage:     "количество рабочих",
		Value:     0,
	}
)

func ordersAcceptCmdSetFlags(cmd *cobra.Command) (err error) {
	// инициализируем флаги
	cmd.Flags().Uint64P(flagOrderID.Unzip())
	cmd.Flags().Uint64P(flagUserID.Unzip())
	cmd.Flags().StringP(flagExpiry.Unzip())
	cmd.Flags().Float64P(flagCost.Unzip())
	cmd.Flags().Float64P(flagWeight.Unzip())
	cmd.Flags().StringP(flagPackaging.Unzip())

	// помечаем флаги как обязательные
	err = cmd.MarkFlagRequired(flagOrderID.Name)
	if err != nil {
		return err
	}
	err = cmd.MarkFlagRequired(flagUserID.Name)
	if err != nil {
		return err
	}
	err = cmd.MarkFlagRequired(flagExpiry.Name)
	if err != nil {
		return err
	}
	err = cmd.MarkFlagRequired(flagCost.Name)
	if err != nil {
		return err
	}
	err = cmd.MarkFlagRequired(flagWeight.Name)
	if err != nil {
		return err
	}
	err = cmd.MarkFlagRequired(flagPackaging.Name)
	if err != nil {
		return err
	}

	return nil
}

func ordersDeliverCmdSetFlags(cmd *cobra.Command) (err error) {
	// инициализируем флаги
	cmd.Flags().StringP(flagOrderIDs.Name, flagOrderIDs.Shorthand, "", flagOrderIDs.Usage)

	// помечаем флаги как обязательные
	err = cmd.MarkFlagRequired(flagOrderIDs.Name)
	if err != nil {
		return err
	}

	return nil
}

func ordersListCmdSetFlags(cmd *cobra.Command) (err error) {
	// инициализируем флаги
	cmd.Flags().Uint64P(flagUserID.Unzip())
	cmd.Flags().Uint64P(flagLimit.Unzip())  // опциональный флаг
	cmd.Flags().Uint64P(flagOffset.Unzip()) // опциональный флаг
	cmd.Flags().BoolP(flagIsStored.Unzip()) // опциональный флаг

	// помечаем флаги как обязательные
	err = cmd.MarkFlagRequired(flagUserID.Name)
	if err != nil {
		return err
	}

	return nil
}

func ordersReturnCmdSetFlags(cmd *cobra.Command) (err error) {
	// инициализируем флаги
	cmd.Flags().Uint64P(flagOrderID.Unzip())

	// помечаем флаги как обязательные
	err = cmd.MarkFlagRequired(flagOrderID.Name)
	if err != nil {
		return err
	}

	return nil
}

func returnsAcceptCmdSetFlags(cmd *cobra.Command) (err error) {
	// инициализируем флаги
	cmd.Flags().Uint64P(flagOrderID.Unzip())
	cmd.Flags().Uint64P(flagUserID.Unzip())

	// помечаем флаги как обязательные
	err = cmd.MarkFlagRequired(flagOrderID.Name)
	if err != nil {
		return err
	}
	err = cmd.MarkFlagRequired(flagUserID.Name)
	if err != nil {
		return err
	}

	return nil
}

func returnsListCmdSetFlags(cmd *cobra.Command) (_ error) {
	// инициализируем флаги
	cmd.Flags().Uint64P(flagLimit.Unzip())  // опциональный флаг
	cmd.Flags().Uint64P(flagOffset.Unzip()) // опциональный флаг

	return nil
}

func workersCmdSetFlags(cmd *cobra.Command) (err error) {
	// инициализируем флаги
	cmd.Flags().IntP(flagWorkersNum.Unzip())

	// помечаем флаги как обязательные
	err = cmd.MarkFlagRequired(flagWorkersNum.Name)
	if err != nil {
		return err
	}

	return nil
}
