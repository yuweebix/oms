package commands

import (
	f "gitlab.ozon.dev/yuweebix/homework-1/internal/cli/flags"
)

// флаги для сеттера кобры
var (
	flagOrderID = f.Flag[int]{
		Name:      "order_id",
		Shorthand: "o",
		Usage:     "ID заказа(*)",
		Value:     0,
	}
	flagUserID = f.Flag[int]{
		Name:      "user_id",
		Shorthand: "u",
		Usage:     "ID клиента(*)",
		Value:     0,
	}
	flagStart = f.Flag[int]{
		Name:      "flagStart",
		Shorthand: "s",
		Usage:     "Нижняя граница по количеству заказов в списке",
		Value:     0,
	}
	flagFinish = f.Flag[int]{
		Name:      "flagFinish",
		Shorthand: "f",
		Usage:     "Верхняя граница по количеству заказов в списке",
		Value:     0,
	}
)
