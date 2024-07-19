package utils

import "math"

// convertToMicrocurrency переводит полученную валюту в формате float в bigint для бдшки
func ConvertToMicrocurrency(amount float64) uint64 {
	// 1 рубль = 1_000_000 микрорублей / 1 доллар = 1_000_000 микродолларов
	microamount := amount * 1_000_000
	return uint64(math.Round(microamount))
}

// convertToMicrocurrency переводит из микровалюты в нормальный вид
func ConvertFromMicrocurrency(microamount uint64) float64 {
	// обратно
	amount := float64(microamount) / 1_000_000.0
	return math.Floor(amount)
}
