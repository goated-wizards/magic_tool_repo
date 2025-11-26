package csvHandler

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"

	"magic/types"
)

func WriteCSV(path string, cards []types.Card, withPrice, withImage bool) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatalf("failed to create file: %v", err)
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Data to write
	records := [][]string{
		{
			"inventory", "name", "set", "setnumber", "foil",
		},
	}
	if withPrice {

		records[0] = append(records[0], "price")
	}
	if withImage {

		records[0] = append(records[0], "image")
	}
	sort.Slice(cards, func(i, j int) bool {
		if cards[i].Set != cards[j].Set {
			return cards[i].Set < cards[j].Set // primary sort by Age
		}
		if cards[i].Number != cards[j].Number {
			return cards[i].Number < cards[j].Number // primary sort by Age
		}
		return cards[i].Foil
	})
	for idx, card := range cards {
		var foilString string
		if card.Foil {
			foilString = "*F*"
		}
		newEntry := []string{
			fmt.Sprintf("%v", card.Inventory),
			card.Name,
			card.Set,
			fmt.Sprintf("%v", card.Number),
			foilString,
		}
		if withPrice {
			newEntry = append(newEntry, fmt.Sprintf("%.2f", cards[idx].Price()))
		}
		if withImage {
			newEntry = append(newEntry, cards[idx].Image)
		}
		records = append(records, newEntry)
	}
	// Write each record
	for _, record := range records {
		if err := writer.Write(record); err != nil {
			log.Fatalf("error writing record to csv: %v", err)
		}
	}

	log.Println("CSV file written successfully.")
}
