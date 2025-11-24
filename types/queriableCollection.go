package types

import (
	"fmt"
	"strings"
)

type QueriableCollection []Collection

// displays a queriable collection summary
func (qCol QueriableCollection) String() string {
	total := []string{}
	for _, col := range qCol {
		for key, value := range col {
			total = append(total, fmt.Sprintf("Set: %v number of cards:%v unique Cards: %v\n", key, value.GetLen(), len(value.Cards)))
		}
	}
	return strings.Join(total, "")
}

// creates a query string to call scryfall without problems
func (qCol QueriableCollection) QueryStrings() []string {
	total := []string{}
	for _, col := range qCol {
		for key, value := range col {
			total = append(total, value.SearchString(key))
		}
	}
	return total
}
