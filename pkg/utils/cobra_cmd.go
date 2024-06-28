package utils

import "strings"

// ContainsHelpFlag проверяет имеется флаг help в списке аргументов
func ContainsHelpFlag(cmd []string) bool {
	if strings.Contains(strings.Join(cmd, " "), "help") {
		return true
	}
	if strings.Contains(strings.Join(cmd, " "), "-h") {
		return true
	}
	return false
}
