package types

import (
	"fmt"
	"strings"
)

type Set struct {
	Cards   []CardId
	CardMap map[uint]map[bool]uint
}

func MakesNewSet() Set {
	cardMap := make(map[uint]map[bool]uint)

	return Set{CardMap: cardMap}
}

func (set Set) GetLen() uint {
	var total uint
	for _, value := range set.CardMap {
		val, ok := value[true]
		if ok {
			total += uint(val)
		}
		val, ok = value[false]
		if ok {
			total += uint(val)
		}
	}
	return total
}

func (set *Set) AddCard(card CardId) {
	set.Cards = append(set.Cards, card)
	numCards, ok := set.CardMap[card.Number][card.IsFoil]
	if !ok {
		if set.CardMap[card.Number] == nil {
			set.CardMap[card.Number] = make(map[bool]uint)
		}
		set.CardMap[card.Number][card.IsFoil] = card.Amount
	} else {
		if set.CardMap[card.Number] == nil {
			set.CardMap[card.Number] = make(map[bool]uint)
		}
		set.CardMap[card.Number][card.IsFoil] = numCards + card.Amount
	}
}

// creates search string for a set
func (s Set) SearchString(set string) string {
	if len(s.Cards) == 1 {
		return fmt.Sprintf("set:%v+(number:%v)", set, s.Cards[0].Number)
	}
	total := []string{}
	for _, card := range s.Cards {
		total = append(total, fmt.Sprintf("number:%v OR ", card.Number))
	}
	total[len(total)-1] = strings.ReplaceAll(total[len(total)-1], " OR ", "")

	return fmt.Sprintf("set:%v+(%v)", set, strings.Join(total, ""))
}
