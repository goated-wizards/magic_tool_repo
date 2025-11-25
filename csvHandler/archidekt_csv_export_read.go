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

func ArchidektReadCSV(path string) []types.Card {
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
		number, err := strconv.Atoi(line[4])
		if err != nil {
			fmt.Printf("got %v and expected uint", line[3])
			continue
		}
		isFoil := line[3] == "Foil"
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
