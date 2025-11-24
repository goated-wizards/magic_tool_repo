package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"magic/types"
)

func GetCardData(col types.QueriableCollection) *map[string]map[uint64]types.Price {

	menu := make(map[string]map[uint64]types.Price)
	for _, collection := range col {
		results := makeCall(collection.SearchString())
		time.Sleep(50 * time.Millisecond)
		for _, each := range results.Data {
			errorList := []error{}
			priceEur, err := strconv.ParseFloat(each.Prices.Eur, 32)
			if err != nil {
				fmt.Println("Error:", err)
				fmt.Println(results)
				errorList = append(errorList, errors.New("not possible to parse nonfoil price"))
			}
			priceEurFoil, err := strconv.ParseFloat(each.Prices.Eur, 32)
			if err != nil {
				fmt.Println("Error:", err)
				errorList = append(errorList, errors.New("not possible to parse foil price"))
			}
			if len(errorList) > 1 {
				fmt.Println("invalid cards")
				priceEurFoil = 100000000.0
				priceEur = 100000000.0
			}
			colNumber, err := strconv.ParseUint(each.CollectorNumber, 10, 64)
			if err != nil {
				fmt.Println("Error:", err)
				panic("not possible to parse collector number")
			}
			set := strings.ToUpper(each.Set)
			if menu[set] == nil {
				menu[set] = make(map[uint64]types.Price)
			}
			menu[set][colNumber] = types.Price{
				NormalPrice: float32(priceEur),
				FoilPrice:   float32(priceEurFoil),
				Rarity:      each.Rarity,
				Name:        each.Name,
				Set:         strings.ToUpper(each.Set),
			}

		}
		// fmt.Println(menu)
	}
	return &menu
}

func makeCall(query string) garbageStruct {
	encoded := url.QueryEscape(query)
	url := fmt.Sprintf("https://api.scryfall.com/cards/search?q=%v", encoded)
	fmt.Println("Making call ", url)

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("status code error: %d %s", resp.StatusCode, resp.Status))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var data garbageStruct
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}
	return data
}

type garbageStruct struct {
	TotalCards uint   `json:"total_cards"`
	Data       []data `json:"data"`
}

type data struct {
	Name            string   `json:"name"`
	Colors          []string `json:"colors"`
	ColorIdentity   []string `json:"color_identity"`
	Set             string   `json:"set"`
	CollectorNumber string   `json:"collector_number"`
	Rarity          string   `json:"rarity"`
	Prices          struct {
		Eur     string `json:"eur"`
		EurFoil string `json:"eur_foil"`
	} `json:"prices"`
}
