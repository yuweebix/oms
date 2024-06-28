package cli

import "github.com/spf13/cobra"

// rootCmd стоковая команда
func (c *CLI) rootCmd() (cmd *cobra.Command, err error) {
	cmd = &cobra.Command{
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

	return cmd, nil
}

// ordersCmd команда операций с заказами
func (c *CLI) ordersCmd() (cmd *cobra.Command, err error) {
	cmd = &cobra.Command{Use: "orders",
		Short: "Совершить операцию с заказом",
		Long: `Команда "orders" содержит перечень команд для обработки заказа.

Для большей информации вызовите команду:
  orders [command] --help`,
	}

	return cmd, nil
}

// returnsCmd команда операций с возвратами
func (c *CLI) returnsCmd() (cmd *cobra.Command, err error) {
	cmd = &cobra.Command{Use: "returns",
		Short: "Совершить операцию с возвратом",
		Long: `Команда "returns" содержит перечень команд для обработки возврата.

Для большей информации вызовите команду:
  returns [command] --help`,
	}

	return cmd, nil
}

// workersCmd команда операций с рабочими
func (c *CLI) workersCmd() (cmd *cobra.Command, err error) {
	cmd = &cobra.Command{Use: "workers",
		Short: "Изменить количество выполняющих команды рабочих горутин",
		Long: `Команда "workers" изменяет количество рабочих [горутин].

Пример использования:
  workers --number 5
	workers --number -5`,
		RunE: c.workersCmdRunE,
	}

	err = workersCmdSetFlags(cmd)
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

// orders подкоманды

// ordersAcceptCmd команда принятия заказа от курьера
func (c *CLI) ordersAcceptCmd() (cmd *cobra.Command, err error) {
	cmd = &cobra.Command{Use: "accept",
		Short: "Принять заказ от курьера",
		Long: `Команда "accept" используется для принятия заказа от курьера. 

Пример использования:
  orders accept --order_id=420 --user_id=69 --expiry=2025-05-05 --cost=1337.69 --weight=0.69 --packaging=bag

Условия:
  - Заказ не может быть принят дважды.
  - Срок хранения не может быть в прошлом.`,
		RunE:          c.ordersAcceptCmdRunE,
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	err = ordersAcceptCmdSetFlags(cmd)
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

// ordersDeliverCmd команда выдачи заказов клиенту
func (c *CLI) ordersDeliverCmd() (cmd *cobra.Command, err error) {
	cmd = &cobra.Command{Use: "deliver",
		Short: "Выдать заказы клиенту",
		Long: `Команда "deliver" используется для выдачи заказов клиенту.

Пример использования:
  orders deliver --order_ids=228,322,420

Условия:
  - Все заказы должны принадлежать одному клиенту и быть приняты от курьера.
  - Срок хранения заказов должен быть больше текущей даты.`,
		RunE:          c.ordersDeliverCmdRunE,
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	err = ordersDeliverCmdSetFlags(cmd)
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

// ordersListCmd команда получения списка заказов
func (c *CLI) ordersListCmd() (cmd *cobra.Command, err error) {
	cmd = &cobra.Command{Use: "list",
		Short: "Получить список заказов",
		Long: `Команда "list" используется для получения списка заказов.

Пример использования:
  orders list --user_id=123 --limit=10 --offset=0 --is_stored=true

Команда возвращает заказы клиента, которые находятся в ПВЗ, с возможностью ограничить количество возвращаемых заказов и задать смещение.`,
		RunE:          c.ordersListCmdRunE,
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	err = ordersListCmdSetFlags(cmd)
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

// ordersReturnCmd команда возврата заказа курьеру
func (c *CLI) ordersReturnCmd() (cmd *cobra.Command, err error) {
	cmd = &cobra.Command{Use: "return",
		Short: "Вернуть заказ курьеру",
		Long: `Команда "return" используется для возврата заказа курьеру.

	Пример использования:
  orders return --order_id=1337

Условия:
  - Заказ может быть возвращен только если истек срок хранения и он не был выдан клиенту.`,
		RunE:          c.ordersReturnCmdRunE,
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	err = ordersReturnCmdSetFlags(cmd)
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

// returns подкоманды

// returnsAcceptCmd represents the Accept command
func (c *CLI) returnsAcceptCmd() (cmd *cobra.Command, err error) {
	cmd = &cobra.Command{Use: "accept",
		Short: "Принять возврат от клиента",
		Long: `Команда "accept" используется для принятия возврата заказа от клиента.

Пример использования:
  returns accept --order_id=420 --user_id=69

Условия:
  - Возврат может быть принят в течение двух дней с момента выдачи заказа.
  - Заказ должен быть выдан из этого ПВЗ.`,
		RunE:          c.returnsAcceptCmdRunE,
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	err = returnsAcceptCmdSetFlags(cmd)
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

// returnsListCmd represents the List command
func (c *CLI) returnsListCmd() (cmd *cobra.Command, err error) {
	cmd = &cobra.Command{Use: "list",
		Short: "Получить список возвратов",
		Long: `Команда "list" используется для получения списка возвратов.

Пример использования:
  returns list --limit=10 --offset=0

Команда возвращает список возвратов с возможностью пагинации.`,
		RunE:          c.returnsListCmdRunE,
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	err = returnsListCmdSetFlags(cmd)
	if err != nil {
		return nil, err
	}

	return cmd, nil
}
