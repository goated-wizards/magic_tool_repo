package batcher

import "magic/types"

func DivideCollection(collection types.Collection) types.QueriableCollection {

	collections := types.QueriableCollection{}
	for key, value := range collection {
		i, j := 0, 0
		batchSize := 41
		newSmallCollection := types.Collection{}
		for i < len(value.CardMap) {

			if j > batchSize {
				collections = append(collections, newSmallCollection)
				newSmallCollection = types.Collection{}
				j = 0
			}
			if j == 0 {
				newSet := types.MakesNewSet()

				newSet.AddCard(types.CardId{
					Name:   collection[key].Cards[i].Name,
					Number: collection[key].Cards[i].Number,
					IsFoil: collection[key].Cards[i].IsFoil,
					Amount: collection[key].Cards[i].Amount,
				})

				newSmallCollection[key] = &newSet

			} else {
				newSmallCollection[key].AddCard(types.CardId{
					Name:   collection[key].Cards[i].Name,
					Number: collection[key].Cards[i].Number,
					IsFoil: collection[key].Cards[i].IsFoil,
					Amount: collection[key].Cards[i].Amount,
				})
			}
			i++
			j++
		}
		if j < batchSize {
			collections = append(collections, newSmallCollection)
		}
	}
	return collections
}
