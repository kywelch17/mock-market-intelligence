package market

type Status string
type Type string

// Statuses
const (
	StatusOpen Status = "Open"
	StatusClosed Status = "Closed"
	StatusSettled Status = "Settled"
)

// Types
const (
	TypeGame Type = "Game"
	TypeMoney Type = "Money"
	TypeSpread Type = "Spread"
	TypeTotal Type = "Total"
	TypePlayer Type = "Player"
)

type Market struct {
	ID			string
	EventID		string
	Description string
	League		string

	Status		Status
	Type 		Type

	Volume		float64

	Player		*Player // Unless it's a player prop, this can be null

	Outcomes	[]Outcome
}