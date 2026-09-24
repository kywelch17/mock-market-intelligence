package market

const (
	LeagueMLB = "MLB"
	LeagueNFL = "NFL"
	LeagueNBA = "NBA"
	LeagueNHL = "NHL"
	LeagueMLS = "MLS"
)

type Event struct {
	ID			string
	Type		string
	League		string
	Description	string
	
	Markets		[]Market
}