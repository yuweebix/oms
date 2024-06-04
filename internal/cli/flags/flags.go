package flags

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// Flag структура параметров, необходимых для объявления флага в кобре
type Flag[T any] struct {
	Name      string
	Shorthand string
	Usage     string
	Value     T
}

// Unzip метод для распаковки значений в сеттер
func (f *Flag[any]) Unzip() (string, string, any, string) {
	return f.Name, f.Shorthand, f.Value, f.Usage
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
				flag.Value.Set("0")
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
