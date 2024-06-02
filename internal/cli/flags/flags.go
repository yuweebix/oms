package flags

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// наименования флагов (long)
const (
	OrderIDL  = "order_id"
	OrderIDsL = "order_ids"
	UserIDL   = "user_id"
	ExpiryL   = "expiry"
	LimitL    = "limit"
	IsStoredL = "is_stored"
	StartL    = "start"
	FinishL   = "finish"
)

// shorthands
const (
	OrderIDS  = "o"
	OrderIDsS = "o"
	UserIDS   = "u"
	ExpiryS   = "e"
	LimitS    = "l"
	IsStoredS = "i"
	StartS    = "s"
	FinishS   = "f"
)

// usage
const (
	OrderIDU  = "ID заказа(*)"
	OrderIDsU = "IDs заказов(*)"
	UserIDU   = "ID клиента(*)"
	ExpiryU   = "Срок хранения в формате YYYY-MM-DD(*)"
	LimitU    = "Ограничение по количеству заказов в списке"
	IsStoredU = "Показать заказы клиента, находящиеся в нашем ПВЗ"
	StartU    = "Нижняя граница по количеству заказов в списке"
	FinishU   = "Верхняя граница по количеству заказов в списке"
)

// дефолтные значения
const (
	DefaultIntValue    = -1
	DefaultStringValue = ""
	DefaultBoolValue   = false
)

// DefaultIntSliceValue для обнуление значения флага
func DefaultIntSliceValue() []int {
	return []int{}
}

// resetFlags сбрасывает флаги, чтобы они не сохранялись в последующих вызовах
// 1. Проверяем значение флага типа SliceValue или же просто Value
// 2. Поскольку флаги не открыты (i.e. нельзя вызвать pflag.intSlice), нужно использовать строковые типы (.Type())
// Если бы был доступ к pflag.intSliceValue или же pflag.intValue было бы удобнее
func resetFlags(cmd *cobra.Command) {
	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		switch flag.Value.(type) {
		case pflag.SliceValue:
			switch flag.Value.Type() {
			case "intSlice":
				if val, ok := flag.Value.(pflag.SliceValue); ok {
					val.Replace([]string{})
				}
			}
		case pflag.Value:
			switch flag.Value.Type() {
			case "int":
				flag.Value.Set("-1")
			case "bool":
				flag.Value.Set("false")
			}
		}
	})
}

// ResetAllFlags resets the flags of a command and all its children recursively
func ResetAllFlags(cmd *cobra.Command) {
	// Reset the flags of the current command
	resetFlags(cmd)

	// Recursively reset the flags of the children
	for _, child := range cmd.Commands() {
		ResetAllFlags(child)
	}
}
