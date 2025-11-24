package main

import (
	"fmt"

	"magic/api"
	"magic/batcher"
	"magic/compare"
	"magic/csvHandler"
	"magic/types"
)

func main() {

	player1 := new(types.Player)

	//
	player1.Cards = ReadCardList("Avatar.csv")
	collection := player1.GenerateCollection()
	results := batcher.DivideCollection(collection)

	currentMarketPrices1 := api.GetCardData(results)
	player1.UpdateCards(*currentMarketPrices1)
	fmt.Println("cardTotal: ", player1.CardTotal())

	// another one
	player2 := new(types.Player)
	player2.Cards = ReadCardList("duque.csv")
	collection2 := player2.GenerateCollection()
	results2 := batcher.DivideCollection(collection2)

	currentMarketPrices2 := api.GetCardData(results2)
	marketPrices := types.MergePriceMaps(currentMarketPrices1, currentMarketPrices2)
	fmt.Println(results2)
	player2.UpdateCards(marketPrices)
	fmt.Println("cardTotal: ", player2.CardTotal())

	goods := compare.AttemptFillSet(*player1, *player2)
	BGives := ""
	AGives := ""
	BGivesA := types.CardList(goods.AtoB).TotalPrice() - types.CardList(goods.BtoA).TotalPrice()
	AGivesB := types.CardList(goods.BtoA).TotalPrice() - types.CardList(goods.AtoB).TotalPrice()
	if BGivesA > 0 {
		AGives = fmt.Sprintf("receives %.2f", BGivesA)
		BGives = fmt.Sprintf("gives %.2f", BGivesA)

	} else {
		AGives = fmt.Sprintf("gives %.2f", AGivesB)
		BGives = fmt.Sprintf("receives %.2f", AGivesB)

	}

	fmt.Printf("Player1 Gives: %v cards and %v euros\nPlayer2 Gives: %v cards and %v euros\n", len(goods.AtoB), AGives, len(goods.BtoA), BGives)

	// checkGoods(goods)

	WriteCardList("./output/duque_gives_joao.csv", goods.BtoA)
	WriteCardList("./output/joao_gives_duque.csv", goods.AtoB)

	for _, card := range goods.AtoB {
		player1.RemoveCard(card)
		player2.AddCard(card)
	}
	for _, card := range goods.BtoA {
		player2.RemoveCard(card)
		player1.AddCard(card)

	}

	WriteCardList("./output/totalList/duque.csv", player2.Cards)
	WriteCardList("./output/totalList/joao.csv", player1.Cards)
}

func checkGoods(goods types.TransactionRecord) {
	println("Player 1 gives:")
	for _, card := range goods.AtoB {
		fmt.Printf("%v : %v %.2f\n", card.Name, card.Rarity, card.Price())
	}
	println("Player 2 gives:")
	for _, card := range goods.BtoA {
		fmt.Printf("%v : %v %.2f\n", card.Name, card.Rarity, card.Price())

	}
}

func ReadCardList(listPath string) []types.Card {
	fmt.Println(listPath)
	return csvHandler.ReadCSV(listPath)
}

func WriteCardList(listPath string, cards []types.Card) {

	csvHandler.WriteCSV(listPath, cards)
}
