package utils

import "math"

// convertToMicrocurrency переводит полученную валюту в формате float в bigint для бдшки
func ConvertToMicrocurrency(amount float64) uint64 {
	// 1 рубль = 1_000_000 микрорублей / 1 доллар = 1_000_000 микродолларов
	microamount := amount * 1_000_000
	return uint64(math.Round(microamount))
}
