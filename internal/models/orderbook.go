package models

type OrderBookLevel struct {
	Price    float64
	Quantity int
}

type OrderBook struct {
	OutcomeID string
	Bids      []OrderBookLevel
	Asks      []OrderBookLevel
}
