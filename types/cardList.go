package types

type CardList []Card

func (cards CardList) TotalPrice() float64 {
	sum := 0.0

	for _, card := range cards {
		sum += card.TotalPrice()
	}
	return sum

}

func (cards CardList) GenerateCollection() Collection {
	collection := make(Collection)
	for _, card := range cards {
		nonEmptySet, ok := collection[card.Set]
		if !ok {
			newSet := MakesNewSet()
			newSet.AddCard(CardId{
				Name:   card.Name,
				Number: card.Number,
				IsFoil: card.Foil,
				Amount: card.Inventory,
			})
			collection[card.Set] = &newSet
		} else {

			nonEmptySet.AddCard(CardId{
				Name:   card.Name,
				Number: card.Number,
				IsFoil: card.Foil,
				Amount: card.Inventory,
			})

		}
	}
	return collection
}
