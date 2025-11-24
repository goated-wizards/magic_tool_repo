package types

import (
	"fmt"
	"strings"
)

type Collection map[string]*Set

// displays collection in a nice way
func (c Collection) String() string {
	var total []string
	for set, cards := range c {
		total = append(total, fmt.Sprintf("%v:\n", set))
		for _, card := range cards.Cards {
			total = append(total, fmt.Sprintf("\t%v\n", card))
		}

	}
	return strings.Join(total, "")

}

// creates searchstring to search in scryfall api
func (c Collection) SearchString() string {
	if len(c) == 1 {
		for key, value := range c {
			return fmt.Sprintf("%v", value.SearchString(key))
		}
	}
	total := []string{}
	for key, value := range c {
		total = append(total, fmt.Sprintf("(%v) OR ", value.SearchString(key)))

	}

	total[len(total)-1] = strings.ReplaceAll(total[len(total)-1], " OR ", "")

	return fmt.Sprintf("%v", strings.Join(total, ""))
}
