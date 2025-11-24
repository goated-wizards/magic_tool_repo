package compare_test

import (
	"magic/types"
)

type compareInputs struct {
	player1 types.Player
	player2 types.Player
}
type expectedOutput struct {
	expected types.TransactionRecord
}

type test struct {
	name   string
	inputs compareInputs
	output expectedOutput
}

// func TestAttemptFillSet(t *testing.T) {
// 	tests := []test{
// 		{
// 			name: "both players have overlaping collections with no extras",
// 			inputs: compareInputs{
// 				player1: types.Player{
// 					Name: "Duque",
// 					Cards: []types.Card{
// 						{
// 							Number:    3,
// 							Set:       "WAR",
// 							Inventory: 3,
// 						},
// 						{
// 							Number:    1,
// 							Set:       "WAR",
// 							Inventory: 3,
// 						},
// 					},
// 				},
// 				player2: types.Player{
// 					Name: "Jorge",
// 					Cards: []types.Card{
// 						{
// 							Number:    1,
// 							Set:       "WAR",
// 							Inventory: 3,
// 						},
// 						{
// 							Number:    3,
// 							Set:       "WAR",
// 							Inventory: 2,
// 						},
// 					},
// 				},
// 			},
// 			output: expectedOutput{
// 				expected: types.TransactionRecord{
// 					AtoB: types.GoodsPackage{
// 						DeltaPrice:    0,
// 						GoodsGiven:    types.Collection{},
// 						GoodsReceived: types.Collection{},
// 					},
// 					BtoA: types.GoodsPackage{
// 						DeltaPrice:    0,
// 						GoodsGiven:    types.Collection{},
// 						GoodsReceived: types.Collection{},
// 					},
// 				},
// 			},
// 		},
// 		{
// 			name: "one player has more cards than other but no extras",
// 			inputs: compareInputs{
// 				player1: types.Player{
// 					Name: "Duque",
// 					Cards: []types.Card{
// 						{
// 							Number:    3,
// 							Set:       "WAR",
// 							Inventory: 3,
// 						},
// 						{
// 							Number:    1,
// 							Set:       "WAR",
// 							Inventory: 3,
// 						},
// 						{
// 							Number:    2,
// 							Set:       "WAR",
// 							Inventory: 1,
// 						},
// 					},
// 				},
// 				player2: types.Player{
// 					Name: "Jorge",
// 					Cards: []types.Card{
// 						{
// 							Number:    1,
// 							Set:       "WAR",
// 							Inventory: 3,
// 						},
// 						{
// 							Number:    3,
// 							Set:       "WAR",
// 							Inventory: 2,
// 						},
// 					},
// 				},
// 			},
// 			output: expectedOutput{
// 				expected: types.TransactionRecord{
// 					AtoB: types.GoodsPackage{
// 						DeltaPrice:    0,
// 						GoodsGiven:    types.Collection{},
// 						GoodsReceived: types.Collection{},
// 					},
// 					BtoA: types.GoodsPackage{
// 						DeltaPrice:    0,
// 						GoodsGiven:    types.Collection{},
// 						GoodsReceived: types.Collection{},
// 					},
// 				},
// 			},
// 		},
// 		{
// 			name: "one player has more cards than other and has extras",
// 			inputs: compareInputs{
// 				player1: types.Player{
// 					Name: "Duque",
// 					Cards: []types.Card{
// 						{
// 							Number:    3,
// 							Set:       "WAR",
// 							Inventory: 3,
// 						},
// 						{
// 							Number:    1,
// 							Set:       "WAR",
// 							Inventory: 3,
// 						},
// 						{
// 							Number:    2,
// 							Set:       "WAR",
// 							Inventory: 2,
// 						},
// 					},
// 				},
// 				player2: types.Player{
// 					Name: "Jorge",
// 					Cards: []types.Card{
// 						{
// 							Number:    1,
// 							Set:       "WAR",
// 							Inventory: 3,
// 						},
// 						{
// 							Number:    3,
// 							Set:       "WAR",
// 							Inventory: 2,
// 						},
// 					},
// 				},
// 			},
// 			output: expectedOutput{
// 				expected: types.TransactionRecord{
// 					AtoB: types.GoodsPackage{
// 						DeltaPrice:    0,
// 						GoodsGiven:    types.Collection{},
// 						GoodsReceived: types.Collection{},
// 					},
// 					BtoA: types.GoodsPackage{
// 						DeltaPrice:    0,
// 						GoodsGiven:    types.Collection{},
// 						GoodsReceived: types.Collection{},
// 					},
// 				},
// 			},
// 		},
// 	}
// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			got := compare.AttemptFillSet(test.inputs.player1, test.inputs.player2)
// 			if got != test.output.expected {
// 				t.Fatalf("expected %v, got %v\n", test.output.expected, got)
// 			}
// 		})
// 	}
// }
