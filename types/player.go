package types

import (
	"fmt"
	"math"
	"strings"
)

// calculates total amount of moneys from a player in his card inventory
func (p Player) CardTotal() float64 {

	var sum float64
	for _, card := range p.Cards {
		if card.Foil {
			if card.Rarity == "rare" || card.Rarity == "mythic" {
				sum += float64(card.Inventory) * math.Max(card.TrendFoil, 0.5)
			} else {
				sum += float64(card.Inventory) * card.TrendFoil
			}

		} else {
			if card.Rarity == "rare" || card.Rarity == "mythic" {
				sum += float64(card.Inventory) * math.Max(card.Trend, 0.5)
			} else {
				sum += float64(card.Inventory) * card.Trend
			}
		}
	}
	return sum
}

// updates prices of each individual card in his collection
func (p *Player) UpdateCards(prices map[string]map[uint64]Price) {
	for idx, value := range p.Cards {
		set := strings.ToUpper(value.Set)
		nonfoil := float64(prices[set][uint64(value.Number)].NormalPrice)
		foil := float64(prices[set][uint64(value.Number)].FoilPrice)
		rarity := prices[set][uint64(value.Number)].Rarity
		name := prices[set][uint64(value.Number)].Name
		value.SetPrice(nonfoil, foil, rarity, name)
		p.Cards[idx] = value
	}
}

// creates a simplified collection for processing
func (p Player) GenerateCollection() Collection {

	collection := make(Collection)
	for _, card := range p.Cards {
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

// just a lil guy wanting to collect some cards
type Player struct {
	Cards []Card // list of cards
	Name  string // player name
}

func (player *Player) RemoveCard(removeTarget Card) {
	for idx, card := range player.Cards {
		if card.Number == removeTarget.Number &&
			card.Foil == removeTarget.Foil &&
			card.Inventory > removeTarget.Inventory {
			player.Cards[idx].Inventory = card.Inventory - removeTarget.Inventory
		} else if card.Number == removeTarget.Number &&
			card.Foil == removeTarget.Foil &&
			card.Inventory == removeTarget.Inventory {
			player.Cards = append(player.Cards[:idx], player.Cards[idx+1:]...)
		}
	}
}
func (player *Player) AddCard(addCard Card) {
	player.Cards = append(player.Cards, addCard)
}

// displays player collection
func (p Player) String() string {
	total := []string{}

	for _, card := range p.Cards {
		if card.Foil {
			total = append(total, fmt.Sprintf(
				"%v: Name: %v (%v) *Foil* %v€\n",
				card.Number,
				card.Name,
				card.Set,
				card.TrendFoil,
			))
		} else {
			total = append(total, fmt.Sprintf(
				"%v: Name: %v (%v) %v€\n",
				card.Number,
				card.Name,
				card.Set,
				card.Trend,
			))

		}
	}

	return strings.Join(total, "")
}
