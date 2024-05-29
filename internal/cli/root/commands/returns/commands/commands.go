package commands

import "github.com/spf13/cobra"

// acceptCmd represents the Accept command
var acceptCmd = &cobra.Command{
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

// listCmd represents the List command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Получить список возвратов",
	Long: `Команда "list" используется для получения списка возвратов.

	Пример использования:
  returns list --start=1 --finish=10

Команда возвращает список возвратов с возможностью пагинации.`,
	SilenceUsage:  true,
	SilenceErrors: true,
}
