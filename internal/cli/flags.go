package cli

import (
	f "gitlab.ozon.dev/yuweebix/homework-1/internal/cli/flags"
)

// инициализируем дефолтные флаги

// orders флаги
var (
	flagOrderID = f.Flag[int]{
		Name:      "order_id",
		Shorthand: "o",
		Usage:     "ID заказа(*)",
		Value:     0,
	}
	flagOrderIDs = f.Flag[[]int]{
		Name:      "order_ids",
		Shorthand: "o",
		Usage:     "IDs заказов(*)",
		Value:     []int{},
	}
	flagUserID = f.Flag[int]{
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
	flagLimit = f.Flag[int]{
		Name:      "limit",
		Shorthand: "l",
		Usage:     "Ограничение по количеству заказов в списке",
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
	flagStart = f.Flag[int]{
		Name:      "start",
		Shorthand: "s",
		Usage:     "Нижняя граница по количеству заказов в списке",
		Value:     0,
	}
	flagFinish = f.Flag[int]{
		Name:      "finish",
		Shorthand: "f",
		Usage:     "Верхняя граница по количеству заказов в списке",
		Value:     0,
	}
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
