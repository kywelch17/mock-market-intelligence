package models

type Outcome struct {
	ID          string
	Description string
	Index       int // 0 means Home/Over/Win - 1 means Away/Under/Lose
}
