package models

import "time"

type MarketEvent struct {
    MarketID  string
    OutcomeID string
    Timestamp time.Time

    Price  float64
    Volume float64
    Trades int
}