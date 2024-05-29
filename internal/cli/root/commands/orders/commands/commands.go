package commands

import "github.com/spf13/cobra"

// acceptCmd represents the Accept command
var acceptCmd = &cobra.Command{
	Use:   "accept",
	Short: "Принять заказ от курьера",
	Long: `Команда "accept" используется для принятия заказа от курьера. 

Пример использования:
  orders accept --order_id=420 --user_id=69 --expiry=7

Условия:
  - Заказ не может быть принят дважды.
  - Срок хранения не может быть в прошлом.`,
	SilenceUsage:  true,
	SilenceErrors: true,
}

// deliverCmd represents the Deliver command
var deliverCmd = &cobra.Command{
	Use:   "deliver",
	Short: "Выдать заказ клиенту",
	Long: `Команда "deliver" используется для выдачи заказов клиенту.

Пример использования:
  orders deliver --order_ids=228,322,420

Условия:
  - Все заказы должны принадлежать одному клиенту и быть приняты от курьера.
  - Срок хранения заказов должен быть больше текущей даты.`,
	SilenceUsage:  true,
	SilenceErrors: true,
}

// listCmd represents the List command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Получить список заказов",
	Long: `Команда "list" используется для получения списка заказов.

Пример использования:
  orders list --limit=10

Команда возвращает заказы клиента, которые находятся в ПВЗ, с возможностью ограничить количество возвращаемых заказов.`,
	SilenceUsage:  true,
	SilenceErrors: true,
}

// returnCmd represents the Return command
var returnCmd = &cobra.Command{
	Use:   "return",
	Short: "Вернуть заказ курьеру",
	Long: `Команда "return" используется для возврата заказа курьеру.

	Пример использования:
  orders return --order_id=1337

Условия:
  - Заказ может быть возвращен только если истек срок хранения и он не был выдан клиенту.`,
	SilenceUsage:  true,
	SilenceErrors: true,
}
