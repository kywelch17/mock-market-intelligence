package models

import "time"

type Side string

const (
	Buy  Side = "Buy"
	Sell Side = "Sell"
)

type Trade struct {
	MarketID  string
	OutcomeID string
	Side      string
	Price     float64
	Quantity  int
	Timestamp time.Time
}
