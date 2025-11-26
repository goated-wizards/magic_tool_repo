package types

type Price struct {
	NormalPrice float32
	FoilPrice   float32
	Rarity      string
	Name        string
	Set         string
	Image       string
}

func MergePriceMaps(a, b *map[string]map[uint64]Price) map[string]map[uint64]Price {
	result := make(map[string]map[uint64]Price)

	// Copy from a
	if a != nil {
		for sym, inner := range *a {
			if inner == nil {
				continue
			}
			result[sym] = make(map[uint64]Price)
			for ts, price := range inner {
				result[sym][ts] = price
			}
		}
	}

	// Merge b (overwrites if same sym+ts exists)
	if b != nil {
		for sym, inner := range *b {
			if inner == nil {
				continue
			}
			if _, exists := result[sym]; !exists {
				result[sym] = make(map[uint64]Price)
			}
			for ts, price := range inner {
				result[sym][ts] = price
			}
		}
	}

	return result
}
