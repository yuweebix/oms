package cli

import "github.com/spf13/cobra"

// rootCmd стоковая команда
var rootCmd = &cobra.Command{
	Use:   "",
	Short: "Утилита для управления ПВЗ",
	Long: `Утилита содержит перечень команд, что можно производить над заказами и возвратами.

Заказы:
  - "orders accept  [flags]": Принять заказ от курьера
  - "orders return  [flags]": Вернуть заказ курьеру
  - "orders deliver [flags]": Выдать заказ клиенту
  - "orders list    [flags]": Получить список заказов
	
Возвраты:
  - "returns accept [flags]": Принять возврат от клиента
  - "returns list   [flags]": Получить список возвратов`,
}

// ordersCmd команда операций с заказами
var ordersCmd = &cobra.Command{
	Use:   "orders",
	Short: "Совершить операцию с заказом",
	Long: `Команда "orders" содержит перечень команд для обработки заказа.

Для большей информации вызовите команду:
  orders [command] --help
`,
}

// returnsCmd команда операций с возвратами
var returnsCmd = &cobra.Command{
	Use:   "returns",
	Short: "Совершить операцию с возвратом",
	Long: `Команда "returns" содержит перечень команд для обработки возврата.

Для большей информации вызовите команду:
  returns [command] --help`,
}

// workersCmd команда операций с рабочими
var workersCmd = &cobra.Command{
	Use:   "workers",
	Short: "Изменить количество выполняющих команды рабочих горутин",
	Long: `Команда "workers" изменяет количество рабочих [горутин].

Пример использования:
  workers --number 5
	workers --number -5`,
}

// orders подкоманды

// ordersAcceptCmd команда принятия заказа от курьера
var ordersAcceptCmd = &cobra.Command{
	Use:   "accept",
	Short: "Принять заказ от курьера",
	Long: `Команда "accept" используется для принятия заказа от курьера. 

Пример использования:
  orders accept --order_id=420 --user_id=69 --expiry=2025-05-05 --cost=1337.69 --weight=0.69 --packaging=bag

Условия:
  - Заказ не может быть принят дважды.
  - Срок хранения не может быть в прошлом.`,
	SilenceUsage:  true,
	SilenceErrors: true,
}

// ordersDeliverCmd команда выдачи заказов клиенту
var ordersDeliverCmd = &cobra.Command{
	Use:   "deliver",
	Short: "Выдать заказы клиенту",
	Long: `Команда "deliver" используется для выдачи заказов клиенту.

Пример использования:
  orders deliver --order_ids=228,322,420

Условия:
  - Все заказы должны принадлежать одному клиенту и быть приняты от курьера.
  - Срок хранения заказов должен быть больше текущей даты.`,
	SilenceUsage:  true,
	SilenceErrors: true,
}

// ordersListCmd команда получения списка заказов
var ordersListCmd = &cobra.Command{
	Use:   "list",
	Short: "Получить список заказов",
	Long: `Команда "list" используется для получения списка заказов.

Пример использования:
  orders list --user_id=123 --limit=10 --offset=0 --is_stored=true

Команда возвращает заказы клиента, которые находятся в ПВЗ, с возможностью ограничить количество возвращаемых заказов и задать смещение.`,
	SilenceUsage:  true,
	SilenceErrors: true,
}

// ordersReturnCmd команда возврата заказа курьеру
var ordersReturnCmd = &cobra.Command{
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

// returns подкоманды

// returnsAcceptCmd represents the Accept command
var returnsAcceptCmd = &cobra.Command{
	Use:   "accept",
	Short: "Принять возврат от клиента",
	Long: `Команда "accept" используется для принятия возврата заказа от клиента.

Пример использования:
  returns accept --order_id=420 --user_id=69

Условия:
  - Возврат может быть принят в течение двух дней с момента выдачи заказа.
  - Заказ должен быть выдан из этого ПВЗ.`,
	SilenceUsage:  true,
	SilenceErrors: true,
}

// returnsListCmd represents the List command
var returnsListCmd = &cobra.Command{
	Use:   "list",
	Short: "Получить список возвратов",
	Long: `Команда "list" используется для получения списка возвратов.

Пример использования:
  returns list --limit=10 --offset=0

Команда возвращает список возвратов с возможностью пагинации.`,
	SilenceUsage:  true,
	SilenceErrors: true,
}
