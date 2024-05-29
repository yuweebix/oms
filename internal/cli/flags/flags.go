package flags

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// ResetFlags сбрасывает флаги, чтобы они не сохранялись в последующих вызовах
// 1. Проверяем значение флага типа слайс и просто
// 2. Поскольку флаги не открыты, нужно использовать строковые типы
// Если бы был доступ к pflag.intSliceValue или же pflag.intValue было бы удобнее
func ResetFlags(cmd *cobra.Command) {
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
			}
		}
	})
}
