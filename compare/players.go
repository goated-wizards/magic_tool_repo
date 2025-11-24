package compare

import (
	"strings"

	"magic/types"
)

func AttemptFillSet(player1, player2 types.Player) types.TransactionRecord {

	Player1Gives, Player2Gives := ComparaCollection(player1.Cards, player2.Cards)

	return types.TransactionRecord{
		AtoB: Player1Gives,
		BtoA: Player2Gives,
	}
}

func ComparaCollection(collectionA, collectionB []types.Card) (AtoB []types.Card, BtoA []types.Card) {

	AtoB, BtoA = []types.Card{}, []types.Card{}
	Acol := make(map[string]map[uint]map[bool]uint)

	for _, card := range collectionA {
		// Ensure first-level map exists
		if _, ok := Acol[card.Set]; !ok {
			Acol[card.Set] = make(map[uint]map[bool]uint)
		}

		// Ensure second-level map exists
		if _, ok := Acol[card.Set][card.Number]; !ok {
			Acol[card.Set][card.Number] = make(map[bool]uint)
		}

		// Sum the inventory
		Acol[card.Set][card.Number][card.Foil] += card.Inventory
	}
	Bcol := make(map[string]map[uint]map[bool]uint)
	for _, card := range collectionB {
		// Ensure first-level map exists
		if _, ok := Bcol[card.Set]; !ok {
			Bcol[card.Set] = make(map[uint]map[bool]uint)
		}

		// Ensure second-level map exists
		if _, ok := Bcol[card.Set][card.Number]; !ok {
			Bcol[card.Set][card.Number] = make(map[bool]uint)
		}

		// Sum the inventory
		Bcol[card.Set][card.Number][card.Foil] += card.Inventory
	}
	for _, card := range collectionA {

		_, okNonFoil := Bcol[card.Set][card.Number][false]
		_, okFoil := Bcol[card.Set][card.Number][true]
		if !okFoil && !okNonFoil && (card.Inventory > 1) {
			AtoB = append(AtoB, types.Card{
				Number:    card.Number,
				Inventory: 1,
				Trend:     card.Trend,
				Foil:      card.Foil,
				Set:       strings.ToUpper(card.Set),
				Name:      card.Name,
				Rarity:    card.Rarity,
			})
			continue
		}

	}
	for _, card := range collectionB {
		_, okNonFoil := Acol[card.Set][card.Number][false]
		_, okFoil := Acol[card.Set][card.Number][true]

		if !okFoil && !okNonFoil && (card.Inventory > 1) {
			BtoA = append(BtoA, types.Card{
				Number:    card.Number,
				Inventory: 1,
				Trend:     card.Trend,
				Foil:      card.Foil,
				Set:       strings.ToUpper(card.Set),
				Name:      card.Name,
				Rarity:    card.Rarity,
			})
			continue
		}
	}
	return
}
