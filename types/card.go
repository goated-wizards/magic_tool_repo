package types

import (
	"fmt"
	"math"
)

// Card specific struct (if you have both foil and non-foil, create both, one for each and specify Foil to true or false in each)
type Card struct {
	Number    uint    // set number / collector number
	Inventory uint    // number of instances of this card
	Trend     float64 // will eventually be autofilled from api with nonfoil price
	TrendFoil float64 // will eventually be autofilled from api with foil price
	Name      string  // card name
	Set       string  // set card you're inventoring
	Foil      bool    // whether card is foil or not
	Rarity    string  // will eventually be autofilled from api
	Image     string  // will eventually be autofilled from api
}

func (c Card) Price() float64 {
	if c.Foil {
		if c.Rarity == "rare" || c.Rarity == "mythic" {
			return math.Max(c.TrendFoil, 0.5)
		}
		if c.Rarity == "rare" || c.Rarity == "mythic" {
			return math.Max(c.TrendFoil, 0.5)
		}
		return c.TrendFoil
	}
	if c.Rarity == "rare" || c.Rarity == "mythic" {
		return math.Max(c.Trend, 0.5)
	}
	if c.Rarity == "rare" || c.Rarity == "mythic" {
		return math.Max(c.Trend, 0.5)
	}
	return c.Trend
}

func (c Card) TotalPrice() float64 {
	return c.Price() * float64(c.Inventory)
}

// sets prices foil and nonfoil on each individual card
func (c *Card) SetPrice(nonfoil, foil float64, rarity, name, image string) {
	c.Trend = float64(nonfoil)
	c.TrendFoil = float64(foil)
	c.Rarity = rarity
	c.Name = name
	c.Image = image

}

// simplified ids for cards for querying
type CardId struct {
	Name   string
	Number uint
	IsFoil bool
	Amount uint
}

// shows card number and card name
func (cId CardId) String() string {
	if cId.IsFoil {
		return fmt.Sprintf("%v: %v *F*", cId.Number, cId.Name)
	}
	return fmt.Sprintf("%v: %v ", cId.Number, cId.Name)
}
