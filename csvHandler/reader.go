package csvHandler

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"magic/types"
)

func ReadCSV(path string) []types.Card {
	csvFile, _ := os.Open(path)
	reader := csv.NewReader(csvFile)
	var cards []types.Card
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		inventory, err := strconv.Atoi(line[0])
		if err != nil {
			fmt.Printf("Line has invalid value (expected uint): %s\n", line[0])
			continue // skip or handle error
		}
		number, err := strconv.Atoi(line[3])
		if err != nil {
			fmt.Printf("got %v and expected uint", line[3])
			continue
		}
		isFoil := line[4] == "*F*"
		cards = append(cards, types.Card{
			Name:      line[1],
			Inventory: uint(inventory),
			Set:       strings.ToUpper(removeBracketsFromSet(line[2])),
			Foil:      isFoil,
			Number:    uint(number),
		})
	}
	return cards
}

func removeBracketsFromSet(str string) string {
	new := strings.ReplaceAll(str, "(", "")
	new = strings.ReplaceAll(new, ")", "")
	return new
}
