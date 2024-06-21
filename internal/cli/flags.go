package cli

import (
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
