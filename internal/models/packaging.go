package models

type PackagingType string

type Packaging struct {
	Cost        uint64  `json:"cost"`
	WeightLimit float64 `json:"weight_limit"`
}

var packagingTable = map[PackagingType]Packaging{
	"bag":  {Cost: 5, WeightLimit: 10},
	"box":  {Cost: 20, WeightLimit: 30},
	"wrap": {Cost: 1, WeightLimit: 0}, // ограничения по весу нету
}

// GetPackaging передает значение таблицы, если он присутствует
func GetPackaging(t PackagingType) (Packaging, bool) {
	res, ok := packagingTable[t]
	return res, ok
}
