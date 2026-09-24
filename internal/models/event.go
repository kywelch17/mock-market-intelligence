package models

const (
	LeagueMLB = "MLB"
	LeagueNFL = "NFL"
	LeagueNBA = "NBA"
	LeagueNHL = "NHL"
	LeagueMLS = "MLS"
)

type Event struct {
	ID          string
	Type        Type
	League      string
	Description string

	Markets []Market
}
