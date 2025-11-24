package types

// represents a transaction
type TransactionRecord struct {
	AtoB []Card
	BtoA []Card
}

// Structure holding what a player receives
type GoodsPackage struct {
	DeltaPrice    float64
	GoodsGiven    Collection
	GoodsReceived Collection
}
