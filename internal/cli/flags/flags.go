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
func ResetFlags(cmd *cobra.Command) {
	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		switch flag.Value.(type) {
		case pflag.SliceValue:
			flag.Value.(pflag.SliceValue).Replace([]string{})
			flag.Changed = false
		case pflag.Value:
			flag.Value.Set("")
			flag.Changed = false
		}
	})
}

// ResetAllFlags рекурсивно сбрасывает все флаги
func ResetAllFlags(cmd *cobra.Command) {
	// флаги команды
	ResetFlags(cmd)

	// флаги подкоманд
	for _, child := range cmd.Commands() {
		ResetAllFlags(child)
	}
}

// ResetAllHelpFlags рекурсивно сбрасывает все флаги "help"
func ResetAllHelpFlags(cmd *cobra.Command) {
	if f := cmd.Flag("help"); f != nil {
		f.Value.Set("false")
		f.Changed = false
	}

	for _, child := range cmd.Commands() {
		ResetAllHelpFlags(child)
	}
}
