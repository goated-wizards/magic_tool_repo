package csvHandler

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"

	"magic/types"
)

func WriteCSV(path string, cards []types.Card) {
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
	sort.Slice(cards, func(i, j int) bool {
		if cards[i].Set != cards[j].Set {
			return cards[i].Set < cards[j].Set // primary sort by Age
		}
		if cards[i].Number != cards[j].Number {
			return cards[i].Number < cards[j].Number // primary sort by Age
		}
		return cards[i].Foil
	})
	for _, card := range cards {
		var foilString string
		if card.Foil {
			foilString = "*F*"
		}
		records = append(records, []string{
			fmt.Sprintf("%v", card.Inventory),
			card.Name,
			card.Set,
			fmt.Sprintf("%v", card.Number),
			foilString,
		})
	}
	// Write each record
	for _, record := range records {
		if err := writer.Write(record); err != nil {
			log.Fatalf("error writing record to csv: %v", err)
		}
	}

	log.Println("CSV file written successfully.")
}
